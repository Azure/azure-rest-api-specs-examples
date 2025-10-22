package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a9dbb28e788355a47dc5bad3ea5f8da212b4bf6/specification/network/resource-manager/Microsoft.Network/stable/2024-10-01/examples/NetworkManagerSecurityUserRuleCollectionPut.json
func ExampleSecurityUserRuleCollectionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecurityUserRuleCollectionsClient().CreateOrUpdate(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", armnetwork.SecurityUserRuleCollection{
		Properties: &armnetwork.SecurityUserRuleCollectionPropertiesFormat{
			Description: to.Ptr("A sample policy"),
			AppliesToGroups: []*armnetwork.SecurityUserGroupItem{
				{
					NetworkGroupID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/testGroup"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SecurityUserRuleCollection = armnetwork.SecurityUserRuleCollection{
	// 	Name: to.Ptr("myTestSecurityConfig"),
	// 	Type: to.Ptr("Microsoft.Network/networkManagers/securityUserConfigurations/ruleCollections"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkManager/testNetworkManager/securityUserConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection"),
	// 	Properties: &armnetwork.SecurityUserRuleCollectionPropertiesFormat{
	// 		Description: to.Ptr("A sample policy"),
	// 		AppliesToGroups: []*armnetwork.SecurityUserGroupItem{
	// 			{
	// 				NetworkGroupID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/testGroup"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("4224a5e7-d273-43d4-9bb0-b38e9b937ded"),
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
