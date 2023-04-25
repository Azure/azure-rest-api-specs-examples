package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPApplicationServerInstances_List.json
func ExampleSAPApplicationServerInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSAPApplicationServerInstancesClient().NewListPager("test-rg", "X00", nil)
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
		// page.SAPApplicationServerInstanceList = armworkloads.SAPApplicationServerInstanceList{
		// 	Value: []*armworkloads.SAPApplicationServerInstance{
		// 		{
		// 			Name: to.Ptr("app01"),
		// 			Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances"),
		// 			ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/applicationInstances/app01"),
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
		// 			Properties: &armworkloads.SAPApplicationServerProperties{
		// 				GatewayPort: to.Ptr[int64](3300),
		// 				Health: to.Ptr(armworkloads.SAPHealthStateUnknown),
		// 				Hostname: to.Ptr("vh-nw1"),
		// 				IcmHTTPPort: to.Ptr[int64](3312),
		// 				IcmHTTPSPort: to.Ptr[int64](3313),
		// 				InstanceNo: to.Ptr("00"),
		// 				IPAddress: to.Ptr("10.0.0.5"),
		// 				KernelPatch: to.Ptr("patch 300"),
		// 				KernelVersion: to.Ptr("777"),
		// 				ProvisioningState: to.Ptr(armworkloads.SapVirtualInstanceProvisioningStateSucceeded),
		// 				Status: to.Ptr(armworkloads.SAPVirtualInstanceStatus("Unknown")),
		// 				VMDetails: []*armworkloads.ApplicationServerVMDetails{
		// 					{
		// 						Type: to.Ptr(armworkloads.ApplicationServerVirtualMachineTypeActive),
		// 						VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/app01-vm"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("app02"),
		// 			Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances"),
		// 			ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/applicationInstances/app02"),
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
		// 			Properties: &armworkloads.SAPApplicationServerProperties{
		// 				GatewayPort: to.Ptr[int64](3300),
		// 				Health: to.Ptr(armworkloads.SAPHealthStateUnknown),
		// 				Hostname: to.Ptr("vh-nw1"),
		// 				IcmHTTPPort: to.Ptr[int64](3312),
		// 				IcmHTTPSPort: to.Ptr[int64](3313),
		// 				InstanceNo: to.Ptr("01"),
		// 				IPAddress: to.Ptr("10.0.0.5"),
		// 				KernelPatch: to.Ptr("patch 300"),
		// 				KernelVersion: to.Ptr("777"),
		// 				ProvisioningState: to.Ptr(armworkloads.SapVirtualInstanceProvisioningStateSucceeded),
		// 				Status: to.Ptr(armworkloads.SAPVirtualInstanceStatus("Unknown")),
		// 				VMDetails: []*armworkloads.ApplicationServerVMDetails{
		// 					{
		// 						Type: to.Ptr(armworkloads.ApplicationServerVirtualMachineTypeActive),
		// 						VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/app01-vm"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
