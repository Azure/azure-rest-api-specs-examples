package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/220ad9c6554fc7d6d10a89bdb441c1e3b36e3285/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountListObjectReplicationPolicies.json
func ExampleObjectReplicationPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewObjectReplicationPoliciesClient().NewListPager("res6977", "sto2527", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ObjectReplicationPolicies = armstorage.ObjectReplicationPolicies{
		// 	Value: []*armstorage.ObjectReplicationPolicy{
		// 		{
		// 			Name: to.Ptr("c6c23999-fd4e-433a-bcf9-1db69d27cd8a"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/objectReplicationPolicies"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/objectReplicationPolicies/c6c23999-fd4e-433a-bcf9-1db69d27cd8a"),
		// 			Properties: &armstorage.ObjectReplicationPolicyProperties{
		// 				DestinationAccount: to.Ptr("destAccount1"),
		// 				SourceAccount: to.Ptr("sto2527"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("141d23dc-8958-4b48-b6e6-5a40bf1af116"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/objectReplicationPolicies"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/objectReplicationPolicies/141d23dc-8958-4b48-b6e6-5a40bf1af116"),
		// 			Properties: &armstorage.ObjectReplicationPolicyProperties{
		// 				DestinationAccount: to.Ptr("destAccount2"),
		// 				SourceAccount: to.Ptr("sto2527"),
		// 			},
		// 	}},
		// }
	}
}
