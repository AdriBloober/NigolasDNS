package dns

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

const (
	ResponseFlag uint16 = 1 << 15
)

type DNSRecord struct {
	DOMAIN_NAME string
	TYPE uint16
	CLASS uint16
	TTL uint32
	RESOURCE_DATA_LENGTH uint16
	RESOURCE_DATA []byte
}

type A_DNSRecord struct {
	record DNSRecord
	address [4]byte
}

type AAAA_DNSRecord struct {
	record DNSRecord
	address [16]byte
}

func readDNSName (requestBuffer *bytes.Buffer) (string, error) {
	var domainName string

	b, err := requestBuffer.ReadByte()

	for ; b != 0 && err == nil; b, err = requestBuffer.ReadByte() {
		labelLength := int(b)
		labelBytes := requestBuffer.Next(labelLength)
		labelName := string(labelBytes)

		if len(domainName) == 0 {
			domainName = labelName
		} else {
			domainName += "." + labelName
		}
	}

	return domainName, err
}

func writeDNSName (responseBuffer *bytes.Buffer, domainName string) error{
	labels := strings.Split(domainName, ".")
	for _, element := range labels {
		length := len(element)
		elementBytes := []byte(element)
		responseBuffer.WriteByte(byte(length))
		responseBuffer.Write(elementBytes)
	}

	err := responseBuffer.WriteByte(byte(0))
	return err
}

func ReadDNSRecord (requestBuffer *bytes.Buffer) DNSRecord{
	var record DNSRecord
	domainName, err := readDNSName(requestBuffer)
	if err != nil {fmt.Println("Error on reading domain name: ", err)}
	record.DOMAIN_NAME = domainName
	record.TYPE = binary.BigEndian.Uint16(requestBuffer.Next(2))
	record.CLASS = binary.BigEndian.Uint16(requestBuffer.Next(2))
	return record
}

func WriteDNSRecords (queryHeader DNSHeader, queryRecords []DNSRecord, answerRecords []DNSRecord,  authorityRecords []DNSRecord, additionalRecords []DNSRecord) ([]byte, error){
	var responseBuffer = new(bytes.Buffer)
	responseHeader := DNSHeader{
		ID: queryHeader.ID,
		FLAGS: ResponseFlag,
		QDCOUNT: queryHeader.QDCOUNT,
		ANCOUNT: uint16(len(answerRecords)),
		NSCOUNT: uint16(len(authorityRecords)),
		ADCOUNT: uint16(len(additionalRecords)),
	}
	err := binary.Write(responseBuffer, binary.BigEndian, &responseHeader)
	if err != nil {
		fmt.Println("Error on writing to response buffer:", err)
	}
	for _, element := range queryRecords {
		err = writeDNSName(responseBuffer, element.DOMAIN_NAME)
		if err != nil {fmt.Println("Error on query response dns name")}
		binary.Write(responseBuffer, binary.BigEndian, element.TYPE)
		binary.Write(responseBuffer, binary.BigEndian, element.CLASS)
	}
	for _, element := range answerRecords {
		err = writeDNSName(responseBuffer, element.DOMAIN_NAME)
		if err != nil {fmt.Println("Error on answer response dns name")}
		binary.Write(responseBuffer, binary.BigEndian, element.TYPE)
		binary.Write(responseBuffer, binary.BigEndian, element.CLASS)
		binary.Write(responseBuffer, binary.BigEndian, element.TTL)
		binary.Write(responseBuffer, binary.BigEndian, element.RESOURCE_DATA_LENGTH)
		binary.Write(responseBuffer, binary.BigEndian, element.RESOURCE_DATA)
	}
	for _, element := range authorityRecords {
		err = writeDNSName(responseBuffer, element.DOMAIN_NAME)
		if err != nil {fmt.Println("Error on authority response dns name")}
		binary.Write(responseBuffer, binary.BigEndian, element.TYPE)
		binary.Write(responseBuffer, binary.BigEndian, element.CLASS)
		binary.Write(responseBuffer, binary.BigEndian, element.TTL)
		binary.Write(responseBuffer, binary.BigEndian, element.RESOURCE_DATA_LENGTH)
		binary.Write(responseBuffer, binary.BigEndian, element.RESOURCE_DATA)
	}
	for _, element := range additionalRecords {
		err = writeDNSName(responseBuffer, element.DOMAIN_NAME)
		if err != nil {fmt.Println("Error on additional response dns name")}
		binary.Write(responseBuffer, binary.BigEndian, element.TYPE)
		binary.Write(responseBuffer, binary.BigEndian, element.CLASS)
		binary.Write(responseBuffer, binary.BigEndian, element.TTL)
		binary.Write(responseBuffer, binary.BigEndian, element.RESOURCE_DATA_LENGTH)
		binary.Write(responseBuffer, binary.BigEndian, element.RESOURCE_DATA)
	}
	return responseBuffer.Bytes(), err
}