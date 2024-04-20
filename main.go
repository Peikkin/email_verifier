package main

import (
	"bufio"
	"net/mail"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		mail, err := mail.ParseAddress(scanner.Text())
		if err != nil {
			log.Error().Err(err).Msgf("%v адресс не действителен", scanner.Text())
		} else {
			log.Info().Msgf("%v действителен", mail)
		}
	}
}
