package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/SchemaRegistry/SchemaRegistryCreate.json
func ExampleSchemaRegistryClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeventhub.NewSchemaRegistryClient("e8baea74-64ce-459b-bee3-5aa4c47b3ae3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"alitest",
		"ali-ua-test-eh-system-1",
		"testSchemaGroup1",
		armeventhub.SchemaGroup{
			Properties: &armeventhub.SchemaGroupProperties{
				GroupProperties:     map[string]*string{},
				SchemaCompatibility: to.Ptr(armeventhub.SchemaCompatibilityForward),
				SchemaType:          to.Ptr(armeventhub.SchemaTypeAvro),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
