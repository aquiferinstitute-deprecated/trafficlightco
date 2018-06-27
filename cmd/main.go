package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aquiferinstitute/trafficlightco"
)

func main() {

	apiKey := flag.String("apikey", "", "traffic-light.co API key")
	country := flag.String("country", "", "country (FR, US...)")
	companyName := flag.String("name", "", "company name")
	flag.Parse()

	if *apiKey == "" {
		buf := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your API Key: ")
		in, err := buf.ReadBytes('\n')
		if err != nil {
			log.Fatalf("error %v", err)
		}
		*apiKey = string(in[:len(in)-1])
	}
	if *country == "" {
		buf := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the country of the company (FR, US...): ")
		in, err := buf.ReadBytes('\n')
		if err != nil {
			log.Fatalf("error %v", err)
		}
		*country = string(in[:len(in)-1])
	}
	if *companyName == "" {
		buf := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the company name: ")
		in, err := buf.ReadBytes('\n')
		if err != nil {
			log.Fatalf("error %v", err)
		}
		*companyName = string(in[:len(in)-1])
	}

	c, err := trafficlightco.NewClient("https://api.traffic-light.co", *apiKey)
	if err != nil {
		log.Fatalf("trafficlightco: %v", err)
	}
	fmt.Println()
	cp, errS := searchCountries(c, *country, *companyName)
	if errS != nil {
		log.Fatalf("trafficlightco: %v", errS)
	}
	fmt.Println(displayCompanyResult(cp))
	d, errR := displayCountriesRatings(c, cp)
	if errR != nil {
		log.Fatalf("trafficlightco: %v", errR)
	}
	fmt.Println(d)
}

func searchCountries(
	c *trafficlightco.Client,
	country, name string,
) ([]trafficlightco.Company, error) {

	// Search companies
	cp, errS := c.Search(country, name, "", "")
	if errS != nil {
		return nil, fmt.Errorf("searchCountries: %v", errS)
	}
	if len(cp) == 0 {
		return nil,
			fmt.Errorf(
				"searchCountries: found no companies matching your criteria",
			)
	}
	return cp, nil
}

func displayCountriesRatings(
	c *trafficlightco.Client,
	cp []trafficlightco.Company,
) (string, error) {
	var res string
	for _, comp := range cp {
		// Lookup trafficlight
		r, errR := c.GetRating(comp.ID)
		if errR != nil {
			return "", fmt.Errorf("trafficlightco: %v", errR)
		}
		res += fmt.Sprintf(
			"%s - rating: %s on %s\n",
			comp.Name,
			r.Color,
			r.Created.Format("Mon Jan _2 15:04:05 2006"),
		)
	}
	return res, nil
}

func displayCompanyResult(cp []trafficlightco.Company) string {
	var ns []string
	for _, n := range cp {
		ns = append(ns, n.Name)
	}
	return fmt.Sprintf(
		"found %d company(ies): %s",
		len(cp),
		strings.Join(ns, " ,"),
	)
}
