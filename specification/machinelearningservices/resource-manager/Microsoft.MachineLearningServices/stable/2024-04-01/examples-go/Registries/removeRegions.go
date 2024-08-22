package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registries/removeRegions.json
func ExampleRegistriesClient_BeginRemoveRegions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRegistriesClient().BeginRemoveRegions(ctx, "test-rg", "string", armmachinelearning.Registry{
		Location: to.Ptr("string"),
		Tags:     map[string]*string{},
		Identity: &armmachinelearning.ManagedServiceIdentity{
			Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
				"string": {},
			},
		},
		Kind: to.Ptr("string"),
		Properties: &armmachinelearning.RegistryProperties{
			DiscoveryURL:                  to.Ptr("string"),
			IntellectualPropertyPublisher: to.Ptr("string"),
			ManagedResourceGroup: &armmachinelearning.ArmResourceID{
				ResourceID: to.Ptr("string"),
			},
			MlFlowRegistryURI:   to.Ptr("string"),
			PublicNetworkAccess: to.Ptr("string"),
			RegionDetails: []*armmachinelearning.RegistryRegionArmDetails{
				{
					AcrDetails: []*armmachinelearning.AcrDetails{
						{
							SystemCreatedAcrAccount: &armmachinelearning.SystemCreatedAcrAccount{
								AcrAccountName: to.Ptr("string"),
								AcrAccountSKU:  to.Ptr("string"),
								ArmResourceID: &armmachinelearning.ArmResourceID{
									ResourceID: to.Ptr("string"),
								},
							},
							UserCreatedAcrAccount: &armmachinelearning.UserCreatedAcrAccount{
								ArmResourceID: &armmachinelearning.ArmResourceID{
									ResourceID: to.Ptr("string"),
								},
							},
						}},
					Location: to.Ptr("string"),
					StorageAccountDetails: []*armmachinelearning.StorageAccountDetails{
						{
							SystemCreatedStorageAccount: &armmachinelearning.SystemCreatedStorageAccount{
								AllowBlobPublicAccess: to.Ptr(false),
								ArmResourceID: &armmachinelearning.ArmResourceID{
									ResourceID: to.Ptr("string"),
								},
								StorageAccountHnsEnabled: to.Ptr(false),
								StorageAccountName:       to.Ptr("string"),
								StorageAccountType:       to.Ptr("string"),
							},
							UserCreatedStorageAccount: &armmachinelearning.UserCreatedStorageAccount{
								ArmResourceID: &armmachinelearning.ArmResourceID{
									ResourceID: to.Ptr("string"),
								},
							},
						}},
				}},
			RegistryPrivateEndpointConnections: []*armmachinelearning.RegistryPrivateEndpointConnection{
				{
					ID:       to.Ptr("string"),
					Location: to.Ptr("string"),
					Properties: &armmachinelearning.RegistryPrivateEndpointConnectionProperties{
						GroupIDs: []*string{
							to.Ptr("string")},
						PrivateEndpoint: &armmachinelearning.PrivateEndpointResource{
							SubnetArmID: to.Ptr("string"),
						},
						ProvisioningState: to.Ptr("string"),
						RegistryPrivateLinkServiceConnectionState: &armmachinelearning.RegistryPrivateLinkServiceConnectionState{
							Description:     to.Ptr("string"),
							ActionsRequired: to.Ptr("string"),
							Status:          to.Ptr(armmachinelearning.EndpointServiceConnectionStatusApproved),
						},
					},
				}},
		},
		SKU: &armmachinelearning.SKU{
			Name:     to.Ptr("string"),
			Capacity: to.Ptr[int32](1),
			Family:   to.Ptr("string"),
			Size:     to.Ptr("string"),
			Tier:     to.Ptr(armmachinelearning.SKUTierFree),
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
	// res.Registry = armmachinelearning.Registry{
	// 	Name: to.Ptr("string"),
	// 	Type: to.Ptr("string"),
	// 	ID: to.Ptr("string"),
	// 	SystemData: &armmachinelearning.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:33:56.999Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:33:56.999Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("string"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armmachinelearning.ManagedServiceIdentity{
	// 		Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
	// 			"string": &armmachinelearning.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 				PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			},
	// 		},
	// 	},
	// 	Kind: to.Ptr("string"),
	// 	Properties: &armmachinelearning.RegistryProperties{
	// 		DiscoveryURL: to.Ptr("string"),
	// 		IntellectualPropertyPublisher: to.Ptr("string"),
	// 		ManagedResourceGroup: &armmachinelearning.ArmResourceID{
	// 			ResourceID: to.Ptr("string"),
	// 		},
	// 		MlFlowRegistryURI: to.Ptr("string"),
	// 		PublicNetworkAccess: to.Ptr("string"),
	// 		RegionDetails: []*armmachinelearning.RegistryRegionArmDetails{
	// 			{
	// 				AcrDetails: []*armmachinelearning.AcrDetails{
	// 					{
	// 						SystemCreatedAcrAccount: &armmachinelearning.SystemCreatedAcrAccount{
	// 							AcrAccountName: to.Ptr("string"),
	// 							AcrAccountSKU: to.Ptr("string"),
	// 							ArmResourceID: &armmachinelearning.ArmResourceID{
	// 								ResourceID: to.Ptr("string"),
	// 							},
	// 						},
	// 						UserCreatedAcrAccount: &armmachinelearning.UserCreatedAcrAccount{
	// 							ArmResourceID: &armmachinelearning.ArmResourceID{
	// 								ResourceID: to.Ptr("string"),
	// 							},
	// 						},
	// 				}},
	// 				Location: to.Ptr("string"),
	// 				StorageAccountDetails: []*armmachinelearning.StorageAccountDetails{
	// 					{
	// 						SystemCreatedStorageAccount: &armmachinelearning.SystemCreatedStorageAccount{
	// 							AllowBlobPublicAccess: to.Ptr(false),
	// 							ArmResourceID: &armmachinelearning.ArmResourceID{
	// 								ResourceID: to.Ptr("string"),
	// 							},
	// 							StorageAccountHnsEnabled: to.Ptr(false),
	// 							StorageAccountName: to.Ptr("string"),
	// 							StorageAccountType: to.Ptr("string"),
	// 						},
	// 						UserCreatedStorageAccount: &armmachinelearning.UserCreatedStorageAccount{
	// 							ArmResourceID: &armmachinelearning.ArmResourceID{
	// 								ResourceID: to.Ptr("string"),
	// 							},
	// 						},
	// 				}},
	// 		}},
	// 		RegistryPrivateEndpointConnections: []*armmachinelearning.RegistryPrivateEndpointConnection{
	// 			{
	// 				ID: to.Ptr("string"),
	// 				Location: to.Ptr("string"),
	// 				Properties: &armmachinelearning.RegistryPrivateEndpointConnectionProperties{
	// 					GroupIDs: []*string{
	// 						to.Ptr("string")},
	// 						PrivateEndpoint: &armmachinelearning.PrivateEndpointResource{
	// 							ID: to.Ptr("string"),
	// 							SubnetArmID: to.Ptr("string"),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 						RegistryPrivateLinkServiceConnectionState: &armmachinelearning.RegistryPrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("string"),
	// 							ActionsRequired: to.Ptr("string"),
	// 							Status: to.Ptr(armmachinelearning.EndpointServiceConnectionStatusApproved),
	// 						},
	// 					},
	// 			}},
	// 		},
	// 		SKU: &armmachinelearning.SKU{
	// 			Name: to.Ptr("string"),
	// 			Capacity: to.Ptr[int32](1),
	// 			Family: to.Ptr("string"),
	// 			Size: to.Ptr("string"),
	// 			Tier: to.Ptr(armmachinelearning.SKUTierFree),
	// 		},
	// 	}
}
