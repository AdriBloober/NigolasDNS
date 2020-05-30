## NigolasDNS

I wanted to learn GoLang and DNS Mangement. So i have developed by own minimal dns server.  
**DO NOT USE IN PRODUCTION**
This is a development project, so do not use it in production environment!

### Installation

``go build`` for building the server  
Configure the db in database/database.go and configure the port in the main.go

It works with the mongo database.

### Configure Records

In the collections records on the database "nigolas_dns" you found all record documents. Create it with this format:  
1. Domain_Name: The name of the domain (string)
2. Type: The Type of the domain -> https://en.wikipedia.org/wiki/List_of_DNS_record_types (int)
3. TTL (int)
4. Data (string -> base64 string of the data)

For A is the data for example 4 bytes (b.b.b.b)

### Holy Fuck: Why NigolasDNS?

Yes, it's a good question: Do you know Nicolas fucking Janzen? No? Ok, he is the owner of the very very good german hosting
company [prohosting24](https://prohosting24.de). Please rent servers there! Nicolas Janzen needs money for his nice
hosting project. So support him!

### Author

My name is AdriBloober, it is my first golang project and my first contact with dns. So don't expect too much!
