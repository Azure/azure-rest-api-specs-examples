```go
package armpostgresqlhsc_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupCreate.json
func ExampleServerGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpostgresqlhsc.NewServerGroupsClient("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"TestGroup",
		"hsctestsg",
		armpostgresqlhsc.ServerGroup{
			Location: to.Ptr("westus"),
			Tags: map[string]*string{
				"ElasticServer": to.Ptr("1"),
			},
			Properties: &armpostgresqlhsc.ServerGroupProperties{
				AdministratorLogin:         to.Ptr("citus"),
				AdministratorLoginPassword: to.Ptr("password"),
				AvailabilityZone:           to.Ptr("1"),
				BackupRetentionDays:        to.Ptr[int32](35),
				CitusVersion:               to.Ptr(armpostgresqlhsc.CitusVersionNine5),
				DelegatedSubnetArguments: &armpostgresqlhsc.ServerGroupPropertiesDelegatedSubnetArguments{
					SubnetArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/test-vnet-subnet"),
				},
				EnableMx:          to.Ptr(true),
				EnableZfs:         to.Ptr(false),
				PostgresqlVersion: to.Ptr(armpostgresqlhsc.PostgreSQLVersionTwelve),
				PrivateDNSZoneArguments: &armpostgresqlhsc.ServerGroupPropertiesPrivateDNSZoneArguments{
					PrivateDNSZoneArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/privateDnsZones/test-private-dns-zone"),
				},
				ServerRoleGroups: []*armpostgresqlhsc.ServerRoleGroup{
					{
						EnableHa:         to.Ptr(true),
						ServerEdition:    to.Ptr(armpostgresqlhsc.ServerEditionGeneralPurpose),
						StorageQuotaInMb: to.Ptr[int64](524288),
						VCores:           to.Ptr[int64](4),
						Name:             to.Ptr(""),
						Role:             to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
						ServerCount:      to.Ptr[int32](1),
					},
					{
						EnableHa:         to.Ptr(false),
						ServerEdition:    to.Ptr(armpostgresqlhsc.ServerEditionMemoryOptimized),
						StorageQuotaInMb: to.Ptr[int64](524288),
						VCores:           to.Ptr[int64](4),
						Name:             to.Ptr(""),
						Role:             to.Ptr(armpostgresqlhsc.ServerRoleWorker),
						ServerCount:      to.Ptr[int32](3),
					}},
				StandbyAvailabilityZone: to.Ptr("2"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpostgresqlhsc%2Farmpostgresqlhsc%2Fv0.5.0/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc/README.md) on how to add the SDK to your project and authenticate.
