package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiWithOpenIdConnect.json
func ExampleAPIClient_BeginCreateOrUpdate_apiManagementCreateApiWithOpenIdConnect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAPIClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "tempgroup", armapimanagement.APICreateOrUpdateParameter{
		Properties: &armapimanagement.APICreateOrUpdateProperties{
			Description: to.Ptr("This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters."),
			AuthenticationSettings: &armapimanagement.AuthenticationSettingsContract{
				Openid: &armapimanagement.OpenIDAuthenticationSettingsContract{
					BearerTokenSendingMethods: []*armapimanagement.BearerTokenSendingMethods{
						to.Ptr(armapimanagement.BearerTokenSendingMethodsAuthorizationHeader)},
					OpenidProviderID: to.Ptr("testopenid"),
				},
			},
			SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
				Header: to.Ptr("Ocp-Apim-Subscription-Key"),
				Query:  to.Ptr("subscription-key"),
			},
			Path:        to.Ptr("petstore"),
			DisplayName: to.Ptr("Swagger Petstore"),
			Protocols: []*armapimanagement.Protocol{
				to.Ptr(armapimanagement.ProtocolHTTPS)},
			ServiceURL: to.Ptr("http://petstore.swagger.io/v2"),
		},
	}, &armapimanagement.APIClientBeginCreateOrUpdateOptions{IfMatch: nil})
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
	// 	Name: to.Ptr("58da4c4ccdae970a08121230"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/58da4c4ccdae970a08121230"),
	// 	Properties: &armapimanagement.APIContractProperties{
	// 		Description: to.Ptr("This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters."),
	// 		APIRevision: to.Ptr("1"),
	// 		AuthenticationSettings: &armapimanagement.AuthenticationSettingsContract{
	// 			Openid: &armapimanagement.OpenIDAuthenticationSettingsContract{
	// 				BearerTokenSendingMethods: []*armapimanagement.BearerTokenSendingMethods{
	// 					to.Ptr(armapimanagement.BearerTokenSendingMethodsAuthorizationHeader)},
	// 					OpenidProviderID: to.Ptr("testopenid"),
	// 				},
	// 			},
	// 			IsCurrent: to.Ptr(true),
	// 			SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
	// 				Header: to.Ptr("Ocp-Apim-Subscription-Key"),
	// 				Query: to.Ptr("subscription-key"),
	// 			},
	// 			Path: to.Ptr("petstore"),
	// 			DisplayName: to.Ptr("Swagger Petstore"),
	// 			Protocols: []*armapimanagement.Protocol{
	// 				to.Ptr(armapimanagement.ProtocolHTTPS)},
	// 				ServiceURL: to.Ptr("http://petstore.swagger.io/v2"),
	// 			},
	// 		}
}
