Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvisualstudio%2Farmvisualstudio%2Fv0.1.0/sdk/resourcemanager/visualstudio/armvisualstudio/README.md) on how to add the SDK to your project and authenticate.

```go
package armvisualstudio_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/visualstudio/armvisualstudio"
)

// x-ms-original-file: specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/UpdateProjectResource.json
func ExampleProjectsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvisualstudio.NewProjectsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<root-resource-name>",
		"<resource-name>",
		armvisualstudio.ProjectResource{
			Resource: armvisualstudio.Resource{
				Name:     to.StringPtr("<name>"),
				Type:     to.StringPtr("<type>"),
				ID:       to.StringPtr("<id>"),
				Location: to.StringPtr("<location>"),
				Tags:     map[string]*string{},
			},
			Properties: map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ProjectResource.ID: %s\n", *res.ID)
}
```
