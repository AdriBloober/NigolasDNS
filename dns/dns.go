package dns

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func HandleDNSRequest(requestBytes []byte, server *net.UDPConn, clientAddr *net.UDPAddr)  {
	requestBuffer := bytes.NewBuffer(requestBytes)

	var header DNSHeader

	err := binary.Read(requestBuffer, binary.BigEndian, &header)
	if err != nil {
		println("Error on header encoding:", err)
	}

	var questions = make([]DNSRecord, header.QDCOUNT)
	for index, _ := range questions {
		questions[index] = ReadDNSRecord(requestBuffer)
	}
	var answerResourceRecords = make([]DNSRecord, 0)
	var authorityResourceRecords = make([]DNSRecord, 0)
	var additionalResourceRecords = make([]DNSRecord, 0)
	for _, element := range questions {
		answerRecords, authorityRecords, additionalRecords := DNSLookup(element)
		answerResourceRecords = append(answerResourceRecords, answerRecords...)
		authorityResourceRecords = append(authorityResourceRecords, authorityRecords...)
		additionalResourceRecords = append(additionalResourceRecords, additionalRecords...)

	}
	data, err := WriteDNSRecords(header, questions, answerResourceRecords, authorityResourceRecords, additionalResourceRecords)
	if err != nil {
		fmt.Println("Error on writing dns records (calculation)")
	}
	server.WriteToUDP(data, clientAddr)

}
