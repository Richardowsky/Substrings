# Substrings

## Prerequisites:
You need to install - [go 1.13](https://golang.org/dl/)

## How to test:
1. Clone this repo: `git clone https://github.com/Richardowsky/Substrings.git`
2. `make test`

## How to use:
```go
package example

import str "Substrings/src"

func example()  {
	input := "ABABA"
	
	str.CheckString(input)
    str.CheckString2(input)
}

```