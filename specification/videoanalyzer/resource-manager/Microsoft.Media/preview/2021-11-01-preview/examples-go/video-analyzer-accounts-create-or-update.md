Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvideoanalyzer%2Farmvideoanalyzer%2Fv0.4.1/sdk/resourcemanager/videoanalyzer/armvideoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```go
package armvideoanalyzer_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-accounts-create-or-update.json
func ExampleVideoAnalyzersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvideoanalyzer.NewVideoAnalyzersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armvideoanalyzer.VideoAnalyzer{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
			Identity: &armvideoanalyzer.Identity{
				Type: to.Ptr("<type>"),
				UserAssignedIdentities: map[string]*armvideoanalyzer.UserAssignedManagedIdentity{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": {},
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id3": {},
				},
			},
			Properties: &armvideoanalyzer.Properties{
				Encryption: &armvideoanalyzer.AccountEncryption{
					Type: to.Ptr(armvideoanalyzer.AccountEncryptionKeyTypeSystemKey),
				},
				IotHubs: []*armvideoanalyzer.IotHub{
					{
						ID: to.Ptr("<id>"),
						Identity: &armvideoanalyzer.ResourceIdentity{
							UserAssignedIdentity: to.Ptr("<user-assigned-identity>"),
						},
					},
					{
						ID: to.Ptr("<id>"),
						Identity: &armvideoanalyzer.ResourceIdentity{
							UserAssignedIdentity: to.Ptr("<user-assigned-identity>"),
						},
					}},
				StorageAccounts: []*armvideoanalyzer.StorageAccount{
					{
						ID: to.Ptr("<id>"),
						Identity: &armvideoanalyzer.ResourceIdentity{
							UserAssignedIdentity: to.Ptr("<user-assigned-identity>"),
						},
					}},
			},
		},
		&armvideoanalyzer.VideoAnalyzersClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
