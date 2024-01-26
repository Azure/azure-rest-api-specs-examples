package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Input_Create_Reference_Blob_CSV.json
func ExampleInputsClient_CreateOrReplace_createAReferenceBlobInputWithCsvSerialization() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInputsClient().CreateOrReplace(ctx, "sjrg8440", "sj9597", "input7225", armstreamanalytics.Input{
		Properties: &armstreamanalytics.ReferenceInputProperties{
			Type: to.Ptr("Reference"),
			Serialization: &armstreamanalytics.CSVSerialization{
				Type: to.Ptr(armstreamanalytics.EventSerializationTypeCSV),
				Properties: &armstreamanalytics.CSVSerializationProperties{
					Encoding:       to.Ptr(armstreamanalytics.EncodingUTF8),
					FieldDelimiter: to.Ptr(","),
				},
			},
			Datasource: &armstreamanalytics.BlobReferenceInputDataSource{
				Type: to.Ptr("Microsoft.Storage/Blob"),
				Properties: &armstreamanalytics.BlobReferenceInputDataSourceProperties{
					Container:   to.Ptr("state"),
					DateFormat:  to.Ptr("yyyy/MM/dd"),
					PathPattern: to.Ptr("{date}/{time}"),
					StorageAccounts: []*armstreamanalytics.StorageAccount{
						{
							AccountKey:  to.Ptr("someAccountKey=="),
							AccountName: to.Ptr("someAccountName"),
						}},
					TimeFormat:               to.Ptr("HH"),
					BlobName:                 to.Ptr("myblobinput"),
					DeltaPathPattern:         to.Ptr("/testBlob/{date}/delta/{time}/"),
					DeltaSnapshotRefreshRate: to.Ptr("16:14:30"),
					FullSnapshotRefreshRate:  to.Ptr("16:14:30"),
					SourcePartitionCount:     to.Ptr[int32](16),
				},
			},
		},
	}, &armstreamanalytics.InputsClientCreateOrReplaceOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Input = armstreamanalytics.Input{
	// 	Name: to.Ptr("input7225"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg8440/providers/Microsoft.StreamAnalytics/streamingjobs/sj9597/inputs/input7225"),
	// 	Properties: &armstreamanalytics.ReferenceInputProperties{
	// 		Type: to.Ptr("Reference"),
	// 		Serialization: &armstreamanalytics.CSVSerialization{
	// 			Type: to.Ptr(armstreamanalytics.EventSerializationTypeCSV),
	// 			Properties: &armstreamanalytics.CSVSerializationProperties{
	// 				Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
	// 				FieldDelimiter: to.Ptr(","),
	// 			},
	// 		},
	// 		Datasource: &armstreamanalytics.BlobReferenceInputDataSource{
	// 			Type: to.Ptr("Microsoft.Storage/Blob"),
	// 			Properties: &armstreamanalytics.BlobReferenceInputDataSourceProperties{
	// 				Container: to.Ptr("state"),
	// 				DateFormat: to.Ptr("yyyy/MM/dd"),
	// 				PathPattern: to.Ptr("{date}/{time}"),
	// 				StorageAccounts: []*armstreamanalytics.StorageAccount{
	// 					{
	// 						AccountName: to.Ptr("someAccountName"),
	// 				}},
	// 				TimeFormat: to.Ptr("HH"),
	// 			},
	// 		},
	// 	},
	// }
}
