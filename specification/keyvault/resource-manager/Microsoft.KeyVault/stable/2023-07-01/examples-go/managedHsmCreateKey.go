package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/managedHsmCreateKey.json
func ExampleManagedHsmKeysClient_CreateIfNotExist() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedHsmKeysClient().CreateIfNotExist(ctx, "sample-group", "sample-managedhsm-name", "sample-key-name", armkeyvault.ManagedHsmKeyCreateParameters{
		Properties: &armkeyvault.ManagedHsmKeyProperties{
			Kty: to.Ptr(armkeyvault.JSONWebKeyTypeRSA),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedHsmKey = armkeyvault.ManagedHsmKey{
	// 	Name: to.Ptr("sample-key-name"),
	// 	Type: to.Ptr("Microsoft.KeyVault/managedHSMs/keys"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedHSMs/sample-managedhsm-name/keys/sample-key-name"),
	// 	Properties: &armkeyvault.ManagedHsmKeyProperties{
	// 		Attributes: &armkeyvault.ManagedHsmKeyAttributes{
	// 			Created: to.Ptr[int64](1598533051),
	// 			Enabled: to.Ptr(true),
	// 			RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
	// 			Updated: to.Ptr[int64](1598533051),
	// 		},
	// 		KeyOps: []*armkeyvault.JSONWebKeyOperation{
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationEncrypt),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationDecrypt),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationSign),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationVerify),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationWrapKey),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationUnwrapKey)},
	// 			KeySize: to.Ptr[int32](2048),
	// 			KeyURI: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name"),
	// 			KeyURIWithVersion: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name/fd618d9519b74f9aae94ade66b876acc"),
	// 			Kty: to.Ptr(armkeyvault.JSONWebKeyTypeRSA),
	// 		},
	// 	}
}
