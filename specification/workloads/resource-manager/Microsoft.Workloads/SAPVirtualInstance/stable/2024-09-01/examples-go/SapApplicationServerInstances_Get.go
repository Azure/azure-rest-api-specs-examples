package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: 2024-09-01/SapApplicationServerInstances_Get.json
func ExampleSAPApplicationServerInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("6d875e77-e412-4d7d-9af4-8895278b4443", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSAPApplicationServerInstancesClient().Get(ctx, "test-rg", "X00", "app01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientGetResponse{
	// 	SAPApplicationServerInstance: &armworkloadssapvirtualinstance.SAPApplicationServerInstance{
	// 		Name: to.Ptr("app01"),
	// 		Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances"),
	// 		ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/applicationInstances/app01"),
	// 		Location: to.Ptr("westcentralus"),
	// 		Properties: &armworkloadssapvirtualinstance.SAPApplicationServerProperties{
	// 			DispatcherStatus: to.Ptr("Running"),
	// 			GatewayPort: to.Ptr[int64](3300),
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			Hostname: to.Ptr("vh-nw1"),
	// 			IcmHTTPPort: to.Ptr[int64](3312),
	// 			IcmHTTPSPort: to.Ptr[int64](3313),
	// 			InstanceNo: to.Ptr("00"),
	// 			IPAddress: to.Ptr("10.0.0.5"),
	// 			KernelPatch: to.Ptr("patch 300"),
	// 			KernelVersion: to.Ptr("777"),
	// 			ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
	// 			Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatusRunning),
	// 			Subnet: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/SAP-E2ETest-rg/providers/Microsoft.Network/virtualNetworks/loop-test-vnet/subnets/loopsubnet"),
	// 			VMDetails: []*armworkloadssapvirtualinstance.ApplicationServerVMDetails{
	// 				{
	// 					Type: to.Ptr(armworkloadssapvirtualinstance.ApplicationServerVirtualMachineTypeActive),
	// 					VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/app01-vm"),
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
