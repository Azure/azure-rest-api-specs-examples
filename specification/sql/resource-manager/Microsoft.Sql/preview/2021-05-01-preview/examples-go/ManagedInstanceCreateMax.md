Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.3.1/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceCreateMax.json
func ExampleManagedInstancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewManagedInstancesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<managed-instance-name>",
		armsql.ManagedInstance{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"tagKey1": to.StringPtr("TagValue1"),
			},
			Properties: &armsql.ManagedInstanceProperties{
				AdministratorLogin:         to.StringPtr("<administrator-login>"),
				AdministratorLoginPassword: to.StringPtr("<administrator-login-password>"),
				Administrators: &armsql.ManagedInstanceExternalAdministrator{
					AzureADOnlyAuthentication: to.BoolPtr(true),
					Login:                     to.StringPtr("<login>"),
					PrincipalType:             armsql.PrincipalType("User").ToPtr(),
					Sid:                       to.StringPtr("<sid>"),
					TenantID:                  to.StringPtr("<tenant-id>"),
				},
				Collation:                        to.StringPtr("<collation>"),
				DNSZonePartner:                   to.StringPtr("<dnszone-partner>"),
				InstancePoolID:                   to.StringPtr("<instance-pool-id>"),
				LicenseType:                      armsql.ManagedInstanceLicenseType("LicenseIncluded").ToPtr(),
				MaintenanceConfigurationID:       to.StringPtr("<maintenance-configuration-id>"),
				MinimalTLSVersion:                to.StringPtr("<minimal-tlsversion>"),
				ProxyOverride:                    armsql.ManagedInstanceProxyOverride("Redirect").ToPtr(),
				PublicDataEndpointEnabled:        to.BoolPtr(false),
				RequestedBackupStorageRedundancy: armsql.BackupStorageRedundancy("Geo").ToPtr(),
				ServicePrincipal: &armsql.ServicePrincipal{
					Type: armsql.ServicePrincipalType("SystemAssigned").ToPtr(),
				},
				StorageSizeInGB: to.Int32Ptr(1024),
				SubnetID:        to.StringPtr("<subnet-id>"),
				TimezoneID:      to.StringPtr("<timezone-id>"),
				VCores:          to.Int32Ptr(8),
			},
			SKU: &armsql.SKU{
				Name: to.StringPtr("<name>"),
				Tier: to.StringPtr("<tier>"),
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
	log.Printf("Response result: %#v\n", res.ManagedInstancesClientCreateOrUpdateResult)
}
```
