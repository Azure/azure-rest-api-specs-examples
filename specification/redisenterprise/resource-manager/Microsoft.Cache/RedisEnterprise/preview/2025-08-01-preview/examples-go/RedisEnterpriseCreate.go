package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v4"
)

// Generated from example definition: 2025-08-01-preview/RedisEnterpriseCreate.json
func ExampleClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginCreate(ctx, "rg1", "cache1", armredisenterprise.Cluster{
		Identity: &armredisenterprise.ManagedServiceIdentity{
			Type: to.Ptr(armredisenterprise.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armredisenterprise.UserAssignedIdentity{
				"/subscriptions/your-subscription/resourceGroups/your-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/your-identity": {},
			},
		},
		Location: to.Ptr("West US"),
		Properties: &armredisenterprise.ClusterCreateProperties{
			Encryption: &armredisenterprise.ClusterPropertiesEncryption{
				CustomerManagedKeyEncryption: &armredisenterprise.ClusterPropertiesEncryptionCustomerManagedKeyEncryption{
					KeyEncryptionKeyIdentity: &armredisenterprise.ClusterPropertiesEncryptionCustomerManagedKeyEncryptionKeyIdentity{
						IdentityType:                   to.Ptr(armredisenterprise.CmkIdentityTypeUserAssignedIdentity),
						UserAssignedIdentityResourceID: to.Ptr("/subscriptions/your-subscription/resourceGroups/your-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/your-identity"),
					},
					KeyEncryptionKeyURL: to.Ptr("https://your-kv.vault.azure.net/keys/your-key/your-key-version"),
				},
			},
			MinimumTLSVersion:   to.Ptr(armredisenterprise.TLSVersionOne2),
			PublicNetworkAccess: to.Ptr(armredisenterprise.PublicNetworkAccessDisabled),
			MaintenanceConfiguration: &armredisenterprise.MaintenanceConfiguration{
				MaintenanceWindows: []*armredisenterprise.MaintenanceWindow{
					{
						Type:         to.Ptr(armredisenterprise.MaintenanceWindowTypeWeekly),
						Duration:     to.Ptr("PT6H"),
						StartHourUTC: to.Ptr[int32](3),
						Schedule: &armredisenterprise.MaintenanceWindowSchedule{
							DayOfWeek: to.Ptr(armredisenterprise.MaintenanceDayOfWeekMonday),
						},
					},
					{
						Type:         to.Ptr(armredisenterprise.MaintenanceWindowTypeWeekly),
						Duration:     to.Ptr("PT6H"),
						StartHourUTC: to.Ptr[int32](3),
						Schedule: &armredisenterprise.MaintenanceWindowSchedule{
							DayOfWeek: to.Ptr(armredisenterprise.MaintenanceDayOfWeekTuesday),
						},
					},
					{
						Type:         to.Ptr(armredisenterprise.MaintenanceWindowTypeWeekly),
						Duration:     to.Ptr("PT6H"),
						StartHourUTC: to.Ptr[int32](3),
						Schedule: &armredisenterprise.MaintenanceWindowSchedule{
							DayOfWeek: to.Ptr(armredisenterprise.MaintenanceDayOfWeekWednesday),
						},
					},
				},
			},
		},
		SKU: &armredisenterprise.SKU{
			Name:     to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
			Capacity: to.Ptr[int32](3),
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
		},
		Zones: []*string{
			to.Ptr("1"),
			to.Ptr("2"),
			to.Ptr("3"),
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
	// res = armredisenterprise.ClientCreateResponse{
	// 	Cluster: &armredisenterprise.Cluster{
	// 		Name: to.Ptr("cache1"),
	// 		Type: to.Ptr("Microsoft.Cache/redisEnterprise"),
	// 		ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1"),
	// 		Identity: &armredisenterprise.ManagedServiceIdentity{
	// 			Type: to.Ptr(armredisenterprise.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armredisenterprise.UserAssignedIdentity{
	// 				"/subscriptions/your-subscription/resourceGroups/your-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/your-identity": &armredisenterprise.UserAssignedIdentity{
	// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 			},
	// 		},
	// 		Kind: to.Ptr(armredisenterprise.KindV1),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armredisenterprise.ClusterCreateProperties{
	// 			HostName: to.Ptr("cache1.westus.something.azure.net"),
	// 			MinimumTLSVersion: to.Ptr(armredisenterprise.TLSVersionOne2),
	// 			ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
	// 			PublicNetworkAccess: to.Ptr(armredisenterprise.PublicNetworkAccessDisabled),
	// 			MaintenanceConfiguration: &armredisenterprise.MaintenanceConfiguration{
	// 				MaintenanceWindows: []*armredisenterprise.MaintenanceWindow{
	// 					{
	// 						Type: to.Ptr(armredisenterprise.MaintenanceWindowTypeWeekly),
	// 						Duration: to.Ptr("PT6H"),
	// 						StartHourUTC: to.Ptr[int32](3),
	// 						Schedule: &armredisenterprise.MaintenanceWindowSchedule{
	// 							DayOfWeek: to.Ptr(armredisenterprise.MaintenanceDayOfWeekMonday),
	// 						},
	// 					},
	// 					{
	// 						Type: to.Ptr(armredisenterprise.MaintenanceWindowTypeWeekly),
	// 						Duration: to.Ptr("PT6H"),
	// 						StartHourUTC: to.Ptr[int32](3),
	// 						Schedule: &armredisenterprise.MaintenanceWindowSchedule{
	// 							DayOfWeek: to.Ptr(armredisenterprise.MaintenanceDayOfWeekTuesday),
	// 						},
	// 					},
	// 					{
	// 						Type: to.Ptr(armredisenterprise.MaintenanceWindowTypeWeekly),
	// 						Duration: to.Ptr("PT6H"),
	// 						StartHourUTC: to.Ptr[int32](3),
	// 						Schedule: &armredisenterprise.MaintenanceWindowSchedule{
	// 							DayOfWeek: to.Ptr(armredisenterprise.MaintenanceDayOfWeekWednesday),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			RedisVersion: to.Ptr("5"),
	// 			ResourceState: to.Ptr(armredisenterprise.ResourceStateRunning),
	// 		},
	// 		SKU: &armredisenterprise.SKU{
	// 			Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
	// 			Capacity: to.Ptr[int32](3),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 			to.Ptr("2"),
	// 			to.Ptr("3"),
	// 		},
	// 	},
	// }
}
