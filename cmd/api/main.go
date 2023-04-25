package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/crm-service/pkg/handlers"
	"github.com/softclub-go-0-0/crm-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DBInit(user, password, dbname, port string) (*gorm.DB, error) {
	dsn := "host=localhost" +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" port=" + port +
		" sslmode=disable" +
		" TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Teacher{},
		&models.Course{},
		&models.Timetable{},
		&models.Group{},
		&models.Student{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	DBName := flag.String("dbname", "crm_service", "Enter the name of DB")
	DBUser := flag.String("dbuser", "crm_service", "Enter the name of a DB user")
	DBPassword := flag.String("dbpassword", "crm_service", "Enter the password of user")
	DBPort := flag.String("dbport", "5432", "Enter the port of DB")
	flag.Parse()

	db, err := DBInit(*DBUser, *DBPassword, *DBName, *DBPort)
	if err != nil {
		log.Fatal("db connection error:", err)
	}

	log.Println("successfully connected to DB")

	h := handlers.NewHandler(db)

	router := gin.Default()

	router.GET("/teachers", h.GetAllTeachers)
	router.POST("/teachers", h.CreateTeacher)
	router.GET("/teachers/:teacherID", h.GetOneTeacher)
	router.PUT("/teachers/:teacherID", h.UpdateTeacher)
	router.DELETE("/teachers/:teacherID", h.DeleteTeacher)

	router.Run(":4000")
}
