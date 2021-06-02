package helper

import (
	"math/rand"

	uuid "github.com/satori/go.uuid"
	word "github.com/zhexuany/wordGenerator"
)

// returns a number
func NewNumber() uint64 {
	return rand.Uint64()
}

// returns a new uuid
func NewID() string {
	u2 := uuid.NewV4()
	return u2.String()
}

// returns a name
func NewName() string {
	return randomStringFromSet("Albert", "Husan", "Samuel", "Ilyos", "Hoshim", "Asadbek", "Davlatbek", "Javohir", "Muxsin", "Farhod", "Jo'rabek", "Ismoil", "Lola", "Malika", "Nodira", "Maftuna", "Mohichehra", "Gullola", "Ann", "Ixtiyor")
}

// returns an address
func NewAddress() string {
	return randomStringFromSet("Mustaqillik k. 10-14", "Amir Temur k. 12-10", "Istiqlol k. 12-14", "Do'stlik k. 22-19", "Cho'l Guli k. 45-10", "Islom Karimov k. 45-19", "Navoiy k. 65-12")
}

// returns an age in [20: 65] range
func NewAge() int {
	min := 20
	max := 65
	return rand.Intn(max-min+1) + min
}

func NewJob() string {
	return randomStringFromSet("teacher", "engineer", "driver", "machenist", "dresser", "baker", "taxi driver", "coder", "software engineer", "sportsman", "football player", "barman", "investor", "businessman")
}

func randomBody() string {
	return ""
}

// returns random string
func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}

	return a[rand.Intn(n)]
}

// returns max 10 / min 5 char.d word
func RandomTask() []string {
	return word.GetWords(5, 10)
}
