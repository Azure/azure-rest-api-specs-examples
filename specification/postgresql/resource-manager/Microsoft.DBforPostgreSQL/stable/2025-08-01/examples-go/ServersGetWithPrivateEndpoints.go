package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/ServersGetWithPrivateEndpoints.json
func ExampleServersClient_Get_getInformationAboutAnExistingServerThatIsntIntegratedIntoAVirtualNetworkProvidedByCustomerAndHasPrivateEndpointConnections() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServersClient().Get(ctx, "exampleresourcegroup", "exampleserver", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Server = armpostgresqlflexibleservers.Server{
	// 	Name: to.Ptr("exampleserver"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armpostgresqlflexibleservers.ServerProperties{
	// 		AdministratorLogin: to.Ptr("exampleadministratorlogin"),
	// 		AuthConfig: &armpostgresqlflexibleservers.AuthConfig{
	// 			ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.MicrosoftEntraAuthDisabled),
	// 			PasswordAuth: to.Ptr(armpostgresqlflexibleservers.PasswordBasedAuthEnabled),
	// 		},
	// 		AvailabilityZone: to.Ptr("1"),
	// 		Backup: &armpostgresqlflexibleservers.Backup{
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T18:35:22.123Z"); return t}()),
	// 			GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeographicallyRedundantBackupDisabled),
	// 		},
	// 		FullyQualifiedDomainName: to.Ptr("exampleserver.postgres.database.azure.com"),
	// 		HighAvailability: &armpostgresqlflexibleservers.HighAvailability{
	// 			Mode: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityMode("Disabled")),
	// 			State: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityStateNotEnabled),
	// 		},
	// 		MaintenanceWindow: &armpostgresqlflexibleservers.MaintenanceWindow{
	// 			CustomWindow: to.Ptr("Disabled"),
	// 			DayOfWeek: to.Ptr[int32](0),
	// 			StartHour: to.Ptr[int32](0),
	// 			StartMinute: to.Ptr[int32](0),
	// 		},
	// 		MinorVersion: to.Ptr("5"),
	// 		Network: &armpostgresqlflexibleservers.Network{
	// 			PublicNetworkAccess: to.Ptr(armpostgresqlflexibleservers.ServerPublicNetworkAccessStateEnabled),
	// 		},
	// 		PrivateEndpointConnections: []*armpostgresqlflexibleservers.PrivateEndpointConnection{
	// 			{
	// 				Name: to.Ptr("exampleprivateendpoint.40c899c7-5847-493e-9c9e-d0a0c90aaf92"),
	// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/privateEndpointConnections/exampleprivateendpoint.40c899c7-5847-493e-9c9e-d0a0c90aaf92"),
	// 				Properties: &armpostgresqlflexibleservers.PrivateEndpointConnectionProperties{
	// 					GroupIDs: []*string{
	// 						to.Ptr("postgresqlServer")},
	// 						PrivateEndpoint: &armpostgresqlflexibleservers.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleprivateendpointresourcegroup/providers/Microsoft.Network/privateEndpoints/exampleprivateendpoint"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armpostgresqlflexibleservers.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("Auto-Approved"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr(armpostgresqlflexibleservers.PrivateEndpointServiceConnectionStatusApproved),
	// 						},
	// 						ProvisioningState: to.Ptr(armpostgresqlflexibleservers.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 					},
	// 			}},
	// 			State: to.Ptr(armpostgresqlflexibleservers.ServerStateReady),
	// 			Storage: &armpostgresqlflexibleservers.Storage{
	// 				AutoGrow: to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowDisabled),
	// 				Iops: to.Ptr[int32](2300),
	// 				StorageSizeGB: to.Ptr[int32](512),
	// 				Tier: to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTierP20),
	// 			},
	// 			Version: to.Ptr(armpostgresqlflexibleservers.PostgresMajorVersionSeventeen),
	// 		},
	// 		SKU: &armpostgresqlflexibleservers.SKU{
	// 			Name: to.Ptr("Standard_D4ds_v5"),
	// 			Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
	// 		},
	// 	}
}
