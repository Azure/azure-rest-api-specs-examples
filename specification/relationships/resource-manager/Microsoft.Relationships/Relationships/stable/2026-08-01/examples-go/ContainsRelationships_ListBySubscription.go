package armrelationships_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relationships/armrelationships"
)

// Generated from example definition: 2026-08-01/ContainsRelationships_ListBySubscription.json
func ExampleContainsRelationshipsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelationships.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainsRelationshipsClient("a925f2f7-5c63-4b7b-8799-25a5f97bc3b2").NewListBySubscriptionPager(nil)
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
		// page = armrelationships.ContainsRelationshipsClientListBySubscriptionResponse{
		// 	ContainsRelationshipListResult: armrelationships.ContainsRelationshipListResult{
		// 		Value: []*armrelationships.ContainsRelationship{
		// 			{
		// 				ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Relationships/contains/contains1"),
		// 				Type: to.Ptr("Microsoft.Relationships/contains"),
		// 				Properties: &armrelationships.ContainsRelationshipProperties{
		// 					SourceID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
		// 					TargetID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/test-vm"),
		// 					TargetTenant: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 					Metadata: &armrelationships.RelationshipMetadata{
		// 						SourceType: to.Ptr("Microsoft.Resources/subscriptions"),
		// 						TargetType: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 					},
		// 					ProvisioningState: to.Ptr(armrelationships.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Relationships/contains/contains2"),
		// 				Type: to.Ptr("Microsoft.Relationships/contains"),
		// 				Properties: &armrelationships.ContainsRelationshipProperties{
		// 					SourceID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
		// 					TargetID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.Sql/servers/sql1"),
		// 					TargetTenant: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 					Metadata: &armrelationships.RelationshipMetadata{
		// 						SourceType: to.Ptr("Microsoft.Resources/subscriptions"),
		// 						TargetType: to.Ptr("Microsoft.Sql/servers"),
		// 					},
		// 					ProvisioningState: to.Ptr(armrelationships.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Relationships/contains?api-version=2026-08-01&$skipToken=abc123"),
		// 	},
		// }
	}
}
