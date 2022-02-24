Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevtestlabs%2Farmdevtestlabs%2Fv0.2.1/sdk/resourcemanager/devtestlabs/armdevtestlabs/README.md) on how to add the SDK to your project and authenticate.

```go
package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArtifactSources_CreateOrUpdate.json
func ExampleArtifactSourcesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewArtifactSourcesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<name>",
		armdevtestlabs.ArtifactSource{
			Tags: map[string]*string{
				"tagName1": to.StringPtr("tagValue1"),
			},
			Properties: &armdevtestlabs.ArtifactSourceProperties{
				ArmTemplateFolderPath: to.StringPtr("<arm-template-folder-path>"),
				BranchRef:             to.StringPtr("<branch-ref>"),
				DisplayName:           to.StringPtr("<display-name>"),
				FolderPath:            to.StringPtr("<folder-path>"),
				SecurityToken:         to.StringPtr("<security-token>"),
				SourceType:            armdevtestlabs.SourceControlType("{VsoGit|GitHub|StorageAccount}").ToPtr(),
				Status:                armdevtestlabs.EnableStatus("{Enabled|Disabled}").ToPtr(),
				URI:                   to.StringPtr("<uri>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ArtifactSourcesClientCreateOrUpdateResult)
}
```
