package mongo

//TODO: add customer address to save on mongo db.
//TODO: error handling in converts.

import (
	"context"
	"domain/customer"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoRepository is data repository.
type MongoRepository struct {
	db        *mongo.Database
	customers *mongo.Collection
}

const (
	dataBaseName   = "DeliveryWithGoLangAndDDD"
	collectionName = "Customers"
)

/* mongoCustomer is an internal type that is used to store a Customer
we make an internal struct for this to avoid coupling this mongo implementation to the customer.
Mongo uses bson so we add tags for that
*/
type mongoCustomer struct {
	id          uuid.UUID `bson:"id"`
	firstName   string    `bson:"first_name"`
	lastName    string    `bson:"last_name"`
	phoneNumber string    `bson:"phone_number"`
	email       string    `bson:"email"`
}

func (mongoRepository *MongoRepository) findCustomerItemByID(id uuid.UUID) (*mongo.Cursor, error) {
	Context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mongoRepository.customers.Find(Context, bson.M{"id": id})
}

//convertMongoItemToCustomer takes in a internal structure and converts into customer.
func (mongocustumer mongoCustomer) convertMongoItemToCustomer() customer.Customer {
	Customer := customer.Customer{}
	Customer.SetID(mongocustumer.id)
	Customer.SetFirstName(mongocustumer.firstName)
	Customer.SetLastName(mongocustumer.lastName)
	Customer.SetPhoneNumber(mongocustumer.phoneNumber)
	Customer.SetEmail(mongocustumer.email)
	return Customer
}

func New(connectionContext context.Context, connectionString string) (*MongoRepository, error) {
	db, colletcion, err := connectDataBaseAndCollection(connectionContext, connectionString, dataBaseName, collectionName)
	if err != nil {
		return nil, err
	}
	return &MongoRepository{
		db:        db,
		customers: colletcion,
	}, nil
}
func (mongoRepository *MongoRepository) Get(id uuid.UUID) (customer.Customer, error) {
	mongoCursor, err := mongoRepository.findCustomerItemByID(id)
	if err != nil {
		return customer.Customer{}, err
	}
	mongocustomer, err := decodeMongoCursorToMongoCustomer(mongoCursor)
	if err != nil {
		return customer.Customer{}, err
	}
	return mongocustomer.convertMongoItemToCustomer(), nil
}

func (mongoRepository *MongoRepository) Add(newCustomer customer.Customer) error {
	mongocustomer := convertCustomerToMongoItem(newCustomer)
	err := mongoRepository.insertMongoCustomerIntoDataBase(mongocustomer)
	return err
}

func (mongoRepository *MongoRepository) insertMongoCustomerIntoDataBase(mongocustomer mongoCustomer) error {
	Context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := mongoRepository.customers.InsertOne(Context, mongocustomer)
	return err
}

func convertCustomerToMongoItem(customerStuct customer.Customer) mongoCustomer {
	return mongoCustomer{
		id:          customerStuct.GetID(),
		firstName:   customerStuct.GetFirstName(),
		lastName:    customerStuct.GetLasttName(),
		phoneNumber: customerStuct.GetPhoneNumber(),
		email:       customerStuct.GetEmail(),
	}
}

func decodeMongoCursorToMongoCustomer(mongoCursor *mongo.Cursor) (mongocustomer mongoCustomer, err error) {
	err = mongoCursor.Decode(&mongocustomer)
	return
}

func connectDataBaseAndCollection(connectionContext context.Context, connectionString, dbName, collectName string) (db *mongo.Database, colletcion *mongo.Collection, err error) {
	mongoClient, err := tryToConnectDataBase(connectionContext, connectionString)
	if err != nil {
		return db, colletcion, err
	}
	db = mongoClient.Database(dbName)
	colletcion = db.Collection(collectName)
	return
}

func tryToConnectDataBase(connectionContext context.Context, connectionString string) (*mongo.Client, error) {
	return mongo.Connect(connectionContext, options.Client().ApplyURI(connectionString))
}
