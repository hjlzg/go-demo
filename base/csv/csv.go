package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func writeCsv(file *os.File) {
	w := csv.NewWriter(file)
	w.Write([]string{"123", "456"})
	w.Flush()
}

func readCsv(file *os.File) {
	c := csv.NewReader(file)
	strs, _ := c.Read()
	for _, str := range strs {
		fmt.Println(str)
	}
}
