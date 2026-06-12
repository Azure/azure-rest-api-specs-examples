package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkManagerSecurityUserConfigurationPut.json
func ExampleSecurityUserConfigurationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecurityUserConfigurationsClient().CreateOrUpdate(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", armnetwork.SecurityUserConfiguration{
		Properties: &armnetwork.SecurityUserConfigurationPropertiesFormat{
			Description: to.Ptr("A sample policy"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.SecurityUserConfigurationsClientCreateOrUpdateResponse{
	// 	SecurityUserConfiguration: armnetwork.SecurityUserConfiguration{
	// 		Name: to.Ptr("myTestSecurityConfig"),
	// 		Type: to.Ptr("Microsoft.Network/networkManagers/securityConfigurations"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManager/testNetworkManager/securityConfigurations/myTestSecurityConfig"),
	// 		Properties: &armnetwork.SecurityUserConfigurationPropertiesFormat{
	// 			Description: to.Ptr("A sample policy"),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("e74045ed-c817-48df-adb2-a06753ad4fff"),
	// 		},
	// 		SystemData: &armnetwork.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
	// 			CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
