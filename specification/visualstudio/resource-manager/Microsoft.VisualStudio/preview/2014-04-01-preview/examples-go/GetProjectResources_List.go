package armvisualstudio_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/visualstudio/armvisualstudio"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/GetProjectResources_List.json
func ExampleProjectsClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvisualstudio.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectsClient().ListByResourceGroup(ctx, "VS-Example-Group", "ExampleAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProjectResourceListResult = armvisualstudio.ProjectResourceListResult{
	// 	Value: []*armvisualstudio.ProjectResource{
	// 		{
	// 			Name: to.Ptr("ExampleProject"),
	// 			Type: to.Ptr("Microsoft.VisualStudio/account/project"),
	// 			ID: to.Ptr("/subscriptions/0de7f055-dbea-498d-8e9e-da287eedca90/resourceGroups/VS-Example-Group/providers/Microsoft.VisualStudio/account/ExampleAccount/project/ExampleProject"),
	// 			Location: to.Ptr("Central US"),
	// 			Tags: map[string]*string{
	// 			},
	// 			Properties: map[string]*string{
	// 				"AzureResourceName": to.Ptr("ExampleProject"),
	// 				"TfsUniqueIdentifier": to.Ptr("vstfs:///Classification/TeamProject/d921d601-0600-4e45-8600-3914ffc8811a"),
	// 			},
	// 	}},
	// }
}
