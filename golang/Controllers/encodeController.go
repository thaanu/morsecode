package controllers

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string `json:"message"`
}

type ReturnResponse struct {
	Status string `json:"status"`
	Data string `json:"encoded_message"`
}

func Home(c *fiber.Ctx) error {
	return c.SendString("Welcome to morse code translator")
}

func EncodeMorseCode(c *fiber.Ctx) error {
	c.Accepts("application/json")

	majorResponse := "Success"

	// todo: Validations (isset, number of characters)
	// todo: Decode morse to plain text

	r := new(Response)

	if err := c.BodyParser(r); err != nil {
		return err
	}

	if r.Message == "" {
		// majorResponse = fiber.NewError(fiber.StatusUnprocessableEntity, "Message is empty")
		majorResponse = "Message is empty"
	}

	encoded := morseCodeEncode( r.Message )

	log.Println(encoded)

	if majorResponse != "Success" {
		c.Status(400)
	}

	return c.SendString(SendResponse(majorResponse, encoded ))
	
}

func DecodeMorseCode(c *fiber.Ctx) error {
	c.Accepts("application/json")

	majorResponse := "Success"

	r := new(Response)

	if err := c.BodyParser(r); err != nil {
		return err
	}

	if r.Message == "" {
		// majorResponse = fiber.NewError(fiber.StatusUnprocessableEntity, "Message is empty")
		majorResponse = "Message is empty"
	}

	decoded := morseCodeDecode( r.Message )

	if majorResponse != "Success" {
		c.Status(400)
	}

	return c.SendString(SendResponse(majorResponse, decoded ))

	
}

func SendResponse(status string, data string) string {
	returnResponse := ReturnResponse{
		Status: status,
		Data: data,
	}
	u, _ := json.Marshal(returnResponse)

	return string(u)
	
}

func morseCodeEncode( plainText string ) string {
	morseCodeCodes := map[string]string{
		"a":".-",
		"b":"-...", 
		"c":"-.-.", 
		"d":"-..", 
		"e":".", 
		"f":"..-.", 
		"g":"--.", 
		"h":"....", 
		"i":"..", 
		"j":".---", 
		"k":"-.-", 
		"l":".-..", 
		"m":"--", 
		"n":"-.", 
		"o":"---", 
		"p":".--.", 
		"q":"--.-", 
		"r":".-.", 
		"s":"...", 
		"t":"-", 
		"u":"..-", 
		"v":"...-", 
		"w":".--", 
		"x":"-..-", 
		"y":"-.--", 
		"z":"--..", 
		"0":"-----",
		"1":".----", 
		"2":"..---", 
		"3":"...--", 
		"4":"....-", 
		"5":".....", 
		"6":"-....", 
		"7":"--...", 
		"8":"---..", 
		"9":"----.",
		".":".-.-.-",
		",":"--..--",
		"?":"..--..",
		"/":"-..-.",
		" ":"/",
	}

	lookupString := strings.Split(strings.ToLower(plainText), "")
	
	encodedMorseCode := ""

	for _, char := range lookupString {
		sMorseCode := morseCodeCodes[char]
		encodedMorseCode = encodedMorseCode+" "+sMorseCode
	}

	return encodedMorseCode

}

func morseCodeDecode( morseCode string ) string {
	
	decodedPlainText := ""
	allWords := strings.Split(morseCode, "  ")

	// Loop all the words
	for _, word := range allWords {
		// Split all the words
		letters := strings.Split(word, " ")

		decodedWord := ""

		for _, letter := range letters {
			decodedWord = decodedWord + findLetter(letter)
		}

		decodedPlainText = decodedPlainText + decodedWord

	}

	return decodedPlainText

}

func findLetter(letter string) string {
	morseCodeCodes := map[string]string{
		"a":".-",
		"b":"-...", 
		"c":"-.-.", 
		"d":"-..", 
		"e":".", 
		"f":"..-.", 
		"g":"--.", 
		"h":"....", 
		"i":"..", 
		"j":".---", 
		"k":"-.-", 
		"l":".-..", 
		"m":"--", 
		"n":"-.", 
		"o":"---", 
		"p":".--.", 
		"q":"--.-", 
		"r":".-.", 
		"s":"...", 
		"t":"-", 
		"u":"..-", 
		"v":"...-", 
		"w":".--", 
		"x":"-..-", 
		"y":"-.--", 
		"z":"--..", 
		"0":"-----",
		"1":".----", 
		"2":"..---", 
		"3":"...--", 
		"4":"....-", 
		"5":".....", 
		"6":"-....", 
		"7":"--...", 
		"8":"---..", 
		"9":"----.",
		".":".-.-.-",
		",":"--..--",
		"?":"..--..",
		"/":"-..-.",
		" ":"/",
	}

	for key, value := range morseCodeCodes {
		if ( value == letter ) {
			return key
		}
	}
	return ""
}