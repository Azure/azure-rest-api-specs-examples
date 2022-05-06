Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fautomation%2Farmautomation%2Fv0.5.0/sdk/resourcemanager/automation/armautomation/README.md) on how to add the SDK to your project and authenticate.

```go
package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControlSyncJob/createSourceControlSyncJob.json
func ExampleSourceControlSyncJobClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armautomation.NewSourceControlSyncJobClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Create(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<source-control-name>",
		"<source-control-sync-job-id>",
		armautomation.SourceControlSyncJobCreateParameters{
			Properties: &armautomation.SourceControlSyncJobCreateProperties{
				CommitID: to.Ptr("<commit-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
```
