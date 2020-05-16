package dns

import (
	"NigolasDNS/database"
	"context"
	"encoding/base64"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	A_RECORD_TYPE = 1
	NS_RECORD_TYPE = 2
	CNAME_RECORD_TYPE = 5
	MX_RECORD_TYPE = 15
	AAAA_RECORD_TYPE = 128
)

type Setting struct {
	Name string
	Value string
}

type DNSRecordDB struct {
	Domain_Name string
	Type uint16
	TTL uint32
	Data string
}

func parseRecord(record DNSRecordDB) DNSRecord {
	data, _ := base64.StdEncoding.DecodeString(record.Data)
		return DNSRecord{
			record.Domain_Name,
			record.Type,
			1,
			31337,
			4,
			data,
		}
}

func DNSLookup(query DNSRecord) ([]DNSRecord, []DNSRecord, []DNSRecord)  {
	answerRecords := make([]DNSRecord, 0)
	authorityRecords := make([]DNSRecord, 0)
	additionalRecords := make([]DNSRecord, 0)
	var record DNSRecordDB
	_ = database.Client.Database("nigolas_dns").Collection("records").FindOne(context.TODO(), bson.M{"Domain_Name": query.DOMAIN_NAME, "Type": query.TYPE}).Decode(&record)
	answerRecords =
		append(answerRecords, parseRecord(record))

	return answerRecords, authorityRecords, additionalRecords
}