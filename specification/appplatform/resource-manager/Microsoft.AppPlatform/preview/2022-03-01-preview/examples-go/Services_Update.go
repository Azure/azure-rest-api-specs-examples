package armappplatform_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-03-01-preview/examples/Services_Update.json
func ExampleServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armappplatform.NewServicesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		armappplatform.ServiceResource{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			Properties: &armappplatform.ClusterResourceProperties{},
			SKU: &armappplatform.SKU{
				Name: to.Ptr("<name>"),
				Tier: to.Ptr("<tier>"),
			},
		},
		&armappplatform.ServicesClientBeginUpdateOptions{ResumeToken: ""})
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
