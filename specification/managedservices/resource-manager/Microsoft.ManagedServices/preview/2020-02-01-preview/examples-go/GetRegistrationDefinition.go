package armmanagedservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
)

// x-ms-original-file: specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2020-02-01-preview/examples/GetRegistrationDefinition.json
func ExampleRegistrationDefinitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagedservices.NewRegistrationDefinitionsClient(cred, nil)
	res, err := client.Get(ctx,
		"<scope>",
		"<registration-definition-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RegistrationDefinitionsClientGetResult)
}
