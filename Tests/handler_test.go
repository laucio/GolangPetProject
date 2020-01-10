package Tests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/laucio/WebApi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllProjects(t *testing.T) {

	// Build our expected body
	body := gin.H{
		"hello": "world",
	}

	router := SetupRouter()
	w := performRequest(router, "GET", "/getallprojects")

	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code) // Convert the JSON response to a map

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response) // Grab the value & whether or not it exists
	value, exists := response["hello"]                        // Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["hello"], value)

}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func SetupRouter() *gin.Engine {

	r := gin.Default()

	//Public group
	public := r.Group("/public")
	public.GET("/publictest", WebApi.GetAllProjects)

	//Private group
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"laucio": "lauciolaucio2",
	}))
	authorized.GET("/privatetest", WebApi.GetAllProjects)
	authorized.GET("/getallprojects", WebApi.GetAllProjects)
	authorized.GET("/getWrongNameProjects/:pattern", WebApi.GetWrongNameProjects)
	authorized.GET("/getTimeWindowProjects/:startdate/:enddate", WebApi.GetTimeWindowProjects)
	authorized.GET("/getReadmeProjects", WebApi.GetReadmeProjects)

	return r
}
