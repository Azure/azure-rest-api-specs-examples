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

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/DedicatedHosts_Update_MaximumSet_Gen.json
func ExampleDedicatedHostsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewDedicatedHostsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<host-group-name>",
		"<host-name>",
		armcompute.DedicatedHostUpdate{
			Tags: map[string]*string{
				"key8813": to.StringPtr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
			},
			Properties: &armcompute.DedicatedHostProperties{
				AutoReplaceOnFailure: to.BoolPtr(true),
				InstanceView: &armcompute.DedicatedHostInstanceView{
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
				},
				LicenseType:         armcompute.DedicatedHostLicenseTypesWindowsServerHybrid.ToPtr(),
				PlatformFaultDomain: to.Int32Ptr(1),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DedicatedHostsClientUpdateResult)
}
```
