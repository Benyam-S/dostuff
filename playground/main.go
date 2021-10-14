package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	// query := gountries.New()
	// country, _ := query.FindCountryByName("Spain")

	// // x, _ := json.Marshal(usa)
	// // fmt.Println(string(x))
	// // return

	// targetLanguage := "am"
	// for key, value := range country.Languages {
	// 	fmt.Println(key, value)
	// 	// targetLanguage = key[:2]
	// 	// break
	// }

	// text := "I have done work"
	// translated, err := gtranslate.TranslateWithParams(
	// 	text,
	// 	gtranslate.TranslationParams{
	// 		From: language.English.String(),
	// 		To:   "am",
	// 	},
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("en: %s | am: %s \n", text, translated)

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

	url := "https://libretranslate.com/translate"

	payload, err := json.Marshal(map[string]string{"q": "Hello", "source": "en", "target": "es"})
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	req.Header.Add("content-type", "application/json")
	// req.Header.Add("x-rapidapi-key", "SIGN-UP-FOR-KEY")
	// req.Header.Add("x-rapidapi-host", "YandexTranslatezakutynskyV1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
