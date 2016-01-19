package models
import(
	"gopkg.in/mgo.v2"
	"webapi/webgo"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	UserId int `json:"userId" bson:"userId"`
	Name string
}

func (u *User) GetUser(id int){
	DB:=webgo.MODULES("db").GetInstance().(*mgo.Session)
	db:=DB.Copy()
	defer func() {
		db.Close()
	}()

	query:=db.DB("local").C("users").Find(bson.M{"userId":id})
	query.One(u)
}

