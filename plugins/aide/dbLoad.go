package aide

import (
	"encoding/json"
	aUtils "github.com/chainHero/heroes-service/plugins/utils"
)

func (jDB *JsonDB) loadJDB() (error){
	var err error
	err = json.Unmarshal(jDB.RawData, &jDB.Jdb)
	return err
}

func NewJDB(dbPath string) (*JsonDB, error) {
	var err error
	var jDB = new(JsonDB)
	jDB.FilePath = dbPath
	jDB.RawData, err = aUtils.GetBufFromFile(jDB.FilePath)

	if err != nil {
		return jDB, err
	}

	err = jDB.loadJDB()
	return jDB, err
}
