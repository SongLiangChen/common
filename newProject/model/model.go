package model

import (
	"bytes"
	"fmt"
	"github.com/becent/golang-common"
	"os"
	"strings"
)

func G_model(projectName string) error {
	if err := os.Mkdir(projectName+"/model", 755); err != nil {
		return err
	}

	db := common.GetDB("database")
	if db == nil {
		file, err := os.OpenFile(projectName+"/model/user.go", os.O_CREATE|os.O_RDWR, 755)
		if err != nil {
			return err
		}

		if _, err := file.WriteString(model_temple); err != nil {
			return err
		}
		return nil
	}

	curDatabase := db.Dialect().CurrentDatabase()

	rows, err := db.DB().Query("select table_name from information_schema.tables where table_schema=? and table_type=?", curDatabase, "base table")
	if err != nil {
		return err
	}
	defer rows.Close()

	buf := bytes.NewBuffer(nil)

	for rows.Next() {
		buf.Reset()

		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return err
		}

		camelTableName := CamelCase([]byte(tableName))

		file, err := os.OpenFile(fmt.Sprintf(projectName+"/model/%s.go", camelTableName), os.O_CREATE|os.O_RDWR, 755)
		if err != nil {
			return err
		}

		buf.WriteString("package model\n\n")

		buf.WriteString(fmt.Sprintf("type %s struct {\n", strings.Title(camelTableName)))

		structRows, err := db.DB().Query("select column_name,COLUMN_TYPE,COLUMN_KEY,COLUMN_COMMENT from information_schema.columns where table_schema=? and table_name=?", curDatabase, tableName)
		if err != nil {
			return err
		}

		for structRows.Next() {
			var columnName, columnType, columnKey, columnComment string
			if err = structRows.Scan(&columnName, &columnType, &columnKey, &columnComment); err != nil {
				return err
			}

			columnName = CamelCase([]byte(columnName))

			// println(columnName, columnType, columnKey)
			buf.WriteString(fmt.Sprintf("	%s ", strings.Title(columnName)))
			if strings.Index(columnType, "bigint") != -1 {
				buf.WriteString("int64 ")
			} else if strings.Index(columnType, "int") != -1 {
				buf.WriteString("int ")
			} else if strings.Index(columnType, "tinyint") != -1 {
				buf.WriteString("int ")
			} else if strings.Index(columnType, "varchar") != -1 {
				buf.WriteString("string ")
			} else if strings.Index(columnType, "char") != -1 {
				buf.WriteString("string ")
			} else if strings.Index(columnType, "text") != -1 {
				buf.WriteString("string ")
			} else if strings.Index(columnType, "decimal") != -1 {
				buf.WriteString("float64 ")
			}

			temp := ""
			if columnKey == "PRI" {
				temp = ";primary_key;AUTO_INCREMENT"
			}

			buf.WriteString(fmt.Sprintf("`gorm:\"column:%s%s\" json:\"%s\"` // %s\n", columnName, temp, columnName, columnComment))
		}
		structRows.Close()

		buf.WriteString("}\n\n")

		buf.WriteString(fmt.Sprintf("func (*%s) TableName() string {\n", strings.Title(camelTableName)))
		buf.WriteString(fmt.Sprintf("	return \"%s\"\n", tableName))
		buf.WriteString("}\n")

		if _, err := file.WriteString(buf.String()); err != nil {
			return err
		}
	}

	return nil
}

func CamelCase(v []byte) string {
	buf := bytes.NewBuffer(nil)
	t := false
	for _, c := range v {
		if c == '_' {
			t = true
			continue
		}

		if c >= '0' && c <= '9' {
			buf.WriteByte(c)
			continue
		}

		if t {
			t = false
			c -= 'a' - 'A'
		}
		buf.WriteByte(c)
	}
	return buf.String()
}

var model_temple = `package model

type User struct {
	Id     int    ` + "`gorm:\"column:id;primary_key;AUTO_INCREMENT\"`" + `
	UserId int64  ` + "`gorm:\"column:userId\"`" + `
	Name   string  ` + "`gorm:\"column:string\"`" + `
}

func (u *User) TableName() string {
	return "user"
}
`
