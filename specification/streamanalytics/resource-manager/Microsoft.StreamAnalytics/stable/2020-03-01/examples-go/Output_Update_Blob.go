package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Update_Blob.json
func ExampleOutputsClient_Update_updateABlobOutputWithCsvSerialization() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOutputsClient().Update(ctx, "sjrg5023", "sj900", "output1623", armstreamanalytics.Output{
		Properties: &armstreamanalytics.OutputProperties{
			Datasource: &armstreamanalytics.BlobOutputDataSource{
				Type: to.Ptr("Microsoft.Storage/Blob"),
				Properties: &armstreamanalytics.BlobOutputDataSourceProperties{
					Container: to.Ptr("differentContainer"),
				},
			},
			Serialization: &armstreamanalytics.CSVSerialization{
				Type: to.Ptr(armstreamanalytics.EventSerializationTypeCSV),
				Properties: &armstreamanalytics.CSVSerializationProperties{
					Encoding:       to.Ptr(armstreamanalytics.EncodingUTF8),
					FieldDelimiter: to.Ptr("|"),
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
	// 	Name: to.Ptr("output1623"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg5023/providers/Microsoft.StreamAnalytics/streamingjobs/sj900/outputs/output1623"),
	// 	Properties: &armstreamanalytics.OutputProperties{
	// 		Datasource: &armstreamanalytics.BlobOutputDataSource{
	// 			Type: to.Ptr("Microsoft.Storage/Blob"),
	// 			Properties: &armstreamanalytics.BlobOutputDataSourceProperties{
	// 				Container: to.Ptr("differentContainer"),
	// 				DateFormat: to.Ptr("yyyy/MM/dd"),
	// 				PathPattern: to.Ptr("{date}/{time}"),
	// 				StorageAccounts: []*armstreamanalytics.StorageAccount{
	// 					{
	// 						AccountName: to.Ptr("someAccountName"),
	// 				}},
	// 				TimeFormat: to.Ptr("HH"),
	// 			},
	// 		},
	// 		Serialization: &armstreamanalytics.CSVSerialization{
	// 			Type: to.Ptr(armstreamanalytics.EventSerializationTypeCSV),
	// 			Properties: &armstreamanalytics.CSVSerializationProperties{
	// 				Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
	// 				FieldDelimiter: to.Ptr("|"),
	// 			},
	// 		},
	// 	},
	// }
}
