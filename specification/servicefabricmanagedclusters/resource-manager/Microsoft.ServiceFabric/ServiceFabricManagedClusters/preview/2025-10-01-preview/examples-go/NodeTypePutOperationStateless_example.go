package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/NodeTypePutOperationStateless_example.json
func ExampleNodeTypesClient_BeginCreateOrUpdate_putAnStatelessNodeTypeWithTemporaryDiskForServiceFabric() {
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
			EnableEncryptionAtHost:  to.Ptr(true),
			IsPrimary:               to.Ptr(false),
			IsStateless:             to.Ptr(true),
			MultiplePlacementGroups: to.Ptr(true),
			UseTempDataDisk:         to.Ptr(true),
			VMExtensions: []*armservicefabricmanagedclusters.VMSSExtension{
				{
					Name: to.Ptr("Microsoft.Azure.Geneva.GenevaMonitoring"),
					Properties: &armservicefabricmanagedclusters.VMSSExtensionProperties{
						Type:                    to.Ptr("GenevaMonitoring"),
						AutoUpgradeMinorVersion: to.Ptr(true),
						Publisher:               to.Ptr("Microsoft.Azure.Geneva"),
						Settings:                map[string]any{},
						TypeHandlerVersion:      to.Ptr("2.0"),
					},
				},
			},
			VMImageOffer:     to.Ptr("WindowsServer"),
			VMImagePublisher: to.Ptr("MicrosoftWindowsServer"),
			VMImageSKU:       to.Ptr("2016-Datacenter-Server-Core"),
			VMImageVersion:   to.Ptr("latest"),
			VMInstanceCount:  to.Ptr[int32](10),
			VMSize:           to.Ptr("Standard_DS3"),
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
	// 			EnableEncryptionAtHost: to.Ptr(true),
	// 			IsPrimary: to.Ptr(false),
	// 			IsStateless: to.Ptr(true),
	// 			MultiplePlacementGroups: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armservicefabricmanagedclusters.ManagedResourceProvisioningStateCreating),
	// 			UseTempDataDisk: to.Ptr(true),
	// 			VMExtensions: []*armservicefabricmanagedclusters.VMSSExtension{
	// 				{
	// 					Name: to.Ptr("Microsoft.Azure.Geneva.GenevaMonitoring"),
	// 					Properties: &armservicefabricmanagedclusters.VMSSExtensionProperties{
	// 						Type: to.Ptr("GenevaMonitoring"),
	// 						AutoUpgradeMinorVersion: to.Ptr(true),
	// 						Publisher: to.Ptr("Microsoft.Azure.Geneva"),
	// 						Settings: map[string]any{
	// 						},
	// 						TypeHandlerVersion: to.Ptr("2.0"),
	// 					},
	// 				},
	// 			},
	// 			VMImageOffer: to.Ptr("WindowsServer"),
	// 			VMImagePublisher: to.Ptr("MicrosoftWindowsServer"),
	// 			VMImageSKU: to.Ptr("2016-Datacenter-Server-Core"),
	// 			VMImageVersion: to.Ptr("latest"),
	// 			VMInstanceCount: to.Ptr[int32](10),
	// 			VMSize: to.Ptr("Standard_DS3"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
