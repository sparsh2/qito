package passwordgenerator

import (
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

type PasswordGeneratorService struct {
	lowerList   string
	upperList   string
	numbersList string
	specialList string
}

func NewPasswordGeneratorService() *PasswordGeneratorService {
	return &PasswordGeneratorService{
		lowerList:   "abcdefghijklmnopqrstuvwxyz",
		upperList:   "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		numbersList: "0123456789",
		specialList: "!@#$%^&*",
	}
}

func (p *PasswordGeneratorService) Generate(options PasswordOptions) string {
	psw := stringWithCharset(options.Lower, p.lowerList) +
		stringWithCharset(options.Upper, p.upperList) +
		stringWithCharset(options.Numbers, p.numbersList) +
		stringWithCharset(options.Special, p.specialList)

	perm := rand.Perm(len(psw))
	final := make([]rune, len(psw))

	for i, v := range perm {
		final[v] = rune(psw[i])
	}

	return string(final)
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
