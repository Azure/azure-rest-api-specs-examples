package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/DedicatedHosts_Update_MaximumSet_Gen.json
func ExampleDedicatedHostsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewDedicatedHostsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<host-group-name>",
		"<host-name>",
		armcompute.DedicatedHostUpdate{
			Tags: map[string]*string{
				"key8813": to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
			},
			Properties: &armcompute.DedicatedHostProperties{
				AutoReplaceOnFailure: to.Ptr(true),
				InstanceView: &armcompute.DedicatedHostInstanceView{
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
				},
				LicenseType:         to.Ptr(armcompute.DedicatedHostLicenseTypesWindowsServerHybrid),
				PlatformFaultDomain: to.Ptr[int32](1),
			},
		},
		&armcompute.DedicatedHostsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
