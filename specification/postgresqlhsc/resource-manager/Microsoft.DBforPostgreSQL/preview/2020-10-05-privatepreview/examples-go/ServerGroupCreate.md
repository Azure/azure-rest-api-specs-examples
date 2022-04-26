Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpostgresqlhsc%2Farmpostgresqlhsc%2Fv0.4.0/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupCreate.json
func ExampleServerGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpostgresqlhsc.NewServerGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-group-name>",
		armpostgresqlhsc.ServerGroup{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"ElasticServer": to.Ptr("1"),
			},
			Properties: &armpostgresqlhsc.ServerGroupProperties{
				AdministratorLogin:         to.Ptr("<administrator-login>"),
				AdministratorLoginPassword: to.Ptr("<administrator-login-password>"),
				AvailabilityZone:           to.Ptr("<availability-zone>"),
				BackupRetentionDays:        to.Ptr[int32](35),
				CitusVersion:               to.Ptr(armpostgresqlhsc.CitusVersionNine5),
				DelegatedSubnetArguments: &armpostgresqlhsc.ServerGroupPropertiesDelegatedSubnetArguments{
					SubnetArmResourceID: to.Ptr("<subnet-arm-resource-id>"),
				},
				EnableMx:          to.Ptr(true),
				EnableZfs:         to.Ptr(false),
				PostgresqlVersion: to.Ptr(armpostgresqlhsc.PostgreSQLVersionTwelve),
				PrivateDNSZoneArguments: &armpostgresqlhsc.ServerGroupPropertiesPrivateDNSZoneArguments{
					PrivateDNSZoneArmResourceID: to.Ptr("<private-dnszone-arm-resource-id>"),
				},
				ServerRoleGroups: []*armpostgresqlhsc.ServerRoleGroup{
					{
						EnableHa:         to.Ptr(true),
						ServerEdition:    to.Ptr(armpostgresqlhsc.ServerEditionGeneralPurpose),
						StorageQuotaInMb: to.Ptr[int64](524288),
						VCores:           to.Ptr[int64](4),
						Name:             to.Ptr("<name>"),
						Role:             to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
						ServerCount:      to.Ptr[int32](1),
					},
					{
						EnableHa:         to.Ptr(false),
						ServerEdition:    to.Ptr(armpostgresqlhsc.ServerEditionMemoryOptimized),
						StorageQuotaInMb: to.Ptr[int64](524288),
						VCores:           to.Ptr[int64](4),
						Name:             to.Ptr("<name>"),
						Role:             to.Ptr(armpostgresqlhsc.ServerRoleWorker),
						ServerCount:      to.Ptr[int32](3),
					}},
				StandbyAvailabilityZone: to.Ptr("<standby-availability-zone>"),
			},
		},
		&armpostgresqlhsc.ServerGroupsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
