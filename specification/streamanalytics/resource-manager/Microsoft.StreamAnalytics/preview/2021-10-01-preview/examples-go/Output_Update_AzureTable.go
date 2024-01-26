package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Output_Update_AzureTable.json
func ExampleOutputsClient_Update_updateAnAzureTableOutput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOutputsClient().Update(ctx, "sjrg5176", "sj2790", "output958", armstreamanalytics.Output{
		Properties: &armstreamanalytics.OutputProperties{
			Datasource: &armstreamanalytics.AzureTableOutputDataSource{
				Type: to.Ptr("Microsoft.Storage/Table"),
				Properties: &armstreamanalytics.AzureTableOutputDataSourceProperties{
					PartitionKey: to.Ptr("differentPartitionKey"),
				},
			},
		},
	}, &armstreamanalytics.OutputsClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Output = armstreamanalytics.Output{
	// 	Name: to.Ptr("output958"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg5176/providers/Microsoft.StreamAnalytics/streamingjobs/sj2790/outputs/output958"),
	// 	Properties: &armstreamanalytics.OutputProperties{
	// 		Datasource: &armstreamanalytics.AzureTableOutputDataSource{
	// 			Type: to.Ptr("Microsoft.Storage/Table"),
	// 			Properties: &armstreamanalytics.AzureTableOutputDataSourceProperties{
	// 				AccountName: to.Ptr("someAccountName"),
	// 				BatchSize: to.Ptr[int32](25),
	// 				ColumnsToRemove: []*string{
	// 					to.Ptr("column1"),
	// 					to.Ptr("column2")},
	// 					PartitionKey: to.Ptr("differentPartitionKey"),
	// 					RowKey: to.Ptr("rowKey"),
	// 					Table: to.Ptr("samples"),
	// 				},
	// 			},
	// 		},
	// 	}
}
