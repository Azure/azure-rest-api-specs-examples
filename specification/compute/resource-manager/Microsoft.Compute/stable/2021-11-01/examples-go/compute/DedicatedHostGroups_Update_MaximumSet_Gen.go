package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/DedicatedHostGroups_Update_MaximumSet_Gen.json
func ExampleDedicatedHostGroupsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewDedicatedHostGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<host-group-name>",
		armcompute.DedicatedHostGroupUpdate{
			Tags: map[string]*string{
				"key9921": to.Ptr("aaaaaaaaaa"),
			},
			Properties: &armcompute.DedicatedHostGroupProperties{
				InstanceView: &armcompute.DedicatedHostGroupInstanceView{
					Hosts: []*armcompute.DedicatedHostInstanceViewWithName{
						{
							AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
								AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
									{
										Count:  to.Ptr[float64](26),
										VMSize: to.Ptr("<vmsize>"),
									}},
							},
							Statuses: []*armcompute.InstanceViewStatus{
								{
									Code:          to.Ptr("<code>"),
									DisplayStatus: to.Ptr("<display-status>"),
									Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
									Message:       to.Ptr("<message>"),
									Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
								}},
						}},
				},
				PlatformFaultDomainCount:  to.Ptr[int32](3),
				SupportAutomaticPlacement: to.Ptr(true),
			},
			Zones: []*string{
				to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
