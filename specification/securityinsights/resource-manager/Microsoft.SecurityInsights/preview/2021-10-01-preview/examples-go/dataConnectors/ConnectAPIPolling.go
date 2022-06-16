package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/dataConnectors/ConnectAPIPolling.json
func ExampleDataConnectorsClient_Connect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewDataConnectorsClient("<subscription-id>", cred, nil)
	_, err = client.Connect(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<data-connector-id>",
		armsecurityinsights.DataConnectorConnectBody{
			APIKey: to.StringPtr("<apikey>"),
			Kind:   armsecurityinsights.ConnectAuthKind("APIKey").ToPtr(),
			RequestConfigUserInputValues: []interface{}{
				map[string]interface{}{
					"displayText":      "Organization Name",
					"placeHolderName":  "{{placeHolder1}}",
					"placeHolderValue": "somePlaceHolderValue",
					"requestObjectKey": "apiEndpoint",
				}},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
