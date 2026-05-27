package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementGetApiContract.json
func ExampleAPIClient_Get_apiManagementGetApiContract() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIClient().Get(ctx, "rg1", "apimService1", "57d1f7558aa04f15146d9d8a", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.APIClientGetResponse{
	// 	APIContract: armapimanagement.APIContract{
	// 		Name: to.Ptr("57d1f7558aa04f15146d9d8a"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/apis"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/57d1f7558aa04f15146d9d8a"),
	// 		Properties: &armapimanagement.APIContractProperties{
	// 			APIType: to.Ptr(armapimanagement.APITypeSoap),
	// 			Path: to.Ptr("schulte"),
	// 			APIRevision: to.Ptr("1"),
	// 			AuthenticationSettings: &armapimanagement.AuthenticationSettingsContract{
	// 				OAuth2: &armapimanagement.OAuth2AuthenticationSettingsContract{
	// 					AuthorizationServerID: to.Ptr("authorizationServerId2283"),
	// 					Scope: to.Ptr("oauth2scope2580"),
	// 				},
	// 				OAuth2AuthenticationSettings: []*armapimanagement.OAuth2AuthenticationSettingsContract{
	// 					{
	// 						AuthorizationServerID: to.Ptr("authorizationServerId2283"),
	// 						Scope: to.Ptr("oauth2scope2580"),
	// 					},
	// 					{
	// 						AuthorizationServerID: to.Ptr("authorizationServerId2284"),
	// 						Scope: to.Ptr("oauth2scope2581"),
	// 					},
	// 				},
	// 			},
	// 			DisplayName: to.Ptr("Service"),
	// 			IsCurrent: to.Ptr(true),
	// 			IsOnline: to.Ptr(true),
	// 			Protocols: []*armapimanagement.Protocol{
	// 				to.Ptr(armapimanagement.ProtocolHTTPS),
	// 			},
	// 			ServiceURL: to.Ptr("https://api.plexonline.com/DataSource/Service.asmx"),
	// 			SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
	// 				Header: to.Ptr("Ocp-Apim-Subscription-Key"),
	// 				Query: to.Ptr("subscription-key"),
	// 			},
	// 		},
	// 	},
	// }
}
