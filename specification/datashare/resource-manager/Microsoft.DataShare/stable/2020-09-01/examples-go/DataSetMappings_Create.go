package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_Create.json
func ExampleDataSetMappingsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatashare.NewDataSetMappingsClient("433a8dfd-e5d5-4e77-ad86-90acdc75eb1a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"SampleResourceGroup",
		"Account1",
		"ShareSubscription1",
		"DatasetMapping1",
		&armdatashare.BlobDataSetMapping{
			Kind: to.Ptr(armdatashare.DataSetMappingKindBlob),
			Properties: &armdatashare.BlobMappingProperties{
				ContainerName:      to.Ptr("C1"),
				DataSetID:          to.Ptr("a08f184b-0567-4b11-ba22-a1199336d226"),
				FilePath:           to.Ptr("file21"),
				ResourceGroup:      to.Ptr("SampleResourceGroup"),
				StorageAccountName: to.Ptr("storage2"),
				SubscriptionID:     to.Ptr("433a8dfd-e5d5-4e77-ad86-90acdc75eb1a"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
