package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Output_Create_AzureDataExplorer.json
func ExampleOutputsClient_CreateOrReplace_createAnAzureDataExplorerOutput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOutputsClient().CreateOrReplace(ctx, "sjrg", "sjName", "adxOutput", armstreamanalytics.Output{
		Properties: &armstreamanalytics.OutputProperties{
			Datasource: &armstreamanalytics.AzureDataExplorerOutputDataSource{
				Type: to.Ptr("Microsoft.Kusto/clusters/databases"),
				Properties: &armstreamanalytics.AzureDataExplorerOutputDataSourceProperties{
					AuthenticationMode: to.Ptr(armstreamanalytics.AuthenticationModeMsi),
					Cluster:            to.Ptr("https://asakustotest.eastus.kusto.windows.net"),
					Database:           to.Ptr("dbname"),
					Table:              to.Ptr("mytable"),
				},
			},
		},
	}, &armstreamanalytics.OutputsClientCreateOrReplaceOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Output = armstreamanalytics.Output{
	// 	Name: to.Ptr("adxOutput"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
	// 	ID: to.Ptr("/subscriptions/7f31cba8-b597-4129-b158-8f21a7395bd0/resourceGroups/sjrg/providers/Microsoft.StreamAnalytics/streamingjobs/sjName/outputs/adxOutput"),
	// 	Properties: &armstreamanalytics.OutputProperties{
	// 		Datasource: &armstreamanalytics.AzureDataExplorerOutputDataSource{
	// 			Type: to.Ptr("Microsoft.Kusto/clusters/databases"),
	// 			Properties: &armstreamanalytics.AzureDataExplorerOutputDataSourceProperties{
	// 				AuthenticationMode: to.Ptr(armstreamanalytics.AuthenticationModeMsi),
	// 				Cluster: to.Ptr("https://asakustotest.eastus.kusto.windows.net"),
	// 				Database: to.Ptr("dbname"),
	// 				Table: to.Ptr("mytable"),
	// 			},
	// 		},
	// 		Etag: to.Ptr("7b912929-346d-432e-9495-6972dbd63179"),
	// 	},
	// }
}
