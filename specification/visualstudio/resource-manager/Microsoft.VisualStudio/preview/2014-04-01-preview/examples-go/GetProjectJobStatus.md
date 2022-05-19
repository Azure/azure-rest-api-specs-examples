Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvisualstudio%2Farmvisualstudio%2Fv0.4.0/sdk/resourcemanager/visualstudio/armvisualstudio/README.md) on how to add the SDK to your project and authenticate.

```go
package armvisualstudio_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/visualstudio/armvisualstudio"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/GetProjectJobStatus.json
func ExampleProjectsClient_GetJobStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvisualstudio.NewProjectsClient("0de7f055-dbea-498d-8e9e-da287eedca90", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetJobStatus(ctx,
		"VS-Example-Group",
		"ExampleAccount",
		"ExampleProject",
		"ExampleProject",
		"put",
		&armvisualstudio.ProjectsClientGetJobStatusOptions{JobID: to.Ptr("126167d2-d710-4b5d-80a8-a1d58717142d")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
