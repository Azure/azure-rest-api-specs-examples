```go
package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArtifactSources_CreateOrUpdate.json
func ExampleArtifactSourcesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewArtifactSourcesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"resourceGroupName",
		"{labName}",
		"{artifactSourceName}",
		armdevtestlabs.ArtifactSource{
			Tags: map[string]*string{
				"tagName1": to.Ptr("tagValue1"),
			},
			Properties: &armdevtestlabs.ArtifactSourceProperties{
				ArmTemplateFolderPath: to.Ptr("{armTemplateFolderPath}"),
				BranchRef:             to.Ptr("{branchRef}"),
				DisplayName:           to.Ptr("{displayName}"),
				FolderPath:            to.Ptr("{folderPath}"),
				SecurityToken:         to.Ptr("{securityToken}"),
				SourceType:            to.Ptr(armdevtestlabs.SourceControlType("{VsoGit|GitHub|StorageAccount}")),
				Status:                to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
				URI:                   to.Ptr("{artifactSourceUri}"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevtestlabs%2Farmdevtestlabs%2Fv1.0.0/sdk/resourcemanager/devtestlabs/armdevtestlabs/README.md) on how to add the SDK to your project and authenticate.
