package utils

import (
	"os"
	"strings"
	"testing"
)

func TestWriteToEnvFile(t *testing.T) {
	var (
		filePath string
		content  []byte
		err      error
	)

	envVars := getEnvArray()
	if filePath, err = WriteToEnvFile(envVars); err != nil {
		t.Fatalf("Getting an error when writting file. Err %s", err.Error())
	}

	if content, err = os.ReadFile(filePath); err != nil || content == nil {
		t.Fatalf("Env file was not created. Err: ")
	}

	if strings.Join(envVars, "\n") != string(content) {
		t.Error("Result not properly formatted")
	}
	if err = os.Remove(filePath); err != nil {
		t.Error("Can't remove test file")
	}
}

func getEnvArray() []string {
	return []string{
		"INTERFACE=eth0",
		"DHCP_START=2",
		"DNSMASQ_LISTENING=1",
		"DHCP_END=false",
		"DHCP_ROUTER=1.2.3.4",
		"ServerIP=1.2.3.4",
		"WEBPASSWORD=single",
		"DHCP_ACTIVE=single",
		"FOO_VAR=true",
		"TZ=America/Argentina/Buenos_Aires",
		"phpver=single",
		"PIHOLE_DOCKER_TAG=2022.01.1",
		"PIHOLE_INSTALL=/etc/.pihole/automated install/basic-install.sh",
		"SOME_JSON={\n    \"glossary\": {\n        \"title\": \"example glossary\",\n\t\t\"GlossDiv\": {\n            \"title\": \"S\",\n\t\t\t\"GlossList\": {\n                \"GlossEntry\": {\n                    \"ID\": \"SGML\",\n\t\t\t\t\t\"SortAs\": \"SGML\",\n\t\t\t\t\t\"GlossTerm\": \"Standard Generalized Markup Language\",\n\t\t\t\t\t\"Acronym\": \"SGML\",\n\t\t\t\t\t\"Abbrev\": \"ISO 8879:1986\",\n\t\t\t\t\t\"GlossDef\": {\n                        \"para\": \"A meta-markup language, used to create markup languages such as DocBook.\",\n\t\t\t\t\t\t\"GlossSeeAlso\": [\"GML\", \"XML\"]\n                    },\n\t\t\t\t\t\"GlossSee\": \"markup\"\n                }\n            }\n        }\n    }\n}",
	}
}
