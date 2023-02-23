package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPCentralInstances_List.json
func ExampleSAPCentralInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armworkloads.NewSAPCentralInstancesClient("6d875e77-e412-4d7d-9af4-8895278b4443", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("test-rg", "X00", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SAPCentralInstanceList = armworkloads.SAPCentralInstanceList{
		// 	Value: []*armworkloads.SAPCentralServerInstance{
		// 		{
		// 			Name: to.Ptr("centralServer"),
		// 			Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances"),
		// 			ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/centralInstances/centralServer"),
		// 			SystemData: &armworkloads.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@xyz.com"),
		// 				CreatedByType: to.Ptr(armworkloads.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@xyz.com"),
		// 				LastModifiedByType: to.Ptr(armworkloads.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("westcentralus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armworkloads.SAPCentralServerProperties{
		// 				EnqueueReplicationServerProperties: &armworkloads.EnqueueReplicationServerProperties{
		// 					ErsVersion: to.Ptr(armworkloads.EnqueueReplicationServerTypeEnqueueReplicator1),
		// 					Health: to.Ptr(armworkloads.SAPHealthStateHealthy),
		// 					Hostname: to.Ptr("vh-ers1"),
		// 					InstanceNo: to.Ptr("00"),
		// 					IPAddress: to.Ptr("10.0.0.5"),
		// 					KernelPatch: to.Ptr("patch 300"),
		// 					KernelVersion: to.Ptr("777"),
		// 				},
		// 				EnqueueServerProperties: &armworkloads.EnqueueServerProperties{
		// 					Health: to.Ptr(armworkloads.SAPHealthStateHealthy),
		// 					Hostname: to.Ptr("vh-ascs1"),
		// 					IPAddress: to.Ptr("10.0.0.5"),
		// 					Port: to.Ptr[int64](3600),
		// 				},
		// 				GatewayServerProperties: &armworkloads.GatewayServerProperties{
		// 					Health: to.Ptr(armworkloads.SAPHealthStateHealthy),
		// 					Port: to.Ptr[int64](3300),
		// 				},
		// 				Health: to.Ptr(armworkloads.SAPHealthStateUnknown),
		// 				InstanceNo: to.Ptr("00"),
		// 				KernelPatch: to.Ptr("patch 300"),
		// 				KernelVersion: to.Ptr("777"),
		// 				MessageServerProperties: &armworkloads.MessageServerProperties{
		// 					Health: to.Ptr(armworkloads.SAPHealthStateHealthy),
		// 					Hostname: to.Ptr("vh-ascs1"),
		// 					HTTPPort: to.Ptr[int64](8100),
		// 					HTTPSPort: to.Ptr[int64](44400),
		// 					InternalMsPort: to.Ptr[int64](3900),
		// 					IPAddress: to.Ptr("10.0.0.5"),
		// 					MSPort: to.Ptr[int64](3600),
		// 				},
		// 				ProvisioningState: to.Ptr(armworkloads.SapVirtualInstanceProvisioningStateSucceeded),
		// 				Status: to.Ptr(armworkloads.SAPVirtualInstanceStatusRunning),
		// 				VMDetails: []*armworkloads.CentralServerVMDetails{
		// 					{
		// 						Type: to.Ptr(armworkloads.CentralServerVirtualMachineTypePrimary),
		// 						VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/cs-vm"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
