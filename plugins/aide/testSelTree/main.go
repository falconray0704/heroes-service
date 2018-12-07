package main

import (
	"github.com/chainHero/heroes-service/plugins/aide"
)

const (
	//jdbPath = "/usr/local/etc/aideDB/aide.db.new.json"
	jdbOldPath = "./testData/aide.db.old.json"
	jdbDiskPath = "./testData/aide.db.new.json"
)


func main() {

	var dbCfg = aide.New_DB_confg(jdbOldPath, "", jdbDiskPath)

	dbCfg.Load_JSON_DB(aide.DB_OLD)
	//dbCfg.Load_JSON_DB(aide.DB_DISK)

	slist, nlist, elist := dbCfg.JdbOld.GetRxList()

	dbCfg.Tree = aide.Gen_tree(slist, nlist, elist)
	dbCfg.Tree.Print_tree_rx(" |")

}
