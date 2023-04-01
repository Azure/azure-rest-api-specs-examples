package armvisualstudio_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/visualstudio/armvisualstudio"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/UpdateExtensionResource.json
func ExampleExtensionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvisualstudio.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExtensionsClient().Update(ctx, "VS-Example-Group", "ExampleAccount", "Example", armvisualstudio.ExtensionResourceRequest{
		Location: to.Ptr("Central US"),
		Plan: &armvisualstudio.ExtensionResourcePlan{
			Name:          to.Ptr("ExamplePlan"),
			Product:       to.Ptr("ExampleExtensionName"),
			PromotionCode: to.Ptr(""),
			Publisher:     to.Ptr("ExampleExtensionPublisher"),
			Version:       to.Ptr("1.0"),
		},
		Properties: map[string]*string{},
		Tags:       map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExtensionResource = armvisualstudio.ExtensionResource{
	// 	Name: to.Ptr("ms.example"),
	// 	Type: to.Ptr("Microsoft.VisualStudio/account/extension"),
	// 	ID: to.Ptr("/subscriptions/0de7f055-dbea-498d-8e9e-da287eedca90/resourceGroups/VS-Example-Group/providers/Microsoft.VisualStudio/account/ExampleAccount/extension/ms.example"),
	// 	Location: to.Ptr("Central US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Plan: &armvisualstudio.ExtensionResourcePlan{
	// 		Name: to.Ptr("ExamplePlan"),
	// 		Product: to.Ptr("ExampleExtensionName"),
	// 		PromotionCode: to.Ptr(""),
	// 		Publisher: to.Ptr("ExampleExtensionPublisher"),
	// 		Version: to.Ptr("1.0"),
	// 	},
	// 	Properties: map[string]*string{
	// 	},
	// }
}
