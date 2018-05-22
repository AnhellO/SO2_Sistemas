package main

import (
	"fmt"
	"log"
	"net/http"
	"math"
	"strconv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/RyanCarrier/dijkstra"
)

type Polygon [][2]float64

type Geometrystruct struct {
	Type  string `bson:"type"`
	Coordinates [][2]float64 `bson:"coordinates" json:"coordinates"`
}

type Vialidad struct {
   	Tipovial string  `bson:"tipovial"`
   	Sentido string  `bson:"sentido"`
   	Geometry Geometrystruct `bson:"geometry"`
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
  // must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func main() {

    dialInfo, err :=      mgo.ParseURL("mongodb://root:$jefr8041@10.0.0.153:27017,10.0.0.154:27017,10.0.0.155:27017,10.0.0.156:27017/admin?replicaSet=nlared")
     session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
            panic(err)
    }
    defer session.Close()
    graph:=dijkstra.NewGraph()
    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)
    c := session.DB("geo").C("mapa")

    //result := Vialidad{}
    //err = c.Find(bson.M{"geografico":"Vialidad"}).One(&result) //busqueda

//http://icchan.github.io/2014/10/18/geospatial-querying-with-go-and-mongodb/
    var results []Vialidad
    err = c.Find(bson.M{"geografico":"Vialidad"}).All(&results)
    if err != nil {
        log.Fatal(err)
    }
   strDict:= map[string]int{}
   var conteo int
	//var strA,strB string
	//var noA,noB int
	var D float64
	//var ant int

	var A [2]float64
	//strA, _ :=""
	//strB:=""
	//noB:=0

   	for _,vialidad:=range results {
 		fmt.Printf("Tipo %s cantidad coords %d \n",vialidad.Geometry.Type,len(vialidad.Geometry.Coordinates))
	
		//var cantidad, _ =len(vialidad.Geometry.Coordinates)
		var cont int
		for noB,B:= range vialidad.Geometry.Coordinates {
	    		
			if cont>0{
				fmt.Printf("Lat %l Lng %l \n",A[0],A[1])
				noA:=0
				noB:=noA+1
				//B,ok:=vialidad.Geometry.Coordinates[ant]
				strA:=strconv.FormatFloat(A[0], 'f', -1, 64)+strconv.FormatFloat(A[1],'f',-1,64)
				strB:=strconv.FormatFloat(B[0], 'f', -1, 64)+strconv.FormatFloat(B[1],'f',-1,64)
				intA,ok:=strDict[strA]
				if !ok{
					strDict[strA]=conteo
					intA=conteo
 					graph.AddVertex(conteo)
					conteo++
				}
				intB,ok:=strDict[strB]
				if !ok{
					strDict[strB]=conteo
					intB=conteo
					graph.AddVertex(conteo)
					conteo++
				}
				D=Distance(A[0],A[1],B[0],B[1])
				fmt.Printf("Distancia %f",D)
				//graph.AddArc(intA,intB,D)
				//graph.AddArc(intB,intA,D)
			}
			A:=B
			noA:=noB
						
		}
	}

 

http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)

}