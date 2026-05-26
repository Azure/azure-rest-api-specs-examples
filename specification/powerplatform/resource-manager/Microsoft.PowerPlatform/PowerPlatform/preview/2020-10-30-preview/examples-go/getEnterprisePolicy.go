package armpowerplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform"
)

// Generated from example definition: 2020-10-30-preview/getEnterprisePolicy.json
func ExampleEnterprisePoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerplatform.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnterprisePoliciesClient().Get(ctx, "enterprisePolicy", "rg", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpowerplatform.EnterprisePoliciesClientGetResponse{
	// 	EnterprisePolicy: armpowerplatform.EnterprisePolicy{
	// 		Name: to.Ptr("enterprisePolicy"),
	// 		Type: to.Ptr("Microsoft.PowerPlatform/enterprisePolicies"),
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.PowerPlatform/enterprisePolicies/enterprisePolicy"),
	// 		Identity: &armpowerplatform.EnterprisePolicyIdentity{
	// 			Type: to.Ptr(armpowerplatform.ResourceIdentityTypeSystemAssigned),
	// 			SystemAssignedIdentityPrincipalID: to.Ptr("principalId"),
	// 			TenantID: to.Ptr("tenantId"),
	// 		},
	// 		Kind: to.Ptr(armpowerplatform.EnterprisePolicyKindLockbox),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armpowerplatform.Properties{
	// 			Encryption: &armpowerplatform.PropertiesEncryption{
	// 				KeyVault: &armpowerplatform.KeyVaultProperties{
	// 					ID: to.Ptr("nnn"),
	// 					Key: &armpowerplatform.KeyProperties{
	// 						Name: to.Ptr("name"),
	// 						Version: to.Ptr("1.0"),
	// 					},
	// 				},
	// 			},
	// 			HealthStatus: to.Ptr(armpowerplatform.HealthStatusUndetermined),
	// 			Lockbox: &armpowerplatform.PropertiesLockbox{
	// 				State: to.Ptr(armpowerplatform.State("succeeded")),
	// 			},
	// 			NetworkInjection: &armpowerplatform.PropertiesNetworkInjection{
	// 				VirtualNetworks: []*armpowerplatform.VirtualNetworkProperties{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/virtualNetwork"),
	// 						Subnet: &armpowerplatform.SubnetProperties{
	// 							Name: to.Ptr("testSubnet"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armpowerplatform.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armpowerplatform.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armpowerplatform.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
