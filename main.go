package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/NotAliAhmad/learningGO/helpers"
)

// did i expect to unmarshall - meaning how i will take json info an put it into a struct object
type Movie struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Length      int    `json:"length"`
	Is_Seen     bool   `json:"is_seen"`
}

type Movies struct {
	Name        string 
	Description string 
	Length      int    
	Is_Seen     bool   
}

type person struct {
	firstName   string
	lastName    string
	phoneNumber string
	age         int
	birthDate   time.Time
}

type fru struct {
	fruit string
	size  int
	color string
	perf  bool
}

type fru_maker struct {
	fruType string
}

// this is a receiver
func (m *person) changeFirstName() string {
	m.firstName = "razi"
	return m.firstName
}

// creating an interface
type electronic interface {
	description() string
	year() int
}

type gameboy struct {
	button    int
	enjoyment string
}

type playstation struct {
	controller int
	memory     string
	yearsOwned string
}

func main() {
	fmt.Println("hello-world")

	// declaring a variable
	var whattosay string
	var orderNum int = 20

	whattosay = "yo yo what up"

	fmt.Println(whattosay)

	fmt.Println(orderNum)

	//pointers
	num := 20
	fmt.Println(num)
	fmt.Println("the memory location of num is", &num)
	num2 := 40
	var s *int = &num
	*s = num2
	fmt.Println("the changed value of num is", num)

	// go can declare 2 returned variables and put that into 2 variables as well
	fruitType1, fruitType2 := Fruit()
	fmt.Println("my favorite fruit is "+fruitType1+" and "+fruitType2+" and i want", orderNum)

	// pointing to the reference
	log.Println("my old fruit is", whattosay)
	changeFruit(&whattosay)
	log.Println("my new fruit is", whattosay)

	me := person{
		firstName: "ali",
		lastName:  "ahmad",
	}
	log.Println(me.birthDate)
	log.Println(me.changeFirstName(), "-using a function to alter the data stucture")

	// creating a map that stores strings and at the place reference int, this can also be string or other primitive data types
	// int is the reference that it will grab
	// string is the data type that the map stores
	myMap := make(map[int]string)

	myMap[2] = "bubbles"
	log.Println("this is an index of 2 getting the value", myMap[2])

	tree := make(map[string]fru)

	tree["Malus domestica"] = fru{
		fruit: "apple",
		size:  15,
		color: "brown",
		perf:  true,
	}

	log.Println(tree["Malus domestica"], "is the produced of the fruit", tree["Malus domestica"].fruit)

	// slices of strings
	var myFamily []string

	myFamily = append(myFamily, "ali")
	myFamily = append(myFamily, "mahad")
	myFamily = append(myFamily, "razi")
	myFamily = append(myFamily, "jozi")
	myFamily = append(myFamily, "dunzy")
	myFamily = append(myFamily, "mumzy")

	// printing out all my brothers from the first index to the 3th since they being at 0
	log.Println(myFamily[1:4])

	// conditional statements

	cat := true
	amount := 51

	if cat == false && amount <= 50 {
		log.Println("he better be gentle or im going chinese on his ass")
	} else if amount > 50 && cat == true {
		log.Println("I will think about it")
	}

	// If there are more then 1 conditions then use the SWITCH statement
	myFruit := "apple"

	switch myFruit {
	case "apple":
		log.Println("an apple a day keeps the doctor away")

	case "banana":
		log.Println("bananas make good milkshakes")

	case "watermelon":
		log.Println("alot of seeds but watermelons are still da sheet")
	// if none of these conditions are met ^^ then it will go to default
	default:
		log.Println("fruits are still awesome")
	}

	// Loops

	for i := 0; i < 5; i++ {
		log.Println(i)
	}
	// range in loops
	veggies := []string{"brocolli", "celery", "pumpkin", "sprouts", "carrots"}

	for _, veg := range veggies {
		log.Println(veg)
	}

	furniture := make(map[string]string)
	furniture["bed"] = "bedroom"
	furniture["dining table"] = "kitchen"
	furniture["table"] = "bedroom"

	var fruits []fru

	fruits = append(fruits, fru{"apple", 5, "red", true})
	fruits = append(fruits, fru{"banana", 8, "yellow", true})
	fruits = append(fruits, fru{"papaya", 15, "orange/green", false})
	fruits = append(fruits, fru{"mango", 5, "yellow", true})

	/* another way to do the above is the following  ~~~~~SLICES~~~~~

	fruits := []fru{
		{"apple",5,"red",true} ,
		{"banana",8,"yellow",true} ,
		{"papaya",15,"orange/green",false} ,
		{"mango",5,"yellow",true} ,
		}
	*/
	for _, fruit := range fruits {
		log.Println(fruit.color, fruit.fruit, fruit.perf, fruit.size)
	}

	for furniture, room := range furniture {
		log.Println(furniture, "belongs in the", room)
	}

	// interface
	gbAdvance := gameboy{
		button:    4,
		enjoyment: "this was my first game so it has a special place in my heart",
	}

	ps2 := playstation{
		controller: 4,
		memory:     "n/a",
		yearsOwned: "5",
	}

	bestGame(&gbAdvance)
	bestGame(&ps2)

	// using a struct from the helpers file
	var something helpers.SomeType
	something.TypeName = "sfdsf"
	something.TypeNumber = 5

	log.Println(something)

	// using a channel
	initChan := make(chan int)
	// this will close the connection as soon as the function is done
	// this is good practice because of security and scalability
	defer close(initChan)

	// making it a go routine
	go CaculateVal(initChan)

	// whatever value the channel has it will put it into the variable
	randNum := <-initChan
	log.Println(randNum)

	// declaring an unmarshal slice of type movie
	var unmarshalled []Movie

	// raw json info
	myJson := `
	[
		{
			"name": "batman",
			"description": "a bored billionair flying around with a bat suit",
			"length": 120,
			"is_seen": false
		},
		{
			"name": "superman",
			"description": "an alien who has super powers cus of the yellow sun and lives on earth",
			"length": 90,
			"is_seen": true
		}
	]`
	// reading json into a struct 
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("error unmarshalling json", err)
	}
	log.Println("unmarshalled:", unmarshalled)

	// writing json from a struct 
	maze_runner := Movies{
		Name: "Maze Runner",
		Description: "post shitty earth virus and kids can survive it",
		Length: 100,
		Is_Seen: true,
	}

	var myMovies []Movies

	maze_runner2 := Movies{
		Name: "Maze Runner 2",
		Description: "they win the maze and run out only to get fked by the government",
		Length: 100,
		Is_Seen: true,
	}
	myMovies = append(myMovies, maze_runner)
	myMovies = append(myMovies, maze_runner2)

	// marshalling data into json
	myJsonMovies, err := json.MarshalIndent(&myMovies, "", "   ")
	// it returns an error if something goes wrong
	// for the positive path it should be nil 
	if err != nil {
		log.Println("error unmarshalling json", err)
	}
	log.Println("marshalled:", string(myJsonMovies))

	// testing
	dividend,err := divide(5,0)
	if err !=nil{
		log.Println(err)
	}
	log.Println(dividend)

	
	
}

func divide(x, y float32) (float32, error){
	result := x/y

	if y == 0{
		return result, errors.New("cannot divide by 0")
	}
	return result, nil
}

// declaring a const variable
const randNumPoolSize = 1000

// func for my channel
func CaculateVal(intChan chan int) {
	perRandNum := helpers.RandomNum(randNumPoolSize)
	intChan <- perRandNum
}

// interface type func
func bestGame(a electronic) {
	log.Println("This electronic came out on", a.year(), "and", a.description())
}

//interface func 1
func (a *playstation) description() string {
	return "The PlayStation 2 (PS2) is a home video game console developed and marketed by Sony Computer Entertainment"
}

//interface func 2
func (a *playstation) year() int {
	return 1994
}

//interface func 1
func (a *gameboy) description() string {
	return "It was a handheld video game console that combined aspects of Nintendo"
}

//interface func 2
func (b *gameboy) year() int {
	return 2003
}

// go can return multiple types
// this method starts with a uppercase therefore its available outside this package
func Fruit() (string, string) {
	return "apple", "banana"
}

// this method starts with a lowercase therefore its only available to this package
func changeFruit(f *string) {

	newFruit := "avocado"
	*f = newFruit
}
