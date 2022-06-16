package armaad_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
)

// x-ms-original-file: specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/preview/2017-04-01-preview/examples/getDiagnosticSetting.json
func ExampleDiagnosticSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armaad.NewDiagnosticSettingsClient(cred, nil)
	res, err := client.Get(ctx,
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DiagnosticSettingsClientGetResult)
}
