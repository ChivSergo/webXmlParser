package xmlforparse

import "encoding/xml"

type MProcessor struct {
	Name string `xml:"name,attr"`
	File string `xml:"file,attr"`
	Clas string `xml:"class,attr"`
}

type MServer struct {
	Name string `xml:"name,attr"`
	File string `xml:"file,attr"`
	Clas string `xml:"class,attr"`
}

type MComponents struct {
	XMLName xml.Name `xml:"MComponents"`
	//Manager    MManager     `xml:"MManager"`
	Servers    []MServer    `xml:"MServers>MServer"`
	Processors []MProcessor `xml:"MProcessors>MProcessor"`
}
