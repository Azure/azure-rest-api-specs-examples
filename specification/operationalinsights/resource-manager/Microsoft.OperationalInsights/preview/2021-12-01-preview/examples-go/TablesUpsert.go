package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/TablesUpsert.json
func ExampleTablesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armoperationalinsights.NewTablesClient("00000000-0000-0000-0000-00000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"oiautorest6685",
		"oiautorest6685",
		"AzureNetworkFlow",
		armoperationalinsights.Table{
			Properties: &armoperationalinsights.TableProperties{
				Schema: &armoperationalinsights.Schema{
					Name: to.Ptr("AzureNetworkFlow"),
					Columns: []*armoperationalinsights.Column{
						{
							Name: to.Ptr("MyNewColumn"),
							Type: to.Ptr(armoperationalinsights.ColumnTypeEnumGUID),
						}},
				},
				RetentionInDays:      to.Ptr[int32](45),
				TotalRetentionInDays: to.Ptr[int32](70),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
