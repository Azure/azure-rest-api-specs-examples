package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/AuthorizationPoliciesListByHub.json
func ExampleAuthorizationPoliciesClient_NewListByHubPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAuthorizationPoliciesClient().NewListByHubPager("TestHubRG", "azSdkTestHub", nil)
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
		// page.AuthorizationPolicyListResult = armcustomerinsights.AuthorizationPolicyListResult{
		// 	Value: []*armcustomerinsights.AuthorizationPolicyResourceFormat{
		// 		{
		// 			Name: to.Ptr("azSdkTestHub/RootManageSharedAccessKey"),
		// 			Type: to.Ptr("Microsoft.CustomerInsights/hubs/AuthorizationPolicies"),
		// 			ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/azSdkTestHub/AuthorizationPolicies/RootManageSharedAccessKey"),
		// 			Properties: &armcustomerinsights.AuthorizationPolicy{
		// 				Permissions: []*armcustomerinsights.PermissionTypes{
		// 					to.Ptr(armcustomerinsights.PermissionTypesRead),
		// 					to.Ptr(armcustomerinsights.PermissionTypesWrite),
		// 					to.Ptr(armcustomerinsights.PermissionTypesManage)},
		// 					PolicyName: to.Ptr("RootManageSharedAccessKey"),
		// 					PrimaryKey: to.Ptr("<primaryKey>"),
		// 					SecondaryKey: to.Ptr("<secondaryKey>"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("azSdkTestHub/testPolicy1125"),
		// 				Type: to.Ptr("Microsoft.CustomerInsights/hubs/AuthorizationPolicies"),
		// 				ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/azSdkTestHub/AuthorizationPolicies/testPolicy1125"),
		// 				Properties: &armcustomerinsights.AuthorizationPolicy{
		// 					Permissions: []*armcustomerinsights.PermissionTypes{
		// 						to.Ptr(armcustomerinsights.PermissionTypesRead),
		// 						to.Ptr(armcustomerinsights.PermissionTypesWrite),
		// 						to.Ptr(armcustomerinsights.PermissionTypesManage)},
		// 						PolicyName: to.Ptr("testPolicy1125"),
		// 						PrimaryKey: to.Ptr("<primaryKey>"),
		// 						SecondaryKey: to.Ptr("<secondaryKey>"),
		// 					},
		// 			}},
		// 		}
	}
}
