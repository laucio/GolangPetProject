package Repository

import (
	"encoding/json"
	"fmt"
	"github.com/laucio/Entity"
	"io/ioutil"
	"net/http"
)

const url = "https://gitlab.lagash.com/"

func GetAllProjects() ([]Entity.Project, error) {
	const endpoint = "api/v4/projects"
	resp, err := http.Get(url + endpoint)
	if err != nil {
		return nil, fmt.Errorf("HTTP error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status OK but got %q instead", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error: %s\n", err)
	}

	var projects []Entity.Project

	json.Unmarshal(data, &projects)

	return projects, nil
}
