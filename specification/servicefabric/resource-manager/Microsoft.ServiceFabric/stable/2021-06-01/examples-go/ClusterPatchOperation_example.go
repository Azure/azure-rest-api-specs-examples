package armservicefabric_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterPatchOperation_example.json
func ExampleClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicefabric.NewClustersClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"resRg",
		"myCluster",
		armservicefabric.ClusterUpdateParameters{
			Properties: &armservicefabric.ClusterPropertiesUpdateParameters{
				EventStoreServiceEnabled: to.Ptr(true),
				NodeTypes: []*armservicefabric.NodeTypeDescription{
					{
						Name: to.Ptr("nt1vm"),
						ApplicationPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Ptr[int32](30000),
							StartPort: to.Ptr[int32](20000),
						},
						ClientConnectionEndpointPort: to.Ptr[int32](19000),
						DurabilityLevel:              to.Ptr(armservicefabric.DurabilityLevelBronze),
						EphemeralPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Ptr[int32](64000),
							StartPort: to.Ptr[int32](49000),
						},
						HTTPGatewayEndpointPort: to.Ptr[int32](19007),
						IsPrimary:               to.Ptr(true),
						VMInstanceCount:         to.Ptr[int32](5),
					},
					{
						Name: to.Ptr("testnt1"),
						ApplicationPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Ptr[int32](2000),
							StartPort: to.Ptr[int32](1000),
						},
						ClientConnectionEndpointPort: to.Ptr[int32](0),
						DurabilityLevel:              to.Ptr(armservicefabric.DurabilityLevelBronze),
						EphemeralPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Ptr[int32](4000),
							StartPort: to.Ptr[int32](3000),
						},
						HTTPGatewayEndpointPort: to.Ptr[int32](0),
						IsPrimary:               to.Ptr(false),
						VMInstanceCount:         to.Ptr[int32](3),
					}},
				ReliabilityLevel:              to.Ptr(armservicefabric.ReliabilityLevelBronze),
				UpgradeMode:                   to.Ptr(armservicefabric.UpgradeModeAutomatic),
				UpgradePauseEndTimestampUTC:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-25T22:00:00Z"); return t }()),
				UpgradePauseStartTimestampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-21T22:00:00Z"); return t }()),
				UpgradeWave:                   to.Ptr(armservicefabric.ClusterUpgradeCadence("Wave")),
			},
			Tags: map[string]*string{
				"a": to.Ptr("b"),
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
