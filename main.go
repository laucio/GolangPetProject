package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/laucio/WebApi"
	"math/rand"
	"runtime"
	"sync"
	"time"
	// "io"
	// "os"
)

// Global
var wg sync.WaitGroup

func main() {

	// Logging to a file.
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// r := SetupRouter()
	// r.Run()

	wg.Add(2)
	go mostrar0()
	go mostrar1()
	wg.Wait()
	fmt.Println("Cantidad de procesadores:", runtime.NumCPU())

	//runtime.GOMAXPROCS(1)

	// aleatorio := rand.New(rand.NewSource(time.Now().UnixNano()))
	// var vec1 [50000]int
	// var vec2 [50000]int
	// cargar(&vec1, aleatorio)
	// cargar(&vec2, aleatorio)

	// var hora1, hora2 time.Time
	// hora1 = time.Now()
	// wg.Add(2)
	// go ordenar(&vec1)
	// go ordenar(&vec2)
	// wg.Wait()
	// hora2 = time.Now()
	// di := diferenciaTiempo(hora1, hora2)
	// fmt.Println("Cantidad de segundos de diferencia:", di.Seconds())

	canal := make(chan int)

	go imprimir(canal)
	go contar(canal)
	var fin string
	fmt.Scanln(&fin)
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

// docker build -t petprojectfinalversion .
// Docker docker run -d -p 5000:8080 petprojectfinalversion
// Corre en localhost:5000
func mostrar0() {
	for index := 0; index < 200; index++ {
		fmt.Println("0")
	}
	wg.Done()
}

func mostrar1() {
	for index := 0; index < 200; index++ {
		fmt.Println("1")
	}
	wg.Done()
}

func cargar(vec *[50000]int, aleatorio *rand.Rand) {
	for f := 0; f < len(vec); f++ {
		vec[f] = aleatorio.Intn(100)
	}
}

func ordenar(vec *[50000]int) {
	for k := 1; k < len(vec); k++ {
		for f := 0; f < len(vec)-k; f++ {
			if vec[f] > vec[f+1] {
				aux := vec[f]
				vec[f] = vec[f+1]
				vec[f+1] = aux
			}
		}
	}
	wg.Done()
}

func diferenciaTiempo(hora1, hora2 time.Time) time.Duration {
	diferencia := hora2.Sub(hora1)
	return diferencia
}

func contar(canal chan int) {
	x := 0
	for index := 0; index < 10; index++ {
		canal <- x
		x++
	}
}

func imprimir(canal chan int) {
	var valor int
	for {
		valor = <-canal
		fmt.Println(valor)
	}
}
