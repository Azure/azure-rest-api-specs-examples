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

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/InstanceFailoverGroupCreateOrUpdate.json
func ExampleInstanceFailoverGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewInstanceFailoverGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<location-name>",
		"<failover-group-name>",
		armsql.InstanceFailoverGroup{
			Properties: &armsql.InstanceFailoverGroupProperties{
				ManagedInstancePairs: []*armsql.ManagedInstancePairInfo{
					{
						PartnerManagedInstanceID: to.StringPtr("<partner-managed-instance-id>"),
						PrimaryManagedInstanceID: to.StringPtr("<primary-managed-instance-id>"),
					}},
				PartnerRegions: []*armsql.PartnerRegionInfo{
					{
						Location: to.StringPtr("<location>"),
					}},
				ReadOnlyEndpoint: &armsql.InstanceFailoverGroupReadOnlyEndpoint{
					FailoverPolicy: armsql.ReadOnlyEndpointFailoverPolicy("Disabled").ToPtr(),
				},
				ReadWriteEndpoint: &armsql.InstanceFailoverGroupReadWriteEndpoint{
					FailoverPolicy:                         armsql.ReadWriteEndpointFailoverPolicy("Automatic").ToPtr(),
					FailoverWithDataLossGracePeriodMinutes: to.Int32Ptr(480),
				},
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
	log.Printf("Response result: %#v\n", res.InstanceFailoverGroupsClientCreateOrUpdateResult)
}
```
