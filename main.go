package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

type (
	user struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Group       string `json:"group"`
		EmployeeNum int    `json:"employeeNum"`
		UID         int    `json:"uid"`
		GID         int    `json:"gid"`
	}
)

var dB *sql.DB

func main() {
	var err error
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbHost := os.Getenv("DATABASE_HOSTNAME")
	dbConnect := fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbHost)

	dB, err = sql.Open("mysql", dbConnect)
	if err != nil {
		fmt.Printf("ERROR %v", err)
	}
	if err := dB.Ping(); err != nil {
		log.Fatal(err)
	}
	defer dB.Close()
	e := echo.New()
	//e.File("/", InputData()
	e.POST("/", inputData)

	e.Run(standard.New(":8083"))
}

func inputData(c echo.Context) (err error) {
	//	dbUser := os.Getenv("DATABASE_USER")
	//	dbPass := os.Getenv("DATABASE_PASS")
	//	dbHost := os.Getenv("DATABASE_HOSTNAME")
	//	dbConnect := fmt.Printf("%s:%s@%s", dbUser, dbPass, dbHost)
	user := new(user)
	if err = c.Bind(&user); err != nil {
		fmt.Printf("ERROR: %v", err)
		return c.NoContent(http.StatusNoContent)
	}
	//fmt.Printf("ID:%d Name:%v Group:%v EmployeeNum:%v UID:%v GID:%v", user.ID, user.Name, user.Group, user.EmployeeNum, user.UID, user.GID)
	//	db, _ := sql.Open("mysql", dbConnect)

	//	defer db.Close()

	//I chose to use Exec here instead of a transaction as I want to try and push the database as hard as possible. Exec opens a connection per transaction.
	result, err := dB.Exec(
		"INSERT INTO users (id, name, group, employeenum, uid, gid) VALUES ( $1, $2, $3, $4, $5 $6",
		user.ID,
		user.Name,
		user.Group,
		user.EmployeeNum,
		user.UID,
		user.GID,
	)
	if err != nil {
		fmt.Printf("ERROR EXEC: %v\n", err)
	}
	fmt.Printf("RESULTS:%v\n", result)
	return c.NoContent(http.StatusNoContent)

}

/*
func getClusters() []inv.Cluster {

	/*
		k8sConfig, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		k8sclient, err := kubernetes.NewForConfig(k8sConfig)
		if err != nil {
			panic(err.Error())
		}


	clusters := inventory.Get(qmConfig)

	return clusters
}
*/
