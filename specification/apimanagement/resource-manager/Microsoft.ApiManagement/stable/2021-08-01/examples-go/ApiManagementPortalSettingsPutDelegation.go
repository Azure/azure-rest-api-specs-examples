package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementPortalSettingsPutDelegation.json
func ExampleDelegationSettingsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewDelegationSettingsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg1",
		"apimService1",
		armapimanagement.PortalDelegationSettings{
			Properties: &armapimanagement.PortalDelegationSettingsProperties{
				Subscriptions: &armapimanagement.SubscriptionsDelegationSettingsProperties{
					Enabled: to.Ptr(true),
				},
				URL: to.Ptr("http://contoso.com/delegation"),
				UserRegistration: &armapimanagement.RegistrationDelegationSettingsProperties{
					Enabled: to.Ptr(true),
				},
				ValidationKey: to.Ptr("<validationKey>"),
			},
		},
		&armapimanagement.DelegationSettingsClientCreateOrUpdateOptions{IfMatch: to.Ptr("*")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
