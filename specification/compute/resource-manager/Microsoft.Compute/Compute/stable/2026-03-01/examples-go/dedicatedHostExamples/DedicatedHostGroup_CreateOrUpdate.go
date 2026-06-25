package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/dedicatedHostExamples/DedicatedHostGroup_CreateOrUpdate.json
func ExampleDedicatedHostGroupsClient_CreateOrUpdate_createOrUpdateADedicatedHostGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHostGroupsClient().CreateOrUpdate(ctx, "myResourceGroup", "myDedicatedHostGroup", armcompute.DedicatedHostGroup{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"department": to.Ptr("finance"),
		},
		Zones: []*string{
			to.Ptr("1"),
		},
		Properties: &armcompute.DedicatedHostGroupProperties{
			PlatformFaultDomainCount:  to.Ptr[int32](3),
			SupportAutomaticPlacement: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.DedicatedHostGroupsClientCreateOrUpdateResponse{
	// 	DedicatedHostGroup: armcompute.DedicatedHostGroup{
	// 		Properties: &armcompute.DedicatedHostGroupProperties{
	// 			PlatformFaultDomainCount: to.Ptr[int32](3),
	// 			SupportAutomaticPlacement: to.Ptr(true),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"department": to.Ptr("finance"),
	// 			"owner": to.Ptr("myCompany"),
	// 		},
	// 		Name: to.Ptr("myDedicatedHostGroup"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 		},
	// 	},
	// }
}
