package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementUpdateWorkspaceApi.json
func ExampleWorkspaceAPIClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceAPIClient().Update(ctx, "rg1", "apimService1", "wks1", "echo-api", "*", armapimanagement.APIUpdateContract{
		Properties: &armapimanagement.APIContractUpdateProperties{
			Path:        to.Ptr("newecho"),
			DisplayName: to.Ptr("Echo API New"),
			ServiceURL:  to.Ptr("http://echoapi.cloudapp.net/api2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.WorkspaceAPIClientUpdateResponse{
	// 	APIContract: armapimanagement.APIContract{
	// 		Name: to.Ptr("echo-api"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/pis"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/echo-api"),
	// 		Properties: &armapimanagement.APIContractProperties{
	// 			Path: to.Ptr("newecho"),
	// 			APIRevision: to.Ptr("1"),
	// 			DisplayName: to.Ptr("Echo API New"),
	// 			IsCurrent: to.Ptr(true),
	// 			IsOnline: to.Ptr(true),
	// 			Protocols: []*armapimanagement.Protocol{
	// 				to.Ptr(armapimanagement.ProtocolHTTPS),
	// 			},
	// 			ServiceURL: to.Ptr("http://echoapi.cloudapp.net/api2"),
	// 			SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
	// 				Header: to.Ptr("Ocp-Apim-Subscription-Key"),
	// 				Query: to.Ptr("subscription-key"),
	// 			},
	// 		},
	// 	},
	// }
}
