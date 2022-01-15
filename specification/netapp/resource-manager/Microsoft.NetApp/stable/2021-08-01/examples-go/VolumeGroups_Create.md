Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetapp%2Farmnetapp%2Fv0.2.0/sdk/resourcemanager/netapp/armnetapp/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetapp_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp"
)

// x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/VolumeGroups_Create.json
func ExampleVolumeGroupsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetapp.NewVolumeGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<volume-group-name>",
		armnetapp.VolumeGroupDetails{
			Location: to.StringPtr("<location>"),
			Properties: &armnetapp.VolumeGroupProperties{
				GroupMetaData: &armnetapp.VolumeGroupMetaData{
					ApplicationIdentifier: to.StringPtr("<application-identifier>"),
					ApplicationType:       armnetapp.ApplicationType("SAP-HANA").ToPtr(),
					DeploymentSpecID:      to.StringPtr("<deployment-spec-id>"),
					GroupDescription:      to.StringPtr("<group-description>"),
				},
				Volumes: []*armnetapp.VolumeGroupVolumeProperties{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetapp.VolumeProperties{
							CapacityPoolResourceID:  to.StringPtr("<capacity-pool-resource-id>"),
							CreationToken:           to.StringPtr("<creation-token>"),
							ProximityPlacementGroup: to.StringPtr("<proximity-placement-group>"),
							ServiceLevel:            armnetapp.ServiceLevel("Premium").ToPtr(),
							SubnetID:                to.StringPtr("<subnet-id>"),
							ThroughputMibps:         to.Float32Ptr(10),
							UsageThreshold:          to.Int64Ptr(107374182400),
							VolumeSpecName:          to.StringPtr("<volume-spec-name>"),
						},
					},
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetapp.VolumeProperties{
							CapacityPoolResourceID:  to.StringPtr("<capacity-pool-resource-id>"),
							CreationToken:           to.StringPtr("<creation-token>"),
							ProximityPlacementGroup: to.StringPtr("<proximity-placement-group>"),
							ServiceLevel:            armnetapp.ServiceLevel("Premium").ToPtr(),
							SubnetID:                to.StringPtr("<subnet-id>"),
							ThroughputMibps:         to.Float32Ptr(10),
							UsageThreshold:          to.Int64Ptr(107374182400),
							VolumeSpecName:          to.StringPtr("<volume-spec-name>"),
						},
					},
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetapp.VolumeProperties{
							CapacityPoolResourceID:  to.StringPtr("<capacity-pool-resource-id>"),
							CreationToken:           to.StringPtr("<creation-token>"),
							ProximityPlacementGroup: to.StringPtr("<proximity-placement-group>"),
							ServiceLevel:            armnetapp.ServiceLevel("Premium").ToPtr(),
							SubnetID:                to.StringPtr("<subnet-id>"),
							ThroughputMibps:         to.Float32Ptr(10),
							UsageThreshold:          to.Int64Ptr(107374182400),
							VolumeSpecName:          to.StringPtr("<volume-spec-name>"),
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
