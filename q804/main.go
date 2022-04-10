package main

//https://leetcode-cn.com/problems/unique-morse-code-words/

var morse = []string{
	".-", "-...", "-.-.", "-..", ".", "..-.", "--.",
	"....", "..", ".---", "-.-", ".-..", "--", "-.",
	"---", ".--.", "--.-", ".-.", "...", "-", "..-",
	"...-", ".--", "-..-", "-.--", "--..",
}

func uniqueMorseRepresentations(words []string) int {
	m := make(map[string]struct{})
	for _, str := range words {
		var morseStr string
		for _, char := range str {
			morseStr += morse[char-'a']
		}
		m[morseStr] = struct{}{}
	}
	return len(m)
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	uniqueMorseRepresentations(words)
}
