package armiotoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
)

// Generated from example definition: 2024-11-01/Dataflow_CreateOrUpdate_MaximumSet_Gen.json
func ExampleDataflowClient_BeginCreateOrUpdate_dataflowCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataflowClient().BeginCreateOrUpdate(ctx, "rgiotoperations", "resource-name123", "resource-name123", "resource-name123", armiotoperations.DataflowResource{
		Properties: &armiotoperations.DataflowProperties{
			Mode: to.Ptr(armiotoperations.OperationalModeEnabled),
			Operations: []*armiotoperations.DataflowOperation{
				{
					OperationType: to.Ptr(armiotoperations.OperationTypeSource),
					Name:          to.Ptr("knnafvkwoeakm"),
					SourceSettings: &armiotoperations.DataflowSourceOperationSettings{
						EndpointRef:         to.Ptr("iixotodhvhkkfcfyrkoveslqig"),
						AssetRef:            to.Ptr("zayyykwmckaocywdkohmu"),
						SerializationFormat: to.Ptr(armiotoperations.SourceSerializationFormatJSON),
						SchemaRef:           to.Ptr("pknmdzqll"),
						DataSources: []*string{
							to.Ptr("chkkpymxhp"),
						},
					},
					BuiltInTransformationSettings: &armiotoperations.DataflowBuiltInTransformationSettings{
						SerializationFormat: to.Ptr(armiotoperations.TransformationSerializationFormatDelta),
						SchemaRef:           to.Ptr("mcdc"),
						Datasets: []*armiotoperations.DataflowBuiltInTransformationDataset{
							{
								Key:         to.Ptr("qsfqcgxaxnhfumrsdsokwyv"),
								Description: to.Ptr("Lorem ipsum odor amet, consectetuer adipiscing elit."),
								SchemaRef:   to.Ptr("n"),
								Inputs: []*string{
									to.Ptr("mosffpsslifkq"),
								},
								Expression: to.Ptr("aatbwomvflemsxialv"),
							},
						},
						Filter: []*armiotoperations.DataflowBuiltInTransformationFilter{
							{
								Type:        to.Ptr(armiotoperations.FilterTypeFilter),
								Description: to.Ptr("Lorem ipsum odor amet, consectetuer adipiscing elit."),
								Inputs: []*string{
									to.Ptr("sxmjkbntgb"),
								},
								Expression: to.Ptr("n"),
							},
						},
						Map: []*armiotoperations.DataflowBuiltInTransformationMap{
							{
								Type:        to.Ptr(armiotoperations.DataflowMappingTypeNewProperties),
								Description: to.Ptr("Lorem ipsum odor amet, consectetuer adipiscing elit."),
								Inputs: []*string{
									to.Ptr("xsbxuk"),
								},
								Expression: to.Ptr("txoiltogsarwkzalsphvlmt"),
								Output:     to.Ptr("nvgtmkfl"),
							},
						},
					},
					DestinationSettings: &armiotoperations.DataflowDestinationOperationSettings{
						EndpointRef:     to.Ptr("kybkchnzimerguekuvqlqiqdvvrt"),
						DataDestination: to.Ptr("cbrh"),
					},
				},
			},
		},
		ExtendedLocation: &armiotoperations.ExtendedLocation{
			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
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
	// res = armiotoperations.DataflowClientCreateOrUpdateResponse{
	// 	DataflowResource: &armiotoperations.DataflowResource{
	// 		Properties: &armiotoperations.DataflowProperties{
	// 			Mode: to.Ptr(armiotoperations.OperationalModeEnabled),
	// 			Operations: []*armiotoperations.DataflowOperation{
	// 				{
	// 					OperationType: to.Ptr(armiotoperations.OperationTypeSource),
	// 					Name: to.Ptr("knnafvkwoeakm"),
	// 					SourceSettings: &armiotoperations.DataflowSourceOperationSettings{
	// 						EndpointRef: to.Ptr("iixotodhvhkkfcfyrkoveslqig"),
	// 						AssetRef: to.Ptr("zayyykwmckaocywdkohmu"),
	// 						SerializationFormat: to.Ptr(armiotoperations.SourceSerializationFormatJSON),
	// 						SchemaRef: to.Ptr("pknmdzqll"),
	// 						DataSources: []*string{
	// 							to.Ptr("chkkpymxhp"),
	// 						},
	// 					},
	// 					BuiltInTransformationSettings: &armiotoperations.DataflowBuiltInTransformationSettings{
	// 						SerializationFormat: to.Ptr(armiotoperations.TransformationSerializationFormatDelta),
	// 						SchemaRef: to.Ptr("mcdc"),
	// 						Datasets: []*armiotoperations.DataflowBuiltInTransformationDataset{
	// 							{
	// 								Key: to.Ptr("qsfqcgxaxnhfumrsdsokwyv"),
	// 								Description: to.Ptr("Lorem ipsum odor amet, consectetuer adipiscing elit."),
	// 								SchemaRef: to.Ptr("n"),
	// 								Inputs: []*string{
	// 									to.Ptr("mosffpsslifkq"),
	// 								},
	// 								Expression: to.Ptr("aatbwomvflemsxialv"),
	// 							},
	// 						},
	// 						Filter: []*armiotoperations.DataflowBuiltInTransformationFilter{
	// 							{
	// 								Type: to.Ptr(armiotoperations.FilterTypeFilter),
	// 								Description: to.Ptr("Lorem ipsum odor amet, consectetuer adipiscing elit."),
	// 								Inputs: []*string{
	// 									to.Ptr("sxmjkbntgb"),
	// 								},
	// 								Expression: to.Ptr("n"),
	// 							},
	// 						},
	// 						Map: []*armiotoperations.DataflowBuiltInTransformationMap{
	// 							{
	// 								Type: to.Ptr(armiotoperations.DataflowMappingTypeNewProperties),
	// 								Description: to.Ptr("Lorem ipsum odor amet, consectetuer adipiscing elit."),
	// 								Inputs: []*string{
	// 									to.Ptr("xsbxuk"),
	// 								},
	// 								Expression: to.Ptr("txoiltogsarwkzalsphvlmt"),
	// 								Output: to.Ptr("nvgtmkfl"),
	// 							},
	// 						},
	// 					},
	// 					DestinationSettings: &armiotoperations.DataflowDestinationOperationSettings{
	// 						EndpointRef: to.Ptr("kybkchnzimerguekuvqlqiqdvvrt"),
	// 						DataDestination: to.Ptr("cbrh"),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/dataflowProfiles/resource-name123/dataflows/resource-name123"),
	// 		Name: to.Ptr("jxhcpwgfkxqasbexkookvxk"),
	// 		Type: to.Ptr("zkuozvgjseokfchkscoswthzjdry"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("gnicpuszwd"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}
