package armtemplatespecs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armtemplatespecs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Resources/stable/2022-02-01/examples/BuiltInTemplateSpecsGet.json
func ExampleClient_GetBuiltIn() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtemplatespecs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().GetBuiltIn(ctx, "nameOfTheBuiltIn", &armtemplatespecs.ClientGetBuiltInOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TemplateSpec = armtemplatespecs.TemplateSpec{
	// 	Name: to.Ptr("nameOfTheBuiltIn"),
	// 	Type: to.Ptr("Microsoft.Resources/builtInTemplateSpecs"),
	// 	ID: to.Ptr("/providers/Microsoft.Resources/builtInTemplateSpecs/nameOfTheBuiltIn"),
	// 	SystemData: &armtemplatespecs.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-01T01:01:01.107Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-02T02:03:01.197Z"); return t}()),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armtemplatespecs.TemplateSpecProperties{
	// 		Description: to.Ptr("The description of the built-in template spec"),
	// 	},
	// }
}
