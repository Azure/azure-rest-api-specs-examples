package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Output_Get_DeltaLake.json
func ExampleOutputsClient_Get_getADeltaLakeOutput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOutputsClient().Get(ctx, "sjrg", "sjName", "output1221", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Output = armstreamanalytics.Output{
	// 	Name: to.Ptr("output1221"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg/providers/Microsoft.StreamAnalytics/streamingjobs/sjName/outputs/output1221"),
	// 	Properties: &armstreamanalytics.OutputProperties{
	// 		Datasource: &armstreamanalytics.BlobOutputDataSource{
	// 			Type: to.Ptr("Microsoft.Storage/Blob"),
	// 			Properties: &armstreamanalytics.BlobOutputDataSourceProperties{
	// 				Container: to.Ptr("deltaoutput"),
	// 				StorageAccounts: []*armstreamanalytics.StorageAccount{
	// 					{
	// 						AccountName: to.Ptr("someAccountName"),
	// 				}},
	// 			},
	// 		},
	// 		Serialization: &armstreamanalytics.DeltaSerialization{
	// 			Type: to.Ptr(armstreamanalytics.EventSerializationTypeDelta),
	// 			Properties: &armstreamanalytics.DeltaSerializationProperties{
	// 				DeltaTablePath: to.Ptr("/folder1/table1"),
	// 				PartitionColumns: []*string{
	// 					to.Ptr("column1")},
	// 				},
	// 			},
	// 		},
	// 	}
}
