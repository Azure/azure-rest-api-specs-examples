package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registries/list-UserCreated.json
func ExampleRegistriesClient_NewListPager_listRegistriesWithUserCreatedAccounts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRegistriesClient().NewListPager("test-rg", nil)
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
		// page.RegistryTrackedResourceArmPaginatedResult = armmachinelearning.RegistryTrackedResourceArmPaginatedResult{
		// 	Value: []*armmachinelearning.Registry{
		// 		{
		// 			Name: to.Ptr("string"),
		// 			Type: to.Ptr("string"),
		// 			ID: to.Ptr("string"),
		// 			SystemData: &armmachinelearning.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T11:52:56.999Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T11:52:56.999Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("string"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armmachinelearning.ManagedServiceIdentity{
		// 				Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeUserAssigned),
		// 				PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
		// 					"string": &armmachinelearning.UserAssignedIdentity{
		// 						ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 						PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 					},
		// 				},
		// 			},
		// 			Kind: to.Ptr("string"),
		// 			Properties: &armmachinelearning.RegistryProperties{
		// 				DiscoveryURL: to.Ptr("string"),
		// 				IntellectualPropertyPublisher: to.Ptr("string"),
		// 				ManagedResourceGroup: &armmachinelearning.ArmResourceID{
		// 					ResourceID: to.Ptr("string"),
		// 				},
		// 				MlFlowRegistryURI: to.Ptr("string"),
		// 				PublicNetworkAccess: to.Ptr("string"),
		// 				RegionDetails: []*armmachinelearning.RegistryRegionArmDetails{
		// 					{
		// 						AcrDetails: []*armmachinelearning.AcrDetails{
		// 							{
		// 								UserCreatedAcrAccount: &armmachinelearning.UserCreatedAcrAccount{
		// 									ArmResourceID: &armmachinelearning.ArmResourceID{
		// 										ResourceID: to.Ptr("string"),
		// 									},
		// 								},
		// 						}},
		// 						Location: to.Ptr("string"),
		// 						StorageAccountDetails: []*armmachinelearning.StorageAccountDetails{
		// 							{
		// 								UserCreatedStorageAccount: &armmachinelearning.UserCreatedStorageAccount{
		// 									ArmResourceID: &armmachinelearning.ArmResourceID{
		// 										ResourceID: to.Ptr("string"),
		// 									},
		// 								},
		// 						}},
		// 				}},
		// 				RegistryPrivateEndpointConnections: []*armmachinelearning.RegistryPrivateEndpointConnection{
		// 					{
		// 						ID: to.Ptr("string"),
		// 						Location: to.Ptr("string"),
		// 						Properties: &armmachinelearning.RegistryPrivateEndpointConnectionProperties{
		// 							GroupIDs: []*string{
		// 								to.Ptr("string")},
		// 								PrivateEndpoint: &armmachinelearning.PrivateEndpointResource{
		// 									ID: to.Ptr("string"),
		// 									SubnetArmID: to.Ptr("string"),
		// 								},
		// 								ProvisioningState: to.Ptr("Succeeded"),
		// 								RegistryPrivateLinkServiceConnectionState: &armmachinelearning.RegistryPrivateLinkServiceConnectionState{
		// 									Description: to.Ptr("string"),
		// 									ActionsRequired: to.Ptr("string"),
		// 									Status: to.Ptr(armmachinelearning.EndpointServiceConnectionStatusApproved),
		// 								},
		// 							},
		// 					}},
		// 				},
		// 				SKU: &armmachinelearning.SKU{
		// 					Name: to.Ptr("string"),
		// 					Capacity: to.Ptr[int32](1),
		// 					Family: to.Ptr("string"),
		// 					Size: to.Ptr("string"),
		// 					Tier: to.Ptr(armmachinelearning.SKUTierFree),
		// 				},
		// 		}},
		// 	}
	}
}
