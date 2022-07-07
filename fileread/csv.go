package fileread

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func CsvReadByline(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	br := csv.NewReader(file)
	//bw := bufio.NewWriter(file1)
	for {
		record, err := br.Read()
		if err == io.EOF {
			fmt.Println("eof")
			break
		}
		fmt.Println(record[0])
	}
}

func CsvWriteByline(path string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	//追加写
	file.Seek(0, io.SeekEnd)
	br := csv.NewWriter(file)
	err = br.Write([]string{"vid", "1"})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = br.Write([]string{"vid", "2"})
	if err != nil {
		fmt.Println(err)
		return
	}
	br.Flush()
	fmt.Println("over")
}
