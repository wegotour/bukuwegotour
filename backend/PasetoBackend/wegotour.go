package PasetoBackend

import (
	"context"
	"fmt"
	"os"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// crud
func GetAllDocs(db *mongo.Database, col string, docs interface{}) interface{} {
	collection := db.Collection(col)
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error GetAllDocs %s: %s", col, err)
	}
	err = cursor.All(context.TODO(), &docs)
	if err != nil {
		return err
	}
	return docs
}

func UpdateOneDoc(id primitive.ObjectID, db *mongo.Database, col string, doc interface{}) (err error) {
	filter := bson.M{"_id": id}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		return fmt.Errorf("error update: %v", err)
	}
	if result.ModifiedCount == 0 {
		err = fmt.Errorf("tidak ada data yang diubah")
		return
	}
	return nil
}

func DeleteOneDoc(_id primitive.ObjectID, db *mongo.Database, col string) error {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// user
func CreateNewUserRole(mongoconn *mongo.Database, collection string, userdata User) interface{} {
	// Hash the password before storing it
	hashedPassword, err := HashPass(userdata.Password)
	if err != nil {
		return err
	}
	userdata.Password = hashedPassword

	// Insert the admin data into the database
	return atdb.InsertOneDoc(mongoconn, collection, userdata)
}

// user
func CreateUserAndAddToken(privateKeyEnv string, mongoconn *mongo.Database, collection string, userdata User) error {
	// Hash the password before storing it
	hashedPassword, err := HashPass(userdata.Password)
	if err != nil {
		return err
	}
	userdata.Password = hashedPassword

	// Create a token for the admin
	tokenstring, err := watoken.Encode(userdata.Username, os.Getenv(privateKeyEnv))
	if err != nil {
		return err
	}

	userdata.Token = tokenstring

	// Insert the user data into the MongoDB collection
	if err := atdb.InsertOneDoc(mongoconn, collection, userdata.Username); err != nil {
		return nil // Mengembalikan kesalahan yang dikembalikan oleh atdb.InsertOneDoc
	}

	// Return nil to indicate success
	return nil
}

func CreateResponse(status bool, message string, data interface{}) Response {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return response
}

// admin
func CreateAdmin(mongoconn *mongo.Database, collection string, admindata Admin) interface{} {
	// Hash the password before storing it
	hashedPassword, err := HashPass(admindata.Password)
	if err != nil {
		return err
	}
	privateKey, publicKey := watoken.GenerateKey()
	adminid := admindata.Email
	tokenstring, err := watoken.Encode(adminid, privateKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tokenstring)
	// decode token to get adminid
	adminidstring := watoken.DecodeGetId(publicKey, tokenstring)
	if adminidstring == "" {
		fmt.Println("expire token")
	}
	fmt.Println(adminidstring)
	admindata.Private = privateKey
	admindata.Public = publicKey
	admindata.Password = hashedPassword

	// Insert the admin data into the database
	return atdb.InsertOneDoc(mongoconn, collection, admindata)
}

// user
func CreateUser(mongoconn *mongo.Database, collection string, userdata User) interface{} {
	// Hash the password before storing it
	hashedPassword, err := HashPass(userdata.Password)
	if err != nil {
		return err
	}
	privateKey, publicKey := watoken.GenerateKey()
	userid := userdata.Email
	tokenstring, err := watoken.Encode(userid, privateKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tokenstring)
	// decode token to get userid
	useridstring := watoken.DecodeGetId(publicKey, tokenstring)
	if useridstring == "" {
		fmt.Println("expire token")
	}
	fmt.Println(useridstring)
	userdata.Private = privateKey
	userdata.Public = publicKey
	userdata.Password = hashedPassword

	// Insert the user data into the database
	return atdb.InsertOneDoc(mongoconn, collection, userdata)
}

// ticket
func CreateNewTicket(mongoconn *mongo.Database, collection string, ticketdata Ticket) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, ticketdata)
}

// ticket function
func insertTicket(mongoconn *mongo.Database, collection string, ticketdata Ticket) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, ticketdata)
}

func DeleteTicket(mongoconn *mongo.Database, collection string, ticketdata Ticket) interface{} {
	filter := bson.M{"nomorid": ticketdata.Nomorid}
	return atdb.DeleteOneDoc(mongoconn, collection, filter)
}

func UpdatedTicket(mongoconn *mongo.Database, collection string, filter bson.M, ticketdata Ticket) interface{} {
	updatedFilter := bson.M{"nomorid": ticketdata.Nomorid}
	return atdb.ReplaceOneDoc(mongoconn, collection, updatedFilter, ticketdata)
}

func GetAllTicket(mongoconn *mongo.Database, collection string) []Ticket {
	ticket := atdb.GetAllDoc[[]Ticket](mongoconn, collection)
	return ticket
}

func GetAllTicketID(mongoconn *mongo.Database, collection string, ticketdata Ticket) Ticket {
	filter := bson.M{
		"nomorid":     ticketdata.Nomorid,
		"title":       ticketdata.Title,
		"description": ticketdata.Description,
		"image":       ticketdata.Image,
	}
	ticketID := atdb.GetOneDoc[Ticket](mongoconn, collection, filter)
	return ticketID
}

// about function

func InsertAbout(mongoconn *mongo.Database, collection string, aboutdata About) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, aboutdata)
}

func DeleteAbout(mongoconn *mongo.Database, collection string, aboutdata About) interface{} {
	filter := bson.M{"id": aboutdata.ID}
	return atdb.DeleteOneDoc(mongoconn, collection, filter)
}

func UpdatedAbout(mongoconn *mongo.Database, collection string, filter bson.M, aboutdata About) interface{} {
	updatedFilter := bson.M{"id": aboutdata.ID}
	return atdb.ReplaceOneDoc(mongoconn, collection, updatedFilter, aboutdata)
}

func GetAllAbout(mongoconn *mongo.Database, collection string) []About {
	about := atdb.GetAllDoc[[]About](mongoconn, collection)
	return about
}

// contact function

func InsertContact(mongoconn *mongo.Database, collection string, contactdata Contact) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, contactdata)
}

func GetAllContact(mongoconn *mongo.Database, collection string) []Contact {
	contact := atdb.GetAllDoc[[]Contact](mongoconn, collection)
	return contact
}

func GetIdContact(mongoconn *mongo.Database, collection string, contactdata Contact) Contact {
	filter := bson.M{"id": contactdata.ID}
	return atdb.GetOneDoc[Contact](mongoconn, collection, filter)
}
