package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86c6306649b02e542117adb46c61e8019dbd78e9/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/StorageAccountCreateObjectReplicationPolicyOnDestination.json
func ExampleObjectReplicationPoliciesClient_CreateOrUpdate_storageAccountCreateObjectReplicationPolicyOnDestination() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewObjectReplicationPoliciesClient().CreateOrUpdate(ctx, "res7687", "dst112", "default", armstorage.ObjectReplicationPolicy{
		Properties: &armstorage.ObjectReplicationPolicyProperties{
			DestinationAccount: to.Ptr("dst112"),
			Metrics: &armstorage.ObjectReplicationPolicyPropertiesMetrics{
				Enabled: to.Ptr(true),
			},
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ObjectReplicationPolicy = armstorage.ObjectReplicationPolicy{
	// 	Name: to.Ptr("2a20bb73-5717-4635-985a-5d4cf777438f"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/objectReplicationPolicies"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7687/providers/Microsoft.Storage/storageAccounts/dst112/objectReplicationPolicies/2a20bb73-5717-4635-985a-5d4cf777438f"),
	// 	Properties: &armstorage.ObjectReplicationPolicyProperties{
	// 		DestinationAccount: to.Ptr("dst112"),
	// 		Metrics: &armstorage.ObjectReplicationPolicyPropertiesMetrics{
	// 			Enabled: to.Ptr(true),
	// 		},
	// 		PolicyID: to.Ptr("2a20bb73-5717-4635-985a-5d4cf777438f"),
	// 		Rules: []*armstorage.ObjectReplicationPolicyRule{
	// 			{
	// 				DestinationContainer: to.Ptr("destContainer1"),
	// 				Filters: &armstorage.ObjectReplicationPolicyFilter{
	// 					PrefixMatch: []*string{
	// 						to.Ptr("blobA"),
	// 						to.Ptr("blobB")},
	// 					},
	// 					RuleID: to.Ptr("d5d18a48-8801-4554-aeaa-74faf65f5ef9"),
	// 					SourceContainer: to.Ptr("sourceContainer1"),
	// 			}},
	// 			SourceAccount: to.Ptr("src1122"),
	// 		},
	// 	}
}
