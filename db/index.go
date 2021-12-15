package db

import (
	"fmt"
	"gore/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SetAppVersion(res model.Release) {
	fmt.Println("Setting new app version in DB...")
	_, err := Db.Query(UpdateLatestVersionOfAppQuery, res.App, res.Major, res.Minor, res.Patch)
	if err != nil {
		log.Fatalf("Error querying database: %+v", err)
	}
}

func GetLatestVersionOfApp(app string) (model.Release, bool) {
	rows, err := Db.Query(GetLatestVersionOfAppQuery, app)
	if err != nil {
		log.Fatalf("Error querying database: %+v", err)
	}

	if !rows.Next() {
		return model.Release{}, false
	}

	res := model.Release{}
	err = rows.Scan(&res.App, &res.Major, &res.Minor, &res.Patch)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	return res, true
}
