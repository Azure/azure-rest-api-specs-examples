package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateWorkspaceApi.json
func ExampleWorkspaceAPIClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspaceAPIClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "wks1", "tempgroup", armapimanagement.APICreateOrUpdateParameter{
		Properties: &armapimanagement.APICreateOrUpdateProperties{
			Description: to.Ptr("apidescription5200"),
			AuthenticationSettings: &armapimanagement.AuthenticationSettingsContract{
				OAuth2: &armapimanagement.OAuth2AuthenticationSettingsContract{
					AuthorizationServerID: to.Ptr("authorizationServerId2283"),
					Scope:                 to.Ptr("oauth2scope2580"),
				},
			},
			SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
				Header: to.Ptr("header4520"),
				Query:  to.Ptr("query3037"),
			},
			Path:        to.Ptr("newapiPath"),
			DisplayName: to.Ptr("apiname1463"),
			Protocols: []*armapimanagement.Protocol{
				to.Ptr(armapimanagement.ProtocolHTTPS),
				to.Ptr(armapimanagement.ProtocolHTTP)},
			ServiceURL: to.Ptr("http://newechoapi.cloudapp.net/api"),
		},
	}, &armapimanagement.WorkspaceAPIClientBeginCreateOrUpdateOptions{IfMatch: nil})
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
	// res.APIContract = armapimanagement.APIContract{
	// 	Name: to.Ptr("apiid9419"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/apis"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/apiid9419"),
	// 	Properties: &armapimanagement.APIContractProperties{
	// 		Description: to.Ptr("apidescription5200"),
	// 		APIRevision: to.Ptr("1"),
	// 		AuthenticationSettings: &armapimanagement.AuthenticationSettingsContract{
	// 			OAuth2: &armapimanagement.OAuth2AuthenticationSettingsContract{
	// 				AuthorizationServerID: to.Ptr("authorizationServerId2283"),
	// 				Scope: to.Ptr("oauth2scope2580"),
	// 			},
	// 			OAuth2AuthenticationSettings: []*armapimanagement.OAuth2AuthenticationSettingsContract{
	// 				{
	// 					AuthorizationServerID: to.Ptr("authorizationServerId2283"),
	// 					Scope: to.Ptr("oauth2scope2580"),
	// 			}},
	// 		},
	// 		IsCurrent: to.Ptr(true),
	// 		IsOnline: to.Ptr(true),
	// 		SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
	// 			Header: to.Ptr("header4520"),
	// 			Query: to.Ptr("query3037"),
	// 		},
	// 		Path: to.Ptr("newapiPath"),
	// 		DisplayName: to.Ptr("apiname1463"),
	// 		Protocols: []*armapimanagement.Protocol{
	// 			to.Ptr(armapimanagement.ProtocolHTTP),
	// 			to.Ptr(armapimanagement.ProtocolHTTPS)},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ServiceURL: to.Ptr("http://newechoapi.cloudapp.net/api"),
	// 		},
	// 	}
}
