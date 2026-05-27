package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementCreateApiRevisionFromExistingApi.json
func ExampleAPIClient_BeginCreateOrUpdate_apiManagementCreateApiRevisionFromExistingApi() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAPIClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "echo-api;rev=3", armapimanagement.APICreateOrUpdateParameter{
		Properties: &armapimanagement.APICreateOrUpdateProperties{
			Path:                   to.Ptr("echo"),
			APIRevisionDescription: to.Ptr("Creating a Revision of an existing API"),
			ServiceURL:             to.Ptr("http://echoapi.cloudapp.net/apiv3"),
			SourceAPIID:            to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echo-api"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.APIClientCreateOrUpdateResponse{
	// 	APIContract: armapimanagement.APIContract{
	// 		Name: to.Ptr("echo-api;rev=3"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/apis"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echo-api;rev=3"),
	// 		Properties: &armapimanagement.APIContractProperties{
	// 			Path: to.Ptr("echo"),
	// 			APIRevision: to.Ptr("3"),
	// 			APIRevisionDescription: to.Ptr("Creating a Revision of an existing API"),
	// 			DisplayName: to.Ptr("Echo API"),
	// 			Protocols: []*armapimanagement.Protocol{
	// 				to.Ptr(armapimanagement.ProtocolHTTPS),
	// 			},
	// 			ProvisioningState: to.Ptr("InProgress"),
	// 			ServiceURL: to.Ptr("http://echoapi.cloudapp.net/apiv3"),
	// 			SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
	// 				Header: to.Ptr("Ocp-Apim-Subscription-Key"),
	// 				Query: to.Ptr("subscription-key"),
	// 			},
	// 			SubscriptionRequired: to.Ptr(true),
	// 		},
	// 	},
	// }
}
