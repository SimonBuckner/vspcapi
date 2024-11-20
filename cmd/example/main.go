package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SimonBuckner/vspcapi"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env files
	// .env.local takes precedence (if present)
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	baseUrl := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	// customerId := goquadac.StringtoI64(os.Getenv("SAAS_CUSTOMER_ID"))
	vspc := vspcapi.NewVspcApi(baseUrl)
	vspc.Authenticate(apiKey, apiSecret)

	// -------------------------------------------------------------------------
	//   Endpoints starting
	// -------------------------------------------------------------------------

	// printBanner("GET:/v1/saas/domains")
	// allDomains, err := datto.GetSaasDomains().GetAll()
	// if err != nil {
	// 	panic(err)
	// }
	// dump.Println(allDomains)

}

func printBanner(message string) {
	fmt.Println()
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("    " + message)
	fmt.Println("-------------------------------------------------------------")
	fmt.Println()
}
