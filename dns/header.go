package dns

type DNSHeader struct {
	ID uint16
	FLAGS uint16
	QDCOUNT uint16
	ANCOUNT uint16
	NSCOUNT uint16
	ADCOUNT uint16
}
