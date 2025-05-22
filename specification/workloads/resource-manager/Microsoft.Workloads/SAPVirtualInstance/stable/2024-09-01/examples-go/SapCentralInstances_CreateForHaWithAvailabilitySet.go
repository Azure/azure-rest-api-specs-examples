package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: 2024-09-01/SapCentralInstances_CreateForHaWithAvailabilitySet.json
func ExampleSAPCentralServerInstancesClient_BeginCreate_createSapCentralInstancesForHaSystemWithAvailabilitySet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("6d875e77-e412-4d7d-9af4-8895278b4443", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPCentralServerInstancesClient().BeginCreate(ctx, "test-rg", "X00", "centralServer", armworkloadssapvirtualinstance.SAPCentralServerInstance{
		Location:   to.Ptr("westcentralus"),
		Properties: &armworkloadssapvirtualinstance.SAPCentralServerProperties{},
		Tags:       map[string]*string{},
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
	// res = armworkloadssapvirtualinstance.SAPCentralServerInstancesClientCreateResponse{
	// 	SAPCentralServerInstance: &armworkloadssapvirtualinstance.SAPCentralServerInstance{
	// 		Name: to.Ptr("centralServer"),
	// 		Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances"),
	// 		ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/centralInstances/centralServer"),
	// 		Location: to.Ptr("westcentralus"),
	// 		Properties: &armworkloadssapvirtualinstance.SAPCentralServerProperties{
	// 			EnqueueReplicationServerProperties: &armworkloadssapvirtualinstance.EnqueueReplicationServerProperties{
	// 				ErsVersion: to.Ptr(armworkloadssapvirtualinstance.EnqueueReplicationServerTypeEnqueueReplicator1),
	// 				Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 				Hostname: to.Ptr("vh-ers1"),
	// 				InstanceNo: to.Ptr("00"),
	// 				IPAddress: to.Ptr("10.0.0.5"),
	// 				KernelPatch: to.Ptr("patch 300"),
	// 				KernelVersion: to.Ptr("777"),
	// 			},
	// 			EnqueueServerProperties: &armworkloadssapvirtualinstance.EnqueueServerProperties{
	// 				Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 				Hostname: to.Ptr("vh-ascs1"),
	// 				IPAddress: to.Ptr("10.0.0.5"),
	// 				Port: to.Ptr[int64](3600),
	// 			},
	// 			GatewayServerProperties: &armworkloadssapvirtualinstance.GatewayServerProperties{
	// 				Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 				Port: to.Ptr[int64](3300),
	// 			},
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			InstanceNo: to.Ptr("00"),
	// 			KernelPatch: to.Ptr("patch 300"),
	// 			KernelVersion: to.Ptr("777"),
	// 			LoadBalancerDetails: &armworkloadssapvirtualinstance.LoadBalancerDetails{
	// 				ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Network/loadBalancers/cs-ASCS-loadBalancer"),
	// 			},
	// 			MessageServerProperties: &armworkloadssapvirtualinstance.MessageServerProperties{
	// 				Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 				Hostname: to.Ptr("vh-ascs1"),
	// 				HTTPPort: to.Ptr[int64](8100),
	// 				HTTPSPort: to.Ptr[int64](44400),
	// 				InternalMsPort: to.Ptr[int64](3900),
	// 				IPAddress: to.Ptr("10.0.0.5"),
	// 				MSPort: to.Ptr[int64](3600),
	// 			},
	// 			ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
	// 			Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatusRunning),
	// 			Subnet: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/SAP-E2ETest-rg/providers/Microsoft.Network/virtualNetworks/loop-test-vnet/subnets/loopsubnet"),
	// 			VMDetails: []*armworkloadssapvirtualinstance.CentralServerVMDetails{
	// 				{
	// 					Type: to.Ptr(armworkloadssapvirtualinstance.CentralServerVirtualMachineTypePrimary),
	// 					StorageDetails: []*armworkloadssapvirtualinstance.StorageInformation{
	// 						{
	// 							ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/nfsstorageaccount"),
	// 						},
	// 					},
	// 					VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/cs-vm1"),
	// 				},
	// 				{
	// 					Type: to.Ptr(armworkloadssapvirtualinstance.CentralServerVirtualMachineTypePrimary),
	// 					StorageDetails: []*armworkloadssapvirtualinstance.StorageInformation{
	// 						{
	// 							ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/nfsstorageaccount"),
	// 						},
	// 					},
	// 					VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/cs-vm2"),
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armworkloadssapvirtualinstance.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 			CreatedBy: to.Ptr("user@xyz.com"),
	// 			CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@xyz.com"),
	// 			LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
