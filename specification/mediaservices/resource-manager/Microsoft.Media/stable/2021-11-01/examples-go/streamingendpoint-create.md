Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv0.6.0/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streamingendpoint-create.json
func ExampleStreamingEndpointsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingEndpointsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<streaming-endpoint-name>",
		armmediaservices.StreamingEndpoint{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
			Properties: &armmediaservices.StreamingEndpointProperties{
				Description: to.Ptr("<description>"),
				AccessControl: &armmediaservices.StreamingEndpointAccessControl{
					Akamai: &armmediaservices.AkamaiAccessControl{
						AkamaiSignatureHeaderAuthenticationKeyList: []*armmediaservices.AkamaiSignatureHeaderAuthenticationKey{
							{
								Base64Key:  to.Ptr("<base64key>"),
								Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2029-12-31T16:00:00-08:00"); return t }()),
								Identifier: to.Ptr("<identifier>"),
							},
							{
								Base64Key:  to.Ptr("<base64key>"),
								Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2030-12-31T16:00:00-08:00"); return t }()),
								Identifier: to.Ptr("<identifier>"),
							}},
					},
					IP: &armmediaservices.IPAccessControl{
						Allow: []*armmediaservices.IPRange{
							{
								Name:    to.Ptr("<name>"),
								Address: to.Ptr("<address>"),
							}},
					},
				},
				AvailabilitySetName: to.Ptr("<availability-set-name>"),
				CdnEnabled:          to.Ptr(false),
				ScaleUnits:          to.Ptr[int32](1),
			},
		},
		&armmediaservices.StreamingEndpointsClientBeginCreateOptions{AutoStart: nil,
			ResumeToken: "",
		})
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
