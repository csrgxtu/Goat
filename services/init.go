package services

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var Session *mgo.Session
var DB = beego.AppConfig.String("MongoDataBase")

func init() {
	SessionInit()
}

func SessionInit() (err error) {
	Session, err = mgo.Dial(beego.AppConfig.String("MongoDBUrl"))
	if err != nil {
		// panic(err)
    return
	}

	Session.SetMode(mgo.Monotonic, true)

  return
}

func CheckAndReconnect() (err error) {
  if Session == nil || Session.Ping() != nil {
    beego.Info("Session is <nil>, Trying 2 Reconnect")
    err = SessionInit()
  }

  return
}
