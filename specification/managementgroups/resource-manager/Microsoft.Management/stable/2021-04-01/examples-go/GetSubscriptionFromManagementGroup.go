package armmanagementgroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetSubscriptionFromManagementGroup.json
func ExampleManagementGroupSubscriptionsClient_GetSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementGroupSubscriptionsClient().GetSubscription(ctx, "Group", "728bcbe4-8d56-4510-86c2-4921b8beefbc", &armmanagementgroups.ManagementGroupSubscriptionsClientGetSubscriptionOptions{CacheControl: to.Ptr("no-cache")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SubscriptionUnderManagementGroup = armmanagementgroups.SubscriptionUnderManagementGroup{
	// 	Name: to.Ptr("728bcbe4-8d56-4510-86c2-4921b8beefbc"),
	// 	Type: to.Ptr("Microsoft.Management/managementGroups/subscriptions"),
	// 	ID: to.Ptr(" /providers/Microsoft.Management/managementGroups/Group/subscriptions/728bcbe4-8d56-4510-86c2-4921b8beefbc"),
	// 	Properties: &armmanagementgroups.SubscriptionUnderManagementGroupProperties{
	// 		DisplayName: to.Ptr("Group"),
	// 		Parent: &armmanagementgroups.DescendantParentGroupInfo{
	// 			ID: to.Ptr("/providers/Microsoft.Management/managementGroups/Group"),
	// 		},
	// 		State: to.Ptr("Active"),
	// 		Tenant: to.Ptr("e751ac82-623b-4913-8d74-22637c832373"),
	// 	},
	// }
}
