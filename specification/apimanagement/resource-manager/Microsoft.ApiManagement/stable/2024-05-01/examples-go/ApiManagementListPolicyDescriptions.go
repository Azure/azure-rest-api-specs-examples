package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListPolicyDescriptions.json
func ExamplePolicyDescriptionClient_ListByService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyDescriptionClient().ListByService(ctx, "rg1", "apimService1", &armapimanagement.PolicyDescriptionClientListByServiceOptions{Scope: to.Ptr(armapimanagement.PolicyScopeContractAPI)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PolicyDescriptionCollection = armapimanagement.PolicyDescriptionCollection{
	// 	Count: to.Ptr[int64](2),
	// 	Value: []*armapimanagement.PolicyDescriptionContract{
	// 		{
	// 			Name: to.Ptr("authentication-basic"),
	// 			Type: to.Ptr("Microsoft.ApiManagement/service/policyDescriptions"),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/policyDescriptions/authentication-basic"),
	// 			Properties: &armapimanagement.PolicyDescriptionContractProperties{
	// 				Description: to.Ptr("Authenticate with the backend service using Basic authentication. Use in the inbound section at API scope."),
	// 				Scope: to.Ptr[int64](268435471),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("authentication-certificate"),
	// 			Type: to.Ptr("Microsoft.ApiManagement/service/policyDescriptions"),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/policyDescriptions/authentication-certificate"),
	// 			Properties: &armapimanagement.PolicyDescriptionContractProperties{
	// 				Description: to.Ptr("Authenticate with the backend service using a client certificate. Use in the inbound section at API scope."),
	// 				Scope: to.Ptr[int64](268435471),
	// 			},
	// 	}},
	// }
}
