package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/AuthorizationPoliciesCreateOrUpdate.json
func ExampleAuthorizationPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcustomerinsights.NewAuthorizationPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"TestHubRG",
		"azSdkTestHub",
		"testPolicy4222",
		armcustomerinsights.AuthorizationPolicyResourceFormat{
			Properties: &armcustomerinsights.AuthorizationPolicy{
				Permissions: []*armcustomerinsights.PermissionTypes{
					to.Ptr(armcustomerinsights.PermissionTypesRead),
					to.Ptr(armcustomerinsights.PermissionTypesWrite),
					to.Ptr(armcustomerinsights.PermissionTypesManage)},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
