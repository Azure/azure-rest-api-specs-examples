```go
package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PlacementPolicies_CreateOrUpdate.json
func ExamplePlacementPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewPlacementPoliciesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"group1",
		"cloud1",
		"cluster1",
		"policy1",
		armavs.PlacementPolicy{
			Properties: &armavs.VMHostPlacementPolicyProperties{
				Type:         to.Ptr(armavs.PlacementPolicyTypeVMHost),
				AffinityType: to.Ptr(armavs.AffinityTypeAntiAffinity),
				HostMembers: []*string{
					to.Ptr("fakehost22.nyc1.kubernetes.center"),
					to.Ptr("fakehost23.nyc1.kubernetes.center"),
					to.Ptr("fakehost24.nyc1.kubernetes.center")},
				VMMembers: []*string{
					to.Ptr("/subscriptions/{subscription-id}/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-128"),
					to.Ptr("/subscriptions/{subscription-id}/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-256")},
			},
		},
		nil)
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Favs%2Farmavs%2Fv1.0.0/sdk/resourcemanager/avs/armavs/README.md) on how to add the SDK to your project and authenticate.
