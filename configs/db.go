package Config

import (
	"fmt"
	"log"

	mod "WanderGo/models"
	"strconv"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var GLOBAL_DB *gorm.DB
var GLOBAL_RDB *redis.Client
var GlobalConfig = ReadConfig()
var DBConf = GlobalConfig.Mysql
var RedisConf = GlobalConfig.Redis

func ConnectToDb() {
	DBDSN := DBConf.Username + ":" + DBConf.Password + "@tcp(" + DBConf.Host + ":" + strconv.Itoa(DBConf.Port) + ")/" + DBConf.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               DBDSN,
		DefaultStringSize: 171,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Println(err)
		return
	}
	GLOBAL_DB = db
	if (!GLOBAL_DB.Migrator().HasTable(&mod.User{})) {
		err := GLOBAL_DB.AutoMigrate(&mod.User{})
		if err != nil {
			log.Println(err)
			return
		}
	}
	if (!GLOBAL_DB.Migrator().HasTable(&mod.Avatar{})) {
		err := GLOBAL_DB.AutoMigrate(&mod.Avatar{})
		if err != nil {
			log.Println(err)
			return
		}
	}
	if (!GLOBAL_DB.Migrator().HasTable(&mod.Photo{})) {
		err := GLOBAL_DB.AutoMigrate(&mod.Photo{})
		if err != nil {
			log.Println(err)
			return
		}
	}
	if (!GLOBAL_DB.Migrator().HasTable(&mod.Place{})) {
		err := GLOBAL_DB.AutoMigrate(&mod.Place{})
		if err != nil {
			log.Println(err)
			return
		}
	}
	if (!GLOBAL_DB.Migrator().HasTable(&mod.Comment{})) {
		err := GLOBAL_DB.AutoMigrate(&mod.Comment{})
		if err != nil {
			log.Println(err)
			return
		}
	}
	if (!GLOBAL_DB.Migrator().HasTable(&mod.Star{})) {
		err := GLOBAL_DB.AutoMigrate(&mod.Star{})
		if err != nil {
			log.Println(err)
			return
		}
	}
	//new test
	// if (!GLOBAL_DB.Migrator().HasTable(&mod.Place{})) {
	// 	err := GLOBAL_DB.AutoMigrate(&mod.Place{})
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// }
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisConf.Host + ":" + RedisConf.Port,
		Password: RedisConf.Password,
		DB:       1,
	})
	GLOBAL_RDB = rdb
	pong, err := rdb.Ping().Result()
	fmt.Println(pong)
	if err != nil {
		log.Print(err)
		return
	}
	if GLOBAL_RDB.Exists("NCU:Buildings").Val() == 0 {
		AddGeoInfo()
	}
	AddGeoInfoMysql()
}
