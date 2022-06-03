Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv1.0.0/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceCreateMax.json
func ExampleManagedInstancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewManagedInstancesClient("20D7082A-0FC7-4468-82BD-542694D5042B", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testrg",
		"testinstance",
		armsql.ManagedInstance{
			Location: to.Ptr("Japan East"),
			Tags: map[string]*string{
				"tagKey1": to.Ptr("TagValue1"),
			},
			Properties: &armsql.ManagedInstanceProperties{
				AdministratorLogin:         to.Ptr("dummylogin"),
				AdministratorLoginPassword: to.Ptr("PLACEHOLDER"),
				Administrators: &armsql.ManagedInstanceExternalAdministrator{
					AzureADOnlyAuthentication: to.Ptr(true),
					Login:                     to.Ptr("bob@contoso.com"),
					PrincipalType:             to.Ptr(armsql.PrincipalTypeUser),
					Sid:                       to.Ptr("00000011-1111-2222-2222-123456789111"),
					TenantID:                  to.Ptr("00000011-1111-2222-2222-123456789111"),
				},
				Collation:                        to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
				DNSZonePartner:                   to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance"),
				InstancePoolID:                   to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Sql/instancePools/pool1"),
				LicenseType:                      to.Ptr(armsql.ManagedInstanceLicenseTypeLicenseIncluded),
				MaintenanceConfigurationID:       to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1"),
				MinimalTLSVersion:                to.Ptr("1.2"),
				ProxyOverride:                    to.Ptr(armsql.ManagedInstanceProxyOverrideRedirect),
				PublicDataEndpointEnabled:        to.Ptr(false),
				RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
				ServicePrincipal: &armsql.ServicePrincipal{
					Type: to.Ptr(armsql.ServicePrincipalTypeSystemAssigned),
				},
				StorageSizeInGB: to.Ptr[int32](1024),
				SubnetID:        to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
				TimezoneID:      to.Ptr("UTC"),
				VCores:          to.Ptr[int32](8),
			},
			SKU: &armsql.SKU{
				Name: to.Ptr("GP_Gen5"),
				Tier: to.Ptr("GeneralPurpose"),
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
