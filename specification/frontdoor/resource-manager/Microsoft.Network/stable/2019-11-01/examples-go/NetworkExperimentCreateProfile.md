Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ffrontdoor%2Farmfrontdoor%2Fv0.2.1/sdk/resourcemanager/frontdoor/armfrontdoor/README.md) on how to add the SDK to your project and authenticate.

```go
package armfrontdoor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentCreateProfile.json
func ExampleNetworkExperimentProfilesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewNetworkExperimentProfilesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<profile-name>",
		"<resource-group-name>",
		armfrontdoor.Profile{
			Location: to.StringPtr("<location>"),
			Properties: &armfrontdoor.ProfileProperties{
				EnabledState: armfrontdoor.State("Enabled").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.NetworkExperimentProfilesClientCreateOrUpdateResult)
}
```
