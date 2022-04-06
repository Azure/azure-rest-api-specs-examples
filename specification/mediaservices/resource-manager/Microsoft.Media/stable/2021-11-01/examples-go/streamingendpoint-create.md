Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv0.4.0/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmediaservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streamingendpoint-create.json
func ExampleStreamingEndpointsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewStreamingEndpointsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<streaming-endpoint-name>",
		armmediaservices.StreamingEndpoint{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
			Properties: &armmediaservices.StreamingEndpointProperties{
				Description: to.StringPtr("<description>"),
				AccessControl: &armmediaservices.StreamingEndpointAccessControl{
					Akamai: &armmediaservices.AkamaiAccessControl{
						AkamaiSignatureHeaderAuthenticationKeyList: []*armmediaservices.AkamaiSignatureHeaderAuthenticationKey{
							{
								Base64Key:  to.StringPtr("<base64key>"),
								Expiration: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2029-12-31T16:00:00-08:00"); return t }()),
								Identifier: to.StringPtr("<identifier>"),
							},
							{
								Base64Key:  to.StringPtr("<base64key>"),
								Expiration: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2030-12-31T16:00:00-08:00"); return t }()),
								Identifier: to.StringPtr("<identifier>"),
							}},
					},
					IP: &armmediaservices.IPAccessControl{
						Allow: []*armmediaservices.IPRange{
							{
								Name:    to.StringPtr("<name>"),
								Address: to.StringPtr("<address>"),
							}},
					},
				},
				AvailabilitySetName: to.StringPtr("<availability-set-name>"),
				CdnEnabled:          to.BoolPtr(false),
				ScaleUnits:          to.Int32Ptr(1),
			},
		},
		&armmediaservices.StreamingEndpointsClientBeginCreateOptions{AutoStart: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.StreamingEndpointsClientCreateResult)
}
```
