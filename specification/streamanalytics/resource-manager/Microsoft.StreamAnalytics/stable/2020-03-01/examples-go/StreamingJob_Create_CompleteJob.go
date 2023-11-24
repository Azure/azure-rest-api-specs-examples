package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Create_CompleteJob.json
func ExampleStreamingJobsClient_BeginCreateOrReplace_createACompleteStreamingJobAStreamingJobWithATransformationAtLeast1InputAndAtLeast1Output() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStreamingJobsClient().BeginCreateOrReplace(ctx, "sjrg3276", "sj7804", armstreamanalytics.StreamingJob{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"key1":      to.Ptr("value1"),
			"key3":      to.Ptr("value3"),
			"randomKey": to.Ptr("randomValue"),
		},
		Properties: &armstreamanalytics.StreamingJobProperties{
			CompatibilityLevel:                 to.Ptr(armstreamanalytics.CompatibilityLevelOne0),
			DataLocale:                         to.Ptr("en-US"),
			EventsLateArrivalMaxDelayInSeconds: to.Ptr[int32](5),
			EventsOutOfOrderMaxDelayInSeconds:  to.Ptr[int32](0),
			EventsOutOfOrderPolicy:             to.Ptr(armstreamanalytics.EventsOutOfOrderPolicyDrop),
			Functions:                          []*armstreamanalytics.Function{},
			Inputs: []*armstreamanalytics.Input{
				{
					Name: to.Ptr("inputtest"),
					Properties: &armstreamanalytics.StreamInputProperties{
						Type: to.Ptr("Stream"),
						Serialization: &armstreamanalytics.JSONSerialization{
							Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
							Properties: &armstreamanalytics.JSONSerializationProperties{
								Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
							},
						},
						Datasource: &armstreamanalytics.BlobStreamInputDataSource{
							Type: to.Ptr("Microsoft.Storage/Blob"),
							Properties: &armstreamanalytics.BlobStreamInputDataSourceProperties{
								Container:   to.Ptr("containerName"),
								PathPattern: to.Ptr(""),
								StorageAccounts: []*armstreamanalytics.StorageAccount{
									{
										AccountKey:  to.Ptr("yourAccountKey=="),
										AccountName: to.Ptr("yourAccountName"),
									}},
							},
						},
					},
				}},
			OutputErrorPolicy: to.Ptr(armstreamanalytics.OutputErrorPolicyDrop),
			Outputs: []*armstreamanalytics.Output{
				{
					Name: to.Ptr("outputtest"),
					Properties: &armstreamanalytics.OutputProperties{
						Datasource: &armstreamanalytics.AzureSQLDatabaseOutputDataSource{
							Type: to.Ptr("Microsoft.Sql/Server/Database"),
							Properties: &armstreamanalytics.AzureSQLDatabaseOutputDataSourceProperties{
								Database: to.Ptr("databaseName"),
								Password: to.Ptr("userPassword"),
								Server:   to.Ptr("serverName"),
								Table:    to.Ptr("tableName"),
								User:     to.Ptr("<user>"),
							},
						},
					},
				}},
			SKU: &armstreamanalytics.SKU{
				Name: to.Ptr(armstreamanalytics.SKUNameStandard),
			},
			Transformation: &armstreamanalytics.Transformation{
				Name: to.Ptr("transformationtest"),
				Properties: &armstreamanalytics.TransformationProperties{
					Query:          to.Ptr("Select Id, Name from inputtest"),
					StreamingUnits: to.Ptr[int32](1),
				},
			},
		},
	}, &armstreamanalytics.StreamingJobsClientBeginCreateOrReplaceOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
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
	// res.StreamingJob = armstreamanalytics.StreamingJob{
	// 	Name: to.Ptr("sj7804"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg3276/providers/Microsoft.StreamAnalytics/streamingjobs/sj7804"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key3": to.Ptr("value3"),
	// 		"randomKey": to.Ptr("randomValue"),
	// 	},
	// 	Properties: &armstreamanalytics.StreamingJobProperties{
	// 		CompatibilityLevel: to.Ptr(armstreamanalytics.CompatibilityLevelOne0),
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-11T04:37:11.540Z"); return t}()),
	// 		DataLocale: to.Ptr("en-US"),
	// 		EventsLateArrivalMaxDelayInSeconds: to.Ptr[int32](5),
	// 		EventsOutOfOrderMaxDelayInSeconds: to.Ptr[int32](0),
	// 		EventsOutOfOrderPolicy: to.Ptr(armstreamanalytics.EventsOutOfOrderPolicyDrop),
	// 		Functions: []*armstreamanalytics.Function{
	// 		},
	// 		Inputs: []*armstreamanalytics.Input{
	// 			{
	// 				Name: to.Ptr("inputtest"),
	// 				Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs"),
	// 				ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg3276/providers/Microsoft.StreamAnalytics/streamingjobs/sj7804/inputs/inputtest"),
	// 				Properties: &armstreamanalytics.StreamInputProperties{
	// 					Type: to.Ptr("Stream"),
	// 					Etag: to.Ptr("ca88f8fa-605b-4c7f-8695-46f5faa60cd0"),
	// 					Serialization: &armstreamanalytics.JSONSerialization{
	// 						Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
	// 						Properties: &armstreamanalytics.JSONSerializationProperties{
	// 							Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
	// 						},
	// 					},
	// 					Datasource: &armstreamanalytics.BlobStreamInputDataSource{
	// 						Type: to.Ptr("Microsoft.Storage/Blob"),
	// 						Properties: &armstreamanalytics.BlobStreamInputDataSourceProperties{
	// 							Container: to.Ptr("containerName"),
	// 							PathPattern: to.Ptr(""),
	// 							StorageAccounts: []*armstreamanalytics.StorageAccount{
	// 								{
	// 									AccountName: to.Ptr("accountName"),
	// 							}},
	// 						},
	// 					},
	// 				},
	// 		}},
	// 		JobID: to.Ptr("732e4b1d-94a7-43ae-8297-3ad04f1540b9"),
	// 		JobState: to.Ptr("Created"),
	// 		OutputErrorPolicy: to.Ptr(armstreamanalytics.OutputErrorPolicyDrop),
	// 		Outputs: []*armstreamanalytics.Output{
	// 			{
	// 				Name: to.Ptr("outputtest"),
	// 				Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
	// 				ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg3276/providers/Microsoft.StreamAnalytics/streamingjobs/sj7804/outputs/outputtest"),
	// 				Properties: &armstreamanalytics.OutputProperties{
	// 					Datasource: &armstreamanalytics.AzureSQLDatabaseOutputDataSource{
	// 						Type: to.Ptr("Microsoft.Sql/Server/Database"),
	// 						Properties: &armstreamanalytics.AzureSQLDatabaseOutputDataSourceProperties{
	// 							Database: to.Ptr("databaseName"),
	// 							Server: to.Ptr("serverName"),
	// 							Table: to.Ptr("tableName"),
	// 							User: to.Ptr("userName"),
	// 						},
	// 					},
	// 					Etag: to.Ptr("62097c3c-b503-41ff-a56f-196a9598ab90"),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SKU: &armstreamanalytics.SKU{
	// 			Name: to.Ptr(armstreamanalytics.SKUNameStandard),
	// 		},
	// 		Transformation: &armstreamanalytics.Transformation{
	// 			Name: to.Ptr("transformationtest"),
	// 			Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/transformations"),
	// 			ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg3276/providers/Microsoft.StreamAnalytics/streamingjobs/sj7804/transformations/transformationtest"),
	// 			Properties: &armstreamanalytics.TransformationProperties{
	// 				Etag: to.Ptr("91d8fcbe-60b3-49c3-9f21-9942b95602b8"),
	// 				Query: to.Ptr("Select Id, Name from inputtest"),
	// 				StreamingUnits: to.Ptr[int32](1),
	// 			},
	// 		},
	// 	},
	// }
}
