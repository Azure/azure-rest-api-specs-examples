Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicefabric%2Farmservicefabric%2Fv0.4.0/sdk/resourcemanager/servicefabric/armservicefabric/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicefabric_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterPatchOperation_example.json
func ExampleClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armservicefabric.ClusterUpdateParameters{
			Properties: &armservicefabric.ClusterPropertiesUpdateParameters{
				EventStoreServiceEnabled: to.BoolPtr(true),
				NodeTypes: []*armservicefabric.NodeTypeDescription{
					{
						Name: to.StringPtr("<name>"),
						ApplicationPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Int32Ptr(30000),
							StartPort: to.Int32Ptr(20000),
						},
						ClientConnectionEndpointPort: to.Int32Ptr(19000),
						DurabilityLevel:              armservicefabric.DurabilityLevel("Bronze").ToPtr(),
						EphemeralPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Int32Ptr(64000),
							StartPort: to.Int32Ptr(49000),
						},
						HTTPGatewayEndpointPort: to.Int32Ptr(19007),
						IsPrimary:               to.BoolPtr(true),
						VMInstanceCount:         to.Int32Ptr(5),
					},
					{
						Name: to.StringPtr("<name>"),
						ApplicationPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Int32Ptr(2000),
							StartPort: to.Int32Ptr(1000),
						},
						ClientConnectionEndpointPort: to.Int32Ptr(0),
						DurabilityLevel:              armservicefabric.DurabilityLevel("Bronze").ToPtr(),
						EphemeralPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Int32Ptr(4000),
							StartPort: to.Int32Ptr(3000),
						},
						HTTPGatewayEndpointPort: to.Int32Ptr(0),
						IsPrimary:               to.BoolPtr(false),
						VMInstanceCount:         to.Int32Ptr(3),
					}},
				ReliabilityLevel:              armservicefabric.ReliabilityLevel("Bronze").ToPtr(),
				UpgradeMode:                   armservicefabric.UpgradeMode("Automatic").ToPtr(),
				UpgradePauseEndTimestampUTC:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-25T22:00:00Z"); return t }()),
				UpgradePauseStartTimestampUTC: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-21T22:00:00Z"); return t }()),
				UpgradeWave:                   armservicefabric.ClusterUpgradeCadence("Wave").ToPtr(),
			},
			Tags: map[string]*string{
				"a": to.StringPtr("b"),
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
	log.Printf("Response result: %#v\n", res.ClustersClientUpdateResult)
}
```
