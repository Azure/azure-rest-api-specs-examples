Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpostgresqlhsc%2Farmpostgresqlhsc%2Fv0.2.0/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc/README.md) on how to add the SDK to your project and authenticate.

```go
package armpostgresqlhsc_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupCreate.json
func ExampleServerGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpostgresqlhsc.NewServerGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-group-name>",
		armpostgresqlhsc.ServerGroup{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"ElasticServer": to.StringPtr("1"),
			},
			Properties: &armpostgresqlhsc.ServerGroupProperties{
				AdministratorLogin:         to.StringPtr("<administrator-login>"),
				AdministratorLoginPassword: to.StringPtr("<administrator-login-password>"),
				AvailabilityZone:           to.StringPtr("<availability-zone>"),
				BackupRetentionDays:        to.Int32Ptr(35),
				CitusVersion:               armpostgresqlhsc.CitusVersion("9.5").ToPtr(),
				DelegatedSubnetArguments: &armpostgresqlhsc.ServerGroupPropertiesDelegatedSubnetArguments{
					SubnetArmResourceID: to.StringPtr("<subnet-arm-resource-id>"),
				},
				EnableMx:          to.BoolPtr(true),
				EnableZfs:         to.BoolPtr(false),
				PostgresqlVersion: armpostgresqlhsc.PostgreSQLVersion("12").ToPtr(),
				PrivateDNSZoneArguments: &armpostgresqlhsc.ServerGroupPropertiesPrivateDNSZoneArguments{
					PrivateDNSZoneArmResourceID: to.StringPtr("<private-dnszone-arm-resource-id>"),
				},
				ServerRoleGroups: []*armpostgresqlhsc.ServerRoleGroup{
					{
						EnableHa:         to.BoolPtr(true),
						ServerEdition:    armpostgresqlhsc.ServerEdition("GeneralPurpose").ToPtr(),
						StorageQuotaInMb: to.Int64Ptr(524288),
						VCores:           to.Int64Ptr(4),
						Name:             to.StringPtr("<name>"),
						Role:             armpostgresqlhsc.ServerRole("Coordinator").ToPtr(),
						ServerCount:      to.Int32Ptr(1),
					},
					{
						EnableHa:         to.BoolPtr(false),
						ServerEdition:    armpostgresqlhsc.ServerEdition("MemoryOptimized").ToPtr(),
						StorageQuotaInMb: to.Int64Ptr(524288),
						VCores:           to.Int64Ptr(4),
						Name:             to.StringPtr("<name>"),
						Role:             armpostgresqlhsc.ServerRole("Worker").ToPtr(),
						ServerCount:      to.Int32Ptr(3),
					}},
				StandbyAvailabilityZone: to.StringPtr("<standby-availability-zone>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServerGroupsClientCreateOrUpdateResult)
}
```
