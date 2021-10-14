package main

import (
	"encoding/json"
	"fmt"

	"github.com/bregydoc/gtranslate"
	"github.com/pariz/gountries"
	"golang.org/x/text/language"
)

// // explicit reads credentials from the specified path.
// func explicit(jsonPath, projectID string) {
// 	ctx := context.Background()
// 	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Close()
// 	fmt.Println("Buckets:")
// 	it := client.Buckets(ctx, projectID)
// 	for {
// 		battrs, err := it.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(battrs.Name)
// 	}
// }

func main() {

	query := gountries.New()
	usa, _ := query.FindCountryByName("ethiopia")

	x, _ := json.Marshal(usa)
	fmt.Println(string(x))
	// return
	text := "I have done work"
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: language.English.String(),
			To:   "amh",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | am: %s \n", text, translated)

	// pwd, _ := os.Getwd()
	// jsonPath := filepath.Join(pwd, "../config/do-staff-e6668985ba67.json")

	// ctx := context.Background()
	// translateService, err := translate.NewService(ctx, option.WithCredentialsFile(jsonPath))
	// if err != nil {
	// 	panic(err)
	// }

	// a := &translate.TranslateTextRequest{
	// 	Q:      []string{"Hello World"},
	// 	Source: "English",
	// 	Target: "Amharic",
	// }

	// translatedList, err := translateService.Translations.Translate(a).Do()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(translatedList)
}
