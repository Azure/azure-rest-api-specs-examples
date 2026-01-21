package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/dedicatedHostExamples/DedicatedHostGroup_CreateOrUpdate.json
func ExampleDedicatedHostGroupsClient_CreateOrUpdate_createOrUpdateADedicatedHostGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHostGroupsClient().CreateOrUpdate(ctx, "myResourceGroup", "myDedicatedHostGroup", armcompute.DedicatedHostGroup{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"department": to.Ptr("finance"),
		},
		Properties: &armcompute.DedicatedHostGroupProperties{
			PlatformFaultDomainCount:  to.Ptr[int32](3),
			SupportAutomaticPlacement: to.Ptr(true),
		},
		Zones: []*string{
			to.Ptr("1")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedHostGroup = armcompute.DedicatedHostGroup{
	// 	Name: to.Ptr("myDedicatedHostGroup"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("finance"),
	// 		"owner": to.Ptr("myCompany"),
	// 	},
	// 	Properties: &armcompute.DedicatedHostGroupProperties{
	// 		PlatformFaultDomainCount: to.Ptr[int32](3),
	// 		SupportAutomaticPlacement: to.Ptr(true),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}
