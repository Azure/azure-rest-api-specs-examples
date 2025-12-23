package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f0a5127d4e8b1ea6007b0bf9570904d9df860b97/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/DataSourcesCreate.json
func ExampleDataSourcesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataSourcesClient().CreateOrUpdate(ctx, "OIAutoRest5123", "AzTest9724", "AzTestDS774", armoperationalinsights.DataSource{
		Kind: to.Ptr(armoperationalinsights.DataSourceKindAzureActivityLog),
		Properties: map[string]any{
			"LinkedResourceId": "/subscriptions/00000000-0000-0000-0000-00000000000/providers/microsoft.insights/eventtypes/management",
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataSource = armoperationalinsights.DataSource{
	// 	Name: to.Ptr("AzTestDS774"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/workspaces/datasources"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourceGroups/OIAutoRest5123/providers/Microsoft.OperationalInsights/workspaces/AzTest9724/datasources/AzTestDS774"),
	// 	Etag: to.Ptr("W/\"datetime'2017-10-01T08%3A01%3A21.2351243Z'\""),
	// 	Kind: to.Ptr(armoperationalinsights.DataSourceKindAzureActivityLog),
	// 	Properties: map[string]any{
	// 		"linkedResourceId": "/subscriptions/00000000-0000-0000-0000-00000000000/providers/microsoft.insights/eventtypes/management",
	// 	},
	// }
}
