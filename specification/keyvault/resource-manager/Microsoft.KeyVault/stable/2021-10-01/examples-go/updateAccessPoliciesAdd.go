package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/updateAccessPoliciesAdd.json
func ExampleVaultsClient_UpdateAccessPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkeyvault.NewVaultsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateAccessPolicy(ctx,
		"sample-group",
		"sample-vault",
		armkeyvault.AccessPolicyUpdateKindAdd,
		armkeyvault.VaultAccessPolicyParameters{
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
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
