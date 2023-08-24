package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiNewVersionUsingExistingApi.json
func ExampleAPIClient_BeginCreateOrUpdate_apiManagementCreateApiNewVersionUsingExistingApi() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAPIClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "echoapiv3", armapimanagement.APICreateOrUpdateParameter{
		Properties: &armapimanagement.APICreateOrUpdateProperties{
			Description:          to.Ptr("Create Echo API into a new Version using Existing Version Set and Copy all Operations."),
			APIVersion:           to.Ptr("v4"),
			APIVersionSetID:      to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apiVersionSets/aa9c59e6-c0cd-4258-9356-9ca7d2f0b458"),
			IsCurrent:            to.Ptr(true),
			SubscriptionRequired: to.Ptr(true),
			Path:                 to.Ptr("echo2"),
			DisplayName:          to.Ptr("Echo API2"),
			Protocols: []*armapimanagement.Protocol{
				to.Ptr(armapimanagement.ProtocolHTTP),
				to.Ptr(armapimanagement.ProtocolHTTPS)},
			ServiceURL:  to.Ptr("http://echoapi.cloudapp.net/api"),
			SourceAPIID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echoPath"),
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
	// 	Name: to.Ptr("echoapiv3"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echoapiv3"),
	// 	Properties: &armapimanagement.APIContractProperties{
	// 		Description: to.Ptr("Create Echo API into a new Version using Existing Version Set and Copy all Operations."),
	// 		APIRevision: to.Ptr("1"),
	// 		APIVersion: to.Ptr("v4"),
	// 		APIVersionSetID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apiVersionSets/aa9c59e6-c0cd-4258-9356-9ca7d2f0b458"),
	// 		IsCurrent: to.Ptr(true),
	// 		SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
	// 			Header: to.Ptr("Ocp-Apim-Subscription-Key"),
	// 			Query: to.Ptr("subscription-key"),
	// 		},
	// 		SubscriptionRequired: to.Ptr(true),
	// 		Path: to.Ptr("echo2"),
	// 		APIVersionSet: &armapimanagement.APIVersionSetContractDetails{
	// 			Name: to.Ptr("Echo API2"),
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apiVersionSets/aa9c59e6-c0cd-4258-9356-9ca7d2f0b458"),
	// 			VersioningScheme: to.Ptr(armapimanagement.APIVersionSetContractDetailsVersioningSchemeSegment),
	// 		},
	// 		DisplayName: to.Ptr("Echo API2"),
	// 		Protocols: []*armapimanagement.Protocol{
	// 			to.Ptr(armapimanagement.ProtocolHTTP),
	// 			to.Ptr(armapimanagement.ProtocolHTTPS)},
	// 			ServiceURL: to.Ptr("http://echoapi.cloudapp.net/api"),
	// 		},
	// 	}
}
