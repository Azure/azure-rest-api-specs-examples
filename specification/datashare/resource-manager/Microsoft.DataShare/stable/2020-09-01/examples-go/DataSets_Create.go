package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSets_Create.json
func ExampleDataSetsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatashare.NewDataSetsClient("433a8dfd-e5d5-4e77-ad86-90acdc75eb1a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"SampleResourceGroup",
		"Account1",
		"Share1",
		"Dataset1",
		&armdatashare.BlobDataSet{
			Kind: to.Ptr(armdatashare.DataSetKindBlob),
			Properties: &armdatashare.BlobProperties{
				ContainerName:      to.Ptr("C1"),
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
