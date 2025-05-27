package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/ApplicationSecurityGroupListAll.json
func ExampleApplicationSecurityGroupsClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationSecurityGroupsClient().NewListAllPager(nil)
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
		// page.ApplicationSecurityGroupListResult = armnetwork.ApplicationSecurityGroupListResult{
		// 	Value: []*armnetwork.ApplicationSecurityGroup{
		// 		{
		// 			Name: to.Ptr("asg1"),
		// 			Type: to.Ptr("Microsoft.Network/applicationSecurityGroups"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationSecurityGroups/asg1"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.ApplicationSecurityGroupPropertiesFormat{
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("asg2"),
		// 			Type: to.Ptr("Microsoft.Network/applicationSecurityGroups"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationSecurityGroups/asg2"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.ApplicationSecurityGroupPropertiesFormat{
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			},
		// 	}},
		// }
	}
}
