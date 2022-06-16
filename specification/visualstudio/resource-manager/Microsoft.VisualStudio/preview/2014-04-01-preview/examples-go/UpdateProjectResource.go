package armvisualstudio_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/visualstudio/armvisualstudio"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/UpdateProjectResource.json
func ExampleProjectsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvisualstudio.NewProjectsClient("0de7f055-dbea-498d-8e9e-da287eedca90", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"VS-Example-Group",
		"ExampleAccount",
		"ExampleProject",
		armvisualstudio.ProjectResource{
			Name:       to.Ptr("ms.example"),
			Type:       to.Ptr("Microsoft.VisualStudio/account/extension"),
			ID:         to.Ptr("/subscriptions/0de7f055-dbea-498d-8e9e-da287eedca90/resourceGroups/VS-Example-Group/providers/microsoft.visualstudio/account/ExampleAccount/project/ExampleProject"),
			Location:   to.Ptr("Central US"),
			Tags:       map[string]*string{},
			Properties: map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
