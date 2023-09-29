package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

var foundPort = make(chan int, 1)

func FetchPort(port int, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("http://10.49.122.144:%d/ping", port)

	resp, err := http.Get(url)
	if err != nil {

		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		// fmt.Printf("Error reading response from %s: %v\n", url, err)
		return
	}
	foundPort <- port

}

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func main() {
	var wg sync.WaitGroup

	for port := 0; port < 10000; port++ {
		wg.Add(1)
		go FetchPort(port, &wg)
	}
	go func() {
		wg.Wait()
		close(foundPort)
	}()
	foundPort := <-foundPort
	fmt.Printf("Found port: %d\n", foundPort)

	user := "agus"
	payload := map[string]string{
		"user": user,
	}
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Adding the user
	resp1 := FetchUrl(fmt.Sprintf("http://10.49.122.144:%d/signup", foundPort), jsonBody)
	fmt.Print(resp1)

	// Check the user
	resp2 := FetchUrl(fmt.Sprintf("http://10.49.122.144:%d/check", foundPort), jsonBody)
	fmt.Print(resp2)

	h := sha256.New()
	h.Write([]byte(user))
	secret := fmt.Sprintf("%x", h.Sum(nil))

	payload = map[string]string{
		"user":   user,
		"secret": secret,
	}

	jsonBody, err = json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// GetUserLevel
	for i := 0; i < 1; i++ {
		resp4 := FetchUrl(fmt.Sprintf("http://10.49.122.144:%d/getUserLevel", foundPort), jsonBody)
		fmt.Print(resp4)
	}

	// GetUserPoints
	for i := 0; i < 1; i++ {
		resp5 := FetchUrl(fmt.Sprintf("http://10.49.122.144:%d/getUserPoints", foundPort), jsonBody)
		fmt.Print(resp5)
	}

	// Get all the Hints
	var slicer []string
	for i := 0; i < 100; i++ {
		resp6 := FetchUrl(fmt.Sprintf("http://10.49.122.144:%d/iNeedAHint", foundPort), jsonBody)
		slicerAns := strings.Trim(resp6, "Coward over here asking for hints...\nHere you go, your random hint:\n")
		if !contains(slicer, slicerAns) {
			slicer = append(slicer, slicerAns)
		}
	}
	fmt.Println("All hints :")
	for _, element := range slicer {
		fmt.Println("|-", element)
	}

	// enterChallenge
	for i := 0; i < 1; i++ {
		resp4 := FetchUrl(fmt.Sprintf("http://10.49.122.144:%d/enterChallenge", foundPort), jsonBody)
		fmt.Print(resp4)
	}

	for i := 0; i < 1; i++ {
		resp5 := FetchUrl(fmt.Sprintf("http://10.49.122.144:%d/submitSolution", foundPort), jsonBody)
		fmt.Print(resp5)
	}

	jsonQuotesData := `[
	{
	  "text": "Always code as if the guy who ends up maintaining your code will be a violent psychopath who knows where you live.",
	  "author": "Martin Golding"
	},
	{
	  "text": "All computers wait at the same speed.",
	  "author": "Unknown"
	},
	{
	  "text": "A misplaced decimal point will always end up where it will do the greatest damage.",
	  "author": "Unknown"
	},
	{
	  "text": "A good programmer looks both ways before crossing a one-way street.",
	  "author": "Unknown"
	},
	{
	  "text": "A computer program does what you tell it to do, not what you want it to do.",
	  "author": "Unknown"
	},
	{
	  "text": "\"Intel Inside\" is a Government Warning required by Law.",
	  "author": "Unknown"
	},
	{
	  "text": "Common sense gets a lot of credit that belongs to cold feet.",
	  "author": "Arthur Godfrey"
	},
	{
	  "text": "Chuck Norris doesn’t go hunting. Chuck Norris goes killing.",
	  "author": "Unknown"
	},
	{
	  "text": "Chuck Norris counted to infinity... twice.",
	  "author": "Unknown"
	},
	{
	  "text": "C is quirky, flawed, and an enormous success.",
	  "author": "Unknown"
	},
	{
	  "text": "Beta is Latin for still doesn’t work.",
	  "author": "Unknown"
	},
	{
	  "text": "ASCII stupid question, get a stupid ANSI!",
	  "author": "Unknown"
	},
	{
	  "text": "Artificial Intelligence usually beats natural stupidity.",
	  "author": "Unknown"
	},
	{
	  "text": "Any fool can use a computer. Many do.",
	  "author": "Ted Nelson"
	},
	{
	  "text": "Hey! It compiles! Ship it!",
	  "author": "Unknown"
	},
	{
	  "text": "Hate cannot drive out hate; only love can do that.",
	  "author": "Martin Luther King Junior"
	},
	{
	  "text": "Guns don’t kill people. Chuck Norris kills people.",
	  "author": "Unknown"
	},
	{
	  "text": "God is real, unless declared integer.",
	  "author": "Unknown"
	},
	{
	  "text": "First, solve the problem. Then, write the code.",
	  "author": "John Johnson"
	},
	{
	  "text": "Experience is the name everyone gives to their mistakes.",
	  "author": "Oscar Wilde"
	},
	{
	  "text": "Every piece of software written today is likely going to infringe on someone else’s patent.",
	  "author": "Miguel de Icaza"
	},
	{
	  "text": "Computers make very fast, very accurate mistakes.",
	  "author": "Unknown"
	},
	{
	  "text": "Computers do not solve problems, they execute solutions.",
	  "author": "Unknown"
	},
	{
	  "text": "I have NOT lost my mind—I have it backed up on tape somewhere.",
	  "author": "Unknown"
	},
	{
	  "text": "If brute force doesn’t solve your problems, then you aren’t using enough.",
	  "author": "Unknown"
	},
	{
	  "text": "It works on my machine.",
	  "author": "Unknown"
	},
	{
	  "text": "Java is, in many ways, C++??.",
	  "author": "Unknown"
	},
	{
	  "text": "Keyboard not found...Press any key to continue.",
	  "author": "Unknown"
	},
	{
	  "text": "Life would be so much easier if we only had the source code.",
	  "author": "Unknown"
	},
	{
	  "text": "Mac users swear by their Mac, PC users swear at their PC.",
	  "author": "Unknown"
	},
	{
	  "text": "Microsoft is not the answer. Microsoft is the question. \"No\" is the answer.",
	  "author": "Unknown"
	},
	{
	  "text": "MS-DOS isn’t dead, it just smells that way.",
	  "author": "Unknown"
	},
	{
	  "text": "Only half of programming is coding. The other 90% is debugging.",
	  "author": "Unknown"
	},
	{
	  "text": "Pasting code from the Internet into production code is like chewing gum found in the street.",
	  "author": "Unknown"
	},
	{
	  "text": "Press any key to continue or any other key to quit.",
	  "author": "Unknown"
	},
	{
	  "text": "Profanity is the one language all programmers know best.",
	  "author": "Unknown"
	},
	{
	  "text": "The best thing about a boolean is even if you are wrong, you are only off by a bit.",
	  "author": "Unknown"
	},
	{
	  "text": "The nice thing about standards is that there are so many to choose from.",
	  "author": "Unknown"
	},
	{
	  "text": "There are 3 kinds of people: those who can count and those who can’t.",
	  "author": "Unknown"
	},
	{
	  "text": "There is no place like 127.0.0.1",
	  "author": "Unknown"
	},
	{
	  "text": "There is nothing quite so permanent as a quick fix.",
	  "author": "Unknown"
	},
	{
	  "text": "There’s no test like production.",
	  "author": "Unknown"
	},
	{
	  "text": "To err is human, but for a real disaster you need a computer.",
	  "author": "Unknown"
	},
	{
	  "text": "Ubuntu is an ancient African word, meaning \"can’t configure Debian\"",
	  "author": "Unknown"
	},
	{
	  "text": "UNIX is the answer, but only if you phrase the question very carefully.",
	  "author": "Unknown"
	},
	{
	  "text": "Usenet is a Mobius strand of spaghetti.",
	  "author": "Unknown"
	},
	{
	  "text": "Weeks of coding can save you hours of planning.",
	  "author": "Unknown"
	},
	{
	  "text": "When your computer starts falling apart, stop hitting it with a Hammer!",
	  "author": "Unknown"
	},
	{
	  "text": "Who is General Failure? And why is he reading my disk?",
	  "author": "Unknown"
	},
	{
	  "text": "You can stand on the shoulders of giants OR a big enough pile of dwarfs, works either way.",
	  "author": "Unknown"
	},
	{
	  "text": "You start coding. I’ll go find out what they want.",
	  "author": "Unknown"
	},
	{
	  "text": "I love deadlines. I like the whooshing sound they make as they fly by.",
	  "author": "Douglas Adams"
	},
	{
	  "text": "I think we agree, the past is over.",
	  "author": "George W. Bush"
	},
	{
	  "text": "In order to be irreplaceable, one must always be different.",
	  "author": "Coco Chanel"
	},
	{
	  "text": "In the future everyone will be famous for fifteen minutes.",
	  "author": "Andy Warhol"
	},
	{
	  "text": "In three words I can sum up everything I’ve learned about life: it goes on.",
	  "author": "Robert Frost"
	},
	{
	  "text": "It is a mistake to think you can solve any major problems just with potatoes.",
	  "author": "Douglas Adams"
	},
	{
	  "text": "It’s kind of fun to do the impossible.",
	  "author": "Walt Disney"
	},
	
	{
	  "text": "Design is choosing how you will fail.",
	  "author": "Ron Fein"
	},
	{
	  "text": "Focus is saying no to 1000 good ideas.",
	  "author": "Steve Jobs"
	},
	{
	  "text": "Code never lies, comments sometimes do.",
	  "author": "Ron Jeffries"
	},
	{
	  "text": "Be careful with each other, so you can be dangerous together.",
	  "author": "Unknown"
	},
	{
	  "text": "When making a PR to a major project, you have to \"sell\" that feature / contribution. You have to be convincing on why it belongs there. The maintainer is going to be the one babysitting that code forever.",
	  "author": "Taylor Otwell"
	},
	{
	  "text": "Be the change you wish was made. Share the lessons you wish you'd been taught. Make the sacrifices you wish others had made.",
	  "author": "David Heinemeier Hansson"
	},
	{
	  "text": "The only way to achieve flexibility is to make things as simple and easy to change as you can.",
	  "author": "Casey Brant"
	},
	{
	  "text": "The computer was born to solve problems that did not exist before.",
	  "author": "Bill Gates"
	},
	{
	  "text": "People don't care about what you say, they care about what you build.",
	  "author": "Mark Zuckerberg"
	},
	{
	  "text": "We will not ship shit.",
	  "author": "Robert C. Martin"
	}
  ]
  `
	// Unmarshal JSON data into a slice of Quote
	var quotes []Quote
	err = json.Unmarshal([]byte(jsonQuotesData), &quotes)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	// for _, quote := range quotes {
	// 	md5Hash := calculateMD5(quote.Text)
	// 	fmt.Println(quote.Text, md5Hash)
	// }
	// The good quote corresponding to Hint 5bc2fb8cff6b14d9c62ea6447da62a4c is :
	// Pasting code from the Internet into production code is like chewing gum found in the street.
}

func calculateMD5(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hashInBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}

func FetchUrl(url string, jsonBody []byte) string {
	qBody := bytes.NewBuffer(jsonBody)
	resp, err := http.Post(url, "application/json", qBody)

	if err != nil {
		// fmt.Printf("Error fetching from %s: %v\n", url, err)
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// fmt.Printf("Error reading response from %s: %v\n", url, err)
		log.Fatal(err)
	}
	// fmt.Printf("Response from %s: %s\n", url, string(body))
	return string(body)
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
