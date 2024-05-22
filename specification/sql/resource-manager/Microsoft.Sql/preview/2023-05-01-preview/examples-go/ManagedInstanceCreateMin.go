package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/ManagedInstanceCreateMin.json
func ExampleManagedInstancesClient_BeginCreateOrUpdate_createManagedInstanceWithMinimalProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedInstancesClient().BeginCreateOrUpdate(ctx, "testrg", "testinstance", armsql.ManagedInstance{
		Location: to.Ptr("Japan East"),
		Properties: &armsql.ManagedInstanceProperties{
			AdministratorLogin:         to.Ptr("dummylogin"),
			AdministratorLoginPassword: to.Ptr("PLACEHOLDER"),
			LicenseType:                to.Ptr(armsql.ManagedInstanceLicenseTypeLicenseIncluded),
			StorageSizeInGB:            to.Ptr[int32](1024),
			SubnetID:                   to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
			VCores:                     to.Ptr[int32](8),
		},
		SKU: &armsql.SKU{
			Name: to.Ptr("GP_Gen5"),
			Tier: to.Ptr("GeneralPurpose"),
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
	// 	ID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance"),
	// 	Location: to.Ptr("japaneast"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armsql.ManagedInstanceProperties{
	// 		AdministratorLogin: to.Ptr("dummylogin"),
	// 		Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
	// 		CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		DatabaseFormat: to.Ptr(armsql.ManagedInstanceDatabaseFormatSQLServer2022),
	// 		DNSZone: to.Ptr("1b4e2caff2530"),
	// 		ExternalGovernanceStatus: to.Ptr(armsql.ExternalGovernanceStatusEnabled),
	// 		FullyQualifiedDomainName: to.Ptr("testinstance.1b4e2caff2530.database.windows.net"),
	// 		HybridSecondaryUsage: to.Ptr(armsql.HybridSecondaryUsagePassive),
	// 		LicenseType: to.Ptr(armsql.ManagedInstanceLicenseTypeLicenseIncluded),
	// 		ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
	// 		ProxyOverride: to.Ptr(armsql.ManagedInstanceProxyOverrideRedirect),
	// 		PublicDataEndpointEnabled: to.Ptr(false),
	// 		RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		State: to.Ptr("Ready"),
	// 		StorageSizeInGB: to.Ptr[int32](1024),
	// 		SubnetID: to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 		VCores: to.Ptr[int32](8),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("GP_Gen5"),
	// 		Capacity: to.Ptr[int32](8),
	// 		Family: to.Ptr("Gen4"),
	// 		Tier: to.Ptr("GeneralPurpose"),
	// 	},
	// }
}
