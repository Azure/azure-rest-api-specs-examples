Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvideoanalyzer%2Farmvideoanalyzer%2Fv0.2.1/sdk/resourcemanager/videoanalyzer/armvideoanalyzer/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-accounts-create-or-update.json
func ExampleVideoAnalyzersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewVideoAnalyzersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armvideoanalyzer.VideoAnalyzer{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
			Identity: &armvideoanalyzer.Identity{
				Type: to.StringPtr("<type>"),
				UserAssignedIdentities: map[string]*armvideoanalyzer.UserAssignedManagedIdentity{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": {},
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id3": {},
				},
			},
			Properties: &armvideoanalyzer.Properties{
				Encryption: &armvideoanalyzer.AccountEncryption{
					Type: armvideoanalyzer.AccountEncryptionKeyType("SystemKey").ToPtr(),
				},
				IotHubs: []*armvideoanalyzer.IotHub{
					{
						ID: to.StringPtr("<id>"),
						Identity: &armvideoanalyzer.ResourceIdentity{
							UserAssignedIdentity: to.StringPtr("<user-assigned-identity>"),
						},
					},
					{
						ID: to.StringPtr("<id>"),
						Identity: &armvideoanalyzer.ResourceIdentity{
							UserAssignedIdentity: to.StringPtr("<user-assigned-identity>"),
						},
					}},
				StorageAccounts: []*armvideoanalyzer.StorageAccount{
					{
						ID: to.StringPtr("<id>"),
						Identity: &armvideoanalyzer.ResourceIdentity{
							UserAssignedIdentity: to.StringPtr("<user-assigned-identity>"),
						},
					}},
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
	log.Printf("Response result: %#v\n", res.VideoAnalyzersClientCreateOrUpdateResult)
}
```
