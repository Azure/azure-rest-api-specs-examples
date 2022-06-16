package armserialconsole_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/serialconsole/armserialconsole"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/EnableConsoleExamples.json
func ExampleMicrosoftSerialConsoleClient_EnableConsole() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armserialconsole.NewMicrosoftSerialConsoleClient("00000000-00000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.EnableConsole(ctx,
		"default",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
