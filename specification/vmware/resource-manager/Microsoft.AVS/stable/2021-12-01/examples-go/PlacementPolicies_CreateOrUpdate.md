Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Favs%2Farmavs%2Fv0.1.0/sdk/resourcemanager/avs/armavs/README.md) on how to add the SDK to your project and authenticate.

```go
package armavs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PlacementPolicies_CreateOrUpdate.json
func ExamplePlacementPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewPlacementPoliciesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<cluster-name>",
		"<placement-policy-name>",
		armavs.PlacementPolicy{
			Properties: &armavs.VMHostPlacementPolicyProperties{
				PlacementPolicyProperties: armavs.PlacementPolicyProperties{
					Type: armavs.PlacementPolicyTypeVMHost.ToPtr(),
				},
				AffinityType: armavs.AffinityTypeAntiAffinity.ToPtr(),
				HostMembers: []*string{
					to.StringPtr("fakehost22.nyc1.kubernetes.center"),
					to.StringPtr("fakehost23.nyc1.kubernetes.center"),
					to.StringPtr("fakehost24.nyc1.kubernetes.center")},
				VMMembers: []*string{
					to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-128"),
					to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-256")},
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
	log.Printf("PlacementPolicy.ID: %s\n", *res.ID)
}
```
