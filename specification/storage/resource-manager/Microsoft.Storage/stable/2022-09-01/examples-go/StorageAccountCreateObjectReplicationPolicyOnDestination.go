package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountCreateObjectReplicationPolicyOnDestination.json
func ExampleObjectReplicationPoliciesClient_CreateOrUpdate_storageAccountCreateObjectReplicationPolicyOnDestination() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewObjectReplicationPoliciesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "res7687", "dst112", "default", armstorage.ObjectReplicationPolicy{
		Properties: &armstorage.ObjectReplicationPolicyProperties{
			DestinationAccount: to.Ptr("dst112"),
			Rules: []*armstorage.ObjectReplicationPolicyRule{
				{
					DestinationContainer: to.Ptr("dcont139"),
					Filters: &armstorage.ObjectReplicationPolicyFilter{
						PrefixMatch: []*string{
							to.Ptr("blobA"),
							to.Ptr("blobB")},
					},
					SourceContainer: to.Ptr("scont139"),
				}},
			SourceAccount: to.Ptr("src1122"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
