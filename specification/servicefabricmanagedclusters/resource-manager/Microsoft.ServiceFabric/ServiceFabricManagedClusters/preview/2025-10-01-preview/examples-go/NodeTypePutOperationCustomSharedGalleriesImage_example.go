package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/NodeTypePutOperationCustomSharedGalleriesImage_example.json
func ExampleNodeTypesClient_BeginCreateOrUpdate_putNodeTypeWithSharedGalleriesCustomVMImage() {
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
			DataDiskSizeGB:         to.Ptr[int32](200),
			IsPrimary:              to.Ptr(false),
			VMInstanceCount:        to.Ptr[int32](10),
			VMSharedGalleryImageID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-custom-image/providers/Microsoft.Compute/sharedGalleries/35349201-a0b3-405e-8a23-9f1450984307-SFSHAREDGALLERY/images/TestNoProdContainerDImage/versions/latest"),
			VMSize:                 to.Ptr("Standard_D3"),
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
	// 			IsPrimary: to.Ptr(false),
	// 			IsStateless: to.Ptr(false),
	// 			PlacementProperties: map[string]*string{
	// 			},
	// 			ProvisioningState: to.Ptr(armservicefabricmanagedclusters.ManagedResourceProvisioningStateCreating),
	// 			VMInstanceCount: to.Ptr[int32](10),
	// 			VMSharedGalleryImageID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-custom-image/providers/Microsoft.Compute/sharedGalleries/35349201-a0b3-405e-8a23-9f1450984307-SFSHAREDGALLERY/images/TestNoProdContainerDImage/versions/latest"),
	// 			VMSize: to.Ptr("Standard_D3"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
