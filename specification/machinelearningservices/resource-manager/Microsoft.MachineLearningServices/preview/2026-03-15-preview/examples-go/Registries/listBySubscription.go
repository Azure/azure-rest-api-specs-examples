package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Registries/listBySubscription.json
func ExampleRegistriesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRegistriesClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armmachinelearning.RegistriesClientListBySubscriptionResponse{
		// 	RegistryTrackedResourceArmPaginatedResult: armmachinelearning.RegistryTrackedResourceArmPaginatedResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.MachineLearningServices/registries?api-version=2025-07-01-preview&$skip=2"),
		// 		Value: []*armmachinelearning.Registry{
		// 			{
		// 				Name: to.Ptr("string"),
		// 				Type: to.Ptr("string"),
		// 				ID: to.Ptr("string"),
		// 				Identity: &armmachinelearning.ManagedServiceIdentity{
		// 					Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 					TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 					UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
		// 						"string": &armmachinelearning.UserAssignedIdentity{
		// 							ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 							PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 						},
		// 					},
		// 				},
		// 				Kind: to.Ptr("string"),
		// 				Location: to.Ptr("string"),
		// 				Properties: &armmachinelearning.RegistryProperties{
		// 					DiscoveryURL: to.Ptr("string"),
		// 					IntellectualPropertyPublisher: to.Ptr("string"),
		// 					ManagedResourceGroup: &armmachinelearning.ArmResourceID{
		// 						ResourceID: to.Ptr("string"),
		// 					},
		// 					MlFlowRegistryURI: to.Ptr("string"),
		// 					PublicNetworkAccess: to.Ptr("string"),
		// 					RegionDetails: []*armmachinelearning.RegistryRegionArmDetails{
		// 						{
		// 							AcrDetails: []*armmachinelearning.AcrDetails{
		// 								{
		// 									SystemCreatedAcrAccount: &armmachinelearning.SystemCreatedAcrAccount{
		// 										AcrAccountName: to.Ptr("string"),
		// 										AcrAccountSKU: to.Ptr("string"),
		// 										ArmResourceID: &armmachinelearning.ArmResourceID{
		// 											ResourceID: to.Ptr("string"),
		// 										},
		// 									},
		// 								},
		// 							},
		// 							Location: to.Ptr("string"),
		// 							StorageAccountDetails: []*armmachinelearning.StorageAccountDetails{
		// 								{
		// 									SystemCreatedStorageAccount: &armmachinelearning.SystemCreatedStorageAccount{
		// 										AllowBlobPublicAccess: to.Ptr(false),
		// 										ArmResourceID: &armmachinelearning.ArmResourceID{
		// 											ResourceID: to.Ptr("string"),
		// 										},
		// 										StorageAccountHnsEnabled: to.Ptr(false),
		// 										StorageAccountName: to.Ptr("string"),
		// 										StorageAccountType: to.Ptr("string"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					RegistryPrivateEndpointConnections: []*armmachinelearning.RegistryPrivateEndpointConnection{
		// 						{
		// 							ID: to.Ptr("string"),
		// 							Location: to.Ptr("string"),
		// 							Properties: &armmachinelearning.RegistryPrivateEndpointConnectionProperties{
		// 								GroupIDs: []*string{
		// 									to.Ptr("string"),
		// 								},
		// 								PrivateEndpoint: &armmachinelearning.PrivateEndpointResource{
		// 									ID: to.Ptr("string"),
		// 									SubnetArmID: to.Ptr("string"),
		// 								},
		// 								ProvisioningState: to.Ptr("string"),
		// 								RegistryPrivateLinkServiceConnectionState: &armmachinelearning.RegistryPrivateLinkServiceConnectionState{
		// 									Description: to.Ptr("string"),
		// 									ActionsRequired: to.Ptr("string"),
		// 									Status: to.Ptr(armmachinelearning.EndpointServiceConnectionStatusApproved),
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 				SKU: &armmachinelearning.SKU{
		// 					Name: to.Ptr("string"),
		// 					Capacity: to.Ptr[int32](1),
		// 					Family: to.Ptr("string"),
		// 					Size: to.Ptr("string"),
		// 					Tier: to.Ptr(armmachinelearning.SKUTierFree),
		// 				},
		// 				SystemData: &armmachinelearning.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999+00:15"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999+00:15"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
