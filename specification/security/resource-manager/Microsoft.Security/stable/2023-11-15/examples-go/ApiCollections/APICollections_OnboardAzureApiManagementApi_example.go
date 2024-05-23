package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/stable/2023-11-15/examples/ApiCollections/APICollections_OnboardAzureApiManagementApi_example.json
func ExampleAPICollectionsClient_BeginOnboardAzureAPIManagementAPI() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAPICollectionsClient().BeginOnboardAzureAPIManagementAPI(ctx, "rg1", "apimService1", "echo-api", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.APICollection = armsecurity.APICollection{
	// 	Name: to.Ptr("echo-api"),
	// 	Type: to.Ptr("Microsoft.Security/apiCollections"),
	// 	ID: to.Ptr("/subscriptions/3fa85f64-5717-4562-b3fc-2c963f66afa6/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/providers/Microsoft.Security/apiCollections/echo-api"),
	// 	Properties: &armsecurity.APICollectionProperties{
	// 		BaseURL: to.Ptr("https://apimservice1.azure-api.net/echo"),
	// 		DiscoveredVia: to.Ptr("/subscriptions/3fa85f64-5717-4562-b3fc-2c963f66afa6/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1"),
	// 		DisplayName: to.Ptr("Echo API"),
	// 		NumberOfAPIEndpoints: to.Ptr[int64](6),
	// 		NumberOfAPIEndpointsWithSensitiveDataExposed: to.Ptr[int64](1),
	// 		NumberOfExternalAPIEndpoints: to.Ptr[int64](3),
	// 		NumberOfInactiveAPIEndpoints: to.Ptr[int64](3),
	// 		NumberOfUnauthenticatedAPIEndpoints: to.Ptr[int64](1),
	// 		ProvisioningState: to.Ptr(armsecurity.ProvisioningStateSucceeded),
	// 		SensitivityLabel: to.Ptr("Highly Confidential"),
	// 	},
	// }
}
