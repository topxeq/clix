package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/topxeq/tk"
)

const clixVersionG = "0.99a"

func main() {
	argsT := os.Args

	argsLenT := len(argsT)

	if argsLenT < 2 {
		tk.Pl("empty command")
		return
	}

	subCmdT := strings.ToLower(argsT[1])

	switch subCmdT {
	case "version":
		tk.Pl("Clix v%v", clixVersionG)
		break

	case "simpleencode":
		if argsLenT < 3 {
			tk.Pl(`not enough parameters`)
			break
		}

		tk.Pl("%v", tk.EncodeStringSimple(argsT[2]))
		break
	case "simpledecode":
		if argsLenT < 3 {
			tk.Pl(`not enough parameters`)
			break
		}

		tk.Pl("%v", tk.DecodeStringSimple(argsT[2]))

		break
	case "sortnumbers":
		if argsLenT < 3 {
			tk.Pl(`please enter the numbers to sort, such as 1,5,6,7,2`)
			break
		}

		s := argsT[2]

		listT := strings.Split(s, ",")

		numberListT := make([]float64, len(listT))

		for i, v := range listT {
			fmt.Sscanf(v, "%f", &numberListT[i])
		}

		tk.Pl("Sorted numbers: %v", numberListT)

		listLenT := len(numberListT)

		for i := 0; i < (listLenT - 1); i++ {
			for j := i + 1; j < listLenT; j++ {
				if numberListT[i] < numberListT[j] {
					numberListT[i], numberListT[j] = numberListT[j], numberListT[i]
				}
			}
		}

		tk.Pl("Sorted numbers: %v", numberListT)

		break

	default:
		tk.Pl("unknown command")
		break
	}

}
