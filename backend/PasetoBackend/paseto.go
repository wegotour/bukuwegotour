package PasetoBackend

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

// <--- ini Login Email & Register Email --->
func Login(Privatekey, MongoEnv, dbname, Colname string, r *http.Request) string {
	var resp Credential
	mconn := SetConnection(MongoEnv, dbname)
	var dataadmin Admin
	err := json.NewDecoder(r.Body).Decode(&dataadmin)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, Colname, dataadmin) {
			tokenstring, err := watoken.Encode(dataadmin.Email, os.Getenv(Privatekey))
			if err != nil {
				resp.Message = "Gagal Encode Token : " + err.Error()
			} else {
				resp.Status = true
				resp.Message = "Selamat Datang SUPERADMIN"
				resp.Token = tokenstring
			}
		} else {
			resp.Message = "Password Salah"
		}
	}
	return GCFReturnStruct(resp)
}

// return struct
func GCFReturnStruct(DataStruct any) string {
	jsondata, _ := json.Marshal(DataStruct)
	return string(jsondata)
}

func ReturnStringStruct(Data any) string {
	jsonee, _ := json.Marshal(Data)
	return string(jsonee)
}

func Register(Mongoenv, dbname string, r *http.Request) string {
	resp := new(Credential)
	admindata := new(Admin)
	resp.Status = false
	conn := SetConnection(Mongoenv, dbname)
	err := json.NewDecoder(r.Body).Decode(&admindata)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		resp.Status = true
		hash, err := HashPass(admindata.Password)
		if err != nil {
			resp.Message = "Gagal Hash Password" + err.Error()
		}
		InsertAdmindata(conn, admindata.Email, admindata.Role, hash)
		resp.Message = "Berhasil Input data"
	}
	response := ReturnStringStruct(resp)
	return response
}

// <--- ini ticket --->

// ticket post
func GCFInsertTicket(publickey, MONGOCONNSTRINGENV, dbname, colladmin, collticket string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var admindata Admin
	gettoken := r.Header.Get("Login")
	if gettoken == "" {
		response.Message = "Header Login Not Exist"
	} else {
		// Process the request with the "Login" token
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		admindata.Email = checktoken
		if checktoken == "" {
			response.Message = "Kamu kayaknya belum punya akun"
		} else {
			user2 := FindAdmin(mconn, colladmin, admindata)
			if user2.Role == "admin" {
				var dataticket Ticket
				err := json.NewDecoder(r.Body).Decode(&dataticket)
				if err != nil {
					response.Message = "Error parsing application/json: " + err.Error()
				} else {
					insertTicket(mconn, collticket, Ticket{
						Nomorid:     dataticket.Nomorid,
						Title:       dataticket.Title,
						Description: dataticket.Description,
						Image:       dataticket.Image,
						Status:      dataticket.Status,
					})
					response.Status = true
					response.Message = "Berhasil Insert Ticket"
				}
			} else {
				response.Message = "Anda tidak dapat Insert data karena bukan admin"
			}
		}
	}
	return GCFReturnStruct(response)
}

// delete ticket
func GCFDeleteTicket(publickey, MONGOCONNSTRINGENV, dbname, colladmin, collticket string, r *http.Request) string {

	var respon Credential
	respon.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var admindata Admin

	gettoken := r.Header.Get("Login")
	if gettoken == "" {
		respon.Message = "Header Login Not Exist"
	} else {
		// Process the request with the "Login" token
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		admindata.Email = checktoken
		if checktoken == "" {
			respon.Message = "Kamu kayaknya belum punya akun"
		} else {
			admin2 := FindAdmin(mconn, colladmin, admindata)
			if admin2.Role == "admin" {
				var dataticket Ticket
				err := json.NewDecoder(r.Body).Decode(&dataticket)
				if err != nil {
					respon.Message = "Error parsing application/json: " + err.Error()
				} else {
					DeleteTicket(mconn, collticket, dataticket)
					respon.Status = true
					respon.Message = "Berhasil Delete Ticket"
				}
			} else {
				respon.Message = "Anda tidak bisa Delete data karena bukan admin"
			}
		}
	}
	return GCFReturnStruct(respon)
}

