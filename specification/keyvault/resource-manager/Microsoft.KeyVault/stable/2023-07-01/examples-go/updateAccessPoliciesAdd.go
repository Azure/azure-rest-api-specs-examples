package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/updateAccessPoliciesAdd.json
func ExampleVaultsClient_UpdateAccessPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVaultsClient().UpdateAccessPolicy(ctx, "sample-group", "sample-vault", armkeyvault.AccessPolicyUpdateKindAdd, armkeyvault.VaultAccessPolicyParameters{
		Properties: &armkeyvault.VaultAccessPolicyProperties{
			AccessPolicies: []*armkeyvault.AccessPolicyEntry{
				{
					ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
					Permissions: &armkeyvault.Permissions{
						Certificates: []*armkeyvault.CertificatePermissions{
							to.Ptr(armkeyvault.CertificatePermissionsGet)},
						Keys: []*armkeyvault.KeyPermissions{
							to.Ptr(armkeyvault.KeyPermissionsEncrypt)},
						Secrets: []*armkeyvault.SecretPermissions{
							to.Ptr(armkeyvault.SecretPermissionsGet)},
					},
					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VaultAccessPolicyParameters = armkeyvault.VaultAccessPolicyParameters{
	// 	Type: to.Ptr("Microsoft.KeyVault/vaults/accessPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault/accessPolicies/"),
	// 	Properties: &armkeyvault.VaultAccessPolicyProperties{
	// 		AccessPolicies: []*armkeyvault.AccessPolicyEntry{
	// 			{
	// 				ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				Permissions: &armkeyvault.Permissions{
	// 					Certificates: []*armkeyvault.CertificatePermissions{
	// 						to.Ptr(armkeyvault.CertificatePermissionsGet)},
	// 						Keys: []*armkeyvault.KeyPermissions{
	// 							to.Ptr(armkeyvault.KeyPermissionsEncrypt)},
	// 							Secrets: []*armkeyvault.SecretPermissions{
	// 								to.Ptr(armkeyvault.SecretPermissionsGet)},
	// 							},
	// 							TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					}},
	// 				},
	// 			}
}
