package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/golang/glog"

	_ "github.com/go-sql-driver/mysql"
	//	"github.com/labstack/echo"
	//	"github.com/labstack/echo/engine/standard"
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
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var createTableMySQL = `CREATE TABLE IF NOT EXISTS users (
		id INT UNSIGNED NOT NULL,
		name VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci,
		groupname VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci,
		employeenum SMALLINT(255) UNSIGNED NULL,
		uid SMALLINT(255) UNSIGNED NULL,
		gid SMALLINT(255) UNSIGNED NULL,
		PRIMARY KEY (id)
	)`

func main() {
	//var err error
	glog.Info("Starting Man-o-war")
	dbType := os.Getenv("DATABASE_TYPE")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbHost := os.Getenv("DATABASE_HOSTNAME")
	testNums := os.Getenv("TEST_AMOUNTS")
	dbConnect := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, "mysql")
	if dbType == "mysql" {
		glog.Infof("Connecting to MySQL Database: %v", dbConnect)
		dB, _ = sql.Open("mysql", dbConnect)
		glog.Infof("Database connection: %v", dB)
		if err := dB.Ping(); err != nil {
			glog.Fatal(err)
		}
		_, err := dB.Exec(createTableMySQL)
		if err != nil {
			glog.Fatalf("TABLE CREATION ERROR: %v", err)
		}
	}
	testAmount, _ := strconv.Atoi(testNums)
	rand.Seed(777)
	for testNum := 1; testNum < testAmount; testNum++ {
		fmt.Printf("Test: %v", testNum)
		testUser := new(user)
		testUser.ID = rand.Intn(20)
		testUser.Name = RandStringRunes(10)
		testUser.Group = RandStringRunes(5)
		testUser.EmployeeNum = rand.Intn(1000)
		testUser.GID = rand.Intn(100)
		testUser.UID = rand.Intn(100)
		randKey := rand.Intn(4000)
		glog.Infof("Rand Key: %v", randKey)
		if randKey <= 1000 {
			glog.Info("Entering into INPUT")
			inputData(testUser)
		} else if (randKey > 1000) && (randKey <= 2000) {
			glog.Info("Entering info SELECT")
			selectData(testUser)
		} else if (randKey > 2000) && (randKey <= 3000) {
			glog.Info("Entering into DELETE")
			deleteData(testUser)
		} else {
			glog.Info("Entering into UPDATE")
			updateData(testUser)
		}
	}

	//	if err != nil {
	//		fmt.Printf("ERROR %v", err)
	//	}
	//	if err := dB.Ping(); err != nil {
	//		log.Fatal(err)
	//	}
	//	defer dB.Close()
	//	e := echo.New()
	//	//e.File("/", InputData()
	//	e.POST("/", inputData)
	//
	//	e.Run(standard.New(":8083"))
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func inputData(testUser *user) {
	fmt.Print(" INPUT ")
	result, err := dB.Exec(
		"INSERT INTO users (id, name, groupname, employeenum, uid, gid) VALUES ( ?, ?, ?, ?, ?, ? )",
		testUser.ID,
		testUser.Name,
		testUser.Group,
		testUser.EmployeeNum,
		testUser.UID,
		testUser.GID,
	)
	if err != nil {
		fmt.Printf(" FAILED: %v\n", err)
	}
	fmt.Printf(" SUCCESS:%v\n", result)
}

func selectData(testUser *user) {
	fmt.Print(" SELECT ")
	result, err := dB.Query(
		"SELECT * FROM users WHERE id = ?",
		testUser.ID,
	)
	if err != nil {
		fmt.Printf(" FAILED: %v\n", err)
	}
	fmt.Printf(" SUCCESS:%v\n", result.Scan())
	result.Close()
}

func updateData(testUser *user) {
	fmt.Print(" UPDATE ")
	result, err := dB.Exec(
		"UPDATE users SET name = ?, employeenum = ? where id = ?",
		testUser.ID,
		testUser.Name,
		testUser.EmployeeNum,
	)
	if err != nil {
		fmt.Printf(" FAILED: %v\n", err)
	}
	fmt.Printf(" SUCCESS:%v\n", result)
}

func deleteData(testUser *user) {
	fmt.Print(" DELETE ")
	result, err := dB.Exec(
		"DELETE FROM users where id = ?",
		testUser.ID,
	)
	if err != nil {
		fmt.Printf(" FAILED: %v\n", err)
	}
	fmt.Printf(" SUCCESS:%v\n", result)
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
