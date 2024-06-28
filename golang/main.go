package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {

	type Response struct {
		PlainText string `json:"plain_text"`
	}

	app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	app.Post("/encode", func (c *fiber.Ctx) error {
		c.Accepts("application/json")

		r := new(Response)

		if err := c.BodyParser(r); err != nil {
			return err
		}

		encoded := morseCodeEncode( r.PlainText )

		log.Println(encoded)

		return c.SendString(encoded)
	})

    log.Fatal(app.Listen(":3000"))

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
		" ":" ",
	}

	lookupString := strings.Split(strings.ToLower(plainText), "")
	
	encodedMorseCode := ""

	for _, char := range lookupString {
		sMorseCode := morseCodeCodes[char]
		encodedMorseCode = encodedMorseCode+" "+sMorseCode
	}

	return encodedMorseCode

}