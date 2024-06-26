package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/WorkloadNetworks_CreateVMGroup.json
func ExampleWorkloadNetworksClient_BeginCreateVMGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadNetworksClient().BeginCreateVMGroup(ctx, "group1", "cloud1", "vmGroup1", armavs.WorkloadNetworkVMGroup{
		Properties: &armavs.WorkloadNetworkVMGroupProperties{
			DisplayName: to.Ptr("vmGroup1"),
			Members: []*string{
				to.Ptr("564d43da-fefc-2a3b-1d92-42855622fa50")},
			Revision: to.Ptr[int64](1),
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
	// res.WorkloadNetworkVMGroup = armavs.WorkloadNetworkVMGroup{
	// 	Name: to.Ptr("vmGroup1"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/workloadNetworks/vmGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/workloadNetworks/default/vmGroups/vmGroup1"),
	// 	Properties: &armavs.WorkloadNetworkVMGroupProperties{
	// 		DisplayName: to.Ptr("vmGroup1"),
	// 		Members: []*string{
	// 			to.Ptr("564d43da-fefc-2a3b-1d92-42855622fa50")},
	// 			Revision: to.Ptr[int64](1),
	// 			Status: to.Ptr(armavs.VMGroupStatusEnumSUCCESS),
	// 		},
	// 	}
}
