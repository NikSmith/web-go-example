package modules
import (
	"webapi/webgo"
	"gopkg.in/mgo.v2"
)
type Mongo struct {
	Instance *mgo.Session
}
func (m *Mongo) Init() {
	var err error

	dbHost := webgo.CFG["dbHost"]
	dbName := webgo.CFG["dbName"]
	dbLogin := webgo.CFG["dbUser"]
	dbPassword := webgo.CFG["dbPassword"]
	isAuth := false

	if dbPassword != "" {
		isAuth = true
	}

	if dbHost == "" {
		webgo.LOGGER.Fatal("Unknow host")
	}

	m.Instance, err = mgo.Dial(dbHost)

	if err != nil {
		webgo.LOGGER.Fatal(err)
	}

	if isAuth{
		err = m.Instance.DB(dbName).Login(dbLogin, dbPassword)
		if err != nil {
			webgo.LOGGER.Fatal(err)
		}
	}

}
func (m *Mongo) GetInstance() interface{}{
	return m.Instance
}