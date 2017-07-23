package goBotUtils

import (
	"os"
	"log"
	"strings"
	"regexp"
	"database/sql"
	"fmt"
	"errors"
	"encoding/json"
	"net/http"
	"io"
	"github.com/tidwall/gjson"
)

func CreateFile(path string) (err error) {
	_, err = os.Stat(path)

	if os.IsNotExist(err) {
		err = nil
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	} else {
		DeleteFile(path)
		return CreateFile(path)
	}
	return
}

func DeleteFile(path string) error {
	return os.Remove(path)
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// функция проверки содержится ли строка в массиве строк
func CheckStrContains(arr []string, str string) bool {
	for _, s := range arr {
		if strings.ToLower(s) == strings.ToLower(str) {
			return true
		}
	}
	return false
}

// функция проверки строки с массивом regex выражений
func CheckStrRegExContains(arr []string, str string) bool {
	for _, s := range arr {
		res, err := regexp.MatchString(strings.ToLower(s), strings.ToLower(str))
		if err != nil {
			log.Printf("CheckStrRegExContains error: %s", err)
		}
		if res {
			return true
		}
	}
	return false
}

// функция удаление элемента из массива
func RemoveElementFromArr(arr []string, str string) ([]string) {
	i := FindIndexInArrStr(arr, str)
	return append(arr[:i], arr[i + 1:]...)
}

// функция поиска индекса элемента в массива
func FindIndexInArrStr(arr []string, str string) (i int) {
	for i, s := range arr {
		if strings.ToLower(s) == strings.ToLower(str) {
			return i
		}
	}
	return -1
}

func CallPgFunc(pgDb *sql.DB, funcName string, jsonStr []byte, res interface{}, metaInfo interface{}) (err error) {

	var queryRes []byte
	var queryStr string

	if len(jsonStr) > 0 {
		queryStr = fmt.Sprintf("select * from %s('%s')", funcName, jsonStr)
	} else {
		queryStr = fmt.Sprintf("select * from %s()", funcName)
	}

	err = pgDb.QueryRow(queryStr).Scan(&queryRes)
	if err != nil {
		return
	}

	//fmt.Printf("queryRes %s\n", queryRes)

	return ParseResponseFromPostgresFunc(queryRes, res, metaInfo)
}

func ParseResponseFromPostgresFunc(queryRes []byte, tempRes interface{}, metaInfo interface{}) (err error) {
	ok := gjson.Get(fmt.Sprintf("%s", queryRes), "ok").Bool()
	if !ok {
		errMsg := gjson.Get(fmt.Sprintf("%s", queryRes), "message").Str
		err = errors.New(errMsg)
		return
	}

	err = json.Unmarshal([]byte(gjson.Get(fmt.Sprintf("%s", queryRes), "result").Raw), &tempRes)
	if err != nil {
		return err
	}
	if metaInfo != nil {
		err = json.Unmarshal([]byte(gjson.Get(fmt.Sprintf("%s", queryRes), "meta_info").Raw), &metaInfo)
		if err != nil {
			return err
		}
	}
	return nil
}

func DownloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

