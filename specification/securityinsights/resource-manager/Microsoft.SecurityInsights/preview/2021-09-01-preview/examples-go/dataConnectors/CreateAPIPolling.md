Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsight%2Farmsecurityinsight%2Fv0.2.1/sdk/resourcemanager/securityinsight/armsecurityinsight/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/CreateAPIPolling.json
func ExampleDataConnectorsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewDataConnectorsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<data-connector-id>",
		&armsecurityinsight.CodelessAPIPollingDataConnector{
			Kind: armsecurityinsight.DataConnectorKind("APIPolling").ToPtr(),
			Properties: &armsecurityinsight.APIPollingParameters{
				ConnectorUIConfig: &armsecurityinsight.CodelessUIConnectorConfigProperties{
					Availability: &armsecurityinsight.Availability{
						IsPreview: to.BoolPtr(true),
						Status:    to.Int32Ptr(1),
					},
					ConnectivityCriteria: []*armsecurityinsight.CodelessUIConnectorConfigPropertiesConnectivityCriteriaItem{
						{
							Type:  armsecurityinsight.ConnectivityType("SentinelKindsV2").ToPtr(),
							Value: []*string{},
						}},
					DataTypes: []*armsecurityinsight.CodelessUIConnectorConfigPropertiesDataTypesItem{
						{
							Name:                  to.StringPtr("<name>"),
							LastDataReceivedQuery: to.StringPtr("<last-data-received-query>"),
						}},
					DescriptionMarkdown: to.StringPtr("<description-markdown>"),
					GraphQueries: []*armsecurityinsight.CodelessUIConnectorConfigPropertiesGraphQueriesItem{
						{
							BaseQuery:  to.StringPtr("<base-query>"),
							Legend:     to.StringPtr("<legend>"),
							MetricName: to.StringPtr("<metric-name>"),
						}},
					GraphQueriesTableName: to.StringPtr("<graph-queries-table-name>"),
					InstructionSteps: []*armsecurityinsight.CodelessUIConnectorConfigPropertiesInstructionStepsItem{
						{
							Description: to.StringPtr("<description>"),
							Instructions: []*armsecurityinsight.InstructionStepsInstructionsItem{
								{
									Type: armsecurityinsight.SettingType("APIKey").ToPtr(),
									Parameters: map[string]interface{}{
										"enable": "true",
										"userRequestPlaceHoldersInput": []interface{}{
											map[string]interface{}{
												"displayText":      "Organization Name",
												"placeHolderName":  "{{placeHolder1}}",
												"placeHolderValue": "",
												"requestObjectKey": "apiEndpoint",
											},
										},
									},
								}},
							Title: to.StringPtr("<title>"),
						}},
					Permissions: &armsecurityinsight.Permissions{
						Customs: []*armsecurityinsight.PermissionsCustomsItem{
							{
								Name:        to.StringPtr("<name>"),
								Description: to.StringPtr("<description>"),
							}},
						ResourceProvider: []*armsecurityinsight.PermissionsResourceProviderItem{
							{
								PermissionsDisplayText: to.StringPtr("<permissions-display-text>"),
								Provider:               armsecurityinsight.ProviderName("Microsoft.OperationalInsights/workspaces").ToPtr(),
								ProviderDisplayName:    to.StringPtr("<provider-display-name>"),
								RequiredPermissions: &armsecurityinsight.RequiredPermissions{
									Delete: to.BoolPtr(true),
									Read:   to.BoolPtr(true),
									Write:  to.BoolPtr(true),
								},
								Scope: armsecurityinsight.PermissionProviderScope("Workspace").ToPtr(),
							}},
					},
					Publisher: to.StringPtr("<publisher>"),
					SampleQueries: []*armsecurityinsight.CodelessUIConnectorConfigPropertiesSampleQueriesItem{
						{
							Description: to.StringPtr("<description>"),
							Query:       to.StringPtr("<query>"),
						}},
					Title: to.StringPtr("<title>"),
				},
				PollingConfig: &armsecurityinsight.CodelessConnectorPollingConfigProperties{
					Auth: &armsecurityinsight.CodelessConnectorPollingAuthProperties{
						APIKeyIdentifier: to.StringPtr("<apikey-identifier>"),
						APIKeyName:       to.StringPtr("<apikey-name>"),
						AuthType:         to.StringPtr("<auth-type>"),
					},
					Paging: &armsecurityinsight.CodelessConnectorPollingPagingProperties{
						PageSizeParaName: to.StringPtr("<page-size-para-name>"),
						PagingType:       to.StringPtr("<paging-type>"),
					},
					Response: &armsecurityinsight.CodelessConnectorPollingResponseProperties{
						EventsJSONPaths: []*string{
							to.StringPtr("$")},
					},
					Request: &armsecurityinsight.CodelessConnectorPollingRequestProperties{
						APIEndpoint: to.StringPtr("<apiendpoint>"),
						Headers: map[string]interface{}{
							"Accept":     "application/json",
							"User-Agent": "Scuba",
						},
						HTTPMethod: to.StringPtr("<httpmethod>"),
						QueryParameters: map[string]interface{}{
							"phrase": "created:{_QueryWindowStartTime}..{_QueryWindowEndTime}",
						},
						QueryTimeFormat:  to.StringPtr("<query-time-format>"),
						QueryWindowInMin: to.Int32Ptr(15),
						RateLimitQPS:     to.Int32Ptr(50),
						RetryCount:       to.Int32Ptr(2),
						TimeoutInSeconds: to.Int32Ptr(60),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DataConnectorsClientCreateOrUpdateResult)
}
```
