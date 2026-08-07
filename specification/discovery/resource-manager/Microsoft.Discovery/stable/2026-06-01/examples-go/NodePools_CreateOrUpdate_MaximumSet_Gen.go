package armdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/discovery/armdiscovery"
)

// Generated from example definition: 2026-06-01/NodePools_CreateOrUpdate_MaximumSet_Gen.json
func ExampleNodePoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdiscovery.NewClientFactory("A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNodePoolsClient().BeginCreateOrUpdate(ctx, "rgdiscovery", "2fda614bbdadfee575", "932c7b8d4ff0c243b8", armdiscovery.NodePool{
		Properties: &armdiscovery.NodePoolProperties{
			SubnetID:                 to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Network/virtualNetworks/virtualnetwork1/subnets/subnet1"),
			VMSize:                   to.Ptr(armdiscovery.VMSizeStandardNC24AdsA100V4),
			MaxNodeCount:             to.Ptr[int32](18),
			MinNodeCount:             to.Ptr[int32](0),
			ScaleSetPriority:         to.Ptr(armdiscovery.ScaleSetPriorityRegular),
			OSDiskSizeGb:             to.Ptr[int32](610),
			ImageCacheLowerThreshold: to.Ptr[int32](4),
			ImageCacheUpperThreshold: to.Ptr[int32](92),
		},
		Tags: map[string]*string{
			"key7034": to.Ptr("obcmoprnvrxxcdbeokgwotr"),
		},
		Location: to.Ptr("uksouth"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdiscovery.NodePoolsClientCreateOrUpdateResponse{
	// 	NodePool: armdiscovery.NodePool{
	// 		ID: to.Ptr("/subscriptions/A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/supercomputers/2fda614bbdadfee575/nodePools/932c7b8d4ff0c243b8"),
	// 		Name: to.Ptr("932c7b8d4ff0c243b8"),
	// 		Tags: map[string]*string{
	// 			"key7034": to.Ptr("obcmoprnvrxxcdbeokgwotr"),
	// 		},
	// 		Location: to.Ptr("uksouth"),
	// 		Type: to.Ptr("Microsoft.Discovery/supercomputers/nodePools"),
	// 		SystemData: &armdiscovery.SystemData{
	// 			CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
	// 			CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
	// 			LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 		},
	// 		Properties: &armdiscovery.NodePoolProperties{
	// 			ProvisioningState: to.Ptr(armdiscovery.ProvisioningStateSucceeded),
	// 			SubnetID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Network/virtualNetworks/virtualnetwork1/subnets/subnet1"),
	// 			VMSize: to.Ptr(armdiscovery.VMSizeStandardNC24AdsA100V4),
	// 			MaxNodeCount: to.Ptr[int32](18),
	// 			MinNodeCount: to.Ptr[int32](0),
	// 			ScaleSetPriority: to.Ptr(armdiscovery.ScaleSetPriorityRegular),
	// 			OSDiskSizeGb: to.Ptr[int32](610),
	// 			ImageCacheLowerThreshold: to.Ptr[int32](4),
	// 			ImageCacheUpperThreshold: to.Ptr[int32](92),
	// 		},
	// 	},
	// }
}
