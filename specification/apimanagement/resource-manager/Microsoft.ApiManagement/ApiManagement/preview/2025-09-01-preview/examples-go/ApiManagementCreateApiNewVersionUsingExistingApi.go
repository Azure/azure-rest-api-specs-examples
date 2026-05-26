package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementCreateApiNewVersionUsingExistingApi.json
func ExampleAPIClient_BeginCreateOrUpdate_apiManagementCreateApiNewVersionUsingExistingApi() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAPIClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "echoapiv3", armapimanagement.APICreateOrUpdateParameter{
		Properties: &armapimanagement.APICreateOrUpdateProperties{
			Path:            to.Ptr("echo2"),
			Description:     to.Ptr("Create Echo API into a new Version using Existing Version Set and Copy all Operations."),
			APIVersion:      to.Ptr("v4"),
			APIVersionSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apiVersionSets/aa9c59e6-c0cd-4258-9356-9ca7d2f0b458"),
			DisplayName:     to.Ptr("Echo API2"),
			IsCurrent:       to.Ptr(true),
			Protocols: []*armapimanagement.Protocol{
				to.Ptr(armapimanagement.ProtocolHTTP),
				to.Ptr(armapimanagement.ProtocolHTTPS),
			},
			ServiceURL:           to.Ptr("http://echoapi.cloudapp.net/api"),
			SourceAPIID:          to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echoPath"),
			SubscriptionRequired: to.Ptr(true),
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
	// 		Name: to.Ptr("echoapiv3"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/apis"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echoapiv3"),
	// 		Properties: &armapimanagement.APIContractProperties{
	// 			Path: to.Ptr("echo2"),
	// 			Description: to.Ptr("Create Echo API into a new Version using Existing Version Set and Copy all Operations."),
	// 			APIRevision: to.Ptr("1"),
	// 			APIVersion: to.Ptr("v4"),
	// 			APIVersionSet: &armapimanagement.APIVersionSetContractDetails{
	// 				Name: to.Ptr("Echo API2"),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apiVersionSets/aa9c59e6-c0cd-4258-9356-9ca7d2f0b458"),
	// 				VersioningScheme: to.Ptr(armapimanagement.VersioningSchemeSegment),
	// 			},
	// 			APIVersionSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apiVersionSets/aa9c59e6-c0cd-4258-9356-9ca7d2f0b458"),
	// 			DisplayName: to.Ptr("Echo API2"),
	// 			IsCurrent: to.Ptr(true),
	// 			Protocols: []*armapimanagement.Protocol{
	// 				to.Ptr(armapimanagement.ProtocolHTTP),
	// 				to.Ptr(armapimanagement.ProtocolHTTPS),
	// 			},
	// 			ProvisioningState: to.Ptr("InProgress"),
	// 			ServiceURL: to.Ptr("http://echoapi.cloudapp.net/api"),
	// 			SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
	// 				Header: to.Ptr("Ocp-Apim-Subscription-Key"),
	// 				Query: to.Ptr("subscription-key"),
	// 			},
	// 			SubscriptionRequired: to.Ptr(true),
	// 		},
	// 	},
	// }
}
