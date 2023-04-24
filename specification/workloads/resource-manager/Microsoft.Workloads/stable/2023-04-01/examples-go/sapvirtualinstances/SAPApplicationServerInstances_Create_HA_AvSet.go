package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPApplicationServerInstances_Create_HA_AvSet.json
func ExampleSAPApplicationServerInstancesClient_BeginCreate_createSapApplicationServerInstancesForHaSystemWithAvailabilitySet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPApplicationServerInstancesClient().BeginCreate(ctx, "test-rg", "X00", "app01", armworkloads.SAPApplicationServerInstance{
		Location:   to.Ptr("westcentralus"),
		Tags:       map[string]*string{},
		Properties: &armworkloads.SAPApplicationServerProperties{},
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
	// res.SAPApplicationServerInstance = armworkloads.SAPApplicationServerInstance{
	// 	Name: to.Ptr("app01"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/applicationInstances/app01"),
	// 	SystemData: &armworkloads.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@xyz.com"),
	// 		CreatedByType: to.Ptr(armworkloads.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@xyz.com"),
	// 		LastModifiedByType: to.Ptr(armworkloads.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westcentralus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armworkloads.SAPApplicationServerProperties{
	// 		GatewayPort: to.Ptr[int64](3300),
	// 		Health: to.Ptr(armworkloads.SAPHealthStateUnknown),
	// 		Hostname: to.Ptr("vh-nw1"),
	// 		IcmHTTPPort: to.Ptr[int64](3312),
	// 		IcmHTTPSPort: to.Ptr[int64](3313),
	// 		InstanceNo: to.Ptr("01"),
	// 		IPAddress: to.Ptr("10.0.0.5"),
	// 		KernelPatch: to.Ptr("patch 300"),
	// 		KernelVersion: to.Ptr("777"),
	// 		LoadBalancerDetails: &armworkloads.LoadBalancerDetails{
	// 			ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Network/loadBalancers/cs-loadBalancer"),
	// 		},
	// 		ProvisioningState: to.Ptr(armworkloads.SapVirtualInstanceProvisioningStateSucceeded),
	// 		Status: to.Ptr(armworkloads.SAPVirtualInstanceStatusRunning),
	// 		VMDetails: []*armworkloads.ApplicationServerVMDetails{
	// 			{
	// 				Type: to.Ptr(armworkloads.ApplicationServerVirtualMachineTypeActive),
	// 				StorageDetails: []*armworkloads.StorageInformation{
	// 					{
	// 						ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/nfsstorageaccount"),
	// 				}},
	// 				VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/cs-vm1"),
	// 			},
	// 			{
	// 				Type: to.Ptr(armworkloads.ApplicationServerVirtualMachineTypeStandby),
	// 				StorageDetails: []*armworkloads.StorageInformation{
	// 					{
	// 						ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/nfsstorageaccount"),
	// 				}},
	// 				VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/cs-vm2"),
	// 		}},
	// 	},
	// }
}
