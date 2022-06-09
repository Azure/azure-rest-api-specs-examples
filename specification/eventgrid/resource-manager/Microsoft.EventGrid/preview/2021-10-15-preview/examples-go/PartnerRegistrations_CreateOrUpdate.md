```go
package armeventgrid_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerRegistrations_CreateOrUpdate.json
func ExamplePartnerRegistrationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armeventgrid.NewPartnerRegistrationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<partner-registration-name>",
		armeventgrid.PartnerRegistration{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key1": to.Ptr("value1"),
				"key2": to.Ptr("Value2"),
				"key3": to.Ptr("Value3"),
			},
			Properties: &armeventgrid.PartnerRegistrationProperties{
				AuthorizedAzureSubscriptionIDs: []*string{
					to.Ptr("d48566a8-2428-4a6c-8347-9675d09fb851")},
				LogoURI:                        to.Ptr("<logo-uri>"),
				PartnerName:                    to.Ptr("<partner-name>"),
				PartnerResourceTypeDescription: to.Ptr("<partner-resource-type-description>"),
				PartnerResourceTypeDisplayName: to.Ptr("<partner-resource-type-display-name>"),
				PartnerResourceTypeName:        to.Ptr("<partner-resource-type-name>"),
				SetupURI:                       to.Ptr("<setup-uri>"),
			},
		},
		&armeventgrid.PartnerRegistrationsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Feventgrid%2Farmeventgrid%2Fv0.5.0/sdk/resourcemanager/eventgrid/armeventgrid/README.md) on how to add the SDK to your project and authenticate.
