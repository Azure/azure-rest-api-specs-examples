package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedInstanceListByResourceGroup.json
func ExampleManagedInstancesClient_NewListByResourceGroupPager_listManagedInstancesByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedInstancesClient().NewListByResourceGroupPager("Test1", &armsql.ManagedInstancesClientListByResourceGroupOptions{Expand: nil})
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
		// page.ManagedInstanceListResult = armsql.ManagedInstanceListResult{
		// 	Value: []*armsql.ManagedInstance{
		// 		{
		// 			Name: to.Ptr("testinstance1"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances"),
		// 			ID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/Test1/providers/Microsoft.Sql/managedInstances/testinstance1"),
		// 			Location: to.Ptr("japaneast"),
		// 			Properties: &armsql.ManagedInstanceProperties{
		// 				AdministratorLogin: to.Ptr("dummylogin"),
		// 				Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		// 				CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				DNSZone: to.Ptr("1b4e2caff2530"),
		// 				FullyQualifiedDomainName: to.Ptr("testinstance1.1b4e2caff2530.database.windows.net"),
		// 				LicenseType: to.Ptr(armsql.ManagedInstanceLicenseTypeLicenseIncluded),
		// 				MaintenanceConfigurationID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1"),
		// 				MinimalTLSVersion: to.Ptr("1.2"),
		// 				ProvisioningState: to.Ptr(armsql.ManagedInstancePropertiesProvisioningStateSucceeded),
		// 				ProxyOverride: to.Ptr(armsql.ManagedInstanceProxyOverrideDefault),
		// 				PublicDataEndpointEnabled: to.Ptr(false),
		// 				RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				State: to.Ptr("Ready"),
		// 				StorageSizeInGB: to.Ptr[int32](1024),
		// 				SubnetID: to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
		// 				VCores: to.Ptr[int32](8),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("GP_Gen4"),
		// 				Capacity: to.Ptr[int32](8),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr("GeneralPurpose"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testinstance2"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances"),
		// 			ID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/Test1/providers/Microsoft.Sql/managedInstances/testinstance2"),
		// 			Location: to.Ptr("japaneast"),
		// 			Properties: &armsql.ManagedInstanceProperties{
		// 				AdministratorLogin: to.Ptr("dummylogin"),
		// 				Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		// 				CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				DNSZone: to.Ptr("2c3d1bdae3412"),
		// 				FullyQualifiedDomainName: to.Ptr("testinstance2.2c3d1bdae3412.database.windows.net"),
		// 				LicenseType: to.Ptr(armsql.ManagedInstanceLicenseType("Full")),
		// 				MaintenanceConfigurationID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1"),
		// 				MinimalTLSVersion: to.Ptr("1.2"),
		// 				ProvisioningState: to.Ptr(armsql.ManagedInstancePropertiesProvisioningStateSucceeded),
		// 				ProxyOverride: to.Ptr(armsql.ManagedInstanceProxyOverrideDefault),
		// 				PublicDataEndpointEnabled: to.Ptr(false),
		// 				RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				State: to.Ptr("Ready"),
		// 				StorageSizeInGB: to.Ptr[int32](1024),
		// 				SubnetID: to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/subnet2"),
		// 				VCores: to.Ptr[int32](16),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("BC_Gen5"),
		// 				Capacity: to.Ptr[int32](16),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr("BusinessCritical"),
		// 			},
		// 	}},
		// }
	}
}
