package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/OperationList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.OperationListResult = armnetapp.OperationListResult{
		// 	Value: []*armnetapp.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.NetApp/register/action"),
		// 			Display: &armnetapp.OperationDisplay{
		// 				Description: to.Ptr("Subscription Registration Action"),
		// 				Operation: to.Ptr("Subscription Registration Action"),
		// 				Provider: to.Ptr("Microsoft.NetApp"),
		// 				Resource: to.Ptr("Subscription"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/read"),
		// 			Display: &armnetapp.OperationDisplay{
		// 				Description: to.Ptr("Reads a volume resource."),
		// 				Operation: to.Ptr("Read volume resource"),
		// 				Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 				Resource: to.Ptr("Volumes resource type"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/write"),
		// 			Display: &armnetapp.OperationDisplay{
		// 				Description: to.Ptr("Writes a volume resource."),
		// 				Operation: to.Ptr("Write volume resource"),
		// 				Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 				Resource: to.Ptr("Volumes resource type"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/delete"),
		// 			Display: &armnetapp.OperationDisplay{
		// 				Description: to.Ptr("Deletes a volume resource."),
		// 				Operation: to.Ptr("Delete volume resource"),
		// 				Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 				Resource: to.Ptr("Volumes resource type"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/Revert/action"),
		// 			Display: &armnetapp.OperationDisplay{
		// 				Description: to.Ptr("Revert volume to specific snapshot"),
		// 				Operation: to.Ptr("Revert volume resource"),
		// 				Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 				Resource: to.Ptr("Volumes resource type"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/BreakReplication/action"),
		// 			Display: &armnetapp.OperationDisplay{
		// 				Description: to.Ptr("Break volume replication relations"),
		// 				Operation: to.Ptr("Break volume replication resource"),
		// 				Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 				Resource: to.Ptr("Volumes resource type"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/ReplicationStatus/action"),
		// 			Display: &armnetapp.OperationDisplay{
		// 				Description: to.Ptr("Reads the statuses of the Volume Replication."),
		// 				Operation: to.Ptr("Read Volume Replication Status."),
		// 				Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 				Resource: to.Ptr("Volumes resource type"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/ListReplications/action"),
		// 			Display: &armnetapp.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 				Resource: to.Ptr("Volumes resource type"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/ReInitializeReplication/action"),
		// 			Display: &armnetapp.OperationDisplay{
		// 				Description: to.Ptr("Attempts to re-initialize an uninitialized replication"),
		// 				Operation: to.Ptr("Re-Initialize replication"),
		// 				Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 				Resource: to.Ptr("Volumes resource type"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armnetapp.OperationDisplay{
		// 				Description: to.Ptr("Gets the available metrics for Volume resource."),
		// 				Operation: to.Ptr("Read volume metric definitions."),
		// 				Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 				Resource: to.Ptr("Volumes resource type"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 			Properties: &armnetapp.OperationProperties{
		// 				ServiceSpecification: &armnetapp.ServiceSpecification{
		// 					MetricSpecifications: []*armnetapp.MetricSpecification{
		// 						{
		// 							Name: to.Ptr("AverageReadLatency"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Dimensions: []*armnetapp.Dimension{
		// 							},
		// 							DisplayDescription: to.Ptr("Average read latency in milliseconds per operation"),
		// 							DisplayName: to.Ptr("Average read latency"),
		// 							EnableRegionalMdmAccount: to.Ptr(true),
		// 							FillGapWithZero: to.Ptr(false),
		// 							InternalMetricName: to.Ptr("AverageReadLatency"),
		// 							IsInternal: to.Ptr(false),
		// 							SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 							SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 							SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 								to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 								SupportedTimeGrainTypes: []*string{
		// 									to.Ptr("PT5M"),
		// 									to.Ptr("PT15M"),
		// 									to.Ptr("PT30M"),
		// 									to.Ptr("PT1H"),
		// 									to.Ptr("PT6H"),
		// 									to.Ptr("PT12H"),
		// 									to.Ptr("P1D")},
		// 									Unit: to.Ptr("MilliSeconds"),
		// 								},
		// 								{
		// 									Name: to.Ptr("AverageWriteLatency"),
		// 									AggregationType: to.Ptr("Average"),
		// 									Dimensions: []*armnetapp.Dimension{
		// 									},
		// 									DisplayDescription: to.Ptr("Average write latency in milliseconds per operation"),
		// 									DisplayName: to.Ptr("Average write latency"),
		// 									EnableRegionalMdmAccount: to.Ptr(true),
		// 									FillGapWithZero: to.Ptr(false),
		// 									InternalMetricName: to.Ptr("AverageWriteLatency"),
		// 									IsInternal: to.Ptr(false),
		// 									SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 									SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 									SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 										to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 										SupportedTimeGrainTypes: []*string{
		// 											to.Ptr("PT5M"),
		// 											to.Ptr("PT15M"),
		// 											to.Ptr("PT30M"),
		// 											to.Ptr("PT1H"),
		// 											to.Ptr("PT6H"),
		// 											to.Ptr("PT12H"),
		// 											to.Ptr("P1D")},
		// 											Unit: to.Ptr("MilliSeconds"),
		// 										},
		// 										{
		// 											Name: to.Ptr("VolumeLogicalSize"),
		// 											AggregationType: to.Ptr("Average"),
		// 											Dimensions: []*armnetapp.Dimension{
		// 											},
		// 											DisplayDescription: to.Ptr("Logical size of the volume (used bytes)"),
		// 											DisplayName: to.Ptr("Volume Consumed Size"),
		// 											EnableRegionalMdmAccount: to.Ptr(true),
		// 											FillGapWithZero: to.Ptr(false),
		// 											InternalMetricName: to.Ptr("VolumeLogicalSize"),
		// 											IsInternal: to.Ptr(false),
		// 											SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 											SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 											SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 												to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 												SupportedTimeGrainTypes: []*string{
		// 													to.Ptr("PT5M"),
		// 													to.Ptr("PT15M"),
		// 													to.Ptr("PT30M"),
		// 													to.Ptr("PT1H"),
		// 													to.Ptr("PT6H"),
		// 													to.Ptr("PT12H"),
		// 													to.Ptr("P1D")},
		// 													Unit: to.Ptr("Bytes"),
		// 												},
		// 												{
		// 													Name: to.Ptr("VolumeSnapshotSize"),
		// 													AggregationType: to.Ptr("Average"),
		// 													Dimensions: []*armnetapp.Dimension{
		// 													},
		// 													DisplayDescription: to.Ptr("Size of all snapshots in volume"),
		// 													DisplayName: to.Ptr("Volume snapshot size"),
		// 													EnableRegionalMdmAccount: to.Ptr(true),
		// 													FillGapWithZero: to.Ptr(false),
		// 													InternalMetricName: to.Ptr("VolumeSnapshotSize"),
		// 													IsInternal: to.Ptr(false),
		// 													SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 													SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 													SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 														to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 														SupportedTimeGrainTypes: []*string{
		// 															to.Ptr("PT5M"),
		// 															to.Ptr("PT15M"),
		// 															to.Ptr("PT30M"),
		// 															to.Ptr("PT1H"),
		// 															to.Ptr("PT6H"),
		// 															to.Ptr("PT12H"),
		// 															to.Ptr("P1D")},
		// 															Unit: to.Ptr("Bytes"),
		// 														},
		// 														{
		// 															Name: to.Ptr("ReadIops"),
		// 															AggregationType: to.Ptr("Average"),
		// 															Dimensions: []*armnetapp.Dimension{
		// 															},
		// 															DisplayDescription: to.Ptr("Read In/out operations per second"),
		// 															DisplayName: to.Ptr("Read iops"),
		// 															EnableRegionalMdmAccount: to.Ptr(true),
		// 															FillGapWithZero: to.Ptr(false),
		// 															InternalMetricName: to.Ptr("ReadIops"),
		// 															IsInternal: to.Ptr(false),
		// 															SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 															SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 															SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																SupportedTimeGrainTypes: []*string{
		// 																	to.Ptr("PT5M"),
		// 																	to.Ptr("PT15M"),
		// 																	to.Ptr("PT30M"),
		// 																	to.Ptr("PT1H"),
		// 																	to.Ptr("PT6H"),
		// 																	to.Ptr("PT12H"),
		// 																	to.Ptr("P1D")},
		// 																	Unit: to.Ptr("CountPerSecond"),
		// 																},
		// 																{
		// 																	Name: to.Ptr("WriteIops"),
		// 																	AggregationType: to.Ptr("Average"),
		// 																	Dimensions: []*armnetapp.Dimension{
		// 																	},
		// 																	DisplayDescription: to.Ptr("Write In/out operations per second"),
		// 																	DisplayName: to.Ptr("Write iops"),
		// 																	EnableRegionalMdmAccount: to.Ptr(true),
		// 																	FillGapWithZero: to.Ptr(false),
		// 																	InternalMetricName: to.Ptr("WriteIops"),
		// 																	IsInternal: to.Ptr(false),
		// 																	SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																	SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																	SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																		to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																		SupportedTimeGrainTypes: []*string{
		// 																			to.Ptr("PT5M"),
		// 																			to.Ptr("PT15M"),
		// 																			to.Ptr("PT30M"),
		// 																			to.Ptr("PT1H"),
		// 																			to.Ptr("PT6H"),
		// 																			to.Ptr("PT12H"),
		// 																			to.Ptr("P1D")},
		// 																			Unit: to.Ptr("CountPerSecond"),
		// 																		},
		// 																		{
		// 																			Name: to.Ptr("VolumeAllocatedSize"),
		// 																			AggregationType: to.Ptr("Average"),
		// 																			Dimensions: []*armnetapp.Dimension{
		// 																			},
		// 																			DisplayDescription: to.Ptr("The provisioned size of a volume"),
		// 																			DisplayName: to.Ptr("Volume allocated size"),
		// 																			EnableRegionalMdmAccount: to.Ptr(true),
		// 																			FillGapWithZero: to.Ptr(false),
		// 																			InternalMetricName: to.Ptr("VolumeAllocatedSize"),
		// 																			IsInternal: to.Ptr(false),
		// 																			SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																			SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																			SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																				to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																				SupportedTimeGrainTypes: []*string{
		// 																					to.Ptr("PT5M"),
		// 																					to.Ptr("PT15M"),
		// 																					to.Ptr("PT30M"),
		// 																					to.Ptr("PT1H"),
		// 																					to.Ptr("PT6H"),
		// 																					to.Ptr("PT12H"),
		// 																					to.Ptr("P1D")},
		// 																					Unit: to.Ptr("Bytes"),
		// 																				},
		// 																				{
		// 																					Name: to.Ptr("VolumeCoolTierSize"),
		// 																					AggregationType: to.Ptr("Average"),
		// 																					Dimensions: []*armnetapp.Dimension{
		// 																					},
		// 																					DisplayDescription: to.Ptr("Volume Footprint for Cool Tier"),
		// 																					DisplayName: to.Ptr("Volume cool tier size"),
		// 																					EnableRegionalMdmAccount: to.Ptr(true),
		// 																					FillGapWithZero: to.Ptr(false),
		// 																					InternalMetricName: to.Ptr("VolumeCoolTierSize"),
		// 																					IsInternal: to.Ptr(false),
		// 																					SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																					SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																					SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																						to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																						SupportedTimeGrainTypes: []*string{
		// 																							to.Ptr("PT5M"),
		// 																							to.Ptr("PT15M"),
		// 																							to.Ptr("PT30M"),
		// 																							to.Ptr("PT1H"),
		// 																							to.Ptr("PT6H"),
		// 																							to.Ptr("PT12H"),
		// 																							to.Ptr("P1D")},
		// 																							Unit: to.Ptr("Bytes"),
		// 																						},
		// 																						{
		// 																							Name: to.Ptr("VolumeCoolTierDataReadSize"),
		// 																							AggregationType: to.Ptr("Average"),
		// 																							Dimensions: []*armnetapp.Dimension{
		// 																							},
		// 																							DisplayDescription: to.Ptr("Data read in using GET per volume"),
		// 																							DisplayName: to.Ptr("Volume cool tier data read size"),
		// 																							EnableRegionalMdmAccount: to.Ptr(true),
		// 																							FillGapWithZero: to.Ptr(false),
		// 																							InternalMetricName: to.Ptr("VolumeCoolTierDataReadSize"),
		// 																							IsInternal: to.Ptr(false),
		// 																							SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																							SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																							SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																								to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																								SupportedTimeGrainTypes: []*string{
		// 																									to.Ptr("PT5M"),
		// 																									to.Ptr("PT15M"),
		// 																									to.Ptr("PT30M"),
		// 																									to.Ptr("PT1H"),
		// 																									to.Ptr("PT6H"),
		// 																									to.Ptr("PT12H"),
		// 																									to.Ptr("P1D")},
		// 																									Unit: to.Ptr("Bytes"),
		// 																								},
		// 																								{
		// 																									Name: to.Ptr("VolumeCoolTierDataWriteSize"),
		// 																									AggregationType: to.Ptr("Average"),
		// 																									Dimensions: []*armnetapp.Dimension{
		// 																									},
		// 																									DisplayDescription: to.Ptr("Data tiered out using PUT per volume"),
		// 																									DisplayName: to.Ptr("Volume cool tier data write size"),
		// 																									EnableRegionalMdmAccount: to.Ptr(true),
		// 																									FillGapWithZero: to.Ptr(false),
		// 																									InternalMetricName: to.Ptr("VolumeCoolTierDataWriteSize"),
		// 																									IsInternal: to.Ptr(false),
		// 																									SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																									SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																									SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																										to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																										SupportedTimeGrainTypes: []*string{
		// 																											to.Ptr("PT5M"),
		// 																											to.Ptr("PT15M"),
		// 																											to.Ptr("PT30M"),
		// 																											to.Ptr("PT1H"),
		// 																											to.Ptr("PT6H"),
		// 																											to.Ptr("PT12H"),
		// 																											to.Ptr("P1D")},
		// 																											Unit: to.Ptr("Bytes"),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("XregionReplicationLastTransferDuration"),
		// 																											AggregationType: to.Ptr("Average"),
		// 																											Dimensions: []*armnetapp.Dimension{
		// 																											},
		// 																											DisplayDescription: to.Ptr("The amount of time in seconds it took for the last transfer to complete."),
		// 																											DisplayName: to.Ptr("Volume replication last transfer duration"),
		// 																											EnableRegionalMdmAccount: to.Ptr(true),
		// 																											FillGapWithZero: to.Ptr(false),
		// 																											InternalMetricName: to.Ptr("XregionReplicationLastTransferDuration"),
		// 																											IsInternal: to.Ptr(false),
		// 																											SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																											SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																											SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																												to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																												SupportedTimeGrainTypes: []*string{
		// 																													to.Ptr("PT5M"),
		// 																													to.Ptr("PT15M"),
		// 																													to.Ptr("PT30M"),
		// 																													to.Ptr("PT1H"),
		// 																													to.Ptr("PT6H"),
		// 																													to.Ptr("PT12H"),
		// 																													to.Ptr("P1D")},
		// 																													Unit: to.Ptr("Seconds"),
		// 																												},
		// 																												{
		// 																													Name: to.Ptr("XregionReplicationLastTransferSize"),
		// 																													AggregationType: to.Ptr("Average"),
		// 																													Dimensions: []*armnetapp.Dimension{
		// 																													},
		// 																													DisplayDescription: to.Ptr("The total number of bytes transferred as part of the last transfer."),
		// 																													DisplayName: to.Ptr("Volume replication last transfer size"),
		// 																													EnableRegionalMdmAccount: to.Ptr(true),
		// 																													FillGapWithZero: to.Ptr(false),
		// 																													InternalMetricName: to.Ptr("XregionReplicationLastTransferSize"),
		// 																													IsInternal: to.Ptr(false),
		// 																													SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																													SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																													SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																														to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																														SupportedTimeGrainTypes: []*string{
		// 																															to.Ptr("PT5M"),
		// 																															to.Ptr("PT15M"),
		// 																															to.Ptr("PT30M"),
		// 																															to.Ptr("PT1H"),
		// 																															to.Ptr("PT6H"),
		// 																															to.Ptr("PT12H"),
		// 																															to.Ptr("P1D")},
		// 																															Unit: to.Ptr("Bytes"),
		// 																														},
		// 																														{
		// 																															Name: to.Ptr("XregionReplicationHealthy"),
		// 																															AggregationType: to.Ptr("Average"),
		// 																															Dimensions: []*armnetapp.Dimension{
		// 																															},
		// 																															DisplayDescription: to.Ptr("Condition of the relationship, 1 or 0."),
		// 																															DisplayName: to.Ptr("Is volume replication status healthy"),
		// 																															EnableRegionalMdmAccount: to.Ptr(true),
		// 																															FillGapWithZero: to.Ptr(false),
		// 																															InternalMetricName: to.Ptr("XregionReplicationHealthy"),
		// 																															IsInternal: to.Ptr(false),
		// 																															SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																															SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																															SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																SupportedTimeGrainTypes: []*string{
		// 																																	to.Ptr("PT5M"),
		// 																																	to.Ptr("PT15M"),
		// 																																	to.Ptr("PT30M"),
		// 																																	to.Ptr("PT1H"),
		// 																																	to.Ptr("PT6H"),
		// 																																	to.Ptr("PT12H"),
		// 																																	to.Ptr("P1D")},
		// 																																	Unit: to.Ptr("Count"),
		// 																																},
		// 																																{
		// 																																	Name: to.Ptr("XregionReplicationLagTime"),
		// 																																	AggregationType: to.Ptr("Average"),
		// 																																	Dimensions: []*armnetapp.Dimension{
		// 																																	},
		// 																																	DisplayDescription: to.Ptr("The amount of time in seconds by which the data on the mirror lags behind the source."),
		// 																																	DisplayName: to.Ptr("Volume replication lag time"),
		// 																																	EnableRegionalMdmAccount: to.Ptr(true),
		// 																																	FillGapWithZero: to.Ptr(false),
		// 																																	InternalMetricName: to.Ptr("XregionReplicationLagTime"),
		// 																																	IsInternal: to.Ptr(false),
		// 																																	SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																	SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																	SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																		to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																		SupportedTimeGrainTypes: []*string{
		// 																																			to.Ptr("PT5M"),
		// 																																			to.Ptr("PT15M"),
		// 																																			to.Ptr("PT30M"),
		// 																																			to.Ptr("PT1H"),
		// 																																			to.Ptr("PT6H"),
		// 																																			to.Ptr("PT12H"),
		// 																																			to.Ptr("P1D")},
		// 																																			Unit: to.Ptr("Seconds"),
		// 																																		},
		// 																																		{
		// 																																			Name: to.Ptr("XregionReplicationTotalTransferBytes"),
		// 																																			AggregationType: to.Ptr("Average"),
		// 																																			Dimensions: []*armnetapp.Dimension{
		// 																																			},
		// 																																			DisplayDescription: to.Ptr("Cumulative bytes transferred for the relationship."),
		// 																																			DisplayName: to.Ptr("Volume replication total transfer"),
		// 																																			EnableRegionalMdmAccount: to.Ptr(true),
		// 																																			FillGapWithZero: to.Ptr(false),
		// 																																			InternalMetricName: to.Ptr("XregionReplicationTotalTransferBytes"),
		// 																																			IsInternal: to.Ptr(false),
		// 																																			SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																			SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																			SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																				to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																				SupportedTimeGrainTypes: []*string{
		// 																																					to.Ptr("PT5M"),
		// 																																					to.Ptr("PT15M"),
		// 																																					to.Ptr("PT30M"),
		// 																																					to.Ptr("PT1H"),
		// 																																					to.Ptr("PT6H"),
		// 																																					to.Ptr("PT12H"),
		// 																																					to.Ptr("P1D")},
		// 																																					Unit: to.Ptr("Bytes"),
		// 																																				},
		// 																																				{
		// 																																					Name: to.Ptr("XregionReplicationRelationshipProgress"),
		// 																																					AggregationType: to.Ptr("Average"),
		// 																																					Dimensions: []*armnetapp.Dimension{
		// 																																					},
		// 																																					DisplayDescription: to.Ptr("Total amount of data transferred for the current transfer operation."),
		// 																																					DisplayName: to.Ptr("Volume replication progress"),
		// 																																					EnableRegionalMdmAccount: to.Ptr(true),
		// 																																					FillGapWithZero: to.Ptr(false),
		// 																																					InternalMetricName: to.Ptr("XregionReplicationRelationshipProgress"),
		// 																																					IsInternal: to.Ptr(false),
		// 																																					SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																					SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																					SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																						to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																						SupportedTimeGrainTypes: []*string{
		// 																																							to.Ptr("PT5M"),
		// 																																							to.Ptr("PT15M"),
		// 																																							to.Ptr("PT30M"),
		// 																																							to.Ptr("PT1H"),
		// 																																							to.Ptr("PT6H"),
		// 																																							to.Ptr("PT12H"),
		// 																																							to.Ptr("P1D")},
		// 																																							Unit: to.Ptr("Bytes"),
		// 																																						},
		// 																																						{
		// 																																							Name: to.Ptr("XregionReplicationRelationshipTransferring"),
		// 																																							AggregationType: to.Ptr("Average"),
		// 																																							Dimensions: []*armnetapp.Dimension{
		// 																																							},
		// 																																							DisplayDescription: to.Ptr("Whether the status of the Volume Replication is 'transferring'."),
		// 																																							DisplayName: to.Ptr("Is volume replication transferring"),
		// 																																							EnableRegionalMdmAccount: to.Ptr(true),
		// 																																							FillGapWithZero: to.Ptr(false),
		// 																																							InternalMetricName: to.Ptr("XregionReplicationRelationshipTransferring"),
		// 																																							IsInternal: to.Ptr(false),
		// 																																							SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																							SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																							SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																								to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																								SupportedTimeGrainTypes: []*string{
		// 																																									to.Ptr("PT5M"),
		// 																																									to.Ptr("PT15M"),
		// 																																									to.Ptr("PT30M"),
		// 																																									to.Ptr("PT1H"),
		// 																																									to.Ptr("PT6H"),
		// 																																									to.Ptr("PT12H"),
		// 																																									to.Ptr("P1D")},
		// 																																									Unit: to.Ptr("Count"),
		// 																																								},
		// 																																								{
		// 																																									Name: to.Ptr("CbsVolumeLogicalBackupBytes"),
		// 																																									AggregationType: to.Ptr("Average"),
		// 																																									Dimensions: []*armnetapp.Dimension{
		// 																																									},
		// 																																									DisplayDescription: to.Ptr("Total bytes backed up for this Volume."),
		// 																																									DisplayName: to.Ptr("Volume Backup Bytes"),
		// 																																									EnableRegionalMdmAccount: to.Ptr(true),
		// 																																									FillGapWithZero: to.Ptr(false),
		// 																																									InternalMetricName: to.Ptr("CbsVolumeLogicalBackupBytes"),
		// 																																									IsInternal: to.Ptr(false),
		// 																																									SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																									SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																									SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																										to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																										SupportedTimeGrainTypes: []*string{
		// 																																											to.Ptr("PT5M"),
		// 																																											to.Ptr("PT15M"),
		// 																																											to.Ptr("PT30M"),
		// 																																											to.Ptr("PT1H"),
		// 																																											to.Ptr("PT6H"),
		// 																																											to.Ptr("PT12H"),
		// 																																											to.Ptr("P1D")},
		// 																																											Unit: to.Ptr("Bytes"),
		// 																																										},
		// 																																										{
		// 																																											Name: to.Ptr("CbsVolumeProtected"),
		// 																																											AggregationType: to.Ptr("Average"),
		// 																																											Dimensions: []*armnetapp.Dimension{
		// 																																											},
		// 																																											DisplayDescription: to.Ptr("Is backup enabled for the volume? 1 if yes, 0 if no."),
		// 																																											DisplayName: to.Ptr("Is Volume Backup Enabled"),
		// 																																											EnableRegionalMdmAccount: to.Ptr(true),
		// 																																											FillGapWithZero: to.Ptr(false),
		// 																																											InternalMetricName: to.Ptr("CbsVolumeProtected"),
		// 																																											IsInternal: to.Ptr(false),
		// 																																											SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																											SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																											SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																												to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																												SupportedTimeGrainTypes: []*string{
		// 																																													to.Ptr("PT5M"),
		// 																																													to.Ptr("PT15M"),
		// 																																													to.Ptr("PT30M"),
		// 																																													to.Ptr("PT1H"),
		// 																																													to.Ptr("PT6H"),
		// 																																													to.Ptr("PT12H"),
		// 																																													to.Ptr("P1D")},
		// 																																													Unit: to.Ptr("Count"),
		// 																																												},
		// 																																												{
		// 																																													Name: to.Ptr("CbsVolumeBackupActive"),
		// 																																													AggregationType: to.Ptr("Average"),
		// 																																													Dimensions: []*armnetapp.Dimension{
		// 																																													},
		// 																																													DisplayDescription: to.Ptr("Is the backup policy suspended for the volume? 0 if yes, 1 if no."),
		// 																																													DisplayName: to.Ptr("Is Volume Backup suspended"),
		// 																																													EnableRegionalMdmAccount: to.Ptr(true),
		// 																																													FillGapWithZero: to.Ptr(false),
		// 																																													InternalMetricName: to.Ptr("CbsVolumeBackupActive"),
		// 																																													IsInternal: to.Ptr(false),
		// 																																													SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																													SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																													SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																														to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																														SupportedTimeGrainTypes: []*string{
		// 																																															to.Ptr("PT5M"),
		// 																																															to.Ptr("PT15M"),
		// 																																															to.Ptr("PT30M"),
		// 																																															to.Ptr("PT1H"),
		// 																																															to.Ptr("PT6H"),
		// 																																															to.Ptr("PT12H"),
		// 																																															to.Ptr("P1D")},
		// 																																															Unit: to.Ptr("Count"),
		// 																																														},
		// 																																														{
		// 																																															Name: to.Ptr("CbsVolumeOperationTransferredBytes"),
		// 																																															AggregationType: to.Ptr("Average"),
		// 																																															Dimensions: []*armnetapp.Dimension{
		// 																																															},
		// 																																															DisplayDescription: to.Ptr("Total bytes transferred for last backup or restore operation."),
		// 																																															DisplayName: to.Ptr("Volume Backup Last Transferred Bytes"),
		// 																																															EnableRegionalMdmAccount: to.Ptr(true),
		// 																																															FillGapWithZero: to.Ptr(false),
		// 																																															InternalMetricName: to.Ptr("CbsVolumeOperationTransferredBytes"),
		// 																																															IsInternal: to.Ptr(false),
		// 																																															SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																															SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																															SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																																to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																																SupportedTimeGrainTypes: []*string{
		// 																																																	to.Ptr("PT5M"),
		// 																																																	to.Ptr("PT15M"),
		// 																																																	to.Ptr("PT30M"),
		// 																																																	to.Ptr("PT1H"),
		// 																																																	to.Ptr("PT6H"),
		// 																																																	to.Ptr("PT12H"),
		// 																																																	to.Ptr("P1D")},
		// 																																																	Unit: to.Ptr("Bytes"),
		// 																																																},
		// 																																																{
		// 																																																	Name: to.Ptr("CbsVolumeOperationComplete"),
		// 																																																	AggregationType: to.Ptr("Average"),
		// 																																																	Dimensions: []*armnetapp.Dimension{
		// 																																																	},
		// 																																																	DisplayDescription: to.Ptr("Did the last volume backup or restore operation complete successfully? 1 if yes, 0 if no."),
		// 																																																	DisplayName: to.Ptr("Is Volume Backup Operation Complete"),
		// 																																																	EnableRegionalMdmAccount: to.Ptr(true),
		// 																																																	FillGapWithZero: to.Ptr(false),
		// 																																																	InternalMetricName: to.Ptr("CbsVolumeOperationComplete"),
		// 																																																	IsInternal: to.Ptr(false),
		// 																																																	SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																																	SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																																	SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																																		to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																																		SupportedTimeGrainTypes: []*string{
		// 																																																			to.Ptr("PT5M"),
		// 																																																			to.Ptr("PT15M"),
		// 																																																			to.Ptr("PT30M"),
		// 																																																			to.Ptr("PT1H"),
		// 																																																			to.Ptr("PT6H"),
		// 																																																			to.Ptr("PT12H"),
		// 																																																			to.Ptr("P1D")},
		// 																																																			Unit: to.Ptr("Count"),
		// 																																																		},
		// 																																																		{
		// 																																																			Name: to.Ptr("VolumeConsumedSizePercentage"),
		// 																																																			AggregationType: to.Ptr("Average"),
		// 																																																			Dimensions: []*armnetapp.Dimension{
		// 																																																			},
		// 																																																			DisplayDescription: to.Ptr("The percentage of the volume consumed including snapshots."),
		// 																																																			DisplayName: to.Ptr("Percentage Volume Consumed Size"),
		// 																																																			EnableRegionalMdmAccount: to.Ptr(true),
		// 																																																			FillGapWithZero: to.Ptr(false),
		// 																																																			InternalMetricName: to.Ptr("VolumeConsumedSizePercentage"),
		// 																																																			IsInternal: to.Ptr(false),
		// 																																																			SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																																			SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																																			SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																																				to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																																				SupportedTimeGrainTypes: []*string{
		// 																																																					to.Ptr("PT5M"),
		// 																																																					to.Ptr("PT15M"),
		// 																																																					to.Ptr("PT30M"),
		// 																																																					to.Ptr("PT1H"),
		// 																																																					to.Ptr("PT6H"),
		// 																																																					to.Ptr("PT12H"),
		// 																																																					to.Ptr("P1D")},
		// 																																																					Unit: to.Ptr("Percent"),
		// 																																																				},
		// 																																																				{
		// 																																																					Name: to.Ptr("OtherThroughput"),
		// 																																																					AggregationType: to.Ptr("Average"),
		// 																																																					Dimensions: []*armnetapp.Dimension{
		// 																																																					},
		// 																																																					DisplayDescription: to.Ptr("Other throughput (that is not read or write) in bytes per second"),
		// 																																																					DisplayName: to.Ptr("Other throughput"),
		// 																																																					EnableRegionalMdmAccount: to.Ptr(true),
		// 																																																					FillGapWithZero: to.Ptr(false),
		// 																																																					InternalMetricName: to.Ptr("OtherThroughput"),
		// 																																																					IsInternal: to.Ptr(false),
		// 																																																					SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																																					SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																																					SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																																						to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																																						SupportedTimeGrainTypes: []*string{
		// 																																																							to.Ptr("PT5M"),
		// 																																																							to.Ptr("PT15M"),
		// 																																																							to.Ptr("PT30M"),
		// 																																																							to.Ptr("PT1H"),
		// 																																																							to.Ptr("PT6H"),
		// 																																																							to.Ptr("PT12H"),
		// 																																																							to.Ptr("P1D")},
		// 																																																							Unit: to.Ptr("BytesPerSecond"),
		// 																																																						},
		// 																																																						{
		// 																																																							Name: to.Ptr("ReadThroughput"),
		// 																																																							AggregationType: to.Ptr("Average"),
		// 																																																							Dimensions: []*armnetapp.Dimension{
		// 																																																							},
		// 																																																							DisplayDescription: to.Ptr("Read throughput in bytes per second"),
		// 																																																							DisplayName: to.Ptr("Read throughput"),
		// 																																																							EnableRegionalMdmAccount: to.Ptr(true),
		// 																																																							FillGapWithZero: to.Ptr(false),
		// 																																																							InternalMetricName: to.Ptr("ReadThroughput"),
		// 																																																							IsInternal: to.Ptr(false),
		// 																																																							SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																																							SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																																							SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																																								to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																																								SupportedTimeGrainTypes: []*string{
		// 																																																									to.Ptr("PT5M"),
		// 																																																									to.Ptr("PT15M"),
		// 																																																									to.Ptr("PT30M"),
		// 																																																									to.Ptr("PT1H"),
		// 																																																									to.Ptr("PT6H"),
		// 																																																									to.Ptr("PT12H"),
		// 																																																									to.Ptr("P1D")},
		// 																																																									Unit: to.Ptr("BytesPerSecond"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("TotalThroughput"),
		// 																																																									AggregationType: to.Ptr("Average"),
		// 																																																									Dimensions: []*armnetapp.Dimension{
		// 																																																									},
		// 																																																									DisplayDescription: to.Ptr("Sum of all throughput in bytes per second"),
		// 																																																									DisplayName: to.Ptr("Total throughput"),
		// 																																																									EnableRegionalMdmAccount: to.Ptr(true),
		// 																																																									FillGapWithZero: to.Ptr(false),
		// 																																																									InternalMetricName: to.Ptr("TotalThroughput"),
		// 																																																									IsInternal: to.Ptr(false),
		// 																																																									SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																																									SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																																									SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																																										to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																																										SupportedTimeGrainTypes: []*string{
		// 																																																											to.Ptr("PT5M"),
		// 																																																											to.Ptr("PT15M"),
		// 																																																											to.Ptr("PT30M"),
		// 																																																											to.Ptr("PT1H"),
		// 																																																											to.Ptr("PT6H"),
		// 																																																											to.Ptr("PT12H"),
		// 																																																											to.Ptr("P1D")},
		// 																																																											Unit: to.Ptr("BytesPerSecond"),
		// 																																																										},
		// 																																																										{
		// 																																																											Name: to.Ptr("WriteThroughput"),
		// 																																																											AggregationType: to.Ptr("Average"),
		// 																																																											Dimensions: []*armnetapp.Dimension{
		// 																																																											},
		// 																																																											DisplayDescription: to.Ptr("Write throughput in bytes per second"),
		// 																																																											DisplayName: to.Ptr("Write throughput"),
		// 																																																											EnableRegionalMdmAccount: to.Ptr(true),
		// 																																																											FillGapWithZero: to.Ptr(false),
		// 																																																											InternalMetricName: to.Ptr("WriteThroughput"),
		// 																																																											IsInternal: to.Ptr(false),
		// 																																																											SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																																											SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																																											SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																																												to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																																												SupportedTimeGrainTypes: []*string{
		// 																																																													to.Ptr("PT5M"),
		// 																																																													to.Ptr("PT15M"),
		// 																																																													to.Ptr("PT30M"),
		// 																																																													to.Ptr("PT1H"),
		// 																																																													to.Ptr("PT6H"),
		// 																																																													to.Ptr("PT12H"),
		// 																																																													to.Ptr("P1D")},
		// 																																																													Unit: to.Ptr("BytesPerSecond"),
		// 																																																											}},
		// 																																																										},
		// 																																																									},
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Gets the diagnostic setting for the resource."),
		// 																																																										Operation: to.Ptr("Read diagnostic setting."),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Volumes resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Creates or updates the diagnostic setting for the resource."),
		// 																																																										Operation: to.Ptr("Write diagnostic setting."),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Volumes resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/AuthorizeReplication/action"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Authorize the source volume replication"),
		// 																																																										Operation: to.Ptr("Authorize Replication"),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Volumes resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/ResyncReplication/action"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Resync the replication on the destination volume"),
		// 																																																										Operation: to.Ptr("Resync Replication"),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Volumes resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/DeleteReplication/action"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Delete the replication on the destination volume"),
		// 																																																										Operation: to.Ptr("Delete Replication"),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Volumes resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/ReplicationStatus/read"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Reads the statuses of the Volume Replication."),
		// 																																																										Operation: to.Ptr("Read Volume Replication Status."),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Volumes resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/MountTargets/read"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Reads a mount target resource."),
		// 																																																										Operation: to.Ptr("Read mount target resource"),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Volumes resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/BackupStatus/read"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Volumes resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/RestoreStatus/read"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Volumes resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/PoolChange/action"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Moves volume to another pool."),
		// 																																																										Operation: to.Ptr("Change pool for volume"),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Volumes resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/RelocateVolume/action"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Relocate volume to a new stamp."),
		// 																																																										Operation: to.Ptr("Relocate volume to a new stamp."),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Volumes resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/FinalizeRelocation/action"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Finalize relocation by cleaning up the old volume."),
		// 																																																										Operation: to.Ptr("Finalize relocation of volume."),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Volumes resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/RevertRelocation/action"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Revert the relocation and revert back to the old volume."),
		// 																																																										Operation: to.Ptr("Revert the relocation of volume."),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Volumes resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/read"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Reads an account resource."),
		// 																																																										Operation: to.Ptr("Read account resource"),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Accounts resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/write"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Writes an account resource."),
		// 																																																										Operation: to.Ptr("Write account resource"),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Accounts resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/delete"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Deletes a account resource."),
		// 																																																										Operation: to.Ptr("Delete account resource"),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Accounts resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/RenewCredentials/action"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Renews MSI credentials of account, if account has MSI credentials that are due for renewal."),
		// 																																																										Operation: to.Ptr("Renew MSI credentials, if possible."),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Accounts resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/read"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Reads a pool resource."),
		// 																																																										Operation: to.Ptr("Read pool resource"),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Pools resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/write"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Writes a pool resource."),
		// 																																																										Operation: to.Ptr("Write pool resource"),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Pools resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/delete"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Deletes a pool resource."),
		// 																																																										Operation: to.Ptr("Delete pool resource"),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Pools resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("user,system"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/providers/Microsoft.Insights/metricDefinitions/read"),
		// 																																																									Display: &armnetapp.OperationDisplay{
		// 																																																										Description: to.Ptr("Gets the available metrics for Volume resource."),
		// 																																																										Operation: to.Ptr("Read volume metric definitions."),
		// 																																																										Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																										Resource: to.Ptr("Pools resource type"),
		// 																																																									},
		// 																																																									Origin: to.Ptr("system"),
		// 																																																									Properties: &armnetapp.OperationProperties{
		// 																																																										ServiceSpecification: &armnetapp.ServiceSpecification{
		// 																																																											MetricSpecifications: []*armnetapp.MetricSpecification{
		// 																																																												{
		// 																																																													Name: to.Ptr("VolumePoolAllocatedUsed"),
		// 																																																													AggregationType: to.Ptr("Average"),
		// 																																																													Dimensions: []*armnetapp.Dimension{
		// 																																																													},
		// 																																																													DisplayDescription: to.Ptr("Allocated used size of the pool"),
		// 																																																													DisplayName: to.Ptr("Pool Allocated To Volume Size"),
		// 																																																													EnableRegionalMdmAccount: to.Ptr(true),
		// 																																																													FillGapWithZero: to.Ptr(false),
		// 																																																													InternalMetricName: to.Ptr("VolumePoolAllocatedUsed"),
		// 																																																													IsInternal: to.Ptr(false),
		// 																																																													SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																																													SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																																													SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																																														to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																																														SupportedTimeGrainTypes: []*string{
		// 																																																															to.Ptr("PT5M"),
		// 																																																															to.Ptr("PT15M"),
		// 																																																															to.Ptr("PT30M"),
		// 																																																															to.Ptr("PT1H"),
		// 																																																															to.Ptr("PT6H"),
		// 																																																															to.Ptr("PT12H"),
		// 																																																															to.Ptr("P1D")},
		// 																																																															Unit: to.Ptr("Bytes"),
		// 																																																														},
		// 																																																														{
		// 																																																															Name: to.Ptr("VolumePoolTotalLogicalSize"),
		// 																																																															AggregationType: to.Ptr("Average"),
		// 																																																															Dimensions: []*armnetapp.Dimension{
		// 																																																															},
		// 																																																															DisplayDescription: to.Ptr("Sum of the logical size of all the volumes belonging to the pool"),
		// 																																																															DisplayName: to.Ptr("Pool Consumed Size"),
		// 																																																															EnableRegionalMdmAccount: to.Ptr(true),
		// 																																																															FillGapWithZero: to.Ptr(false),
		// 																																																															InternalMetricName: to.Ptr("VolumePoolTotalLogicalSize"),
		// 																																																															IsInternal: to.Ptr(false),
		// 																																																															SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																																															SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																																															SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																																																to.Ptr(armnetapp.MetricAggregationTypeAverage),
		// 																																																																to.Ptr(armnetapp.MetricAggregationType("Total"))},
		// 																																																																SupportedTimeGrainTypes: []*string{
		// 																																																																	to.Ptr("PT5M"),
		// 																																																																	to.Ptr("PT15M"),
		// 																																																																	to.Ptr("PT30M"),
		// 																																																																	to.Ptr("PT1H"),
		// 																																																																	to.Ptr("PT6H"),
		// 																																																																	to.Ptr("PT12H"),
		// 																																																																	to.Ptr("P1D")},
		// 																																																																	Unit: to.Ptr("Bytes"),
		// 																																																																},
		// 																																																																{
		// 																																																																	Name: to.Ptr("VolumePoolAllocatedSize"),
		// 																																																																	AggregationType: to.Ptr("Average"),
		// 																																																																	Dimensions: []*armnetapp.Dimension{
		// 																																																																	},
		// 																																																																	DisplayDescription: to.Ptr("Provisioned size of this pool"),
		// 																																																																	DisplayName: to.Ptr("Pool Allocated Size"),
		// 																																																																	EnableRegionalMdmAccount: to.Ptr(true),
		// 																																																																	FillGapWithZero: to.Ptr(false),
		// 																																																																	InternalMetricName: to.Ptr("VolumePoolAllocatedSize"),
		// 																																																																	IsInternal: to.Ptr(false),
		// 																																																																	SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																																																	SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																																																	SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																																																		to.Ptr(armnetapp.MetricAggregationTypeAverage),
		// 																																																																		to.Ptr(armnetapp.MetricAggregationType("Total"))},
		// 																																																																		SupportedTimeGrainTypes: []*string{
		// 																																																																			to.Ptr("PT5M"),
		// 																																																																			to.Ptr("PT15M"),
		// 																																																																			to.Ptr("PT30M"),
		// 																																																																			to.Ptr("PT1H"),
		// 																																																																			to.Ptr("PT6H"),
		// 																																																																			to.Ptr("PT12H"),
		// 																																																																			to.Ptr("P1D")},
		// 																																																																			Unit: to.Ptr("Bytes"),
		// 																																																																		},
		// 																																																																		{
		// 																																																																			Name: to.Ptr("VolumePoolTotalSnapshotSize"),
		// 																																																																			AggregationType: to.Ptr("Average"),
		// 																																																																			Dimensions: []*armnetapp.Dimension{
		// 																																																																			},
		// 																																																																			DisplayDescription: to.Ptr("Sum of snapshot size of all volumes in this pool"),
		// 																																																																			DisplayName: to.Ptr("Total Snapshot size for the pool"),
		// 																																																																			EnableRegionalMdmAccount: to.Ptr(true),
		// 																																																																			FillGapWithZero: to.Ptr(false),
		// 																																																																			InternalMetricName: to.Ptr("VolumePoolTotalSnapshotSize"),
		// 																																																																			IsInternal: to.Ptr(false),
		// 																																																																			SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																																																			SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																																																			SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																																																				to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																																																				SupportedTimeGrainTypes: []*string{
		// 																																																																					to.Ptr("PT5M"),
		// 																																																																					to.Ptr("PT15M"),
		// 																																																																					to.Ptr("PT30M"),
		// 																																																																					to.Ptr("PT1H"),
		// 																																																																					to.Ptr("PT6H"),
		// 																																																																					to.Ptr("PT12H"),
		// 																																																																					to.Ptr("P1D")},
		// 																																																																					Unit: to.Ptr("Bytes"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("VolumePoolProvisionedThroughput"),
		// 																																																																					AggregationType: to.Ptr("Average"),
		// 																																																																					Dimensions: []*armnetapp.Dimension{
		// 																																																																					},
		// 																																																																					DisplayDescription: to.Ptr("Provisioned throughput of this pool"),
		// 																																																																					DisplayName: to.Ptr("Provisioned throughput for the pool"),
		// 																																																																					EnableRegionalMdmAccount: to.Ptr(true),
		// 																																																																					FillGapWithZero: to.Ptr(false),
		// 																																																																					InternalMetricName: to.Ptr("VolumePoolProvisionedThroughput"),
		// 																																																																					IsInternal: to.Ptr(false),
		// 																																																																					SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																																																					SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																																																					SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																																																						to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																																																						SupportedTimeGrainTypes: []*string{
		// 																																																																							to.Ptr("PT5M"),
		// 																																																																							to.Ptr("PT15M"),
		// 																																																																							to.Ptr("PT30M"),
		// 																																																																							to.Ptr("PT1H"),
		// 																																																																							to.Ptr("PT6H"),
		// 																																																																							to.Ptr("PT12H"),
		// 																																																																							to.Ptr("P1D")},
		// 																																																																							Unit: to.Ptr("BytesPerSecond"),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("VolumePoolAllocatedToVolumeThroughput"),
		// 																																																																							AggregationType: to.Ptr("Average"),
		// 																																																																							Dimensions: []*armnetapp.Dimension{
		// 																																																																							},
		// 																																																																							DisplayDescription: to.Ptr("Sum of the throughput of all the volumes belonging to the pool"),
		// 																																																																							DisplayName: to.Ptr("Pool allocated throughput"),
		// 																																																																							EnableRegionalMdmAccount: to.Ptr(true),
		// 																																																																							FillGapWithZero: to.Ptr(false),
		// 																																																																							InternalMetricName: to.Ptr("VolumePoolAllocatedToVolumeThroughput"),
		// 																																																																							IsInternal: to.Ptr(false),
		// 																																																																							SourceMdmAccount: to.Ptr("MicrosoftNetAppShoebox2"),
		// 																																																																							SourceMdmNamespace: to.Ptr("NetAppUsageAndMetrics"),
		// 																																																																							SupportedAggregationTypes: []*armnetapp.MetricAggregationType{
		// 																																																																								to.Ptr(armnetapp.MetricAggregationTypeAverage)},
		// 																																																																								SupportedTimeGrainTypes: []*string{
		// 																																																																									to.Ptr("PT5M"),
		// 																																																																									to.Ptr("PT15M"),
		// 																																																																									to.Ptr("PT30M"),
		// 																																																																									to.Ptr("PT1H"),
		// 																																																																									to.Ptr("PT6H"),
		// 																																																																									to.Ptr("PT12H"),
		// 																																																																									to.Ptr("P1D")},
		// 																																																																									Unit: to.Ptr("BytesPerSecond"),
		// 																																																																							}},
		// 																																																																						},
		// 																																																																					},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/providers/Microsoft.Insights/logDefinitions/read"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Gets the log definitions for the resource."),
		// 																																																																						Operation: to.Ptr("Read log definitions."),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Pools resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("system"),
		// 																																																																					Properties: &armnetapp.OperationProperties{
		// 																																																																						ServiceSpecification: &armnetapp.ServiceSpecification{
		// 																																																																							LogSpecifications: []*armnetapp.LogSpecification{
		// 																																																																								{
		// 																																																																									Name: to.Ptr("Autoscale"),
		// 																																																																									DisplayName: to.Ptr("Capacity Pool Autoscaled"),
		// 																																																																							}},
		// 																																																																						},
		// 																																																																					},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Gets the diagnostic setting for the resource."),
		// 																																																																						Operation: to.Ptr("Read diagnostic setting."),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Pools resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Creates or updates the diagnostic setting for the resource."),
		// 																																																																						Operation: to.Ptr("Write diagnostic setting."),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Pools resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots/read"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Reads a snapshot resource."),
		// 																																																																						Operation: to.Ptr("Read snapshot resource"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Snapshots resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots/write"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Writes a snapshot resource."),
		// 																																																																						Operation: to.Ptr("Write snapshot resource"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Snapshots resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots/delete"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Deletes a snapshot resource."),
		// 																																																																						Operation: to.Ptr("Delete snapshot resource"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Snapshots resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots/RestoreFiles/action"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Restores files from a snapshot resource"),
		// 																																																																						Operation: to.Ptr("Single File Snapshot Restore"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Snapshots resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/subvolumes/read"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Subvolume resource type."),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/subvolumes/write"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Write a subvolume resource."),
		// 																																																																						Operation: to.Ptr("Write subvolume Resource"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Subvolume resource type."),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/subvolumes/delete"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Subvolume resource type."),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/subvolumes/GetMetadata/action"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Read subvolume metadata resource."),
		// 																																																																						Operation: to.Ptr("Subvolume Metadata resource."),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Subvolume resource type."),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/snapshotPolicies/read"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Reads a snapshot policy resource."),
		// 																																																																						Operation: to.Ptr("Read snapshot policy resource"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Snapshot Policies resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/snapshotPolicies/write"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Writes a snapshot policy resource."),
		// 																																																																						Operation: to.Ptr("Write snapshot policy resource"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Snapshot Policies resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/snapshotPolicies/delete"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Deletes a snapshot policy resource."),
		// 																																																																						Operation: to.Ptr("Delete snapshot policy resource"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Snapshot Policies resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/snapshotPolicies/Volumes/action"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("List volumes connected to snapshot policy"),
		// 																																																																						Operation: to.Ptr("List connected volumes"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Snapshot Policies resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/snapshotPolicies/ListVolumes/action"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("List volumes connected to snapshot policy"),
		// 																																																																						Operation: to.Ptr("List connected volumes"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Snapshot Policies resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/backups/read"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Reads a backup resource."),
		// 																																																																						Operation: to.Ptr("Read backup resource."),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Backup resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/backups/write"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Writes a backup resource."),
		// 																																																																						Operation: to.Ptr("Write backup resource."),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Backup resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/backups/delete"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Deletes a backup resource."),
		// 																																																																						Operation: to.Ptr("Delete backup resource"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Backup resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/backupPolicies/read"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Reads a backup policy resource."),
		// 																																																																						Operation: to.Ptr("Read backup policy resource."),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Backup Policy resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/backupPolicies/write"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Writes a backup policy resource."),
		// 																																																																						Operation: to.Ptr("Write backup policy resource."),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Backup Policy resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/backupPolicies/delete"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Deletes a backup policy resource."),
		// 																																																																						Operation: to.Ptr("Delete backup policy resource"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Backup Policy resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/vaults/read"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Reads a vault resource."),
		// 																																																																						Operation: to.Ptr("Read vault resource."),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Vault resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/accountBackups/read"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Reads an account backup resource."),
		// 																																																																						Operation: to.Ptr("Read an account backup resource."),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Account backup resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/accountBackups/write"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Writes an account backup resource."),
		// 																																																																						Operation: to.Ptr("Write an account backup resource."),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Account backup resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/accountBackups/delete"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Deletes an account backup resource."),
		// 																																																																						Operation: to.Ptr("Delete an account backup resource."),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Account backup resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/volumeGroups/read"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Reads a volume group resource."),
		// 																																																																						Operation: to.Ptr("Read volume group resource"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Volume Group resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/volumeGroups/write"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Writes a volume group resource."),
		// 																																																																						Operation: to.Ptr("Write volume group resource"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Volume Group resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/netAppAccounts/volumeGroups/delete"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Deletes a volume group resource."),
		// 																																																																						Operation: to.Ptr("Delete volume group resource"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Volume Group resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/locations/quotaLimits/read"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Reads a Quotalimit resource type."),
		// 																																																																						Operation: to.Ptr("Read QuotaLimit resource type"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("QuotaLimit resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/Operations/read"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Reads an operation resources."),
		// 																																																																						Operation: to.Ptr("Read operation resource"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Operations resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/locations/operationresults/read"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Reads an operation result resource."),
		// 																																																																						Operation: to.Ptr("Read operation result resource"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Operation results resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/locations/read"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Reads a location wide operation."),
		// 																																																																						Operation: to.Ptr("Read location wide operation"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Location wide operation"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/locations/checknameavailability/action"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Check if resource name is available"),
		// 																																																																						Operation: to.Ptr("Check if resource name is available"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("NetApp resources"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/locations/checkfilepathavailability/action"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Check if file path is available"),
		// 																																																																						Operation: to.Ptr("Check if file path is available"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Volumes resource type"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/unregister/action"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Unregisters Subscription with Microsoft.NetApp resource provider"),
		// 																																																																						Operation: to.Ptr("Unregister Subscription for Azure NetApp Files"),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("Subscription"),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Microsoft.NetApp/locations/checkinventory/action"),
		// 																																																																					Display: &armnetapp.OperationDisplay{
		// 																																																																						Description: to.Ptr("Checks ReservedCapacity inventory."),
		// 																																																																						Operation: to.Ptr("Checks ReservedCapacity inventory."),
		// 																																																																						Provider: to.Ptr("Microsoft.NetApp Resource Provider"),
		// 																																																																						Resource: to.Ptr("ReservedCapacity reservation resource."),
		// 																																																																					},
		// 																																																																					Origin: to.Ptr("user,system"),
		// 																																																																			}},
		// 																																																																		}
	}
}
