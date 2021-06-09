package utils

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
)

func DealErr(err error) {
	if err != nil {
		panic(err)
	}
}

func DealWrapErr(err *Err)  {
	if err.Err != nil {
		panic(err)
	}
}

func DealErrAndRollback(db *sqlx.Tx, err error) {
	if err != nil {
		var errTmp = err
		err := db.Rollback()
		log.Fatal(errTmp, err)
	}
}

type Err struct {
	Code int
	Msg  string
	Err  error
}

func (e *Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

func NewErr(code int, msg string, err error) *Err {
	return &Err{
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}
