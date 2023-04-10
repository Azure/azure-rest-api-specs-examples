package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/createVault.json
func ExampleVaultsClient_BeginCreateOrUpdate_createANewVaultOrUpdateAnExistingVault() {
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
			AccessPolicies: []*armkeyvault.AccessPolicyEntry{
				{
					ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
					Permissions: &armkeyvault.Permissions{
						Certificates: []*armkeyvault.CertificatePermissions{
							to.Ptr(armkeyvault.CertificatePermissionsGet),
							to.Ptr(armkeyvault.CertificatePermissionsList),
							to.Ptr(armkeyvault.CertificatePermissionsDelete),
							to.Ptr(armkeyvault.CertificatePermissionsCreate),
							to.Ptr(armkeyvault.CertificatePermissionsImport),
							to.Ptr(armkeyvault.CertificatePermissionsUpdate),
							to.Ptr(armkeyvault.CertificatePermissionsManagecontacts),
							to.Ptr(armkeyvault.CertificatePermissionsGetissuers),
							to.Ptr(armkeyvault.CertificatePermissionsListissuers),
							to.Ptr(armkeyvault.CertificatePermissionsSetissuers),
							to.Ptr(armkeyvault.CertificatePermissionsDeleteissuers),
							to.Ptr(armkeyvault.CertificatePermissionsManageissuers),
							to.Ptr(armkeyvault.CertificatePermissionsRecover),
							to.Ptr(armkeyvault.CertificatePermissionsPurge)},
						Keys: []*armkeyvault.KeyPermissions{
							to.Ptr(armkeyvault.KeyPermissionsEncrypt),
							to.Ptr(armkeyvault.KeyPermissionsDecrypt),
							to.Ptr(armkeyvault.KeyPermissionsWrapKey),
							to.Ptr(armkeyvault.KeyPermissionsUnwrapKey),
							to.Ptr(armkeyvault.KeyPermissionsSign),
							to.Ptr(armkeyvault.KeyPermissionsVerify),
							to.Ptr(armkeyvault.KeyPermissionsGet),
							to.Ptr(armkeyvault.KeyPermissionsList),
							to.Ptr(armkeyvault.KeyPermissionsCreate),
							to.Ptr(armkeyvault.KeyPermissionsUpdate),
							to.Ptr(armkeyvault.KeyPermissionsImport),
							to.Ptr(armkeyvault.KeyPermissionsDelete),
							to.Ptr(armkeyvault.KeyPermissionsBackup),
							to.Ptr(armkeyvault.KeyPermissionsRestore),
							to.Ptr(armkeyvault.KeyPermissionsRecover),
							to.Ptr(armkeyvault.KeyPermissionsPurge)},
						Secrets: []*armkeyvault.SecretPermissions{
							to.Ptr(armkeyvault.SecretPermissionsGet),
							to.Ptr(armkeyvault.SecretPermissionsList),
							to.Ptr(armkeyvault.SecretPermissionsSet),
							to.Ptr(armkeyvault.SecretPermissionsDelete),
							to.Ptr(armkeyvault.SecretPermissionsBackup),
							to.Ptr(armkeyvault.SecretPermissionsRestore),
							to.Ptr(armkeyvault.SecretPermissionsRecover),
							to.Ptr(armkeyvault.SecretPermissionsPurge)},
					},
					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
				}},
			EnabledForDeployment:         to.Ptr(true),
			EnabledForDiskEncryption:     to.Ptr(true),
			EnabledForTemplateDeployment: to.Ptr(true),
			PublicNetworkAccess:          to.Ptr("Enabled"),
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
	// 		AccessPolicies: []*armkeyvault.AccessPolicyEntry{
	// 			{
	// 				ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				Permissions: &armkeyvault.Permissions{
	// 					Certificates: []*armkeyvault.CertificatePermissions{
	// 						to.Ptr(armkeyvault.CertificatePermissionsGet),
	// 						to.Ptr(armkeyvault.CertificatePermissionsList),
	// 						to.Ptr(armkeyvault.CertificatePermissionsDelete),
	// 						to.Ptr(armkeyvault.CertificatePermissionsCreate),
	// 						to.Ptr(armkeyvault.CertificatePermissionsImport),
	// 						to.Ptr(armkeyvault.CertificatePermissionsUpdate),
	// 						to.Ptr(armkeyvault.CertificatePermissionsManagecontacts),
	// 						to.Ptr(armkeyvault.CertificatePermissionsGetissuers),
	// 						to.Ptr(armkeyvault.CertificatePermissionsListissuers),
	// 						to.Ptr(armkeyvault.CertificatePermissionsSetissuers),
	// 						to.Ptr(armkeyvault.CertificatePermissionsDeleteissuers),
	// 						to.Ptr(armkeyvault.CertificatePermissionsManageissuers),
	// 						to.Ptr(armkeyvault.CertificatePermissionsRecover),
	// 						to.Ptr(armkeyvault.CertificatePermissionsPurge)},
	// 						Keys: []*armkeyvault.KeyPermissions{
	// 							to.Ptr(armkeyvault.KeyPermissionsEncrypt),
	// 							to.Ptr(armkeyvault.KeyPermissionsDecrypt),
	// 							to.Ptr(armkeyvault.KeyPermissionsWrapKey),
	// 							to.Ptr(armkeyvault.KeyPermissionsUnwrapKey),
	// 							to.Ptr(armkeyvault.KeyPermissionsSign),
	// 							to.Ptr(armkeyvault.KeyPermissionsVerify),
	// 							to.Ptr(armkeyvault.KeyPermissionsGet),
	// 							to.Ptr(armkeyvault.KeyPermissionsList),
	// 							to.Ptr(armkeyvault.KeyPermissionsCreate),
	// 							to.Ptr(armkeyvault.KeyPermissionsUpdate),
	// 							to.Ptr(armkeyvault.KeyPermissionsImport),
	// 							to.Ptr(armkeyvault.KeyPermissionsDelete),
	// 							to.Ptr(armkeyvault.KeyPermissionsBackup),
	// 							to.Ptr(armkeyvault.KeyPermissionsRestore),
	// 							to.Ptr(armkeyvault.KeyPermissionsRecover),
	// 							to.Ptr(armkeyvault.KeyPermissionsPurge)},
	// 							Secrets: []*armkeyvault.SecretPermissions{
	// 								to.Ptr(armkeyvault.SecretPermissionsGet),
	// 								to.Ptr(armkeyvault.SecretPermissionsList),
	// 								to.Ptr(armkeyvault.SecretPermissionsSet),
	// 								to.Ptr(armkeyvault.SecretPermissionsDelete),
	// 								to.Ptr(armkeyvault.SecretPermissionsBackup),
	// 								to.Ptr(armkeyvault.SecretPermissionsRestore),
	// 								to.Ptr(armkeyvault.SecretPermissionsRecover),
	// 								to.Ptr(armkeyvault.SecretPermissionsPurge)},
	// 							},
	// 							TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					}},
	// 					EnabledForDeployment: to.Ptr(true),
	// 					EnabledForDiskEncryption: to.Ptr(true),
	// 					EnabledForTemplateDeployment: to.Ptr(true),
	// 					HsmPoolResourceID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					NetworkACLs: &armkeyvault.NetworkRuleSet{
	// 						Bypass: to.Ptr(armkeyvault.NetworkRuleBypassOptionsAzureServices),
	// 						DefaultAction: to.Ptr(armkeyvault.NetworkRuleActionDeny),
	// 						IPRules: []*armkeyvault.IPRule{
	// 							{
	// 								Value: to.Ptr(""),
	// 						}},
	// 						VirtualNetworkRules: []*armkeyvault.VirtualNetworkRule{
	// 							{
	// 								ID: to.Ptr(""),
	// 								IgnoreMissingVnetServiceEndpoint: to.Ptr(false),
	// 						}},
	// 					},
	// 					PrivateEndpointConnections: []*armkeyvault.PrivateEndpointConnectionItem{
	// 						{
	// 							Etag: to.Ptr(""),
	// 							ID: to.Ptr(""),
	// 							Properties: &armkeyvault.PrivateEndpointConnectionProperties{
	// 								PrivateEndpoint: &armkeyvault.PrivateEndpoint{
	// 									ID: to.Ptr(""),
	// 								},
	// 								PrivateLinkServiceConnectionState: &armkeyvault.PrivateLinkServiceConnectionState{
	// 									ActionsRequired: to.Ptr(armkeyvault.ActionsRequiredNone),
	// 									Status: to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
	// 								},
	// 								ProvisioningState: to.Ptr(armkeyvault.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 							},
	// 					}},
	// 					ProvisioningState: to.Ptr(armkeyvault.VaultProvisioningStateSucceeded),
	// 					PublicNetworkAccess: to.Ptr("Enabled"),
	// 					SKU: &armkeyvault.SKU{
	// 						Name: to.Ptr(armkeyvault.SKUNameStandard),
	// 						Family: to.Ptr(armkeyvault.SKUFamilyA),
	// 					},
	// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					VaultURI: to.Ptr("https://sample-vault.vault.azure.net"),
	// 				},
	// 				SystemData: &armkeyvault.SystemData{
	// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:00:00.0000000Z"); return t}()),
	// 					CreatedBy: to.Ptr("keyVaultUser1"),
	// 					CreatedByType: to.Ptr(armkeyvault.IdentityTypeUser),
	// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:00:00.0000000Z"); return t}()),
	// 					LastModifiedBy: to.Ptr("keyVaultUser2"),
	// 					LastModifiedByType: to.Ptr(armkeyvault.IdentityTypeUser),
	// 				},
	// 				Tags: map[string]*string{
	// 				},
	// 			}
}
