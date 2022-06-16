package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/dataConnectors/CreateAPIPolling.json
func ExampleDataConnectorsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewDataConnectorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<data-connector-id>",
		&armsecurityinsights.CodelessAPIPollingDataConnector{
			Kind: to.Ptr(armsecurityinsights.DataConnectorKindAPIPolling),
			Properties: &armsecurityinsights.APIPollingParameters{
				ConnectorUIConfig: &armsecurityinsights.CodelessUIConnectorConfigProperties{
					Availability: &armsecurityinsights.Availability{
						IsPreview: to.Ptr(true),
						Status:    to.Ptr[int32](1),
					},
					ConnectivityCriteria: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesConnectivityCriteriaItem{
						{
							Type:  to.Ptr(armsecurityinsights.ConnectivityType("SentinelKindsV2")),
							Value: []*string{},
						}},
					DataTypes: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesDataTypesItem{
						{
							Name:                  to.Ptr("<name>"),
							LastDataReceivedQuery: to.Ptr("<last-data-received-query>"),
						}},
					DescriptionMarkdown: to.Ptr("<description-markdown>"),
					GraphQueries: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesGraphQueriesItem{
						{
							BaseQuery:  to.Ptr("<base-query>"),
							Legend:     to.Ptr("<legend>"),
							MetricName: to.Ptr("<metric-name>"),
						}},
					GraphQueriesTableName: to.Ptr("<graph-queries-table-name>"),
					InstructionSteps: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesInstructionStepsItem{
						{
							Description: to.Ptr("<description>"),
							Instructions: []*armsecurityinsights.InstructionStepsInstructionsItem{
								{
									Type: to.Ptr(armsecurityinsights.SettingType("APIKey")),
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
							Title: to.Ptr("<title>"),
						}},
					Permissions: &armsecurityinsights.Permissions{
						Customs: []*armsecurityinsights.PermissionsCustomsItem{
							{
								Name:        to.Ptr("<name>"),
								Description: to.Ptr("<description>"),
							}},
						ResourceProvider: []*armsecurityinsights.PermissionsResourceProviderItem{
							{
								PermissionsDisplayText: to.Ptr("<permissions-display-text>"),
								Provider:               to.Ptr(armsecurityinsights.ProviderNameMicrosoftOperationalInsightsWorkspaces),
								ProviderDisplayName:    to.Ptr("<provider-display-name>"),
								RequiredPermissions: &armsecurityinsights.RequiredPermissions{
									Delete: to.Ptr(true),
									Read:   to.Ptr(true),
									Write:  to.Ptr(true),
								},
								Scope: to.Ptr(armsecurityinsights.PermissionProviderScopeWorkspace),
							}},
					},
					Publisher: to.Ptr("<publisher>"),
					SampleQueries: []*armsecurityinsights.CodelessUIConnectorConfigPropertiesSampleQueriesItem{
						{
							Description: to.Ptr("<description>"),
							Query:       to.Ptr("<query>"),
						}},
					Title: to.Ptr("<title>"),
				},
				PollingConfig: &armsecurityinsights.CodelessConnectorPollingConfigProperties{
					Auth: &armsecurityinsights.CodelessConnectorPollingAuthProperties{
						APIKeyIdentifier: to.Ptr("<apikey-identifier>"),
						APIKeyName:       to.Ptr("<apikey-name>"),
						AuthType:         to.Ptr("<auth-type>"),
					},
					Paging: &armsecurityinsights.CodelessConnectorPollingPagingProperties{
						PageSizeParaName: to.Ptr("<page-size-para-name>"),
						PagingType:       to.Ptr("<paging-type>"),
					},
					Response: &armsecurityinsights.CodelessConnectorPollingResponseProperties{
						EventsJSONPaths: []*string{
							to.Ptr("$")},
					},
					Request: &armsecurityinsights.CodelessConnectorPollingRequestProperties{
						APIEndpoint: to.Ptr("<apiendpoint>"),
						Headers: map[string]interface{}{
							"Accept":     "application/json",
							"User-Agent": "Scuba",
						},
						HTTPMethod: to.Ptr("<httpmethod>"),
						QueryParameters: map[string]interface{}{
							"phrase": "created:{_QueryWindowStartTime}..{_QueryWindowEndTime}",
						},
						QueryTimeFormat:  to.Ptr("<query-time-format>"),
						QueryWindowInMin: to.Ptr[int32](15),
						RateLimitQPS:     to.Ptr[int32](50),
						RetryCount:       to.Ptr[int32](2),
						TimeoutInSeconds: to.Ptr[int32](60),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
