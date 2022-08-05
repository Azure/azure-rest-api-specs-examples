package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/dedicatedHostExamples/DedicatedHostGroup_CreateOrUpdate_WithUltraSSD.json
func ExampleDedicatedHostGroupsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewDedicatedHostGroupsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"myResourceGroup",
		"myDedicatedHostGroup",
		armcompute.DedicatedHostGroup{
			Location: to.Ptr("westus"),
			Tags: map[string]*string{
				"department": to.Ptr("finance"),
			},
			Properties: &armcompute.DedicatedHostGroupProperties{
				AdditionalCapabilities: &armcompute.DedicatedHostGroupPropertiesAdditionalCapabilities{
					UltraSSDEnabled: to.Ptr(true),
				},
				PlatformFaultDomainCount:  to.Ptr[int32](3),
				SupportAutomaticPlacement: to.Ptr(true),
			},
			Zones: []*string{
				to.Ptr("1")},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
