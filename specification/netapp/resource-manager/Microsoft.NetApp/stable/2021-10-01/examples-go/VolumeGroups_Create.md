Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetapp%2Farmnetapp%2Fv0.4.0/sdk/resourcemanager/netapp/armnetapp/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/VolumeGroups_Create.json
func ExampleVolumeGroupsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetapp.NewVolumeGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<volume-group-name>",
		armnetapp.VolumeGroupDetails{
			Location: to.Ptr("<location>"),
			Properties: &armnetapp.VolumeGroupProperties{
				GroupMetaData: &armnetapp.VolumeGroupMetaData{
					ApplicationIdentifier: to.Ptr("<application-identifier>"),
					ApplicationType:       to.Ptr(armnetapp.ApplicationTypeSAPHANA),
					DeploymentSpecID:      to.Ptr("<deployment-spec-id>"),
					GroupDescription:      to.Ptr("<group-description>"),
				},
				Volumes: []*armnetapp.VolumeGroupVolumeProperties{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetapp.VolumeProperties{
							CapacityPoolResourceID:  to.Ptr("<capacity-pool-resource-id>"),
							CreationToken:           to.Ptr("<creation-token>"),
							ProximityPlacementGroup: to.Ptr("<proximity-placement-group>"),
							ServiceLevel:            to.Ptr(armnetapp.ServiceLevelPremium),
							SubnetID:                to.Ptr("<subnet-id>"),
							ThroughputMibps:         to.Ptr[float32](10),
							UsageThreshold:          to.Ptr[int64](107374182400),
							VolumeSpecName:          to.Ptr("<volume-spec-name>"),
						},
					},
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetapp.VolumeProperties{
							CapacityPoolResourceID:  to.Ptr("<capacity-pool-resource-id>"),
							CreationToken:           to.Ptr("<creation-token>"),
							ProximityPlacementGroup: to.Ptr("<proximity-placement-group>"),
							ServiceLevel:            to.Ptr(armnetapp.ServiceLevelPremium),
							SubnetID:                to.Ptr("<subnet-id>"),
							ThroughputMibps:         to.Ptr[float32](10),
							UsageThreshold:          to.Ptr[int64](107374182400),
							VolumeSpecName:          to.Ptr("<volume-spec-name>"),
						},
					},
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetapp.VolumeProperties{
							CapacityPoolResourceID:  to.Ptr("<capacity-pool-resource-id>"),
							CreationToken:           to.Ptr("<creation-token>"),
							ProximityPlacementGroup: to.Ptr("<proximity-placement-group>"),
							ServiceLevel:            to.Ptr(armnetapp.ServiceLevelPremium),
							SubnetID:                to.Ptr("<subnet-id>"),
							ThroughputMibps:         to.Ptr[float32](10),
							UsageThreshold:          to.Ptr[int64](107374182400),
							VolumeSpecName:          to.Ptr("<volume-spec-name>"),
						},
					}},
			},
		},
		&armnetapp.VolumeGroupsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```
