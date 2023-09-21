package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedInstanceUpdateMax.json
func ExampleManagedInstancesClient_BeginUpdate_updateManagedInstanceWithAllProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedInstancesClient().BeginUpdate(ctx, "testrg", "testinstance", armsql.ManagedInstanceUpdate{
		Properties: &armsql.ManagedInstanceProperties{
			AdministratorLogin:               to.Ptr("dummylogin"),
			AdministratorLoginPassword:       to.Ptr("PLACEHOLDER"),
			Collation:                        to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
			LicenseType:                      to.Ptr(armsql.ManagedInstanceLicenseTypeBasePrice),
			MaintenanceConfigurationID:       to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1"),
			MinimalTLSVersion:                to.Ptr("1.2"),
			ProxyOverride:                    to.Ptr(armsql.ManagedInstanceProxyOverrideRedirect),
			PublicDataEndpointEnabled:        to.Ptr(false),
			RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
			StorageSizeInGB:                  to.Ptr[int32](448),
			VCores:                           to.Ptr[int32](8),
		},
		SKU: &armsql.SKU{
			Name:     to.Ptr("GP_Gen4"),
			Capacity: to.Ptr[int32](8),
			Tier:     to.Ptr("GeneralPurpose"),
		},
		Tags: map[string]*string{
			"tagKey1": to.Ptr("TagValue1"),
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
	// res.ManagedInstance = armsql.ManagedInstance{
	// 	Name: to.Ptr("testinstance"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances"),
	// 	ID: to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance"),
	// 	Location: to.Ptr("japaneast"),
	// 	Tags: map[string]*string{
	// 		"tagKey1": to.Ptr("TagValue1"),
	// 	},
	// 	Properties: &armsql.ManagedInstanceProperties{
	// 		AdministratorLogin: to.Ptr("dummylogin"),
	// 		Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
	// 		CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		DNSZone: to.Ptr("1b4e2caff2530"),
	// 		FullyQualifiedDomainName: to.Ptr("testinstance.1b4e2caff2530.database.windows.net"),
	// 		InstancePoolID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Sql/instancePools/instancePool1"),
	// 		LicenseType: to.Ptr(armsql.ManagedInstanceLicenseTypeBasePrice),
	// 		MaintenanceConfigurationID: to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1"),
	// 		MinimalTLSVersion: to.Ptr("1.2"),
	// 		ProxyOverride: to.Ptr(armsql.ManagedInstanceProxyOverrideRedirect),
	// 		PublicDataEndpointEnabled: to.Ptr(false),
	// 		RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		State: to.Ptr("Ready"),
	// 		StorageSizeInGB: to.Ptr[int32](448),
	// 		SubnetID: to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 		VCores: to.Ptr[int32](8),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("GP_Gen4"),
	// 		Capacity: to.Ptr[int32](8),
	// 		Family: to.Ptr("Gen4"),
	// 		Tier: to.Ptr("GeneralPurpose"),
	// 	},
	// }
}
