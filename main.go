package main

import (
	"fmt"
	"io/ioutil"
	mrand "math/rand"
	"os"
	"strings"
	"time"

	"berty.tech/berty/v2/go/pkg/banner"
	"github.com/spf13/pflag"
)

func main() {
	var (
		author string
		text   string
		qotd   bool
		random bool
		help   bool
		quote  banner.Quote
	)

	cmd := pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
	cmd.SortFlags = false

	cmd.StringVarP(&author, "author", "a", "", "author name")
	cmd.StringVarP(&text, "text", "t", "", "text to say")
	cmd.BoolVarP(&qotd, "qotd", "q", false, "say the quote of the day about privacy (override text and random flags)")
	cmd.BoolVarP(&random, "random", "r", false, "say a random quote about privacy (override text flag)")
	cmd.BoolVarP(&help, "help", "h", false, "display this help message (override all other flags)")

	if err := cmd.Parse(os.Args); err != nil || help {
		printUsageAndExit(cmd.FlagUsages(), err)
	}

	if qotd {
		quote = banner.QOTD()
	} else if random {
		mrand.Seed(time.Now().UnixNano())
		quote = banner.RandomQuote()
	} else if text != "" {
		quote.Text = text
	} else {
		stdin, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		quote.Text = strings.TrimSpace(string(stdin))
	}

	if author != "" {
		quote.Author = author
	}

	if quote.Author == "" {
		fmt.Print(banner.Say(fmt.Sprintf(`"%s"`, quote.Text)))
	} else {
		fmt.Print(banner.Say(quote.String()))
	}
}

func printUsageAndExit(flagUsage string, err error) {
	exitCode := 0
	output := os.Stdout
	errText := ""

	if err != nil {
		exitCode = 2
		output = os.Stderr
		errText = fmt.Sprintf("%s\n\n", err.Error())
	}

	fmt.Fprintf(
		output,
		"%sUsage: %s [[-t | -q | -r] [-a] | [-h]]\n\n%s\n%s\n\n%s",
		errText,
		os.Args[0],
		"Bertysay is like cowsay but with a parrot and optional quote about privacy",
		"If none or only the author flag is specified, will read on stdin",
		flagUsage,
	)

	os.Exit(exitCode)
}