// update ticket
func GCFUpdateTicket(publickey, MONGOCONNSTRINGENV, dbname, colladmin, collticket string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var admindata Admin

	gettoken := r.Header.Get("Login")
	if gettoken == "" {
		response.Message = "Header Login Not Exist"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		admindata.Email = checktoken
		if checktoken == "" {
			response.Message = "Kamu kayaknya belum punya akun"
		} else {
			admin2 := FindAdmin(mconn, colladmin, admindata)
			if admin2.Role == "admin" {
				var dataticket Ticket
				err := json.NewDecoder(r.Body).Decode(&dataticket)
				if err != nil {
					response.Message = "Error parsing application/json: " + err.Error()

				} else {
					UpdatedTicket(mconn, collticket, bson.M{"id": dataticket.ID}, dataticket)
					response.Status = true
					response.Message = "Berhasil Update Ticket"
					GCFReturnStruct(CreateResponse(true, "Success Update Ticket", dataticket))
				}
			} else {
				response.Message = "Anda tidak bisa Update data karena bukan admin"
			}

		}
	}
	return GCFReturnStruct(response)
}

// get all ticket
func GCFGetAllTicket(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	dataticket := GetAllTicket(mconn, collectionname)
	if dataticket != nil {
		return GCFReturnStruct(CreateResponse(true, "success Get All Ticket", dataticket))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Get All Ticket", dataticket))
	}
}

func GCFGetAllTickett(publickey, Mongostring, dbname, colname string, r *http.Request) string {
	resp := new(Credential)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		resp.Status = false
		resp.Message = "Header Login Not Exist"
	} else {
		existing := IsExist(tokenlogin, os.Getenv(publickey))
		if !existing {
			resp.Status = false
			resp.Message = "Kamu kayaknya belum punya akun"
		} else {
			koneksyen := SetConnection(Mongostring, dbname)
			dataticket := GetAllTicket(koneksyen, colname)
			dap, _ := json.Marshal(dataticket)
			resp.Status = true
			resp.Message = "Data Berhasil diambil"
			resp.Token = string(dap)
		}
	}
	return ReturnStringStruct(resp)
}

func GetAllDataTickets(PublicKey, MongoEnv, dbname, colname string, r *http.Request) string {
	req := new(Response)
	conn := SetConnection(MongoEnv, dbname)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		req.Status = false
		req.Message = "Header Login Not Found"
	} else {
		// Dekode token untuk mendapatkan
		_, err := DecodeGetTicket(os.Getenv(PublicKey), tokenlogin)
		if err != nil {
			req.Status = false
			req.Message = "Tidak ada data  " + tokenlogin
		} else {
			// Langsung ambil data ticket
			dataticket := GetAllTicket(conn, colname)
			if dataticket == nil {
				req.Status = false
				req.Message = "Data Ticket tidak ada"
			} else {
				req.Status = true
				req.Message = "Data Ticket berhasil di ambil"
				req.Data = dataticket
			}
		}
	}
	return ReturnStringStruct(req)
}

// get all ticket by id
func GCFGetAllTicketID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	var dataticket Ticket
	err := json.NewDecoder(r.Body).Decode(&dataticket)
	if err != nil {
		return err.Error()
	}

	ticket := GetAllTicketID(mconn, collectionname, dataticket)
	if ticket != (Ticket{}) {
		return GCFReturnStruct(CreateResponse(true, "Success: Get ID Ticket", dataticket))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed to Get ID Ticket", dataticket))
	}
}

// about post
func GCFInsertAbout(publickey, MONGOCONNSTRINGENV, dbname, colladmin, collabout string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var authdata Admin

	gettoken := r.Header.Get("Login")

	if gettoken == "" {
		response.Message = "Header Login Not Exist"
	} else {
		// Process the request with the "Login" token
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		authdata.Email = checktoken
		if checktoken == "" {
			response.Message = "Kamu kayaknya belum punya akun"
		} else {
			auth2 := FindAdmin(mconn, colladmin, authdata)
			if auth2.Role == "admin" {

				var dataabout About
				err := json.NewDecoder(r.Body).Decode(&dataabout)
				if err != nil {
					response.Message = "Error parsing application/json: " + err.Error()
				} else {
					InsertAbout(mconn, collabout, About{
						ID:          dataabout.ID,
						Title:       dataabout.Title,
						Description: dataabout.Description,
						Image:       dataabout.Image,
						Status:      dataabout.Status,
					})
					response.Status = true
					response.Message = "Berhasil Insert About"
				}
			} else {
				response.Message = "Anda tidak dapat Insert data karena bukan admin"
			}
		}
	}
	return GCFReturnStruct(response)

}

