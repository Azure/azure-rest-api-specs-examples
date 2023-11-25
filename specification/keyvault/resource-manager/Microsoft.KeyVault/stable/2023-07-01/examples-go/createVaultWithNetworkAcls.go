package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/createVaultWithNetworkAcls.json
func ExampleVaultsClient_BeginCreateOrUpdate_createOrUpdateAVaultWithNetworkAcls() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVaultsClient().BeginCreateOrUpdate(ctx, "sample-resource-group", "sample-vault", armkeyvault.VaultCreateOrUpdateParameters{
		Location: to.Ptr("westus"),
		Properties: &armkeyvault.VaultProperties{
			EnabledForDeployment:         to.Ptr(true),
			EnabledForDiskEncryption:     to.Ptr(true),
			EnabledForTemplateDeployment: to.Ptr(true),
			NetworkACLs: &armkeyvault.NetworkRuleSet{
				Bypass:        to.Ptr(armkeyvault.NetworkRuleBypassOptionsAzureServices),
				DefaultAction: to.Ptr(armkeyvault.NetworkRuleActionDeny),
				IPRules: []*armkeyvault.IPRule{
					{
						Value: to.Ptr("124.56.78.91"),
					},
					{
						Value: to.Ptr("'10.91.4.0/24'"),
					}},
				VirtualNetworkRules: []*armkeyvault.VirtualNetworkRule{
					{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1"),
					}},
			},
			SKU: &armkeyvault.SKU{
				Name:   to.Ptr(armkeyvault.SKUNameStandard),
				Family: to.Ptr(armkeyvault.SKUFamilyA),
			},
			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Vault = armkeyvault.Vault{
	// 	Name: to.Ptr("sample-vault"),
	// 	Type: to.Ptr("Microsoft.KeyVault/vaults"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-resource-group/providers/Microsoft.KeyVault/vaults/sample-vault"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armkeyvault.VaultProperties{
	// 		EnabledForDeployment: to.Ptr(true),
	// 		EnabledForDiskEncryption: to.Ptr(true),
	// 		EnabledForTemplateDeployment: to.Ptr(true),
	// 		HsmPoolResourceID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		NetworkACLs: &armkeyvault.NetworkRuleSet{
	// 			Bypass: to.Ptr(armkeyvault.NetworkRuleBypassOptionsAzureServices),
	// 			DefaultAction: to.Ptr(armkeyvault.NetworkRuleActionDeny),
	// 			IPRules: []*armkeyvault.IPRule{
	// 				{
	// 					Value: to.Ptr("124.56.78.91/32"),
	// 				},
	// 				{
	// 					Value: to.Ptr("'10.91.4.0/24'"),
	// 			}},
	// 			VirtualNetworkRules: []*armkeyvault.VirtualNetworkRule{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.network/virtualnetworks/test-vnet/subnets/subnet1"),
	// 			}},
	// 		},
	// 		SKU: &armkeyvault.SKU{
	// 			Name: to.Ptr(armkeyvault.SKUNameStandard),
	// 			Family: to.Ptr(armkeyvault.SKUFamilyA),
	// 		},
	// 		TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		VaultURI: to.Ptr("https://sample-vault.vault.azure.net"),
	// 	},
	// 	SystemData: &armkeyvault.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:00:00.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("keyVaultUser1"),
	// 		CreatedByType: to.Ptr(armkeyvault.IdentityTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:00:00.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("keyVaultUser2"),
	// 		LastModifiedByType: to.Ptr(armkeyvault.IdentityTypeUser),
	// 	},
	// 	Tags: map[string]*string{
	// 	},
	// }
}
