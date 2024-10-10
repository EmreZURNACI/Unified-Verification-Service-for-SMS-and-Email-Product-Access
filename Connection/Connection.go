package Connection

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	m "ProductService/Models"

	_ "github.com/lib/pq"
)

var ConfigModel m.ConfigFile

func init() {
	bs, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bs, &ConfigModel)
	if err != nil {
		panic(err)
	}
}
func Connection() *sql.DB {
	var dsn string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", ConfigModel.Dbmodel.Host, ConfigModel.Dbmodel.Port, ConfigModel.Dbmodel.User, ConfigModel.Dbmodel.Password, ConfigModel.Dbmodel.Dbname)
	con, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return con
}
