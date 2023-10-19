package main

import (
	"encoding/binary"
	"fmt"
)

func decodePacket(packet []byte) struct {
	Short1      int
	Characters1 string
	SingleByte  int
	Characters2 string
	Short2      int
	Characters3 string
	Long        int
} {
	decodedStruct := struct {
		Short1      int
		Characters1 string
		SingleByte  int
		Characters2 string
		Short2      int
		Characters3 string
		Long        int
	}{}

	if len(packet) != 44 {
		fmt.Println("Invalid packet size. Expected 44 bytes, got", len(packet), "bytes")
		return decodedStruct
	}

	decodedStruct.Short1 = int(binary.BigEndian.Uint16(packet[0:2]))
	decodedStruct.Characters1 = string(packet[2:14])
	decodedStruct.SingleByte = int(packet[14])
	decodedStruct.Characters2 = string(packet[15:23])
	decodedStruct.Short2 = int(binary.BigEndian.Uint16(packet[23:25]))
	decodedStruct.Characters3 = string(packet[25:40])
	decodedStruct.Long = int(binary.BigEndian.Uint32(packet[40:44]))

	return decodedStruct
}

func main() {
	packet := []byte{
		0x04, 0xD2, 0x6B, 0x65, 0x65, 0x70, 0x64, 0x65, 0x63, 0x6F, 0x64, 0x69, 0x6E, 0x67, 0x38, 0x64,
		0x6F, 0x6E, 0x74, 0x73, 0x74, 0x6F, 0x70, 0x03, 0x15, 0x63, 0x6F, 0x6E, 0x67, 0x72, 0x61, 0x74,
		0x75, 0x6C, 0x61, 0x74, 0x69, 0x6F, 0x6E, 0x73, 0x07, 0x5B, 0xCD, 0x15,
	}

	decodedStruct := decodePacket(packet)
	fmt.Printf("Decoded struct: {%d, \"%s\", %d, \"%s\", %d, \"%s\", , %d}\n",
		decodedStruct.Short1, decodedStruct.Characters1,
		decodedStruct.SingleByte, decodedStruct.Characters2,
		decodedStruct.Short2, decodedStruct.Characters3,
		decodedStruct.Long)
}
