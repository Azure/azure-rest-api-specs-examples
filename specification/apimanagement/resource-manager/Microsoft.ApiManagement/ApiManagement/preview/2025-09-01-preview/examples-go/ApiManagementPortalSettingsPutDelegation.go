package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementPortalSettingsPutDelegation.json
func ExampleDelegationSettingsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDelegationSettingsClient().CreateOrUpdate(ctx, "rg1", "apimService1", armapimanagement.PortalDelegationSettings{
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
	}, &armapimanagement.DelegationSettingsClientCreateOrUpdateOptions{
		IfMatch: to.Ptr("*")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.DelegationSettingsClientCreateOrUpdateResponse{
	// 	PortalDelegationSettings: armapimanagement.PortalDelegationSettings{
	// 		Name: to.Ptr("delegation"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/portalsettings"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/portalsettings/delegation"),
	// 		Properties: &armapimanagement.PortalDelegationSettingsProperties{
	// 			Subscriptions: &armapimanagement.SubscriptionsDelegationSettingsProperties{
	// 				Enabled: to.Ptr(true),
	// 			},
	// 			URL: to.Ptr("http://contoso.com/delegation"),
	// 			UserRegistration: &armapimanagement.RegistrationDelegationSettingsProperties{
	// 				Enabled: to.Ptr(true),
	// 			},
	// 		},
	// 	},
	// }
}
