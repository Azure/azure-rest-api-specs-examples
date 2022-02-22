Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Foperationalinsights%2Farmoperationalinsights%2Fv0.3.1/sdk/resourcemanager/operationalinsights/armoperationalinsights/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersUpdate.json
func ExampleClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armoperationalinsights.NewClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armoperationalinsights.ClusterPatch{
			Identity: &armoperationalinsights.Identity{
				Type: armoperationalinsights.IdentityTypeUserAssigned.ToPtr(),
				UserAssignedIdentities: map[string]*armoperationalinsights.UserIdentityProperties{
					"/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/oiautorest6685/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity": {},
				},
			},
			Properties: &armoperationalinsights.ClusterPatchProperties{
				KeyVaultProperties: &armoperationalinsights.KeyVaultProperties{
					KeyName:     to.StringPtr("<key-name>"),
					KeyRsaSize:  to.Int32Ptr(1024),
					KeyVaultURI: to.StringPtr("<key-vault-uri>"),
					KeyVersion:  to.StringPtr("<key-version>"),
				},
			},
			SKU: &armoperationalinsights.ClusterSKU{
				Name:     armoperationalinsights.ClusterSKUNameEnum("CapacityReservation").ToPtr(),
				Capacity: armoperationalinsights.CapacityTenHundred.ToPtr(),
			},
			Tags: map[string]*string{
				"tag1": to.StringPtr("val1"),
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
	log.Printf("Response result: %#v\n", res.ClustersClientUpdateResult)
}
```
