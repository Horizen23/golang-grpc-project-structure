package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"golang-grpc-project-structure/pkg/config"

	"golang-grpc-project-structure/pkg/utils"

	_ "github.com/denisenkom/go-mssqldb"
)

type Result struct {
	RecordSets [][]map[string]any
}

func openDB(dsn string, config *config.Configuration) (*sql.DB, error) {
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(int(config.MSSQL.MAX_OPEN_CONNS))
	db.SetMaxIdleConns(int(config.MSSQL.MAX_IDLE_CONNS))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.MSSQL.CONN_TIME_OUT)*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectSQL(config *config.Configuration) *sql.DB {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		config.MSSQL.IP,
		config.MSSQL.USERNAME,
		config.MSSQL.PASSWORD,
		config.MSSQL.PORT,
		config.MSSQL.DB)
	DBConnection, err := openDB(connectionString, config)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("SQL           connected : %s @user: %s \n", config.MSSQL.IP, config.MSSQL.USERNAME)

	return DBConnection
}

func isDatabaseAlive(db *sql.DB) (context.Context, error) {
	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("Error ping: ", err.Error())
		return ctx, err
	}
	return ctx, nil
}

func ExecuteStoreProcedureWithInput(db *sql.DB, option utils.USPOptions) (*Result, error) {

	result := &Result{}

	ctx, err := isDatabaseAlive(db)
	if err != nil {
		panic(err)
	}

	ctxWC, cancel := context.WithCancel(ctx)
	defer cancel()

	stmt, err := db.Prepare(option.GetProcedureQuery())
	if err != nil {
		log.Fatal("Prepare error: ", err.Error())
		return result, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctxWC, option.GetParams()...)
	if err != nil {
		fmt.Print("query: ")
		panic(err)
	}
	for {
		if rows.Err() != nil {
			panic(rows.Err())
		}
		cols, err := rows.Columns()
		if err != nil {
			return result, err
		}
		if cols == nil {
			continue
		}

		var newTable []map[string]any
		for rows.Next() {
			lenCols := len(cols)
			var colsName []string
			newRow := map[string]any{}
			rowValues := make([]any, lenCols)
			for index, name := range cols {
				rowValues[index] = new(any)
				colsName = append(colsName, name)
			}
			fmt.Println(colsName)

			err = rows.Scan(rowValues...)
			if err != nil {
				fmt.Println(err)
				continue
			}

			for index, item := range rowValues {
				newRow[colsName[index]] = convertType(item.(*any))
			}
			newTable = append(newTable, newRow)

		}

		result.RecordSets = append(result.RecordSets, newTable)
		if !rows.NextResultSet() {
			break
		}
	}

	return result, nil
}

func convertType(pval *any) any {
	switch v := (*pval).(type) {
	case nil:
		return nil
	case bool:
		if v {
			return true
		} else {
			return false
		}
	case []byte:
		return string(v)
	case time.Time:
		return v.Format("2006-01-02 15:04:05.999")
	default:
		return v
	}
}
