package armmanagementgroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups/v2"
)

// Generated from example definition: 2023-04-01/GetAllSubscriptionsFromManagementGroup.json
func ExampleManagementGroupSubscriptionsClient_NewGetSubscriptionsUnderManagementGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagementGroupSubscriptionsClient().NewGetSubscriptionsUnderManagementGroupPager("Group", nil)
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
		// page = armmanagementgroups.ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse{
		// 	ListSubscriptionUnderManagementGroup: armmanagementgroups.ListSubscriptionUnderManagementGroup{
		// 		Value: []*armmanagementgroups.SubscriptionUnderManagementGroup{
		// 			{
		// 				Name: to.Ptr("728bcbe4-8d56-4510-86c2-4921b8beefbc"),
		// 				Type: to.Ptr("Microsoft.Management/managementGroups/subscriptions"),
		// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/Group/subscriptions/728bcbe4-8d56-4510-86c2-4921b8beefbc"),
		// 				Properties: &armmanagementgroups.SubscriptionUnderManagementGroupProperties{
		// 					DisplayName: to.Ptr("S5"),
		// 					Parent: &armmanagementgroups.DescendantParentGroupInfo{
		// 						ID: to.Ptr("/providers/Microsoft.Management/managementGroups/Group"),
		// 					},
		// 					State: to.Ptr("Active"),
		// 					Tenant: to.Ptr("e751ac82-623b-4913-8d74-22637c832373"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
