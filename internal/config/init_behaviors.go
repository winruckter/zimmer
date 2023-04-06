package configure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

//Behavior 行为
type Behavior struct {
	Tag  string `json:"tag_name"`
	Name string `json:"behavior_name"`
}

var (
	mapBehaviors     *map[string]string
	onceForBehaviors sync.Once
)

//GetBehaviors 获取标签及对应行为
func GetBehaviors() *map[string]string {
	onceForBehaviors.Do(initBehaviors)
	return mapBehaviors
}
func initBehaviors() {
	tempMap := map[string]string{}
	behaviors := &[]Behavior{}

	fileName := "behaviors.json"
	filePath := generateFilePath(fileName)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	f := []byte(file)
	println(f)
	err = json.Unmarshal([]byte(file), behaviors)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for _, e := range *behaviors {
		tempMap[e.Tag] = e.Name
	}

	mapBehaviors = &tempMap
}
