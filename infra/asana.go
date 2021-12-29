package infra

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/konan0802/dakoku/config"
	"github.com/konan0802/dakoku/domain/model"
	"github.com/konan0802/dakoku/domain/repository"
)

type asanaRepository struct {
	token string
	pjtID string
}

// NewRepository handlerを生成
func NewAsanaRepository() repository.AsanaRepository {
	asanaToken := config.ASANATOKEN
	asanaPjtID := config.ASANAPJTID
	return &asanaRepository{
		token: asanaToken,
		pjtID: asanaPjtID,
	}
}

func (arp asanaRepository) GetSectionsInMyTask() (model.AsanaSections, error) {
	var receiveSections model.AsanaSectionsReceivedFormat
	request, err := http.NewRequest("GET", "https://app.asana.com/api/1.0/projects/"+arp.pjtID+"/sections", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	request.Header.Set("Authorization", "Bearer "+arp.token)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(responseBody, &receiveSections)
	if err != nil {
		fmt.Println(err.Error())
	}
	return receiveSections.Data, nil
}

func (arp asanaRepository) GetTasksFromSection(targetSectionsGID string) (model.AsanaTasksReceiver, error) {
	var receiveTasks model.AsanaTasksReceivedFormat
	request, err := http.NewRequest("GET", "https://app.asana.com/api/1.0/sections/"+targetSectionsGID+"/tasks?opt_fields=this.name,this.completed,this.parent,this.due_on,this.projects", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	request.Header.Set("Authorization", "Bearer "+arp.token)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(responseBody, &receiveTasks)
	if err != nil {
		fmt.Println(err.Error())
	}
	return receiveTasks.Data, nil
}

func (arp asanaRepository) GetTasksNameFromGID(taskGID string) (string, error) {
	var receiveTaskName model.AsanaTaskNameReceivedFormat
	request, err := http.NewRequest("GET", "https://app.asana.com/api/1.0/tasks/"+taskGID+"?opt_fields=this.name", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	request.Header.Set("Authorization", "Bearer "+arp.token)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(responseBody, &receiveTaskName)
	if err != nil {
		fmt.Println(err.Error())
	}
	return receiveTaskName.Data.Name, nil
}