// delete about
func GCFDeleteAbout(publickey, MONGOCONNSTRINGENV, dbname, colladmin, collabout string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var authdata Admin

	gettoken := r.Header.Get("Login")

	if gettoken == "" {
		response.Message = "Header Login Not Exist"
	} else {
		// Process the request with the "Login" token
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		authdata.Email = checktoken
		if checktoken == "" {
			response.Message = "Kamu kayaknya belum punya akun"
		} else {
			auth2 := FindAdmin(mconn, colladmin, authdata)
			if auth2.Role == "admin" {

				var dataabout About
				err := json.NewDecoder(r.Body).Decode(&dataabout)
				if err != nil {
					response.Message = "Error parsing application/json: " + err.Error()
				} else {
					DeleteAbout(mconn, collabout, dataabout)
					response.Status = true
					response.Message = "Berhasil Delete About"
					CreateResponse(true, "Success Delete About", dataabout)
				}
			} else {
				response.Message = "Anda tidak dapat Delete data karena bukan admin"
			}
		}
	}
	return GCFReturnStruct(response)
}

// update about
func GCFUpdateAbout(publickey, MONGOCONNSTRINGENV, dbname, colladmin, collabout string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var authdata Admin

	gettoken := r.Header.Get("Login")

	if gettoken == "" {
		response.Message = "Header Login Not Exist"
	} else {
		// Process the request with the "Login" token
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		authdata.Email = checktoken
		if checktoken == "" {
			response.Message = "Kamu kayaknya belum punya akun"
		} else {
			auth2 := FindAdmin(mconn, colladmin, authdata)
			if auth2.Role == "admin" {
				var dataabout About
				err := json.NewDecoder(r.Body).Decode(&dataabout)
				if err != nil {
					response.Message = "Error parsing application/json: " + err.Error()
				} else {
					UpdatedAbout(mconn, collabout, bson.M{"id": dataabout.ID}, dataabout)
					response.Status = true
					response.Message = "Berhasil Update Catalog"
					CreateResponse(true, "Success Update About", dataabout)
				}
			} else {
				response.Message = "Anda tidak dapat Update data karena bukan admin"
			}
		}
	}
	return GCFReturnStruct(response)
}

// get all about
func GCFGetAllAbout(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	dataabout := GetAllAbout(mconn, collectionname)
	if dataabout != nil {
		return GCFReturnStruct(CreateResponse(true, "Berhasil Get All About", dataabout))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Gagal Get All About", dataabout))
	}
}

func GCFGetAllAboutt(publickey, Mongostring, dbname, colname string, r *http.Request) string {
	resp := new(Credential)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		resp.Status = false
		resp.Message = "Header Login Not Exist"
	} else {
		existing := IsExist(tokenlogin, os.Getenv(publickey))
		if !existing {
			resp.Status = false
			resp.Message = "Kamu kayaknya belum punya akun"
		} else {
			koneksyen := SetConnection(Mongostring, dbname)
			datacatalog := GetAllAbout(koneksyen, colname)
			yas, _ := json.Marshal(datacatalog)
			resp.Status = true
			resp.Message = "Data Berhasil diambil"
			resp.Token = string(yas)
		}
	}
	return ReturnStringStruct(resp)
}

// <--- ini contact --->

// contact post
func GCFInsertContact(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datacontact Contact
	err := json.NewDecoder(r.Body).Decode(&datacontact)
	if err != nil {
		return err.Error()
	}

	if err := InsertContact(mconn, collectionname, datacontact); err != nil {
		return GCFReturnStruct(CreateResponse(true, "Success Create Contact", datacontact))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Create Contact", datacontact))
	}
}

// get all contact
func GCFGetAllContact(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	datacontact := GetAllContact(mconn, collectionname)
	if datacontact != nil {
		return GCFReturnStruct(CreateResponse(true, "success Get All Contact", datacontact))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Get All Contact", datacontact))
	}
}

func GCFGetAllContactt(publickey, Mongostring, dbname, colname string, r *http.Request) string {
	resp := new(Credential)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		resp.Status = false
		resp.Message = "Header Login Not Exist"
	} else {
		existing := IsExist(tokenlogin, os.Getenv(publickey))
		if !existing {
			resp.Status = false
			resp.Message = "Kamu kayaknya belum punya akun"
		} else {
			koneksyen := SetConnection(Mongostring, dbname)
			datacatalog := GetAllContact(koneksyen, colname)
			yas, _ := json.Marshal(datacatalog)
			resp.Status = true
			resp.Message = "Data Berhasil diambil"
			resp.Token = string(yas)
		}
	}
	return ReturnStringStruct(resp)
}
