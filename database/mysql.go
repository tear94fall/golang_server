package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

type MysqlConn struct {
	Info *MysqlInfo `yaml:"database"`
	Conn *sql.DB
}

type MysqlInfo struct {
	Ip     string `yaml:"ip"`
	Port   int    `yaml:"port"`
	Db     string `yaml:"db"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
}

func InitMysqlConn() *MysqlConn {
	confFile := os.Getenv("DBCONF")
	conn := &MysqlConn{}
	conn.ReadConfFile(confFile)
	conn.PrintConf()
	if conn.Info == nil {
		return nil
	}

	info := conn.Info.User + ":" + conn.Info.Passwd + "@tcp(" + conn.Info.Ip + ")/" + conn.Info.Db
	db, err := sql.Open("mysql", info)
	if err != nil {
		panic(err)
	}

	// Set db conn
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	conn.Conn = db

	return conn
}

func (conn *MysqlConn) ReadConfFile(filename string) error {
	yamlfile, err := ioutil.ReadFile(filename)
	log.Printf("read yaml file : [%s]", filename)
	if err != nil {
		log.Printf("read yaml file : [%s] error #%v", filename, err)
		return nil
	}

	err = yaml.Unmarshal(yamlfile, &conn)

	if err != nil {
		log.Printf("yaml file bind fail : [%s] error #%v", filename, err)
		panic(err)
	}

	return nil
}

func (conn *MysqlConn) RunQuery(query string) error {
	m := &Member{}

	rows, err := conn.Conn.Query(query)
	if err != nil {
		log.Printf("run query fail : [%s] error $%v", query, err)
		return fmt.Errorf("run query fail : [%s] error $%v", query, err)
	}
	defer rows.Close()

	count := 0

	for rows.Next() {
		err := rows.Scan(&m.Email, &m.Passwd)
		if err != nil {
			log.Fatal(err)
		}
		m.PrintMember()
		count += 1
	}

	if count == 0 {
		log.Printf("login fail")
		return fmt.Errorf("login fail")
	}

	return nil
}

func (conn *MysqlConn) PrintConf() {
	whilte := color.New(color.FgWhite)
	boldWhite := whilte.Add(color.BgGreen)
	boldWhite.Print(conn.Info.Ip)
	fmt.Println()
	boldWhite.Print(conn.Info.User)
	fmt.Println()
}
