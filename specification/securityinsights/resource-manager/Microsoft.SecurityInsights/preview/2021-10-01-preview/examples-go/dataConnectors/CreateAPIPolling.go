package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/dataConnectors/CreateAPIPolling.json
func ExampleDataConnectorsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewDataConnectorsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<data-connector-id>",
		&armsecurityinsights.CodelessAPIPollingDataConnector{
			Kind: armsecurityinsights.DataConnectorKind("APIPolling").ToPtr(),
			Properties: &armsecurityinsights.APIPollingParameters{
				ConnectorUIConfig: &armsecurityinsights.CodelessUIConnectorConfigProperties{
					Availability: &armsecurityinsights.Availability{
						IsPreview: to.BoolPtr(true),
						Status:    to.Int32Ptr(1),
					},
					ConnectivityCriteria: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesConnectivityCriteriaItem{
						{
							Type:  armsecurityinsights.ConnectivityType("SentinelKindsV2").ToPtr(),
							Value: []*string{},
						}},
					DataTypes: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesDataTypesItem{
						{
							Name:                  to.StringPtr("<name>"),
							LastDataReceivedQuery: to.StringPtr("<last-data-received-query>"),
						}},
					DescriptionMarkdown: to.StringPtr("<description-markdown>"),
					GraphQueries: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesGraphQueriesItem{
						{
							BaseQuery:  to.StringPtr("<base-query>"),
							Legend:     to.StringPtr("<legend>"),
							MetricName: to.StringPtr("<metric-name>"),
						}},
					GraphQueriesTableName: to.StringPtr("<graph-queries-table-name>"),
					InstructionSteps: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesInstructionStepsItem{
						{
							Description: to.StringPtr("<description>"),
							Instructions: []*armsecurityinsights.InstructionStepsInstructionsItem{
								{
									Type: armsecurityinsights.SettingType("APIKey").ToPtr(),
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
					Permissions: &armsecurityinsights.Permissions{
						Customs: []*armsecurityinsights.PermissionsCustomsItem{
							{
								Name:        to.StringPtr("<name>"),
								Description: to.StringPtr("<description>"),
							}},
						ResourceProvider: []*armsecurityinsights.PermissionsResourceProviderItem{
							{
								PermissionsDisplayText: to.StringPtr("<permissions-display-text>"),
								Provider:               armsecurityinsights.ProviderName("Microsoft.OperationalInsights/workspaces").ToPtr(),
								ProviderDisplayName:    to.StringPtr("<provider-display-name>"),
								RequiredPermissions: &armsecurityinsights.RequiredPermissions{
									Delete: to.BoolPtr(true),
									Read:   to.BoolPtr(true),
									Write:  to.BoolPtr(true),
								},
								Scope: armsecurityinsights.PermissionProviderScope("Workspace").ToPtr(),
							}},
					},
					Publisher: to.StringPtr("<publisher>"),
					SampleQueries: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesSampleQueriesItem{
						{
							Description: to.StringPtr("<description>"),
							Query:       to.StringPtr("<query>"),
						}},
					Title: to.StringPtr("<title>"),
				},
				PollingConfig: &armsecurityinsights.CodelessConnectorPollingConfigProperties{
					Auth: &armsecurityinsights.CodelessConnectorPollingAuthProperties{
						APIKeyIdentifier: to.StringPtr("<apikey-identifier>"),
						APIKeyName:       to.StringPtr("<apikey-name>"),
						AuthType:         to.StringPtr("<auth-type>"),
					},
					Paging: &armsecurityinsights.CodelessConnectorPollingPagingProperties{
						PageSizeParaName: to.StringPtr("<page-size-para-name>"),
						PagingType:       to.StringPtr("<paging-type>"),
					},
					Response: &armsecurityinsights.CodelessConnectorPollingResponseProperties{
						EventsJSONPaths: []*string{
							to.StringPtr("$")},
					},
					Request: &armsecurityinsights.CodelessConnectorPollingRequestProperties{
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
