package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/DeviceCapacityGet.json
func ExampleDeviceCapacityInfoClient_GetDeviceCapacityInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeviceCapacityInfoClient().GetDeviceCapacityInfo(ctx, "GroupForEdgeAutomation", "testedgedevice", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeviceCapacityInfo = armdataboxedge.DeviceCapacityInfo{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/deviceCapacityInfo"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/deviceCapacityInfo/default"),
	// 	Properties: &armdataboxedge.DeviceCapacityInfoProperties{
	// 		ClusterComputeCapacityInfo: &armdataboxedge.ClusterCapacityViewData{
	// 			Fqdn: to.Ptr("3V78B03CL"),
	// 			GpuCapacity: &armdataboxedge.ClusterGpuCapacity{
	// 				GpuFreeUnitsCount: to.Ptr[int32](1),
	// 				GpuReservedForFailoverUnitsCount: to.Ptr[int32](0),
	// 				GpuTotalUnitsCount: to.Ptr[int32](1),
	// 				GpuType: to.Ptr("NvidiaTeslaT4"),
	// 				GpuUsedUnitsCount: to.Ptr[int32](0),
	// 			},
	// 			LastRefreshedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-28T19:35:23.820Z"); return t}()),
	// 			MemoryCapacity: &armdataboxedge.ClusterMemoryCapacity{
	// 				ClusterFailoverMemoryMb: to.Ptr[float64](0),
	// 				ClusterFragmentationMemoryMb: to.Ptr[float64](0),
	// 				ClusterFreeMemoryMb: to.Ptr[float64](20978.64),
	// 				ClusterHypervReserveMemoryMb: to.Ptr[float64](19599.36),
	// 				ClusterInfraVMMemoryMb: to.Ptr[float64](0),
	// 				ClusterMemoryUsedByVMsMb: to.Ptr[float64](90112),
	// 				ClusterNonFailoverVMMb: to.Ptr[float64](0),
	// 				ClusterTotalMemoryMb: to.Ptr[float64](130690),
	// 				ClusterUsedMemoryMb: to.Ptr[float64](109711.36),
	// 			},
	// 			TotalProvisionedNonHpnCores: to.Ptr[int64](5),
	// 		},
	// 		ClusterStorageCapacityInfo: &armdataboxedge.ClusterStorageViewData{
	// 			ClusterFreeStorageMb: to.Ptr[float64](5284151),
	// 			ClusterTotalStorageMb: to.Ptr[float64](5368642),
	// 		},
	// 		NodeCapacityInfos: map[string]*armdataboxedge.HostCapacity{
	// 			"3V78B03": &armdataboxedge.HostCapacity{
	// 				AvailableGpuCount: to.Ptr[int32](1),
	// 				EffectiveAvailableMemoryMbOnHost: to.Ptr[int64](45818),
	// 				GpuType: to.Ptr("NvidiaTeslaT4"),
	// 				HostName: to.Ptr("3V78B03"),
	// 				NumaNodesData: []*armdataboxedge.NumaNodeData{
	// 					{
	// 						EffectiveAvailableMemoryInMb: to.Ptr[int64](30182),
	// 						FreeVCPUIndexesForHpn: []*int32{
	// 						},
	// 						LogicalCoreCountPerCore: to.Ptr[int32](2),
	// 						NumaNodeIndex: to.Ptr[int32](0),
	// 						TotalMemoryInMb: to.Ptr[int64](65154),
	// 						VCPUIndexesForHpn: []*int32{
	// 						},
	// 						VCPUIndexesForRoot: []*int32{
	// 							to.Ptr[int32](0),
	// 							to.Ptr[int32](1),
	// 							to.Ptr[int32](2),
	// 							to.Ptr[int32](3),
	// 							to.Ptr[int32](4),
	// 							to.Ptr[int32](5),
	// 							to.Ptr[int32](6),
	// 							to.Ptr[int32](7),
	// 							to.Ptr[int32](8),
	// 							to.Ptr[int32](9),
	// 							to.Ptr[int32](10),
	// 							to.Ptr[int32](11),
	// 							to.Ptr[int32](12),
	// 							to.Ptr[int32](13),
	// 							to.Ptr[int32](14),
	// 							to.Ptr[int32](15),
	// 							to.Ptr[int32](16),
	// 							to.Ptr[int32](17),
	// 							to.Ptr[int32](18),
	// 							to.Ptr[int32](19)},
	// 						},
	// 						{
	// 							EffectiveAvailableMemoryInMb: to.Ptr[int64](19493),
	// 							FreeVCPUIndexesForHpn: []*int32{
	// 							},
	// 							LogicalCoreCountPerCore: to.Ptr[int32](2),
	// 							NumaNodeIndex: to.Ptr[int32](1),
	// 							TotalMemoryInMb: to.Ptr[int64](65536),
	// 							VCPUIndexesForHpn: []*int32{
	// 							},
	// 							VCPUIndexesForRoot: []*int32{
	// 								to.Ptr[int32](20),
	// 								to.Ptr[int32](21),
	// 								to.Ptr[int32](22),
	// 								to.Ptr[int32](23),
	// 								to.Ptr[int32](24),
	// 								to.Ptr[int32](25),
	// 								to.Ptr[int32](26),
	// 								to.Ptr[int32](27),
	// 								to.Ptr[int32](28),
	// 								to.Ptr[int32](29),
	// 								to.Ptr[int32](30),
	// 								to.Ptr[int32](31),
	// 								to.Ptr[int32](32),
	// 								to.Ptr[int32](33),
	// 								to.Ptr[int32](34),
	// 								to.Ptr[int32](35),
	// 								to.Ptr[int32](36),
	// 								to.Ptr[int32](37),
	// 								to.Ptr[int32](38),
	// 								to.Ptr[int32](39)},
	// 						}},
	// 						VMUsedMemory: map[string]*armdataboxedge.VMMemory{
	// 							"23ddd5ea-bda9-467a-803f-1a7e2888d1fa": &armdataboxedge.VMMemory{
	// 								CurrentMemoryUsageMB: to.Ptr[int64](24576),
	// 								StartupMemoryMB: to.Ptr[int64](24576),
	// 							},
	// 							"2e867196-e44b-49b7-9cb3-46c25973b762": &armdataboxedge.VMMemory{
	// 								CurrentMemoryUsageMB: to.Ptr[int64](3584),
	// 								StartupMemoryMB: to.Ptr[int64](3584),
	// 							},
	// 							"48034ac1-04ce-48e1-89fd-15a6d71fd623": &armdataboxedge.VMMemory{
	// 								CurrentMemoryUsageMB: to.Ptr[int64](2048),
	// 								StartupMemoryMB: to.Ptr[int64](2048),
	// 							},
	// 							"4c4261f0-eac5-4ff6-95ef-2cb70e09d3f6": &armdataboxedge.VMMemory{
	// 								CurrentMemoryUsageMB: to.Ptr[int64](3584),
	// 								StartupMemoryMB: to.Ptr[int64](3584),
	// 							},
	// 							"6b1817d5-4003-4f17-8423-51c33d9de88c": &armdataboxedge.VMMemory{
	// 								CurrentMemoryUsageMB: to.Ptr[int64](0),
	// 								StartupMemoryMB: to.Ptr[int64](24576),
	// 							},
	// 							"853efe82-58f1-4f94-8f9c-7752dab6767c": &armdataboxedge.VMMemory{
	// 								CurrentMemoryUsageMB: to.Ptr[int64](24576),
	// 								StartupMemoryMB: to.Ptr[int64](24576),
	// 							},
	// 							"c1023a57-d6a0-4a77-9308-48b320bfedbd": &armdataboxedge.VMMemory{
	// 								CurrentMemoryUsageMB: to.Ptr[int64](3584),
	// 								StartupMemoryMB: to.Ptr[int64](3584),
	// 							},
	// 							"dc3d20b2-8581-41cb-9ca3-1bbc1a8a2fe4": &armdataboxedge.VMMemory{
	// 								CurrentMemoryUsageMB: to.Ptr[int64](3584),
	// 								StartupMemoryMB: to.Ptr[int64](3584),
	// 							},
	// 						},
	// 					},
	// 				},
	// 				TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-28T19:35:34.558Z"); return t}()),
	// 			},
	// 		}
}
