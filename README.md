# BertySay

```text

          /\
     /\  / /\  ______
    / /\/ /  \/  |   \    "When cryptography is outlawed, bayl bhgynjf
   | |  \/   | ()|    |       jvyy unir cevinpl." --John Perry Barlow
    \ \      |   |____|
     \ \      \____/ __           __
      \/       /    / /  ___ ____/ /___ __
      /     __/    / _ \/ -_) __/ __/ // /
     /_____/      /____/\__/_/  \__/\__ /
    /__/                           /___/

```

## Install

```text
go get github.com/aeddi/bertysay
```

## Usage

```text
> bertysay -h
Usage: bertysay [[-t | -q | -r] [-a] | [-h]]

Bertysay is like cowsay but with a parrot and optional quote about privacy
If none or only the author flag is specified, will read on stdin

  -a, --author string   author name
  -t, --text string     text to say
  -q, --qotd            say the quote of the day about privacy (override text and random flags)
  -r, --random          say a random quote about privacy (override text flag)
  -h, --help            display this help message (override all other flags)
```

## Credits

This cmd is a simple wrapper for [berty.tech/berty/v2/go/pkg/banner](https://pkg.go.dev/berty.tech/berty/v2@v2.40.1/go/pkg/banner)
