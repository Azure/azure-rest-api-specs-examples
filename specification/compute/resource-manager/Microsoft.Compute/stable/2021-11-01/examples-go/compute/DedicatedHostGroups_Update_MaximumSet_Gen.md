Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.5.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/DedicatedHostGroups_Update_MaximumSet_Gen.json
func ExampleDedicatedHostGroupsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewDedicatedHostGroupsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<host-group-name>",
		armcompute.DedicatedHostGroupUpdate{
			Tags: map[string]*string{
				"key9921": to.StringPtr("aaaaaaaaaa"),
			},
			Properties: &armcompute.DedicatedHostGroupProperties{
				InstanceView: &armcompute.DedicatedHostGroupInstanceView{
					Hosts: []*armcompute.DedicatedHostInstanceViewWithName{
						{
							AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
								AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
									{
										Count:  to.Float64Ptr(26),
										VMSize: to.StringPtr("<vmsize>"),
									}},
							},
							Statuses: []*armcompute.InstanceViewStatus{
								{
									Code:          to.StringPtr("<code>"),
									DisplayStatus: to.StringPtr("<display-status>"),
									Level:         armcompute.StatusLevelTypesInfo.ToPtr(),
									Message:       to.StringPtr("<message>"),
									Time:          to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
								}},
						}},
				},
				PlatformFaultDomainCount:  to.Int32Ptr(3),
				SupportAutomaticPlacement: to.BoolPtr(true),
			},
			Zones: []*string{
				to.StringPtr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DedicatedHostGroupsClientUpdateResult)
}
```
