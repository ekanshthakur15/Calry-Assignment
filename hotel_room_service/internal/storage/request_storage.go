package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/ekanshthakur15/hotel-service/internal/models"
)

const jsonFilePath = "data/requests.json"
var mutex sync.Mutex

func ReadFromFile() ([]models.RoomServiceRequest, error) {
	mutex.Lock()

	defer mutex.Unlock()

	var requests []models.RoomServiceRequest
	file, err := os.Open(jsonFilePath)

	if err != nil{
		if os.IsNotExist(err){
			return requests, nil
		}
		return nil, err
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)

	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(byteValue, &requests)

	return requests, err

}

func WriteToFile(requests []models.RoomServiceRequest) error {

	mutex.Lock()

	defer mutex.Unlock()

	file, err := os.OpenFile(jsonFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil{
		return err
	}

	defer file.Close()

	jsonData, err := json.Marshal(requests)
	if err != nil {
		return err
	}
	_, err = file.Write(jsonData)
	fmt.Printf("Write error: %s", err)
	return err

}