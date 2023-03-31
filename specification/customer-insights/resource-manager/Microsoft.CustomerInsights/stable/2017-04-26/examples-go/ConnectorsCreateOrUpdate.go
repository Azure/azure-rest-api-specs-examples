package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ConnectorsCreateOrUpdate.json
func ExampleConnectorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectorsClient().BeginCreateOrUpdate(ctx, "TestHubRG", "sdkTestHub", "testConnector", armcustomerinsights.ConnectorResourceFormat{
		Properties: &armcustomerinsights.Connector{
			Description: to.Ptr("Test connector"),
			ConnectorProperties: map[string]any{
				"connectionKeyVaultUrl": map[string]any{
					"organizationId":  "XXX",
					"organizationUrl": "https://XXX.crmlivetie.com/",
				},
			},
			ConnectorType: to.Ptr(armcustomerinsights.ConnectorTypesAzureBlob),
			DisplayName:   to.Ptr("testConnector"),
		},
	}, nil)
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
	// res.ConnectorResourceFormat = armcustomerinsights.ConnectorResourceFormat{
	// 	Name: to.Ptr("sdkTestHub/testConnector"),
	// 	Type: to.Ptr("Microsoft.CustomerInsights/hubs/connectors"),
	// 	ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/sdkTestHub/connectors/testConnector"),
	// 	Properties: &armcustomerinsights.Connector{
	// 		Description: to.Ptr("Test connector"),
	// 		ConnectorID: to.Ptr[int32](0),
	// 		ConnectorName: to.Ptr("testConnector"),
	// 		ConnectorProperties: map[string]any{
	// 			"connectionKeyVaultUrl": map[string]any{
	// 				"organizationId": "XXX",
	// 				"organizationUrl": "https://XXX.crmlivetie.com/",
	// 			},
	// 		},
	// 		ConnectorType: to.Ptr(armcustomerinsights.ConnectorTypesAzureBlob),
	// 		DisplayName: to.Ptr("testConnector"),
	// 		State: to.Ptr(armcustomerinsights.ConnectorStatesCreating),
	// 		TenantID: to.Ptr("sdktesthub"),
	// 	},
	// }
}
