package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/ConnectAPIPolling.json
func ExampleDataConnectorsClient_Connect_connectAnApiPollingDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDataConnectorsClient().Connect(ctx, "myRg", "myWorkspace", "316ec55e-7138-4d63-ab18-90c8a60fd1c8", armsecurityinsights.DataConnectorConnectBody{
		APIKey: to.Ptr("123456789"),
		Kind:   to.Ptr(armsecurityinsights.ConnectAuthKindAPIKey),
		RequestConfigUserInputValues: []any{
			map[string]any{
				"displayText":      "Organization Name",
				"placeHolderName":  "{{placeHolder1}}",
				"placeHolderValue": "somePlaceHolderValue",
				"requestObjectKey": "apiEndpoint",
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
