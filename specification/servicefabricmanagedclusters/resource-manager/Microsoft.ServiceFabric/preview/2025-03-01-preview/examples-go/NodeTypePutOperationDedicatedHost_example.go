package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-03-01-preview/NodeTypePutOperationDedicatedHost_example.json
func ExampleNodeTypesClient_BeginCreateOrUpdate_putNodeTypeWithDedicatedHosts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNodeTypesClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", "BE", armservicefabricmanagedclusters.NodeType{
		Properties: &armservicefabricmanagedclusters.NodeTypeProperties{
			Capacities:          map[string]*string{},
			DataDiskSizeGB:      to.Ptr[int32](200),
			DataDiskType:        to.Ptr(armservicefabricmanagedclusters.DiskTypeStandardSSDLRS),
			HostGroupID:         to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testhostgroupRG/providers/Microsoft.Compute/hostGroups/testHostGroup"),
			IsPrimary:           to.Ptr(false),
			PlacementProperties: map[string]*string{},
			VMImageOffer:        to.Ptr("WindowsServer"),
			VMImagePublisher:    to.Ptr("MicrosoftWindowsServer"),
			VMImageSKU:          to.Ptr("2019-Datacenter"),
			VMImageVersion:      to.Ptr("latest"),
			VMInstanceCount:     to.Ptr[int32](10),
			VMSize:              to.Ptr("Standard_D8s_v3"),
			Zones: []*string{
				to.Ptr("1"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicefabricmanagedclusters.NodeTypesClientCreateOrUpdateResponse{
	// 	NodeType: &armservicefabricmanagedclusters.NodeType{
	// 		Name: to.Ptr("BE"),
	// 		Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/nodeTypes"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedClusters/myCluster/nodeTypes/BE"),
	// 		Properties: &armservicefabricmanagedclusters.NodeTypeProperties{
	// 			Capacities: map[string]*string{
	// 			},
	// 			DataDiskSizeGB: to.Ptr[int32](200),
	// 			DataDiskType: to.Ptr(armservicefabricmanagedclusters.DiskTypeStandardSSDLRS),
	// 			HostGroupID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testhostgroupRG/providers/Microsoft.Compute/hostGroups/testHostGroup"),
	// 			IsPrimary: to.Ptr(false),
	// 			IsStateless: to.Ptr(false),
	// 			PlacementProperties: map[string]*string{
	// 			},
	// 			ProvisioningState: to.Ptr(armservicefabricmanagedclusters.ManagedResourceProvisioningStateCreating),
	// 			VMImageOffer: to.Ptr("WindowsServer"),
	// 			VMImagePublisher: to.Ptr("MicrosoftWindowsServer"),
	// 			VMImageSKU: to.Ptr("2019-Datacenter"),
	// 			VMImageVersion: to.Ptr("latest"),
	// 			VMInstanceCount: to.Ptr[int32](10),
	// 			VMSize: to.Ptr("Standard_D8s_v3"),
	// 			Zones: []*string{
	// 				to.Ptr("1"),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
