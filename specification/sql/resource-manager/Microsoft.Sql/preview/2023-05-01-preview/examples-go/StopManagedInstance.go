package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/StopManagedInstance.json
func ExampleManagedInstancesClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedInstancesClient().BeginStop(ctx, "stoprg", "mitostop", nil)
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
	// 	Name: to.Ptr("mitostop"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/stoprg/providers/Microsoft.Sql/managedInstances/mitostop"),
	// 	Location: to.Ptr("japaneast"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armsql.ManagedInstanceProperties{
	// 		AdministratorLogin: to.Ptr("dummylogin"),
	// 		AuthenticationMetadata: to.Ptr(armsql.AuthMetadataLookupModesWindows),
	// 		Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
	// 		CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		DatabaseFormat: to.Ptr(armsql.ManagedInstanceDatabaseFormatSQLServer2022),
	// 		DNSZone: to.Ptr("1234567891234"),
	// 		ExternalGovernanceStatus: to.Ptr(armsql.ExternalGovernanceStatusEnabled),
	// 		FullyQualifiedDomainName: to.Ptr("mitostop.1234567891234.database.windows.net"),
	// 		HybridSecondaryUsage: to.Ptr(armsql.HybridSecondaryUsagePassive),
	// 		InstancePoolID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/stoprg/providers/Microsoft.Sql/instancePools/instancePool1"),
	// 		LicenseType: to.Ptr(armsql.ManagedInstanceLicenseTypeLicenseIncluded),
	// 		MaintenanceConfigurationID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1"),
	// 		MinimalTLSVersion: to.Ptr("1.2"),
	// 		ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
	// 		ProxyOverride: to.Ptr(armsql.ManagedInstanceProxyOverrideDefault),
	// 		PublicDataEndpointEnabled: to.Ptr(false),
	// 		RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		ServicePrincipal: &armsql.ServicePrincipal{
	// 			Type: to.Ptr(armsql.ServicePrincipalTypeSystemAssigned),
	// 			ClientID: to.Ptr("00000011-1111-2222-2222-123456789111"),
	// 			PrincipalID: to.Ptr("00000011-1111-2222-2222-123456789111"),
	// 			TenantID: to.Ptr("00000011-1111-2222-2222-123456789111"),
	// 		},
	// 		State: to.Ptr("Stopped"),
	// 		StorageSizeInGB: to.Ptr[int32](1024),
	// 		SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/stoprg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 		VCores: to.Ptr[int32](8),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("GP_Gen5"),
	// 		Capacity: to.Ptr[int32](8),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr("GeneralPurpose"),
	// 	},
	// }
}
