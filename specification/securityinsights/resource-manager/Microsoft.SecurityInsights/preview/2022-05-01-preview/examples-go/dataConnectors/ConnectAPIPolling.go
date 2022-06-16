package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-05-01-preview/examples/dataConnectors/ConnectAPIPolling.json
func ExampleDataConnectorsClient_Connect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewDataConnectorsClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Connect(ctx,
		"myRg",
		"myWorkspace",
		"316ec55e-7138-4d63-ab18-90c8a60fd1c8",
		armsecurityinsights.DataConnectorConnectBody{
			APIKey: to.Ptr("123456789"),
			Kind:   to.Ptr(armsecurityinsights.ConnectAuthKindAPIKey),
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
		log.Fatalf("failed to finish the request: %v", err)
	}
}
