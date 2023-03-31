package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementPortalSettingsUpdateDelegation.json
func ExampleDelegationSettingsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDelegationSettingsClient().Update(ctx, "rg1", "apimService1", "*", armapimanagement.PortalDelegationSettings{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
