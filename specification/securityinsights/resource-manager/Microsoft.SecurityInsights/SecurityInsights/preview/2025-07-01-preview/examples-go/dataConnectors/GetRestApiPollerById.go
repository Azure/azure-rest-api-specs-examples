package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/dataConnectors/GetRestApiPollerById.json
func ExampleDataConnectorsClient_Get_getARestApiPollerDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "RestApiPoller_fce27b90-d6f5-4d30-991a-af509a2b50a1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	DataConnectorClassification: &armsecurityinsights.RestAPIPollerDataConnector{
	// 		Name: to.Ptr("RestApiPoller_fce27b90-d6f5-4d30-991a-af509a2b50a1"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/RestApiPoller_afef3743-0c88-469c-84ff-ca2e87dc1e48"),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindRestAPIPoller),
	// 		Properties: &armsecurityinsights.RestAPIPollerDataConnectorProperties{
	// 			Auth: &armsecurityinsights.APIKeyAuthModel{
	// 				Type: to.Ptr(armsecurityinsights.CcpAuthTypeAPIKey),
	// 				APIKey: to.Ptr("6bec40cf957de430a6f1f2baa056b99a4fac9ea0"),
	// 				APIKeyName: to.Ptr("X-Cisco-Meraki-API-Key"),
	// 			},
	// 			ConnectorDefinitionName: to.Ptr("RestApiPollerDefinition"),
	// 			DcrConfig: &armsecurityinsights.DCRConfiguration{
	// 				DataCollectionEndpoint: to.Ptr("data collection Endpoint"),
	// 				DataCollectionRuleImmutableID: to.Ptr("data collection rule immutableId"),
	// 				StreamName: to.Ptr("Meraki"),
	// 			},
	// 			Paging: &armsecurityinsights.RestAPIPollerRequestPagingConfig{
	// 				PagingType: to.Ptr(armsecurityinsights.RestAPIPollerRequestPagingKindLinkHeader),
	// 			},
	// 			Response: &armsecurityinsights.CcpResponseConfig{
	// 				EventsJSONPaths: []*string{
	// 					to.Ptr("$"),
	// 				},
	// 			},
	// 			Request: &armsecurityinsights.RestAPIPollerRequestConfig{
	// 				APIEndpoint: to.Ptr("https://api.meraki.com/api/v1/organizations/573083052582915028/networks"),
	// 				EndTimeAttributeName: to.Ptr("t1"),
	// 				Headers: map[string]*string{
	// 					"Accept": to.Ptr("application/json"),
	// 					"User-Agent": to.Ptr("Scuba"),
	// 				},
	// 				HTTPMethod: to.Ptr(armsecurityinsights.HTTPMethodVerbGET),
	// 				QueryParameters: map[string]any{
	// 					"perPage": 1000,
	// 				},
	// 				QueryTimeFormat: to.Ptr("UnixTimestamp"),
	// 				QueryWindowInMin: to.Ptr[int32](6),
	// 				RateLimitQPS: to.Ptr[int32](10),
	// 				RetryCount: to.Ptr[int32](3),
	// 				StartTimeAttributeName: to.Ptr("t0"),
	// 				TimeoutInSeconds: to.Ptr[int32](60),
	// 			},
	// 		},
	// 	},
	// }
}
