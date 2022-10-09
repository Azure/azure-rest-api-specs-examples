package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/dedicatedHostExamples/DedicatedHosts_Update_MaximumSet_Gen.json
func ExampleDedicatedHostsClient_BeginUpdate_dedicatedHostsUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewDedicatedHostsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "rgcompute", "aaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaa", armcompute.DedicatedHostUpdate{
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
							VMSize: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
						}},
				},
				Statuses: []*armcompute.InstanceViewStatus{
					{
						Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
						DisplayStatus: to.Ptr("aaaaaa"),
						Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
						Message:       to.Ptr("a"),
						Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
					}},
			},
			LicenseType:         to.Ptr(armcompute.DedicatedHostLicenseTypesWindowsServerHybrid),
			PlatformFaultDomain: to.Ptr[int32](1),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
