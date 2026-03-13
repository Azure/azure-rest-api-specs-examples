package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps/v2"
)

// Generated from example definition: 2025-10-01-preview/CreateAccountManagedIdentity.json
func ExampleAccountsClient_CreateOrUpdate_createAccountWithManagedIdentities() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaps.NewClientFactory("21a9967a-e8a9-4656-a70b-96ff1c4d05a0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().CreateOrUpdate(ctx, "myResourceGroup", "myMapsAccount", armmaps.Account{
		Identity: &armmaps.ManagedServiceIdentity{
			Type: to.Ptr(armmaps.ManagedServiceIdentityType("SystemAssigned, UserAssigned")),
			UserAssignedIdentities: map[string]*armmaps.UserAssignedIdentity{
				"/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName": {},
			},
		},
		Kind:     to.Ptr(armmaps.KindGen2),
		Location: to.Ptr("eastus"),
		Properties: &armmaps.AccountProperties{
			DisableLocalAuth: to.Ptr(false),
			LinkedResources: []*armmaps.LinkedResource{
				{
					ID:         to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Storage/accounts/mystorageacc"),
					UniqueName: to.Ptr("myBatchStorageAccount"),
				},
				{
					ID:         to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Storage/accounts/mystorageacc"),
					UniqueName: to.Ptr("myBlobDataSource"),
				},
			},
		},
		SKU: &armmaps.SKU{
			Name: to.Ptr(armmaps.NameG2),
		},
		Tags: map[string]*string{
			"test": to.Ptr("true"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmaps.AccountsClientCreateOrUpdateResponse{
	// 	Account: &armmaps.Account{
	// 		Name: to.Ptr("myMapsAccount"),
	// 		Type: to.Ptr("Microsoft.Maps/accounts"),
	// 		ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Maps/accounts/myMapsAccount"),
	// 		Identity: &armmaps.ManagedServiceIdentity{
	// 			Type: to.Ptr(armmaps.ManagedServiceIdentityType("SystemAssigned, UserAssigned")),
	// 			PrincipalID: to.Ptr("77f72dac-e0aa-484e-9acd-e5e7075310ef"),
	// 			TenantID: to.Ptr("06006684-60c1-4954-a20c-ffd8fbea7276"),
	// 			UserAssignedIdentities: map[string]*armmaps.UserAssignedIdentity{
	// 				"/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName": &armmaps.UserAssignedIdentity{
	// 					ClientID: to.Ptr("b602d315-01b5-4265-af23-859edc4f2431"),
	// 					PrincipalID: to.Ptr("ac287332-364a-41d9-a567-9ad86b9fc299"),
	// 				},
	// 			},
	// 		},
	// 		Kind: to.Ptr(armmaps.KindGen2),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armmaps.AccountProperties{
	// 			DisableLocalAuth: to.Ptr(true),
	// 			LinkedResources: []*armmaps.LinkedResource{
	// 				{
	// 					ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Storage/accounts/mystorageacc"),
	// 					UniqueName: to.Ptr("myBatchStorageAccount"),
	// 				},
	// 				{
	// 					ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Storage/accounts/mystorageacc"),
	// 					UniqueName: to.Ptr("myBlobDataSource"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			UniqueID: to.Ptr("b2e763e6-d6f3-4858-9e2b-7cf8df85c593"),
	// 		},
	// 		SKU: &armmaps.SKU{
	// 			Name: to.Ptr(armmaps.NameG2),
	// 			Tier: to.Ptr("Standard"),
	// 		},
	// 		SystemData: &armmaps.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-02T01:01:01.1075056Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armmaps.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-02T01:01:01.1075056Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armmaps.CreatedByTypeApplication),
	// 		},
	// 		Tags: map[string]*string{
	// 			"test": to.Ptr("true"),
	// 		},
	// 	},
	// }
}
