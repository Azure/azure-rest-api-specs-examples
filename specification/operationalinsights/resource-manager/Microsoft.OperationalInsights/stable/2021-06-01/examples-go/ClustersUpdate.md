Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Foperationalinsights%2Farmoperationalinsights%2Fv0.5.0/sdk/resourcemanager/operationalinsights/armoperationalinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armoperationalinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersUpdate.json
func ExampleClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armoperationalinsights.NewClustersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armoperationalinsights.ClusterPatch{
			Identity: &armoperationalinsights.Identity{
				Type: to.Ptr(armoperationalinsights.IdentityTypeUserAssigned),
				UserAssignedIdentities: map[string]*armoperationalinsights.UserIdentityProperties{
					"/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/oiautorest6685/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity": {},
				},
			},
			Properties: &armoperationalinsights.ClusterPatchProperties{
				KeyVaultProperties: &armoperationalinsights.KeyVaultProperties{
					KeyName:     to.Ptr("<key-name>"),
					KeyRsaSize:  to.Ptr[int32](1024),
					KeyVaultURI: to.Ptr("<key-vault-uri>"),
					KeyVersion:  to.Ptr("<key-version>"),
				},
			},
			SKU: &armoperationalinsights.ClusterSKU{
				Name:     to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
				Capacity: to.Ptr(armoperationalinsights.CapacityTenHundred),
			},
			Tags: map[string]*string{
				"tag1": to.Ptr("val1"),
			},
		},
		&armoperationalinsights.ClustersClientBeginUpdateOptions{ResumeToken: ""})
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
```
