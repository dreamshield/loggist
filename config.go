package loggist

import (
	"encoding/xml"
	"io/ioutil"
)

// log configuration string
const congString = `
<?xml version="1.0" encoding="UTF-8"?>
<Loggist version="1.0">
    <LogRoot>./Logs</LogRoot>
    <LogFormat>
        <Mode fileFmt="%04d%02d" subLogDir="%04d"/>
        <Mode fileFmt="%04d%02d%02d" subLogDir="%04d%02d"/>
        <Mode fileFmt="%04d%02d%02d-%02d" subLogDir="%04d%02d%02d"/>
        <Mode fileFmt="%04d%02d%02d-%02d%02d" subLogDir="%04d%02d%02d-%02d"/>
    </LogFormat>
</Loggist>
`

// log configuration
type LogConf struct {
	XMLName   xml.Name  `xml:"Loggist"`
	Version   string    `xml:"version,attr"`
	LogRoot   string    `xml:"LogRoot"`
	LogFormat LogFormat `xml:"LogFormat"`
}

type LogFormat struct {
	XMLName xml.Name     `xml:"LogFormat"`
	Mode    []LogFmtMode `xml:"Mode"`
}

type LogFmtMode struct {
	XMLName   xml.Name `xml:"Mode"`
	FileFmt   string   `xml:"fileFmt,attr"`
	SubLogDir string   `xml:"subLogDir,attr"`
}

// read log configuration file
func (self *LogConf) initLogConf() {
	data, err := ioutil.ReadFile("./LoggistConf/conf.xml")
	if err != nil {
		data = []byte(congString)
	}
	if err = xml.Unmarshal(data, self); err != nil {
		panic(err)
	}
}
