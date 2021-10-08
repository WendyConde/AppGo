package main

// Importing the necessary libraries
import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/twpayne/go-polyline"
)

//Principal function
func main() {
	fmt.Println("Please enter your path: ")

	/*The Bufio library was used to parse the path
	received from the console and also call the readFile function*/
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		path := scanner.Text()
		fmt.Printf("Input was: %q\n", path)
		readFile(path)
	}
}

func readFile(path string) {
	var coordsList [][]float64

	/* File reading*/
	fptr := flag.String("fpath", path, "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)

	/*Regular expression created to validate the location*/
	regex := regexp.MustCompile("latitude: ?-?\\d+\\.?\\d*\\, longitude: ?-?\\d+\\.?\\d*|-?\\d+\\.?\\d*\\,-?\\d+\\.?\\d*")

	for s.Scan() {
		coord := regex.FindString(s.Text())
		if coord != "" {

			/*Replacing unnecessary characters*/
			if strings.Contains(coord, "latitude") || strings.Contains(coord, "longitude") {
				regex := regexp.MustCompile("[a-zA-Z: ]")
				coord = regex.ReplaceAllString(coord, "")
			}

			coordx := strings.Split(coord, ",")[0]
			coordy := strings.Split(coord, ",")[1]

			//Converts values to float type
			coordxNumber, err := strconv.ParseFloat(coordx, 64)
			coordyNumber, err := strconv.ParseFloat(coordy, 64)

			if err != nil {
				fmt.Println("Error de lectura", err)
			}

			/*Matrix is generated based on the coordinates
			found in the file line*/
			coordPairs := [][]float64{
				{coordxNumber, coordyNumber},
			}
			coordsList = append(coordsList, coordPairs...)
		}
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	/* A matrix of coordinates found is printed
	using the polyline library*/
	if len(coordsList) > 0 {
		fmt.Printf("%s\n", polyline.EncodeCoords(coordsList))
	} else {
		fmt.Println("There are no lines to encode")
	}
}
