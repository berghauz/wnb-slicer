package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func main() {

	var (
		fSize          uint
		maxChunkSize32 uint = 32768
		maxChunkSize64 uint = 65536
		chunkSizeFlag  *bool
		fileName       *string
		chunks         uint
		modulo         uint
		givenChunkSize uint
		rawBuffer      []byte
		csum           uint16
		chunkIter      uint
	)

	chunkSizeFlag = flag.Bool("max-chunk-size", false, "default is 32768, enable if chunk size should be 65536")
	fileName = flag.String("file-name", "", "file to slice")
	flag.Parse()

	file, err := os.Open(*fileName)
	checkError(err)
	defer file.Close()

	fInfo, err := file.Stat()
	checkError(err)

	fSize = uint(fInfo.Size())
	log.Println("file size is:", fSize)

	if *chunkSizeFlag {
		chunks = fSize / maxChunkSize64
		if fSize > maxChunkSize64 {
			modulo = fSize % maxChunkSize64
			// chunks++
		} else {
			chunks = 1
		}
		givenChunkSize = maxChunkSize64
	} else {
		chunks = fSize / maxChunkSize32
		if fSize > maxChunkSize32 {
			modulo = fSize % maxChunkSize32
			// chunks++
		} else {
			chunks = 1
		}
		givenChunkSize = maxChunkSize32
	}

	if chunks == 0 || fSize < 15 {
		log.Panicf("File smaller than it should be: %d bytes", fSize)
	}

	log.Printf("max-chunk: %t, chunk size: %d, raw file name: %s, chunks: %d, modulo: %d\n", *chunkSizeFlag, givenChunkSize, *fileName, chunks, modulo)

	rawBuffer = make([]byte, fSize)
	readed, err := file.Read(rawBuffer)
	checkError(err)
	log.Printf("bytes readed: %d\n", readed)

	copy(rawBuffer, []byte{0xFF, 0xFF, 0x0C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x3E, 0x10, 0x01})
	binary.LittleEndian.PutUint32(rawBuffer[6:], uint32(fSize-2))
	//0c 00 00 00 | fe ff ff ff | 80 3e 10 01
	csum = csum16(rawBuffer[2:])
	binary.BigEndian.PutUint16(rawBuffer, csum)
	//copy(rawBuffer, c)
	log.Printf("Formatted header: %#x, csum16 is: %#x\nSlicing...\n", rawBuffer[:14], csum)

	if modulo > 0 {
		chunks++
	}

	for chunkIter < chunks {
		var (
			bWrited int
			err     error
		)
		newFile := strings.TrimSuffix(path.Base(*fileName), path.Ext(*fileName)) + fmt.Sprintf("_p%.2d.wnb", chunkIter)
		oFile, err := os.OpenFile(newFile, os.O_RDWR|os.O_CREATE, 0644)
		checkError(err)

		if chunkIter+1 == chunks {
			bWrited, err = oFile.Write(rawBuffer[chunkIter*givenChunkSize:])
			checkError(err)
		} else {
			bWrited, err = oFile.Write(rawBuffer[chunkIter*givenChunkSize : (chunkIter+1)*givenChunkSize])
			checkError(err)
		}
		err = oFile.Close()
		checkError(err)
		chunkIter++

		log.Printf("File: %s, size: %d", newFile, bWrited)
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func csum16(data []byte) uint16 {
	var csum uint16

	for index := 0; index < len(data); index++ {
		csum = (csum + uint16(data[index]))
	}
	inverted := csum ^ 0xffff
	return (inverted >> 8) | (inverted<<8&0xff00>>8)<<8
}
