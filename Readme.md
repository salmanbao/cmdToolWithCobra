# Viper(Golang)

This the demonstration of viper to create cli tools and apps in golang

## Run

Install the viper.
```
go get github.com/spf13/cobra
```
`

## Import

```go
import (
	"fmt"
	"os"

	. "github.com/salmanbao/cmdToolWithCobra/cmd"
)

func main() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```
## Usage
```
1- ./binary-file version
2- ./binary-file -v
3- ./binary-file --version
4- ./binary-file print hello
5- ./binary-file print hello
6- ./binary-file print times hello --times 5
```

## License

[MIT](https://choosealicense.com/licenses/mit/)