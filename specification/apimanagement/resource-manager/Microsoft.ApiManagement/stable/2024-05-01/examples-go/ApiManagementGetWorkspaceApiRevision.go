package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetWorkspaceApiRevision.json
func ExampleWorkspaceAPIClient_Get_apiManagementGetWorkspaceApiRevision() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceAPIClient().Get(ctx, "rg1", "apimService1", "wks1", "echo-api;rev=3", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.APIContract = armapimanagement.APIContract{
	// 	Name: to.Ptr("echo-api;rev=3"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/apis"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/echo-api;rev=3"),
	// 	Properties: &armapimanagement.APIContractProperties{
	// 		APIRevision: to.Ptr("3"),
	// 		APIRevisionDescription: to.Ptr("fixed bug in contract"),
	// 		AuthenticationSettings: &armapimanagement.AuthenticationSettingsContract{
	// 			OAuth2: &armapimanagement.OAuth2AuthenticationSettingsContract{
	// 				AuthorizationServerID: to.Ptr("authorizationServerId2283"),
	// 				Scope: to.Ptr("oauth2scope2580"),
	// 			},
	// 			OAuth2AuthenticationSettings: []*armapimanagement.OAuth2AuthenticationSettingsContract{
	// 				{
	// 					AuthorizationServerID: to.Ptr("authorizationServerId2283"),
	// 					Scope: to.Ptr("oauth2scope2580"),
	// 				},
	// 				{
	// 					AuthorizationServerID: to.Ptr("authorizationServerId2284"),
	// 					Scope: to.Ptr("oauth2scope2581"),
	// 			}},
	// 		},
	// 		SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
	// 			Header: to.Ptr("Ocp-Apim-Subscription-Key"),
	// 			Query: to.Ptr("subscription-key"),
	// 		},
	// 		Path: to.Ptr("schulte"),
	// 		DisplayName: to.Ptr("Service"),
	// 		Protocols: []*armapimanagement.Protocol{
	// 			to.Ptr(armapimanagement.ProtocolHTTPS)},
	// 			ServiceURL: to.Ptr("https://api.plexonline.com/DataSource/Service.asmx"),
	// 		},
	// 	}
}
