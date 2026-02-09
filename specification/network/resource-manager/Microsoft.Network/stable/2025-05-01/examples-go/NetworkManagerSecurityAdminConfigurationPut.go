package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerSecurityAdminConfigurationPut.json
func ExampleSecurityAdminConfigurationsClient_CreateOrUpdate_createNetworkManagerSecurityAdminConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecurityAdminConfigurationsClient().CreateOrUpdate(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", armnetwork.SecurityAdminConfiguration{
		Properties: &armnetwork.SecurityAdminConfigurationPropertiesFormat{
			Description: to.Ptr("A sample policy"),
			ApplyOnNetworkIntentPolicyBasedServices: []*armnetwork.NetworkIntentPolicyBasedService{
				to.Ptr(armnetwork.NetworkIntentPolicyBasedServiceNone)},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SecurityAdminConfiguration = armnetwork.SecurityAdminConfiguration{
	// 	Name: to.Ptr("myTestSecurityConfig"),
	// 	Type: to.Ptr("Microsoft.Network/networkManagers/securityAdminConfigurations"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManager/testNetworkManager/securityAdminConfigurations/myTestSecurityConfig"),
	// 	Properties: &armnetwork.SecurityAdminConfigurationPropertiesFormat{
	// 		Description: to.Ptr("A sample policy"),
	// 		ApplyOnNetworkIntentPolicyBasedServices: []*armnetwork.NetworkIntentPolicyBasedService{
	// 			to.Ptr(armnetwork.NetworkIntentPolicyBasedServiceNone)},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		SystemData: &armnetwork.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 			CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		},
	// 	}
}
