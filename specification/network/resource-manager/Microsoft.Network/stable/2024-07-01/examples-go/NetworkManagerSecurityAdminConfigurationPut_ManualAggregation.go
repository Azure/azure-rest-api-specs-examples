package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkManagerSecurityAdminConfigurationPut_ManualAggregation.json
func ExampleSecurityAdminConfigurationsClient_CreateOrUpdate_createManualModeSecurityAdminConfiguration() {
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
			Description: to.Ptr("A configuration which will update any network groups ip addresses at commit times."),
			NetworkGroupAddressSpaceAggregationOption: to.Ptr(armnetwork.AddressSpaceAggregationOptionManual),
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
	// 	ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg1/providers/Microsoft.Network/networkManager/testNetworkManager/securityAdminConfigurations/myTestSecurityConfig"),
	// 	Properties: &armnetwork.SecurityAdminConfigurationPropertiesFormat{
	// 		Description: to.Ptr("A sample policy"),
	// 		NetworkGroupAddressSpaceAggregationOption: to.Ptr(armnetwork.AddressSpaceAggregationOptionManual),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// 	SystemData: &armnetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 		CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 		LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 	},
	// }
}
