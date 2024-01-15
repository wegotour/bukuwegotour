package PasetoBackend

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

// user
func TestCreateNewUserRole(t *testing.T) {
	var userdata User
	userdata.Username = "prisyahaura"
	userdata.Email = "prisyahaura@gmail.com"
	userdata.Password = "picaw"
	userdata.Role = "user"
	mconn := SetConnection("MONGOSTRING", "wegotour")
	CreateNewUserRole(mconn, "user", userdata)
}

// user
func CreateNewUserToken(t *testing.T) {
	var userdata User
	userdata.Username = "prisyahaura"
	userdata.Email = "prisyahaura@gmail.com"
	userdata.Password = "picaw"
	userdata.Role = "user"

	// Create a MongoDB connection
	mconn := SetConnection("MONGOSTRING", "wegotour")

	// Call the function to create a user and generate a token
	err := CreateUserAndAddToken("your_private_key_env", mconn, "user", userdata)

	if err != nil {
		t.Errorf("Error creating user and token: %v", err)
	}
}

// user
func TestGFCPostHandlerUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "wegotour")
	var userdata User
	userdata.Username = "prisyahaura"
	userdata.Email = "prisyahaura@gmail.com"
	userdata.Password = "picaw"
	userdata.Role = "user"
	CreateNewUserRole(mconn, "user", userdata)
}

// Test Insert Ticket
func TestInsertTicket(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "wegotour")
	var ticketdata Ticket
	ticketdata.Nomorid = 1
	ticketdata.Title = "garut"
	ticketdata.Description = "waw garut keren banget"
	ticketdata.Image = "https://images3.alphacoders.com/165/thumb-1920-165265.jpg"
	CreateNewTicket(mconn, "ticket", ticketdata)
}

// Test All Ticket
func TestAllTicket(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "wegotour")
	ticket := GetAllTicket(mconn, "ticket")
	fmt.Println(ticket)
}

func TestGeneratePasswordHash(t *testing.T) {
	password := "picaw"
	hash, _ := HashPass(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)
	match := CompareHashPass(password, hash)
	fmt.Println("Match:   ", match)
}

// pasetokey
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("wegotour", privateKey)
	fmt.Println(hasil, err)
}

// user
func TestHashFunctionUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "wegotour")
	var userdata User
	userdata.Username = "prisyahaura"
	userdata.Email = "prisyahaura@gmail.com"
	userdata.Password = "picaw"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[Admin](mconn, "user", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPass(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CompareHashPass(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)
}

// user
func TestUserFix(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "wegotour")
	var userdata User
	userdata.Username = "prisyahaura"
	userdata.Email = "prisyahaura@gmail.com"
	userdata.Password = "picaw"
	userdata.Role = "user"
	CreateUser(mconn, "user", userdata)
}
