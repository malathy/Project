 package main

  import ( 
	"encoding/json"
	"log"
	//"net/http"
"os"
	"fmt")

//Creating the URL Struct

type URL struct{
	Url string `json:"url"`
	Product  Product  `json:"product"`
}
//Creating Product Struct
type Product struct
	{
	Name         string `json:"name"`
	Imageurl     string `json:"imageurl"`
	Description  string `json: "description"`
	Price        string `json : "price"`
	TotalReviews int    `json:"totalreviews"`
   }

 //Main Function
func main(){

	//Creating instances of Product and URL struct
	
	product:=Product{Name:"PlayStation 4 Pro 1TB Console",Imageurl:"https://images-na.ssl-images-amazon.com/images/I/41GGPRqTZtL._AC_.jpg",Description: "Heighten your experiences.\n Enrich your adventures.",Price: "$348.00",TotalReviews:4436}
	url1:=URL{Url:"https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC/",Product:product}
	
	byteArray,err:=json.Marshal(url1)
	if err!=nil{
		fmt.Println(err)
	}
fmt.Println(string(byteArray))

//creating  response text file
file,e:=os.OpenFile("Response_new.text",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
if e!=nil{
	log.Fatalln("failed")

}
//Writing into text file
log.SetOutput(file)
log.Println(string(byteArray))

}


