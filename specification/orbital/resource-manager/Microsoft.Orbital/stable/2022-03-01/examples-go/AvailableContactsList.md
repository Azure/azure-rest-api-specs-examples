```go
package armorbital_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/AvailableContactsList.json
func ExampleSpacecraftsClient_BeginListAvailableContacts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armorbital.NewSpacecraftsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginListAvailableContacts(ctx,
		"rgName",
		"AQUA",
		armorbital.ContactParameters{
			ContactProfile: &armorbital.ContactParametersContactProfile{
				ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/Microsoft.Orbital/contactProfiles/AQUA_DIRECTPLAYBACK_WITH_UPLINK"),
			},
			EndTime:           to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-17T23:49:40.00Z"); return t }()),
			GroundStationName: to.Ptr("westus_gs1"),
			StartTime:         to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-16T05:40:21.00Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	for res.More() {
		nextResult, err := res.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Forbital%2Farmorbital%2Fv1.0.0/sdk/resourcemanager/orbital/armorbital/README.md) on how to add the SDK to your project and authenticate.
