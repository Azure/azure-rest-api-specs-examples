package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/ManagedInstanceListByResourceGroup.json
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
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Test1/providers/Microsoft.Sql/managedInstances/testinstance1"),
		// 			Location: to.Ptr("japaneast"),
		// 			Properties: &armsql.ManagedInstanceProperties{
		// 				AdministratorLogin: to.Ptr("dummylogin"),
		// 				AuthenticationMetadata: to.Ptr(armsql.AuthMetadataLookupModesWindows),
		// 				Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-01T01:01:01.011Z"); return t}()),
		// 				CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				DatabaseFormat: to.Ptr(armsql.ManagedInstanceDatabaseFormatSQLServer2022),
		// 				DNSZone: to.Ptr("1b4e2caff2530"),
		// 				ExternalGovernanceStatus: to.Ptr(armsql.ExternalGovernanceStatusEnabled),
		// 				FullyQualifiedDomainName: to.Ptr("testinstance1.1b4e2caff2530.database.windows.net"),
		// 				HybridSecondaryUsage: to.Ptr(armsql.HybridSecondaryUsagePassive),
		// 				LicenseType: to.Ptr(armsql.ManagedInstanceLicenseTypeLicenseIncluded),
		// 				MaintenanceConfigurationID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1"),
		// 				MinimalTLSVersion: to.Ptr("1.2"),
		// 				PricingModel: to.Ptr(armsql.FreemiumTypeRegular),
		// 				ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
		// 				ProxyOverride: to.Ptr(armsql.ManagedInstanceProxyOverrideDefault),
		// 				PublicDataEndpointEnabled: to.Ptr(false),
		// 				RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				State: to.Ptr("Ready"),
		// 				StorageSizeInGB: to.Ptr[int32](1024),
		// 				SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
		// 				VCores: to.Ptr[int32](8),
		// 				VirtualClusterID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/virtualClusters/vc-5555555-6666-7777-8888-999999999999"),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("GP_Gen5"),
		// 				Capacity: to.Ptr[int32](8),
		// 				Family: to.Ptr("Gen5"),
		// 				Tier: to.Ptr("GeneralPurpose"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testinstance2"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Test1/providers/Microsoft.Sql/managedInstances/testinstance2"),
		// 			Location: to.Ptr("japaneast"),
		// 			Properties: &armsql.ManagedInstanceProperties{
		// 				AdministratorLogin: to.Ptr("dummylogin"),
		// 				AuthenticationMetadata: to.Ptr(armsql.AuthMetadataLookupModesWindows),
		// 				Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-01T01:01:01.011Z"); return t}()),
		// 				CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				DatabaseFormat: to.Ptr(armsql.ManagedInstanceDatabaseFormatSQLServer2022),
		// 				DNSZone: to.Ptr("2c3d1bdae3412"),
		// 				ExternalGovernanceStatus: to.Ptr(armsql.ExternalGovernanceStatusEnabled),
		// 				FullyQualifiedDomainName: to.Ptr("testinstance2.2c3d1bdae3412.database.windows.net"),
		// 				HybridSecondaryUsage: to.Ptr(armsql.HybridSecondaryUsagePassive),
		// 				LicenseType: to.Ptr(armsql.ManagedInstanceLicenseType("Full")),
		// 				MaintenanceConfigurationID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1"),
		// 				MinimalTLSVersion: to.Ptr("1.2"),
		// 				PricingModel: to.Ptr(armsql.FreemiumTypeRegular),
		// 				ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
		// 				ProxyOverride: to.Ptr(armsql.ManagedInstanceProxyOverrideDefault),
		// 				PublicDataEndpointEnabled: to.Ptr(false),
		// 				RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				State: to.Ptr("Ready"),
		// 				StorageSizeInGB: to.Ptr[int32](1024),
		// 				SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/subnet2"),
		// 				VCores: to.Ptr[int32](16),
		// 				VirtualClusterID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/virtualClusters/vc-5555555-6666-7777-8888-444444444444"),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("BC_Gen5"),
		// 				Capacity: to.Ptr[int32](16),
		// 				Family: to.Ptr("Gen5"),
		// 				Tier: to.Ptr("BusinessCritical"),
		// 			},
		// 	}},
		// }
	}
}
