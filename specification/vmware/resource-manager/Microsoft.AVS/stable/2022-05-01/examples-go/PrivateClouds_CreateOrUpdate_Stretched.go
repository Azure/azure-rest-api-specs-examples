package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/PrivateClouds_CreateOrUpdate_Stretched.json
func ExamplePrivateCloudsClient_BeginCreateOrUpdate_privateCloudsCreateOrUpdateStretched() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewPrivateCloudsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "group1", "cloud1", armavs.PrivateCloud{
		Location: to.Ptr("eastus2"),
		Tags:     map[string]*string{},
		Properties: &armavs.PrivateCloudProperties{
			Availability: &armavs.AvailabilityProperties{
				SecondaryZone: to.Ptr[int32](2),
				Strategy:      to.Ptr(armavs.AvailabilityStrategyDualZone),
				Zone:          to.Ptr[int32](1),
			},
			ManagementCluster: &armavs.ManagementCluster{
				ClusterSize: to.Ptr[int32](4),
			},
			NetworkBlock: to.Ptr("192.168.48.0/22"),
		},
		SKU: &armavs.SKU{
			Name: to.Ptr("AV36"),
		},
	}, nil)
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
