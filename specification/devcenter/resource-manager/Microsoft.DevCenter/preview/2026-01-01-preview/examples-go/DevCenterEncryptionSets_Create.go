package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/DevCenterEncryptionSets_Create.json
func ExampleEncryptionSetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEncryptionSetsClient().BeginCreateOrUpdate(ctx, "rg1", "Contoso", "EncryptionWestUs", armdevcenter.EncryptionSet{
		Identity: &armdevcenter.ManagedServiceIdentity{
			Type: to.Ptr(armdevcenter.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armdevcenter.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1": {},
			},
		},
		Location: to.Ptr("westus"),
		Properties: &armdevcenter.EncryptionSetProperties{
			DevboxDisksEncryptionEnableStatus: to.Ptr(armdevcenter.DevboxDisksEncryptionEnableStatusEnabled),
			KeyEncryptionKeyIdentity: &armdevcenter.KeyEncryptionKeyIdentity{
				Type:                           to.Ptr(armdevcenter.CmkIdentityTypeUserAssigned),
				UserAssignedIdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1"),
			},
			KeyEncryptionKeyURL: to.Ptr("https://contosovaultwestus.vault.azure.net/keys/contosokek"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevcenter.EncryptionSetsClientCreateOrUpdateResponse{
	// 	EncryptionSet: armdevcenter.EncryptionSet{
	// 		Name: to.Ptr("EncryptionWestUs"),
	// 		Type: to.Ptr("Microsoft.DevCenter/devcenters/encryptionSets"),
	// 		ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/encryptionSets/EncryptionWestUs"),
	// 		Identity: &armdevcenter.ManagedServiceIdentity{
	// 			Type: to.Ptr(armdevcenter.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armdevcenter.UserAssignedIdentity{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1": &armdevcenter.UserAssignedIdentity{
	// 					ClientID: to.Ptr("e35621a5-f615-4a20-940e-de8a84b15abc"),
	// 					PrincipalID: to.Ptr("2111b8fc-e123-485a-b408-bf1153189494"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armdevcenter.EncryptionSetProperties{
	// 			DevboxDisksEncryptionEnableStatus: to.Ptr(armdevcenter.DevboxDisksEncryptionEnableStatusEnabled),
	// 			KeyEncryptionKeyIdentity: &armdevcenter.KeyEncryptionKeyIdentity{
	// 				Type: to.Ptr(armdevcenter.CmkIdentityTypeUserAssigned),
	// 				UserAssignedIdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1"),
	// 			},
	// 			KeyEncryptionKeyURL: to.Ptr("https://contosovaultwestus.vault.azure.net/keys/contosokek"),
	// 			ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateAccepted),
	// 		},
	// 		SystemData: &armdevcenter.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
	// 			CreatedBy: to.Ptr("User1"),
	// 			CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("User1"),
	// 			LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
