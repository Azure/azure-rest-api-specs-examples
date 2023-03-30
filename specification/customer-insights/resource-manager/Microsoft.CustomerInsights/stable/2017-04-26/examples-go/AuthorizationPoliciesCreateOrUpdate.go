package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/AuthorizationPoliciesCreateOrUpdate.json
func ExampleAuthorizationPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationPoliciesClient().CreateOrUpdate(ctx, "TestHubRG", "azSdkTestHub", "testPolicy4222", armcustomerinsights.AuthorizationPolicyResourceFormat{
		Properties: &armcustomerinsights.AuthorizationPolicy{
			Permissions: []*armcustomerinsights.PermissionTypes{
				to.Ptr(armcustomerinsights.PermissionTypesRead),
				to.Ptr(armcustomerinsights.PermissionTypesWrite),
				to.Ptr(armcustomerinsights.PermissionTypesManage)},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationPolicyResourceFormat = armcustomerinsights.AuthorizationPolicyResourceFormat{
	// 	Name: to.Ptr("testPolicy4222"),
	// 	Type: to.Ptr("Microsoft.CustomerInsights/hubs/AuthorizationPolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups//TestHubRG/providers/Microsoft.CustomerInsights/hubs/azSdkTestHub/AuthorizationPolicies/testPolicy4222"),
	// 	Properties: &armcustomerinsights.AuthorizationPolicy{
	// 		Permissions: []*armcustomerinsights.PermissionTypes{
	// 			to.Ptr(armcustomerinsights.PermissionTypesRead),
	// 			to.Ptr(armcustomerinsights.PermissionTypesWrite),
	// 			to.Ptr(armcustomerinsights.PermissionTypesManage)},
	// 			PolicyName: to.Ptr("testPolicy4222"),
	// 			PrimaryKey: to.Ptr("<primaryKey>"),
	// 			SecondaryKey: to.Ptr("<secondaryKey>"),
	// 		},
	// 	}
}
