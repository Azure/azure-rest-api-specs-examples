package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/preview/2022-11-20-preview/examples/ApiCollections/APICollectionOnboarding_Create_example.json
func ExampleAPICollectionOnboardingClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPICollectionOnboardingClient().Create(ctx, "rg1", "apimService1", "echo-api", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.APICollectionResponse = armsecurity.APICollectionResponse{
	// 	Name: to.Ptr("echo-api"),
	// 	Type: to.Ptr("Microsoft.Security/apiCollections"),
	// 	ID: to.Ptr("/subscriptions/3fa85f64-5717-4562-b3fc-2c963f66afa6/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/providers/Microsoft.Security/apiCollections/echo-api"),
	// 	Properties: &armsecurity.APICollectionProperties{
	// 		AdditionalData: map[string]*string{
	// 			"apiManagementApiId": to.Ptr("/subscriptions/3fa85f64-5717-4562-b3fc-2c963f66afa6/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echo-api"),
	// 		},
	// 		DisplayName: to.Ptr("echo-api"),
	// 	},
	// }
}
