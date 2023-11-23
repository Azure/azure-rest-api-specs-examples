package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountGetObjectReplicationPolicy.json
func ExampleObjectReplicationPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewObjectReplicationPoliciesClient().Get(ctx, "res6977", "sto2527", "{objectReplicationPolicy-Id}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ObjectReplicationPolicy = armstorage.ObjectReplicationPolicy{
	// 	Name: to.Ptr("{objectReplicationPolicy-Id}"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/objectReplicationPolicies"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/objectReplicationPolicies/{objectReplicationPolicy-Id}"),
	// 	Properties: &armstorage.ObjectReplicationPolicyProperties{
	// 		DestinationAccount: to.Ptr("destAccount1"),
	// 		EnabledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-08T03:01:55.716Z"); return t}()),
	// 		PolicyID: to.Ptr("{objectReplicationPolicy-Id}"),
	// 		Rules: []*armstorage.ObjectReplicationPolicyRule{
	// 			{
	// 				DestinationContainer: to.Ptr("destContainer1"),
	// 				Filters: &armstorage.ObjectReplicationPolicyFilter{
	// 					PrefixMatch: []*string{
	// 						to.Ptr("blobA"),
	// 						to.Ptr("blobB")},
	// 					},
	// 					SourceContainer: to.Ptr("sourceContainer1"),
	// 				},
	// 				{
	// 					DestinationContainer: to.Ptr("destContainer1"),
	// 					Filters: &armstorage.ObjectReplicationPolicyFilter{
	// 						PrefixMatch: []*string{
	// 							to.Ptr("blobC"),
	// 							to.Ptr("blobD")},
	// 						},
	// 						SourceContainer: to.Ptr("sourceContainer1"),
	// 				}},
	// 				SourceAccount: to.Ptr("sto2527"),
	// 			},
	// 		}
}
