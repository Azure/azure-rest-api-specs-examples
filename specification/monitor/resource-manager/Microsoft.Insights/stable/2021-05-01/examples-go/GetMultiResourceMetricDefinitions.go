package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2021-05-01/examples/GetMultiResourceMetricDefinitions.json
func ExampleMetricDefinitionsClient_NewListAtSubscriptionScopePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMetricDefinitionsClient().NewListAtSubscriptionScopePager("westus2", &armmonitor.MetricDefinitionsClientListAtSubscriptionScopeOptions{Metricnamespace: to.Ptr("microsoft.compute/virtualmachines")})
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
		// page.SubscriptionScopeMetricDefinitionCollection = armmonitor.SubscriptionScopeMetricDefinitionCollection{
		// 	Value: []*armmonitor.SubscriptionScopeMetricDefinition{
		// 		{
		// 			Name: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("Percentage CPU"),
		// 				Value: to.Ptr("Percentage CPU"),
		// 			},
		// 			Dimensions: []*armmonitor.LocalizableString{
		// 				{
		// 					LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 					Value: to.Ptr("Microsoft.ResourceId"),
		// 				},
		// 				{
		// 					LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 					Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 			}},
		// 			DisplayDescription: to.Ptr("The percentage of allocated compute units that are currently in use by the Virtual Machine(s)"),
		// 			ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Percentage CPU"),
		// 			IsDimensionRequired: to.Ptr(false),
		// 			MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("PT15M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("PT30M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("PT6H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("PT12H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P93D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 			PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 			SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 				to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 				to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 				to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 				to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 				to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 				to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 				Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 			},
		// 			{
		// 				Name: &armmonitor.LocalizableString{
		// 					LocalizedValue: to.Ptr("Network In Billable (Deprecated)"),
		// 					Value: to.Ptr("Network In"),
		// 				},
		// 				Dimensions: []*armmonitor.LocalizableString{
		// 					{
		// 						LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 						Value: to.Ptr("Microsoft.ResourceId"),
		// 					},
		// 					{
		// 						LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 						Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 				}},
		// 				DisplayDescription: to.Ptr("The number of billable bytes received on all network interfaces by the Virtual Machine(s) (Incoming Traffic) (Deprecated)"),
		// 				ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Network In"),
		// 				IsDimensionRequired: to.Ptr(false),
		// 				MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("PT1M"),
		// 					},
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("PT5M"),
		// 					},
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("PT15M"),
		// 					},
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("PT30M"),
		// 					},
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("PT1H"),
		// 					},
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("PT6H"),
		// 					},
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("PT12H"),
		// 					},
		// 					{
		// 						Retention: to.Ptr("P93D"),
		// 						TimeGrain: to.Ptr("P1D"),
		// 				}},
		// 				Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 				PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 				ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 				SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 					to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 					to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 					to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 					to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 					to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 					to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 					Unit: to.Ptr(armmonitor.MetricUnitBytes),
		// 				},
		// 				{
		// 					Name: &armmonitor.LocalizableString{
		// 						LocalizedValue: to.Ptr("Network Out Billable (Deprecated)"),
		// 						Value: to.Ptr("Network Out"),
		// 					},
		// 					Dimensions: []*armmonitor.LocalizableString{
		// 						{
		// 							LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 							Value: to.Ptr("Microsoft.ResourceId"),
		// 						},
		// 						{
		// 							LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 							Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 					}},
		// 					DisplayDescription: to.Ptr("The number of billable bytes out on all network interfaces by the Virtual Machine(s) (Outgoing Traffic) (Deprecated)"),
		// 					ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Network Out"),
		// 					IsDimensionRequired: to.Ptr(false),
		// 					MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("PT1M"),
		// 						},
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("PT5M"),
		// 						},
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("PT15M"),
		// 						},
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("PT30M"),
		// 						},
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("PT1H"),
		// 						},
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("PT6H"),
		// 						},
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("PT12H"),
		// 						},
		// 						{
		// 							Retention: to.Ptr("P93D"),
		// 							TimeGrain: to.Ptr("P1D"),
		// 					}},
		// 					Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 					PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 					ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 					SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 						to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 						to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 						to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 						to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 						to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 						to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 						Unit: to.Ptr(armmonitor.MetricUnitBytes),
		// 					},
		// 					{
		// 						Name: &armmonitor.LocalizableString{
		// 							LocalizedValue: to.Ptr("Disk Read Bytes"),
		// 							Value: to.Ptr("Disk Read Bytes"),
		// 						},
		// 						Dimensions: []*armmonitor.LocalizableString{
		// 							{
		// 								LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 								Value: to.Ptr("Microsoft.ResourceId"),
		// 							},
		// 							{
		// 								LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 								Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 						}},
		// 						DisplayDescription: to.Ptr("Bytes read from disk during monitoring period"),
		// 						ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Disk Read Bytes"),
		// 						IsDimensionRequired: to.Ptr(false),
		// 						MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("PT1M"),
		// 							},
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("PT5M"),
		// 							},
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("PT15M"),
		// 							},
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("PT30M"),
		// 							},
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("PT1H"),
		// 							},
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("PT6H"),
		// 							},
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("PT12H"),
		// 							},
		// 							{
		// 								Retention: to.Ptr("P93D"),
		// 								TimeGrain: to.Ptr("P1D"),
		// 						}},
		// 						Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 						PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 						ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 						SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 							to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 							to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 							to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 							to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 							to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 							to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 							Unit: to.Ptr(armmonitor.MetricUnitBytes),
		// 						},
		// 						{
		// 							Name: &armmonitor.LocalizableString{
		// 								LocalizedValue: to.Ptr("Disk Write Bytes"),
		// 								Value: to.Ptr("Disk Write Bytes"),
		// 							},
		// 							Dimensions: []*armmonitor.LocalizableString{
		// 								{
		// 									LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 									Value: to.Ptr("Microsoft.ResourceId"),
		// 								},
		// 								{
		// 									LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 									Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Bytes written to disk during monitoring period"),
		// 							ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Disk Write Bytes"),
		// 							IsDimensionRequired: to.Ptr(false),
		// 							MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 								},
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("PT5M"),
		// 								},
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("PT15M"),
		// 								},
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("PT30M"),
		// 								},
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("PT1H"),
		// 								},
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("PT6H"),
		// 								},
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("PT12H"),
		// 								},
		// 								{
		// 									Retention: to.Ptr("P93D"),
		// 									TimeGrain: to.Ptr("P1D"),
		// 							}},
		// 							Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 							PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 							ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 							SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 								to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 								to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 								to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 								to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 								to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 								to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 								Unit: to.Ptr(armmonitor.MetricUnitBytes),
		// 							},
		// 							{
		// 								Name: &armmonitor.LocalizableString{
		// 									LocalizedValue: to.Ptr("Disk Read Operations/Sec"),
		// 									Value: to.Ptr("Disk Read Operations/Sec"),
		// 								},
		// 								Dimensions: []*armmonitor.LocalizableString{
		// 									{
		// 										LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 										Value: to.Ptr("Microsoft.ResourceId"),
		// 									},
		// 									{
		// 										LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 										Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 								}},
		// 								DisplayDescription: to.Ptr("Disk Read IOPS"),
		// 								ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Disk Read Operations/Sec"),
		// 								IsDimensionRequired: to.Ptr(false),
		// 								MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("PT1M"),
		// 									},
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("PT5M"),
		// 									},
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("PT15M"),
		// 									},
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("PT30M"),
		// 									},
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("PT1H"),
		// 									},
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("PT6H"),
		// 									},
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("PT12H"),
		// 									},
		// 									{
		// 										Retention: to.Ptr("P93D"),
		// 										TimeGrain: to.Ptr("P1D"),
		// 								}},
		// 								Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 								PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 								ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 								SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 									to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 									to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 									to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 									to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 									to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 									to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 									Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 								},
		// 								{
		// 									Name: &armmonitor.LocalizableString{
		// 										LocalizedValue: to.Ptr("Disk Write Operations/Sec"),
		// 										Value: to.Ptr("Disk Write Operations/Sec"),
		// 									},
		// 									Dimensions: []*armmonitor.LocalizableString{
		// 										{
		// 											LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 											Value: to.Ptr("Microsoft.ResourceId"),
		// 										},
		// 										{
		// 											LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 											Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 									}},
		// 									DisplayDescription: to.Ptr("Disk Write IOPS"),
		// 									ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Disk Write Operations/Sec"),
		// 									IsDimensionRequired: to.Ptr(false),
		// 									MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("PT1M"),
		// 										},
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("PT5M"),
		// 										},
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("PT15M"),
		// 										},
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("PT30M"),
		// 										},
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("PT1H"),
		// 										},
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("PT6H"),
		// 										},
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("PT12H"),
		// 										},
		// 										{
		// 											Retention: to.Ptr("P93D"),
		// 											TimeGrain: to.Ptr("P1D"),
		// 									}},
		// 									Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 									PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 									ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 									SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 										to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 										to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 										to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 										to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 										to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 										to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 										Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 									},
		// 									{
		// 										Name: &armmonitor.LocalizableString{
		// 											LocalizedValue: to.Ptr("CPU Credits Remaining"),
		// 											Value: to.Ptr("CPU Credits Remaining"),
		// 										},
		// 										Dimensions: []*armmonitor.LocalizableString{
		// 											{
		// 												LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 												Value: to.Ptr("Microsoft.ResourceId"),
		// 											},
		// 											{
		// 												LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 												Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 										}},
		// 										DisplayDescription: to.Ptr("Total number of credits available to burst. Only available on B-series burstable VMs"),
		// 										ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/CPU Credits Remaining"),
		// 										IsDimensionRequired: to.Ptr(false),
		// 										MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("PT1M"),
		// 											},
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("PT5M"),
		// 											},
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("PT15M"),
		// 											},
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("PT30M"),
		// 											},
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("PT1H"),
		// 											},
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("PT6H"),
		// 											},
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("PT12H"),
		// 											},
		// 											{
		// 												Retention: to.Ptr("P93D"),
		// 												TimeGrain: to.Ptr("P1D"),
		// 										}},
		// 										Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 										PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 										ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 										SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 											to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 											to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 											to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 											to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 											to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 											to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 											Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 										},
		// 										{
		// 											Name: &armmonitor.LocalizableString{
		// 												LocalizedValue: to.Ptr("CPU Credits Consumed"),
		// 												Value: to.Ptr("CPU Credits Consumed"),
		// 											},
		// 											Dimensions: []*armmonitor.LocalizableString{
		// 												{
		// 													LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 													Value: to.Ptr("Microsoft.ResourceId"),
		// 												},
		// 												{
		// 													LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 													Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 											}},
		// 											DisplayDescription: to.Ptr("Total number of credits consumed by the Virtual Machine. Only available on B-series burstable VMs"),
		// 											ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/CPU Credits Consumed"),
		// 											IsDimensionRequired: to.Ptr(false),
		// 											MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("PT1M"),
		// 												},
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("PT5M"),
		// 												},
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("PT15M"),
		// 												},
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("PT30M"),
		// 												},
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("PT1H"),
		// 												},
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("PT6H"),
		// 												},
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("PT12H"),
		// 												},
		// 												{
		// 													Retention: to.Ptr("P93D"),
		// 													TimeGrain: to.Ptr("P1D"),
		// 											}},
		// 											Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 											PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 											ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 											SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 												to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 												to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 												to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 												to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 												to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 												to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 												Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 											},
		// 											{
		// 												Name: &armmonitor.LocalizableString{
		// 													LocalizedValue: to.Ptr("Data Disk Read Bytes/Sec"),
		// 													Value: to.Ptr("Data Disk Read Bytes/sec"),
		// 												},
		// 												Dimensions: []*armmonitor.LocalizableString{
		// 													{
		// 														LocalizedValue: to.Ptr("LUN"),
		// 														Value: to.Ptr("LUN"),
		// 													},
		// 													{
		// 														LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 														Value: to.Ptr("Microsoft.ResourceId"),
		// 													},
		// 													{
		// 														LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 														Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 												}},
		// 												DisplayDescription: to.Ptr("Bytes/Sec read from a single disk during monitoring period"),
		// 												ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Data Disk Read Bytes/sec"),
		// 												IsDimensionRequired: to.Ptr(false),
		// 												MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("PT1M"),
		// 													},
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("PT5M"),
		// 													},
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("PT15M"),
		// 													},
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("PT30M"),
		// 													},
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("PT1H"),
		// 													},
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("PT6H"),
		// 													},
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("PT12H"),
		// 													},
		// 													{
		// 														Retention: to.Ptr("P93D"),
		// 														TimeGrain: to.Ptr("P1D"),
		// 												}},
		// 												Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 												PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 												ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 												SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 													to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 													to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 													to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 													to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 													to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 													to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 													Unit: to.Ptr(armmonitor.MetricUnitBytesPerSecond),
		// 												},
		// 												{
		// 													Name: &armmonitor.LocalizableString{
		// 														LocalizedValue: to.Ptr("Data Disk Write Bytes/Sec"),
		// 														Value: to.Ptr("Data Disk Write Bytes/sec"),
		// 													},
		// 													Dimensions: []*armmonitor.LocalizableString{
		// 														{
		// 															LocalizedValue: to.Ptr("LUN"),
		// 															Value: to.Ptr("LUN"),
		// 														},
		// 														{
		// 															LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 															Value: to.Ptr("Microsoft.ResourceId"),
		// 														},
		// 														{
		// 															LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 															Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 													}},
		// 													DisplayDescription: to.Ptr("Bytes/Sec written to a single disk during monitoring period"),
		// 													ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Data Disk Write Bytes/sec"),
		// 													IsDimensionRequired: to.Ptr(false),
		// 													MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("PT1M"),
		// 														},
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("PT5M"),
		// 														},
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("PT15M"),
		// 														},
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("PT30M"),
		// 														},
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("PT1H"),
		// 														},
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("PT6H"),
		// 														},
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("PT12H"),
		// 														},
		// 														{
		// 															Retention: to.Ptr("P93D"),
		// 															TimeGrain: to.Ptr("P1D"),
		// 													}},
		// 													Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 													PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 													ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 													SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 														to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 														to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 														to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 														to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 														to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 														to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 														Unit: to.Ptr(armmonitor.MetricUnitBytesPerSecond),
		// 													},
		// 													{
		// 														Name: &armmonitor.LocalizableString{
		// 															LocalizedValue: to.Ptr("Data Disk Read Operations/Sec"),
		// 															Value: to.Ptr("Data Disk Read Operations/Sec"),
		// 														},
		// 														Dimensions: []*armmonitor.LocalizableString{
		// 															{
		// 																LocalizedValue: to.Ptr("LUN"),
		// 																Value: to.Ptr("LUN"),
		// 															},
		// 															{
		// 																LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																Value: to.Ptr("Microsoft.ResourceId"),
		// 															},
		// 															{
		// 																LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 														}},
		// 														DisplayDescription: to.Ptr("Read IOPS from a single disk during monitoring period"),
		// 														ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Data Disk Read Operations/Sec"),
		// 														IsDimensionRequired: to.Ptr(false),
		// 														MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("PT1M"),
		// 															},
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("PT5M"),
		// 															},
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("PT15M"),
		// 															},
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("PT30M"),
		// 															},
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("PT1H"),
		// 															},
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("PT6H"),
		// 															},
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("PT12H"),
		// 															},
		// 															{
		// 																Retention: to.Ptr("P93D"),
		// 																TimeGrain: to.Ptr("P1D"),
		// 														}},
		// 														Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 														PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 														ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 														SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 															to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 															to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 															to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 															to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 															to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 															to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 															Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 														},
		// 														{
		// 															Name: &armmonitor.LocalizableString{
		// 																LocalizedValue: to.Ptr("Data Disk Write Operations/Sec"),
		// 																Value: to.Ptr("Data Disk Write Operations/Sec"),
		// 															},
		// 															Dimensions: []*armmonitor.LocalizableString{
		// 																{
		// 																	LocalizedValue: to.Ptr("LUN"),
		// 																	Value: to.Ptr("LUN"),
		// 																},
		// 																{
		// 																	LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																	Value: to.Ptr("Microsoft.ResourceId"),
		// 																},
		// 																{
		// 																	LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																	Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 															}},
		// 															DisplayDescription: to.Ptr("Write IOPS from a single disk during monitoring period"),
		// 															ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Data Disk Write Operations/Sec"),
		// 															IsDimensionRequired: to.Ptr(false),
		// 															MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("PT1M"),
		// 																},
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("PT5M"),
		// 																},
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("PT15M"),
		// 																},
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("PT30M"),
		// 																},
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("PT1H"),
		// 																},
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("PT6H"),
		// 																},
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("PT12H"),
		// 																},
		// 																{
		// 																	Retention: to.Ptr("P93D"),
		// 																	TimeGrain: to.Ptr("P1D"),
		// 															}},
		// 															Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 															PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 															ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 															SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 															},
		// 															{
		// 																Name: &armmonitor.LocalizableString{
		// 																	LocalizedValue: to.Ptr("Data Disk Queue Depth"),
		// 																	Value: to.Ptr("Data Disk Queue Depth"),
		// 																},
		// 																Dimensions: []*armmonitor.LocalizableString{
		// 																	{
		// 																		LocalizedValue: to.Ptr("LUN"),
		// 																		Value: to.Ptr("LUN"),
		// 																	},
		// 																	{
		// 																		LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																		Value: to.Ptr("Microsoft.ResourceId"),
		// 																	},
		// 																	{
		// 																		LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																		Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																}},
		// 																DisplayDescription: to.Ptr("Data Disk Queue Depth(or Queue Length)"),
		// 																ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Data Disk Queue Depth"),
		// 																IsDimensionRequired: to.Ptr(false),
		// 																MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("PT1M"),
		// 																	},
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("PT5M"),
		// 																	},
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("PT15M"),
		// 																	},
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("PT30M"),
		// 																	},
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("PT1H"),
		// 																	},
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("PT6H"),
		// 																	},
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("PT12H"),
		// 																	},
		// 																	{
		// 																		Retention: to.Ptr("P93D"),
		// 																		TimeGrain: to.Ptr("P1D"),
		// 																}},
		// 																Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																	to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																	to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																	to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																	to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																	to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																	to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																	Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																},
		// 																{
		// 																	Name: &armmonitor.LocalizableString{
		// 																		LocalizedValue: to.Ptr("Data Disk Bandwidth Consumed Percentage"),
		// 																		Value: to.Ptr("Data Disk Bandwidth Consumed Percentage"),
		// 																	},
		// 																	Dimensions: []*armmonitor.LocalizableString{
		// 																		{
		// 																			LocalizedValue: to.Ptr("LUN"),
		// 																			Value: to.Ptr("LUN"),
		// 																		},
		// 																		{
		// 																			LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																			Value: to.Ptr("Microsoft.ResourceId"),
		// 																		},
		// 																		{
		// 																			LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																			Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																	}},
		// 																	DisplayDescription: to.Ptr("Percentage of data disk bandwidth consumed per minute"),
		// 																	ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Data Disk Bandwidth Consumed Percentage"),
		// 																	IsDimensionRequired: to.Ptr(false),
		// 																	MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("PT1M"),
		// 																		},
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("PT5M"),
		// 																		},
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("PT15M"),
		// 																		},
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("PT30M"),
		// 																		},
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("PT1H"),
		// 																		},
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("PT6H"),
		// 																		},
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("PT12H"),
		// 																		},
		// 																		{
		// 																			Retention: to.Ptr("P93D"),
		// 																			TimeGrain: to.Ptr("P1D"),
		// 																	}},
		// 																	Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																	PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																	ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																	SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																		to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																		to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																		to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																		to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																		to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																		to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																		Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																	},
		// 																	{
		// 																		Name: &armmonitor.LocalizableString{
		// 																			LocalizedValue: to.Ptr("Data Disk IOPS Consumed Percentage"),
		// 																			Value: to.Ptr("Data Disk IOPS Consumed Percentage"),
		// 																		},
		// 																		Dimensions: []*armmonitor.LocalizableString{
		// 																			{
		// 																				LocalizedValue: to.Ptr("LUN"),
		// 																				Value: to.Ptr("LUN"),
		// 																			},
		// 																			{
		// 																				LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																				Value: to.Ptr("Microsoft.ResourceId"),
		// 																			},
		// 																			{
		// 																				LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																				Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																		}},
		// 																		DisplayDescription: to.Ptr("Percentage of data disk I/Os consumed per minute"),
		// 																		ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Data Disk IOPS Consumed Percentage"),
		// 																		IsDimensionRequired: to.Ptr(false),
		// 																		MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("PT1M"),
		// 																			},
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("PT5M"),
		// 																			},
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("PT15M"),
		// 																			},
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("PT30M"),
		// 																			},
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("PT1H"),
		// 																			},
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("PT6H"),
		// 																			},
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("PT12H"),
		// 																			},
		// 																			{
		// 																				Retention: to.Ptr("P93D"),
		// 																				TimeGrain: to.Ptr("P1D"),
		// 																		}},
		// 																		Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																		PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																		ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																		SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																			to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																			to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																			to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																			to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																			to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																			to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																			Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																		},
		// 																		{
		// 																			Name: &armmonitor.LocalizableString{
		// 																				LocalizedValue: to.Ptr("Data Disk Target Bandwidth"),
		// 																				Value: to.Ptr("Data Disk Target Bandwidth"),
		// 																			},
		// 																			Dimensions: []*armmonitor.LocalizableString{
		// 																				{
		// 																					LocalizedValue: to.Ptr("LUN"),
		// 																					Value: to.Ptr("LUN"),
		// 																				},
		// 																				{
		// 																					LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																					Value: to.Ptr("Microsoft.ResourceId"),
		// 																				},
		// 																				{
		// 																					LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																					Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																			}},
		// 																			DisplayDescription: to.Ptr("Baseline bytes per second throughput Data Disk can achieve without bursting"),
		// 																			ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Data Disk Target Bandwidth"),
		// 																			IsDimensionRequired: to.Ptr(false),
		// 																			MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("PT1M"),
		// 																				},
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("PT5M"),
		// 																				},
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("PT15M"),
		// 																				},
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("PT30M"),
		// 																				},
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("PT1H"),
		// 																				},
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("PT6H"),
		// 																				},
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("PT12H"),
		// 																				},
		// 																				{
		// 																					Retention: to.Ptr("P93D"),
		// 																					TimeGrain: to.Ptr("P1D"),
		// 																			}},
		// 																			Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																			PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																			ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																			SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																				to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																				to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																				to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																				to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																				to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																				to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																				Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																			},
		// 																			{
		// 																				Name: &armmonitor.LocalizableString{
		// 																					LocalizedValue: to.Ptr("Data Disk Target IOPS"),
		// 																					Value: to.Ptr("Data Disk Target IOPS"),
		// 																				},
		// 																				Dimensions: []*armmonitor.LocalizableString{
		// 																					{
		// 																						LocalizedValue: to.Ptr("LUN"),
		// 																						Value: to.Ptr("LUN"),
		// 																					},
		// 																					{
		// 																						LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																						Value: to.Ptr("Microsoft.ResourceId"),
		// 																					},
		// 																					{
		// 																						LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																						Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																				}},
		// 																				DisplayDescription: to.Ptr("Baseline IOPS Data Disk can achieve without bursting"),
		// 																				ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Data Disk Target IOPS"),
		// 																				IsDimensionRequired: to.Ptr(false),
		// 																				MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("PT1M"),
		// 																					},
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("PT5M"),
		// 																					},
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("PT15M"),
		// 																					},
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("PT30M"),
		// 																					},
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("PT1H"),
		// 																					},
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("PT6H"),
		// 																					},
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("PT12H"),
		// 																					},
		// 																					{
		// 																						Retention: to.Ptr("P93D"),
		// 																						TimeGrain: to.Ptr("P1D"),
		// 																				}},
		// 																				Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																				PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																				ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																				SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																					to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																					to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																					to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																					to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																					to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																					to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																					Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																				},
		// 																				{
		// 																					Name: &armmonitor.LocalizableString{
		// 																						LocalizedValue: to.Ptr("Data Disk Max Burst Bandwidth"),
		// 																						Value: to.Ptr("Data Disk Max Burst Bandwidth"),
		// 																					},
		// 																					Dimensions: []*armmonitor.LocalizableString{
		// 																						{
		// 																							LocalizedValue: to.Ptr("LUN"),
		// 																							Value: to.Ptr("LUN"),
		// 																						},
		// 																						{
		// 																							LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																							Value: to.Ptr("Microsoft.ResourceId"),
		// 																						},
		// 																						{
		// 																							LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																							Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																					}},
		// 																					DisplayDescription: to.Ptr("Maximum bytes per second throughput Data Disk can achieve with bursting"),
		// 																					ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Data Disk Max Burst Bandwidth"),
		// 																					IsDimensionRequired: to.Ptr(false),
		// 																					MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("PT1M"),
		// 																						},
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("PT5M"),
		// 																						},
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("PT15M"),
		// 																						},
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("PT30M"),
		// 																						},
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("PT1H"),
		// 																						},
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("PT6H"),
		// 																						},
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("PT12H"),
		// 																						},
		// 																						{
		// 																							Retention: to.Ptr("P93D"),
		// 																							TimeGrain: to.Ptr("P1D"),
		// 																					}},
		// 																					Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																					PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																					ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																					SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																						to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																						to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																						to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																						to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																						to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																						to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																						Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																					},
		// 																					{
		// 																						Name: &armmonitor.LocalizableString{
		// 																							LocalizedValue: to.Ptr("Data Disk Max Burst IOPS"),
		// 																							Value: to.Ptr("Data Disk Max Burst IOPS"),
		// 																						},
		// 																						Dimensions: []*armmonitor.LocalizableString{
		// 																							{
		// 																								LocalizedValue: to.Ptr("LUN"),
		// 																								Value: to.Ptr("LUN"),
		// 																							},
		// 																							{
		// 																								LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																								Value: to.Ptr("Microsoft.ResourceId"),
		// 																							},
		// 																							{
		// 																								LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																								Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																						}},
		// 																						DisplayDescription: to.Ptr("Maximum IOPS Data Disk can achieve with bursting"),
		// 																						ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Data Disk Max Burst IOPS"),
		// 																						IsDimensionRequired: to.Ptr(false),
		// 																						MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("PT1M"),
		// 																							},
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("PT5M"),
		// 																							},
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("PT15M"),
		// 																							},
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("PT30M"),
		// 																							},
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("PT1H"),
		// 																							},
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("PT6H"),
		// 																							},
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("PT12H"),
		// 																							},
		// 																							{
		// 																								Retention: to.Ptr("P93D"),
		// 																								TimeGrain: to.Ptr("P1D"),
		// 																						}},
		// 																						Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																						PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																						ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																						SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																							to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																							to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																							to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																							to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																							to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																							to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																							Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																						},
		// 																						{
		// 																							Name: &armmonitor.LocalizableString{
		// 																								LocalizedValue: to.Ptr("Data Disk Used Burst BPS Credits Percentage"),
		// 																								Value: to.Ptr("Data Disk Used Burst BPS Credits Percentage"),
		// 																							},
		// 																							Dimensions: []*armmonitor.LocalizableString{
		// 																								{
		// 																									LocalizedValue: to.Ptr("LUN"),
		// 																									Value: to.Ptr("LUN"),
		// 																								},
		// 																								{
		// 																									LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																									Value: to.Ptr("Microsoft.ResourceId"),
		// 																								},
		// 																								{
		// 																									LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																									Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																							}},
		// 																							DisplayDescription: to.Ptr("Percentage of Data Disk burst bandwidth credits used so far"),
		// 																							ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Data Disk Used Burst BPS Credits Percentage"),
		// 																							IsDimensionRequired: to.Ptr(false),
		// 																							MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("PT1M"),
		// 																								},
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("PT5M"),
		// 																								},
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("PT15M"),
		// 																								},
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("PT30M"),
		// 																								},
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("PT1H"),
		// 																								},
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("PT6H"),
		// 																								},
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("PT12H"),
		// 																								},
		// 																								{
		// 																									Retention: to.Ptr("P93D"),
		// 																									TimeGrain: to.Ptr("P1D"),
		// 																							}},
		// 																							Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																							PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																							ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																							SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																								to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																								to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																								to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																								to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																								to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																								to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																								Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																							},
		// 																							{
		// 																								Name: &armmonitor.LocalizableString{
		// 																									LocalizedValue: to.Ptr("Data Disk Used Burst IO Credits Percentage"),
		// 																									Value: to.Ptr("Data Disk Used Burst IO Credits Percentage"),
		// 																								},
		// 																								Dimensions: []*armmonitor.LocalizableString{
		// 																									{
		// 																										LocalizedValue: to.Ptr("LUN"),
		// 																										Value: to.Ptr("LUN"),
		// 																									},
		// 																									{
		// 																										LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																										Value: to.Ptr("Microsoft.ResourceId"),
		// 																									},
		// 																									{
		// 																										LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																										Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																								}},
		// 																								DisplayDescription: to.Ptr("Percentage of Data Disk burst I/O credits used so far"),
		// 																								ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Data Disk Used Burst IO Credits Percentage"),
		// 																								IsDimensionRequired: to.Ptr(false),
		// 																								MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("PT1M"),
		// 																									},
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("PT5M"),
		// 																									},
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("PT15M"),
		// 																									},
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("PT30M"),
		// 																									},
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("PT1H"),
		// 																									},
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("PT6H"),
		// 																									},
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("PT12H"),
		// 																									},
		// 																									{
		// 																										Retention: to.Ptr("P93D"),
		// 																										TimeGrain: to.Ptr("P1D"),
		// 																								}},
		// 																								Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																								PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																								ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																								SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																									to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																									to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																									to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																									to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																									to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																									to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																									Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																								},
		// 																								{
		// 																									Name: &armmonitor.LocalizableString{
		// 																										LocalizedValue: to.Ptr("OS Disk Read Bytes/Sec"),
		// 																										Value: to.Ptr("OS Disk Read Bytes/sec"),
		// 																									},
		// 																									Dimensions: []*armmonitor.LocalizableString{
		// 																										{
		// 																											LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																											Value: to.Ptr("Microsoft.ResourceId"),
		// 																										},
		// 																										{
		// 																											LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																											Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																									}},
		// 																									DisplayDescription: to.Ptr("Bytes/Sec read from a single disk during monitoring period for OS disk"),
		// 																									ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/OS Disk Read Bytes/sec"),
		// 																									IsDimensionRequired: to.Ptr(false),
		// 																									MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("PT1M"),
		// 																										},
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("PT5M"),
		// 																										},
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("PT15M"),
		// 																										},
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("PT30M"),
		// 																										},
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("PT1H"),
		// 																										},
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("PT6H"),
		// 																										},
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("PT12H"),
		// 																										},
		// 																										{
		// 																											Retention: to.Ptr("P93D"),
		// 																											TimeGrain: to.Ptr("P1D"),
		// 																									}},
		// 																									Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																									PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																									ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																									SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																										to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																										to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																										to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																										to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																										to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																										to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																										Unit: to.Ptr(armmonitor.MetricUnitBytesPerSecond),
		// 																									},
		// 																									{
		// 																										Name: &armmonitor.LocalizableString{
		// 																											LocalizedValue: to.Ptr("OS Disk Write Bytes/Sec"),
		// 																											Value: to.Ptr("OS Disk Write Bytes/sec"),
		// 																										},
		// 																										Dimensions: []*armmonitor.LocalizableString{
		// 																											{
		// 																												LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																												Value: to.Ptr("Microsoft.ResourceId"),
		// 																											},
		// 																											{
		// 																												LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																												Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																										}},
		// 																										DisplayDescription: to.Ptr("Bytes/Sec written to a single disk during monitoring period for OS disk"),
		// 																										ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/OS Disk Write Bytes/sec"),
		// 																										IsDimensionRequired: to.Ptr(false),
		// 																										MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																											{
		// 																												Retention: to.Ptr("P93D"),
		// 																												TimeGrain: to.Ptr("PT1M"),
		// 																											},
		// 																											{
		// 																												Retention: to.Ptr("P93D"),
		// 																												TimeGrain: to.Ptr("PT5M"),
		// 																											},
		// 																											{
		// 																												Retention: to.Ptr("P93D"),
		// 																												TimeGrain: to.Ptr("PT15M"),
		// 																											},
		// 																											{
		// 																												Retention: to.Ptr("P93D"),
		// 																												TimeGrain: to.Ptr("PT30M"),
		// 																											},
		// 																											{
		// 																												Retention: to.Ptr("P93D"),
		// 																												TimeGrain: to.Ptr("PT1H"),
		// 																											},
		// 																											{
		// 																												Retention: to.Ptr("P93D"),
		// 																												TimeGrain: to.Ptr("PT6H"),
		// 																											},
		// 																											{
		// 																												Retention: to.Ptr("P93D"),
		// 																												TimeGrain: to.Ptr("PT12H"),
		// 																											},
		// 																											{
		// 																												Retention: to.Ptr("P93D"),
		// 																												TimeGrain: to.Ptr("P1D"),
		// 																										}},
		// 																										Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																										PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																										ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																										SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																											to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																											to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																											to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																											to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																											to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																											to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																											Unit: to.Ptr(armmonitor.MetricUnitBytesPerSecond),
		// 																										},
		// 																										{
		// 																											Name: &armmonitor.LocalizableString{
		// 																												LocalizedValue: to.Ptr("OS Disk Read Operations/Sec"),
		// 																												Value: to.Ptr("OS Disk Read Operations/Sec"),
		// 																											},
		// 																											Dimensions: []*armmonitor.LocalizableString{
		// 																												{
		// 																													LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																													Value: to.Ptr("Microsoft.ResourceId"),
		// 																												},
		// 																												{
		// 																													LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																													Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																											}},
		// 																											DisplayDescription: to.Ptr("Read IOPS from a single disk during monitoring period for OS disk"),
		// 																											ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/OS Disk Read Operations/Sec"),
		// 																											IsDimensionRequired: to.Ptr(false),
		// 																											MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																												{
		// 																													Retention: to.Ptr("P93D"),
		// 																													TimeGrain: to.Ptr("PT1M"),
		// 																												},
		// 																												{
		// 																													Retention: to.Ptr("P93D"),
		// 																													TimeGrain: to.Ptr("PT5M"),
		// 																												},
		// 																												{
		// 																													Retention: to.Ptr("P93D"),
		// 																													TimeGrain: to.Ptr("PT15M"),
		// 																												},
		// 																												{
		// 																													Retention: to.Ptr("P93D"),
		// 																													TimeGrain: to.Ptr("PT30M"),
		// 																												},
		// 																												{
		// 																													Retention: to.Ptr("P93D"),
		// 																													TimeGrain: to.Ptr("PT1H"),
		// 																												},
		// 																												{
		// 																													Retention: to.Ptr("P93D"),
		// 																													TimeGrain: to.Ptr("PT6H"),
		// 																												},
		// 																												{
		// 																													Retention: to.Ptr("P93D"),
		// 																													TimeGrain: to.Ptr("PT12H"),
		// 																												},
		// 																												{
		// 																													Retention: to.Ptr("P93D"),
		// 																													TimeGrain: to.Ptr("P1D"),
		// 																											}},
		// 																											Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																											PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																											ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																											SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																												to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																												to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																												to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																												to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																												to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																												to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																												Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 																											},
		// 																											{
		// 																												Name: &armmonitor.LocalizableString{
		// 																													LocalizedValue: to.Ptr("OS Disk Write Operations/Sec"),
		// 																													Value: to.Ptr("OS Disk Write Operations/Sec"),
		// 																												},
		// 																												Dimensions: []*armmonitor.LocalizableString{
		// 																													{
		// 																														LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																														Value: to.Ptr("Microsoft.ResourceId"),
		// 																													},
		// 																													{
		// 																														LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																														Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																												}},
		// 																												DisplayDescription: to.Ptr("Write IOPS from a single disk during monitoring period for OS disk"),
		// 																												ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/OS Disk Write Operations/Sec"),
		// 																												IsDimensionRequired: to.Ptr(false),
		// 																												MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																													{
		// 																														Retention: to.Ptr("P93D"),
		// 																														TimeGrain: to.Ptr("PT1M"),
		// 																													},
		// 																													{
		// 																														Retention: to.Ptr("P93D"),
		// 																														TimeGrain: to.Ptr("PT5M"),
		// 																													},
		// 																													{
		// 																														Retention: to.Ptr("P93D"),
		// 																														TimeGrain: to.Ptr("PT15M"),
		// 																													},
		// 																													{
		// 																														Retention: to.Ptr("P93D"),
		// 																														TimeGrain: to.Ptr("PT30M"),
		// 																													},
		// 																													{
		// 																														Retention: to.Ptr("P93D"),
		// 																														TimeGrain: to.Ptr("PT1H"),
		// 																													},
		// 																													{
		// 																														Retention: to.Ptr("P93D"),
		// 																														TimeGrain: to.Ptr("PT6H"),
		// 																													},
		// 																													{
		// 																														Retention: to.Ptr("P93D"),
		// 																														TimeGrain: to.Ptr("PT12H"),
		// 																													},
		// 																													{
		// 																														Retention: to.Ptr("P93D"),
		// 																														TimeGrain: to.Ptr("P1D"),
		// 																												}},
		// 																												Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																												PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																												ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																												SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																													to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																													to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																													to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																													to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																													to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																													to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																													Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 																												},
		// 																												{
		// 																													Name: &armmonitor.LocalizableString{
		// 																														LocalizedValue: to.Ptr("OS Disk Queue Depth"),
		// 																														Value: to.Ptr("OS Disk Queue Depth"),
		// 																													},
		// 																													Dimensions: []*armmonitor.LocalizableString{
		// 																														{
		// 																															LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																															Value: to.Ptr("Microsoft.ResourceId"),
		// 																														},
		// 																														{
		// 																															LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																															Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																													}},
		// 																													DisplayDescription: to.Ptr("OS Disk Queue Depth(or Queue Length)"),
		// 																													ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/OS Disk Queue Depth"),
		// 																													IsDimensionRequired: to.Ptr(false),
		// 																													MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																														{
		// 																															Retention: to.Ptr("P93D"),
		// 																															TimeGrain: to.Ptr("PT1M"),
		// 																														},
		// 																														{
		// 																															Retention: to.Ptr("P93D"),
		// 																															TimeGrain: to.Ptr("PT5M"),
		// 																														},
		// 																														{
		// 																															Retention: to.Ptr("P93D"),
		// 																															TimeGrain: to.Ptr("PT15M"),
		// 																														},
		// 																														{
		// 																															Retention: to.Ptr("P93D"),
		// 																															TimeGrain: to.Ptr("PT30M"),
		// 																														},
		// 																														{
		// 																															Retention: to.Ptr("P93D"),
		// 																															TimeGrain: to.Ptr("PT1H"),
		// 																														},
		// 																														{
		// 																															Retention: to.Ptr("P93D"),
		// 																															TimeGrain: to.Ptr("PT6H"),
		// 																														},
		// 																														{
		// 																															Retention: to.Ptr("P93D"),
		// 																															TimeGrain: to.Ptr("PT12H"),
		// 																														},
		// 																														{
		// 																															Retention: to.Ptr("P93D"),
		// 																															TimeGrain: to.Ptr("P1D"),
		// 																													}},
		// 																													Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																													PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																													ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																													SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																														to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																														to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																														to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																														to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																														to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																														to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																														Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																													},
		// 																													{
		// 																														Name: &armmonitor.LocalizableString{
		// 																															LocalizedValue: to.Ptr("OS Disk Bandwidth Consumed Percentage"),
		// 																															Value: to.Ptr("OS Disk Bandwidth Consumed Percentage"),
		// 																														},
		// 																														Dimensions: []*armmonitor.LocalizableString{
		// 																															{
		// 																																LocalizedValue: to.Ptr("LUN"),
		// 																																Value: to.Ptr("LUN"),
		// 																															},
		// 																															{
		// 																																LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																Value: to.Ptr("Microsoft.ResourceId"),
		// 																															},
		// 																															{
		// 																																LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																														}},
		// 																														DisplayDescription: to.Ptr("Percentage of operating system disk bandwidth consumed per minute"),
		// 																														ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/OS Disk Bandwidth Consumed Percentage"),
		// 																														IsDimensionRequired: to.Ptr(false),
		// 																														MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																															{
		// 																																Retention: to.Ptr("P93D"),
		// 																																TimeGrain: to.Ptr("PT1M"),
		// 																															},
		// 																															{
		// 																																Retention: to.Ptr("P93D"),
		// 																																TimeGrain: to.Ptr("PT5M"),
		// 																															},
		// 																															{
		// 																																Retention: to.Ptr("P93D"),
		// 																																TimeGrain: to.Ptr("PT15M"),
		// 																															},
		// 																															{
		// 																																Retention: to.Ptr("P93D"),
		// 																																TimeGrain: to.Ptr("PT30M"),
		// 																															},
		// 																															{
		// 																																Retention: to.Ptr("P93D"),
		// 																																TimeGrain: to.Ptr("PT1H"),
		// 																															},
		// 																															{
		// 																																Retention: to.Ptr("P93D"),
		// 																																TimeGrain: to.Ptr("PT6H"),
		// 																															},
		// 																															{
		// 																																Retention: to.Ptr("P93D"),
		// 																																TimeGrain: to.Ptr("PT12H"),
		// 																															},
		// 																															{
		// 																																Retention: to.Ptr("P93D"),
		// 																																TimeGrain: to.Ptr("P1D"),
		// 																														}},
		// 																														Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																														PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																														ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																														SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																															to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																															to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																															to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																															to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																															to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																															to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																															Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																														},
		// 																														{
		// 																															Name: &armmonitor.LocalizableString{
		// 																																LocalizedValue: to.Ptr("OS Disk IOPS Consumed Percentage"),
		// 																																Value: to.Ptr("OS Disk IOPS Consumed Percentage"),
		// 																															},
		// 																															Dimensions: []*armmonitor.LocalizableString{
		// 																																{
		// 																																	LocalizedValue: to.Ptr("LUN"),
		// 																																	Value: to.Ptr("LUN"),
		// 																																},
		// 																																{
		// 																																	LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																	Value: to.Ptr("Microsoft.ResourceId"),
		// 																																},
		// 																																{
		// 																																	LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																	Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																															}},
		// 																															DisplayDescription: to.Ptr("Percentage of operating system disk I/Os consumed per minute"),
		// 																															ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/OS Disk IOPS Consumed Percentage"),
		// 																															IsDimensionRequired: to.Ptr(false),
		// 																															MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																{
		// 																																	Retention: to.Ptr("P93D"),
		// 																																	TimeGrain: to.Ptr("PT1M"),
		// 																																},
		// 																																{
		// 																																	Retention: to.Ptr("P93D"),
		// 																																	TimeGrain: to.Ptr("PT5M"),
		// 																																},
		// 																																{
		// 																																	Retention: to.Ptr("P93D"),
		// 																																	TimeGrain: to.Ptr("PT15M"),
		// 																																},
		// 																																{
		// 																																	Retention: to.Ptr("P93D"),
		// 																																	TimeGrain: to.Ptr("PT30M"),
		// 																																},
		// 																																{
		// 																																	Retention: to.Ptr("P93D"),
		// 																																	TimeGrain: to.Ptr("PT1H"),
		// 																																},
		// 																																{
		// 																																	Retention: to.Ptr("P93D"),
		// 																																	TimeGrain: to.Ptr("PT6H"),
		// 																																},
		// 																																{
		// 																																	Retention: to.Ptr("P93D"),
		// 																																	TimeGrain: to.Ptr("PT12H"),
		// 																																},
		// 																																{
		// 																																	Retention: to.Ptr("P93D"),
		// 																																	TimeGrain: to.Ptr("P1D"),
		// 																															}},
		// 																															Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																															PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																															ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																															SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																															},
		// 																															{
		// 																																Name: &armmonitor.LocalizableString{
		// 																																	LocalizedValue: to.Ptr("OS Disk Target Bandwidth"),
		// 																																	Value: to.Ptr("OS Disk Target Bandwidth"),
		// 																																},
		// 																																Dimensions: []*armmonitor.LocalizableString{
		// 																																	{
		// 																																		LocalizedValue: to.Ptr("LUN"),
		// 																																		Value: to.Ptr("LUN"),
		// 																																	},
		// 																																	{
		// 																																		LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																		Value: to.Ptr("Microsoft.ResourceId"),
		// 																																	},
		// 																																	{
		// 																																		LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																		Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																}},
		// 																																DisplayDescription: to.Ptr("Baseline bytes per second throughput OS Disk can achieve without bursting"),
		// 																																ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/OS Disk Target Bandwidth"),
		// 																																IsDimensionRequired: to.Ptr(false),
		// 																																MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																	{
		// 																																		Retention: to.Ptr("P93D"),
		// 																																		TimeGrain: to.Ptr("PT1M"),
		// 																																	},
		// 																																	{
		// 																																		Retention: to.Ptr("P93D"),
		// 																																		TimeGrain: to.Ptr("PT5M"),
		// 																																	},
		// 																																	{
		// 																																		Retention: to.Ptr("P93D"),
		// 																																		TimeGrain: to.Ptr("PT15M"),
		// 																																	},
		// 																																	{
		// 																																		Retention: to.Ptr("P93D"),
		// 																																		TimeGrain: to.Ptr("PT30M"),
		// 																																	},
		// 																																	{
		// 																																		Retention: to.Ptr("P93D"),
		// 																																		TimeGrain: to.Ptr("PT1H"),
		// 																																	},
		// 																																	{
		// 																																		Retention: to.Ptr("P93D"),
		// 																																		TimeGrain: to.Ptr("PT6H"),
		// 																																	},
		// 																																	{
		// 																																		Retention: to.Ptr("P93D"),
		// 																																		TimeGrain: to.Ptr("PT12H"),
		// 																																	},
		// 																																	{
		// 																																		Retention: to.Ptr("P93D"),
		// 																																		TimeGrain: to.Ptr("P1D"),
		// 																																}},
		// 																																Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																	to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																	to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																	to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																	to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																	to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																	to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																	Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																																},
		// 																																{
		// 																																	Name: &armmonitor.LocalizableString{
		// 																																		LocalizedValue: to.Ptr("OS Disk Target IOPS"),
		// 																																		Value: to.Ptr("OS Disk Target IOPS"),
		// 																																	},
		// 																																	Dimensions: []*armmonitor.LocalizableString{
		// 																																		{
		// 																																			LocalizedValue: to.Ptr("LUN"),
		// 																																			Value: to.Ptr("LUN"),
		// 																																		},
		// 																																		{
		// 																																			LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																			Value: to.Ptr("Microsoft.ResourceId"),
		// 																																		},
		// 																																		{
		// 																																			LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																			Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																	}},
		// 																																	DisplayDescription: to.Ptr("Baseline IOPS OS Disk can achieve without bursting"),
		// 																																	ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/OS Disk Target IOPS"),
		// 																																	IsDimensionRequired: to.Ptr(false),
		// 																																	MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																		{
		// 																																			Retention: to.Ptr("P93D"),
		// 																																			TimeGrain: to.Ptr("PT1M"),
		// 																																		},
		// 																																		{
		// 																																			Retention: to.Ptr("P93D"),
		// 																																			TimeGrain: to.Ptr("PT5M"),
		// 																																		},
		// 																																		{
		// 																																			Retention: to.Ptr("P93D"),
		// 																																			TimeGrain: to.Ptr("PT15M"),
		// 																																		},
		// 																																		{
		// 																																			Retention: to.Ptr("P93D"),
		// 																																			TimeGrain: to.Ptr("PT30M"),
		// 																																		},
		// 																																		{
		// 																																			Retention: to.Ptr("P93D"),
		// 																																			TimeGrain: to.Ptr("PT1H"),
		// 																																		},
		// 																																		{
		// 																																			Retention: to.Ptr("P93D"),
		// 																																			TimeGrain: to.Ptr("PT6H"),
		// 																																		},
		// 																																		{
		// 																																			Retention: to.Ptr("P93D"),
		// 																																			TimeGrain: to.Ptr("PT12H"),
		// 																																		},
		// 																																		{
		// 																																			Retention: to.Ptr("P93D"),
		// 																																			TimeGrain: to.Ptr("P1D"),
		// 																																	}},
		// 																																	Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																	PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																	ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																	SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																		to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																		to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																		to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																		to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																		to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																		to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																		Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																																	},
		// 																																	{
		// 																																		Name: &armmonitor.LocalizableString{
		// 																																			LocalizedValue: to.Ptr("OS Disk Max Burst Bandwidth"),
		// 																																			Value: to.Ptr("OS Disk Max Burst Bandwidth"),
		// 																																		},
		// 																																		Dimensions: []*armmonitor.LocalizableString{
		// 																																			{
		// 																																				LocalizedValue: to.Ptr("LUN"),
		// 																																				Value: to.Ptr("LUN"),
		// 																																			},
		// 																																			{
		// 																																				LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																				Value: to.Ptr("Microsoft.ResourceId"),
		// 																																			},
		// 																																			{
		// 																																				LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																				Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																		}},
		// 																																		DisplayDescription: to.Ptr("Maximum bytes per second throughput OS Disk can achieve with bursting"),
		// 																																		ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/OS Disk Max Burst Bandwidth"),
		// 																																		IsDimensionRequired: to.Ptr(false),
		// 																																		MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																			{
		// 																																				Retention: to.Ptr("P93D"),
		// 																																				TimeGrain: to.Ptr("PT1M"),
		// 																																			},
		// 																																			{
		// 																																				Retention: to.Ptr("P93D"),
		// 																																				TimeGrain: to.Ptr("PT5M"),
		// 																																			},
		// 																																			{
		// 																																				Retention: to.Ptr("P93D"),
		// 																																				TimeGrain: to.Ptr("PT15M"),
		// 																																			},
		// 																																			{
		// 																																				Retention: to.Ptr("P93D"),
		// 																																				TimeGrain: to.Ptr("PT30M"),
		// 																																			},
		// 																																			{
		// 																																				Retention: to.Ptr("P93D"),
		// 																																				TimeGrain: to.Ptr("PT1H"),
		// 																																			},
		// 																																			{
		// 																																				Retention: to.Ptr("P93D"),
		// 																																				TimeGrain: to.Ptr("PT6H"),
		// 																																			},
		// 																																			{
		// 																																				Retention: to.Ptr("P93D"),
		// 																																				TimeGrain: to.Ptr("PT12H"),
		// 																																			},
		// 																																			{
		// 																																				Retention: to.Ptr("P93D"),
		// 																																				TimeGrain: to.Ptr("P1D"),
		// 																																		}},
		// 																																		Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																		PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																		ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																		SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																			to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																			to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																			to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																			to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																			to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																			to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																			Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																																		},
		// 																																		{
		// 																																			Name: &armmonitor.LocalizableString{
		// 																																				LocalizedValue: to.Ptr("OS Disk Max Burst IOPS"),
		// 																																				Value: to.Ptr("OS Disk Max Burst IOPS"),
		// 																																			},
		// 																																			Dimensions: []*armmonitor.LocalizableString{
		// 																																				{
		// 																																					LocalizedValue: to.Ptr("LUN"),
		// 																																					Value: to.Ptr("LUN"),
		// 																																				},
		// 																																				{
		// 																																					LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																					Value: to.Ptr("Microsoft.ResourceId"),
		// 																																				},
		// 																																				{
		// 																																					LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																					Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																			}},
		// 																																			DisplayDescription: to.Ptr("Maximum IOPS OS Disk can achieve with bursting"),
		// 																																			ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/OS Disk Max Burst IOPS"),
		// 																																			IsDimensionRequired: to.Ptr(false),
		// 																																			MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																				{
		// 																																					Retention: to.Ptr("P93D"),
		// 																																					TimeGrain: to.Ptr("PT1M"),
		// 																																				},
		// 																																				{
		// 																																					Retention: to.Ptr("P93D"),
		// 																																					TimeGrain: to.Ptr("PT5M"),
		// 																																				},
		// 																																				{
		// 																																					Retention: to.Ptr("P93D"),
		// 																																					TimeGrain: to.Ptr("PT15M"),
		// 																																				},
		// 																																				{
		// 																																					Retention: to.Ptr("P93D"),
		// 																																					TimeGrain: to.Ptr("PT30M"),
		// 																																				},
		// 																																				{
		// 																																					Retention: to.Ptr("P93D"),
		// 																																					TimeGrain: to.Ptr("PT1H"),
		// 																																				},
		// 																																				{
		// 																																					Retention: to.Ptr("P93D"),
		// 																																					TimeGrain: to.Ptr("PT6H"),
		// 																																				},
		// 																																				{
		// 																																					Retention: to.Ptr("P93D"),
		// 																																					TimeGrain: to.Ptr("PT12H"),
		// 																																				},
		// 																																				{
		// 																																					Retention: to.Ptr("P93D"),
		// 																																					TimeGrain: to.Ptr("P1D"),
		// 																																			}},
		// 																																			Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																			PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																			ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																			SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																				to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																				to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																				to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																				to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																				to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																				to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																				Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																																			},
		// 																																			{
		// 																																				Name: &armmonitor.LocalizableString{
		// 																																					LocalizedValue: to.Ptr("OS Disk Used Burst BPS Credits Percentage"),
		// 																																					Value: to.Ptr("OS Disk Used Burst BPS Credits Percentage"),
		// 																																				},
		// 																																				Dimensions: []*armmonitor.LocalizableString{
		// 																																					{
		// 																																						LocalizedValue: to.Ptr("LUN"),
		// 																																						Value: to.Ptr("LUN"),
		// 																																					},
		// 																																					{
		// 																																						LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																						Value: to.Ptr("Microsoft.ResourceId"),
		// 																																					},
		// 																																					{
		// 																																						LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																						Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																				}},
		// 																																				DisplayDescription: to.Ptr("Percentage of OS Disk burst bandwidth credits used so far"),
		// 																																				ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/OS Disk Used Burst BPS Credits Percentage"),
		// 																																				IsDimensionRequired: to.Ptr(false),
		// 																																				MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																					{
		// 																																						Retention: to.Ptr("P93D"),
		// 																																						TimeGrain: to.Ptr("PT1M"),
		// 																																					},
		// 																																					{
		// 																																						Retention: to.Ptr("P93D"),
		// 																																						TimeGrain: to.Ptr("PT5M"),
		// 																																					},
		// 																																					{
		// 																																						Retention: to.Ptr("P93D"),
		// 																																						TimeGrain: to.Ptr("PT15M"),
		// 																																					},
		// 																																					{
		// 																																						Retention: to.Ptr("P93D"),
		// 																																						TimeGrain: to.Ptr("PT30M"),
		// 																																					},
		// 																																					{
		// 																																						Retention: to.Ptr("P93D"),
		// 																																						TimeGrain: to.Ptr("PT1H"),
		// 																																					},
		// 																																					{
		// 																																						Retention: to.Ptr("P93D"),
		// 																																						TimeGrain: to.Ptr("PT6H"),
		// 																																					},
		// 																																					{
		// 																																						Retention: to.Ptr("P93D"),
		// 																																						TimeGrain: to.Ptr("PT12H"),
		// 																																					},
		// 																																					{
		// 																																						Retention: to.Ptr("P93D"),
		// 																																						TimeGrain: to.Ptr("P1D"),
		// 																																				}},
		// 																																				Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																				PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																				ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																				SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																					to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																					to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																					to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																					to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																					to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																					to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																					Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																																				},
		// 																																				{
		// 																																					Name: &armmonitor.LocalizableString{
		// 																																						LocalizedValue: to.Ptr("OS Disk Used Burst IO Credits Percentage"),
		// 																																						Value: to.Ptr("OS Disk Used Burst IO Credits Percentage"),
		// 																																					},
		// 																																					Dimensions: []*armmonitor.LocalizableString{
		// 																																						{
		// 																																							LocalizedValue: to.Ptr("LUN"),
		// 																																							Value: to.Ptr("LUN"),
		// 																																						},
		// 																																						{
		// 																																							LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																							Value: to.Ptr("Microsoft.ResourceId"),
		// 																																						},
		// 																																						{
		// 																																							LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																							Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																					}},
		// 																																					DisplayDescription: to.Ptr("Percentage of OS Disk burst I/O credits used so far"),
		// 																																					ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/OS Disk Used Burst IO Credits Percentage"),
		// 																																					IsDimensionRequired: to.Ptr(false),
		// 																																					MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																						{
		// 																																							Retention: to.Ptr("P93D"),
		// 																																							TimeGrain: to.Ptr("PT1M"),
		// 																																						},
		// 																																						{
		// 																																							Retention: to.Ptr("P93D"),
		// 																																							TimeGrain: to.Ptr("PT5M"),
		// 																																						},
		// 																																						{
		// 																																							Retention: to.Ptr("P93D"),
		// 																																							TimeGrain: to.Ptr("PT15M"),
		// 																																						},
		// 																																						{
		// 																																							Retention: to.Ptr("P93D"),
		// 																																							TimeGrain: to.Ptr("PT30M"),
		// 																																						},
		// 																																						{
		// 																																							Retention: to.Ptr("P93D"),
		// 																																							TimeGrain: to.Ptr("PT1H"),
		// 																																						},
		// 																																						{
		// 																																							Retention: to.Ptr("P93D"),
		// 																																							TimeGrain: to.Ptr("PT6H"),
		// 																																						},
		// 																																						{
		// 																																							Retention: to.Ptr("P93D"),
		// 																																							TimeGrain: to.Ptr("PT12H"),
		// 																																						},
		// 																																						{
		// 																																							Retention: to.Ptr("P93D"),
		// 																																							TimeGrain: to.Ptr("P1D"),
		// 																																					}},
		// 																																					Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																					PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																					ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																					SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																						to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																						to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																						to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																						to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																						to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																						to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																						Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																																					},
		// 																																					{
		// 																																						Name: &armmonitor.LocalizableString{
		// 																																							LocalizedValue: to.Ptr("Inbound Flows"),
		// 																																							Value: to.Ptr("Inbound Flows"),
		// 																																						},
		// 																																						Dimensions: []*armmonitor.LocalizableString{
		// 																																							{
		// 																																								LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																								Value: to.Ptr("Microsoft.ResourceId"),
		// 																																							},
		// 																																							{
		// 																																								LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																								Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																						}},
		// 																																						DisplayDescription: to.Ptr("Inbound Flows are number of current flows in the inbound direction (traffic going into the VM)"),
		// 																																						ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Inbound Flows"),
		// 																																						IsDimensionRequired: to.Ptr(false),
		// 																																						MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																							{
		// 																																								Retention: to.Ptr("P93D"),
		// 																																								TimeGrain: to.Ptr("PT1M"),
		// 																																							},
		// 																																							{
		// 																																								Retention: to.Ptr("P93D"),
		// 																																								TimeGrain: to.Ptr("PT5M"),
		// 																																							},
		// 																																							{
		// 																																								Retention: to.Ptr("P93D"),
		// 																																								TimeGrain: to.Ptr("PT15M"),
		// 																																							},
		// 																																							{
		// 																																								Retention: to.Ptr("P93D"),
		// 																																								TimeGrain: to.Ptr("PT30M"),
		// 																																							},
		// 																																							{
		// 																																								Retention: to.Ptr("P93D"),
		// 																																								TimeGrain: to.Ptr("PT1H"),
		// 																																							},
		// 																																							{
		// 																																								Retention: to.Ptr("P93D"),
		// 																																								TimeGrain: to.Ptr("PT6H"),
		// 																																							},
		// 																																							{
		// 																																								Retention: to.Ptr("P93D"),
		// 																																								TimeGrain: to.Ptr("PT12H"),
		// 																																							},
		// 																																							{
		// 																																								Retention: to.Ptr("P93D"),
		// 																																								TimeGrain: to.Ptr("P1D"),
		// 																																						}},
		// 																																						Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																						PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																						ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																						SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																							to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																							to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																							to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																							to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																							to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																							to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																							Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																																						},
		// 																																						{
		// 																																							Name: &armmonitor.LocalizableString{
		// 																																								LocalizedValue: to.Ptr("Outbound Flows"),
		// 																																								Value: to.Ptr("Outbound Flows"),
		// 																																							},
		// 																																							Dimensions: []*armmonitor.LocalizableString{
		// 																																								{
		// 																																									LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																									Value: to.Ptr("Microsoft.ResourceId"),
		// 																																								},
		// 																																								{
		// 																																									LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																									Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																							}},
		// 																																							DisplayDescription: to.Ptr("Outbound Flows are number of current flows in the outbound direction (traffic going out of the VM)"),
		// 																																							ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Outbound Flows"),
		// 																																							IsDimensionRequired: to.Ptr(false),
		// 																																							MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																								{
		// 																																									Retention: to.Ptr("P93D"),
		// 																																									TimeGrain: to.Ptr("PT1M"),
		// 																																								},
		// 																																								{
		// 																																									Retention: to.Ptr("P93D"),
		// 																																									TimeGrain: to.Ptr("PT5M"),
		// 																																								},
		// 																																								{
		// 																																									Retention: to.Ptr("P93D"),
		// 																																									TimeGrain: to.Ptr("PT15M"),
		// 																																								},
		// 																																								{
		// 																																									Retention: to.Ptr("P93D"),
		// 																																									TimeGrain: to.Ptr("PT30M"),
		// 																																								},
		// 																																								{
		// 																																									Retention: to.Ptr("P93D"),
		// 																																									TimeGrain: to.Ptr("PT1H"),
		// 																																								},
		// 																																								{
		// 																																									Retention: to.Ptr("P93D"),
		// 																																									TimeGrain: to.Ptr("PT6H"),
		// 																																								},
		// 																																								{
		// 																																									Retention: to.Ptr("P93D"),
		// 																																									TimeGrain: to.Ptr("PT12H"),
		// 																																								},
		// 																																								{
		// 																																									Retention: to.Ptr("P93D"),
		// 																																									TimeGrain: to.Ptr("P1D"),
		// 																																							}},
		// 																																							Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																							PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																							ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																							SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																								to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																								to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																								to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																								to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																								to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																								to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																								Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																																							},
		// 																																							{
		// 																																								Name: &armmonitor.LocalizableString{
		// 																																									LocalizedValue: to.Ptr("Inbound Flows Maximum Creation Rate"),
		// 																																									Value: to.Ptr("Inbound Flows Maximum Creation Rate"),
		// 																																								},
		// 																																								Dimensions: []*armmonitor.LocalizableString{
		// 																																									{
		// 																																										LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																										Value: to.Ptr("Microsoft.ResourceId"),
		// 																																									},
		// 																																									{
		// 																																										LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																										Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																								}},
		// 																																								DisplayDescription: to.Ptr("The maximum creation rate of inbound flows (traffic going into the VM)"),
		// 																																								ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Inbound Flows Maximum Creation Rate"),
		// 																																								IsDimensionRequired: to.Ptr(false),
		// 																																								MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																									{
		// 																																										Retention: to.Ptr("P93D"),
		// 																																										TimeGrain: to.Ptr("PT1M"),
		// 																																									},
		// 																																									{
		// 																																										Retention: to.Ptr("P93D"),
		// 																																										TimeGrain: to.Ptr("PT5M"),
		// 																																									},
		// 																																									{
		// 																																										Retention: to.Ptr("P93D"),
		// 																																										TimeGrain: to.Ptr("PT15M"),
		// 																																									},
		// 																																									{
		// 																																										Retention: to.Ptr("P93D"),
		// 																																										TimeGrain: to.Ptr("PT30M"),
		// 																																									},
		// 																																									{
		// 																																										Retention: to.Ptr("P93D"),
		// 																																										TimeGrain: to.Ptr("PT1H"),
		// 																																									},
		// 																																									{
		// 																																										Retention: to.Ptr("P93D"),
		// 																																										TimeGrain: to.Ptr("PT6H"),
		// 																																									},
		// 																																									{
		// 																																										Retention: to.Ptr("P93D"),
		// 																																										TimeGrain: to.Ptr("PT12H"),
		// 																																									},
		// 																																									{
		// 																																										Retention: to.Ptr("P93D"),
		// 																																										TimeGrain: to.Ptr("P1D"),
		// 																																								}},
		// 																																								Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																								PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																								ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																								SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																									to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																									to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																									to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																									to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																									to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																									to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																									Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 																																								},
		// 																																								{
		// 																																									Name: &armmonitor.LocalizableString{
		// 																																										LocalizedValue: to.Ptr("Outbound Flows Maximum Creation Rate"),
		// 																																										Value: to.Ptr("Outbound Flows Maximum Creation Rate"),
		// 																																									},
		// 																																									Dimensions: []*armmonitor.LocalizableString{
		// 																																										{
		// 																																											LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																											Value: to.Ptr("Microsoft.ResourceId"),
		// 																																										},
		// 																																										{
		// 																																											LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																											Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																									}},
		// 																																									DisplayDescription: to.Ptr("The maximum creation rate of outbound flows (traffic going out of the VM)"),
		// 																																									ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Outbound Flows Maximum Creation Rate"),
		// 																																									IsDimensionRequired: to.Ptr(false),
		// 																																									MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																										{
		// 																																											Retention: to.Ptr("P93D"),
		// 																																											TimeGrain: to.Ptr("PT1M"),
		// 																																										},
		// 																																										{
		// 																																											Retention: to.Ptr("P93D"),
		// 																																											TimeGrain: to.Ptr("PT5M"),
		// 																																										},
		// 																																										{
		// 																																											Retention: to.Ptr("P93D"),
		// 																																											TimeGrain: to.Ptr("PT15M"),
		// 																																										},
		// 																																										{
		// 																																											Retention: to.Ptr("P93D"),
		// 																																											TimeGrain: to.Ptr("PT30M"),
		// 																																										},
		// 																																										{
		// 																																											Retention: to.Ptr("P93D"),
		// 																																											TimeGrain: to.Ptr("PT1H"),
		// 																																										},
		// 																																										{
		// 																																											Retention: to.Ptr("P93D"),
		// 																																											TimeGrain: to.Ptr("PT6H"),
		// 																																										},
		// 																																										{
		// 																																											Retention: to.Ptr("P93D"),
		// 																																											TimeGrain: to.Ptr("PT12H"),
		// 																																										},
		// 																																										{
		// 																																											Retention: to.Ptr("P93D"),
		// 																																											TimeGrain: to.Ptr("P1D"),
		// 																																									}},
		// 																																									Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																									PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																									ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																									SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																										to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																										to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																										to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																										to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																										to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																										to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																										Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 																																									},
		// 																																									{
		// 																																										Name: &armmonitor.LocalizableString{
		// 																																											LocalizedValue: to.Ptr("Premium Data Disk Cache Read Hit"),
		// 																																											Value: to.Ptr("Premium Data Disk Cache Read Hit"),
		// 																																										},
		// 																																										Dimensions: []*armmonitor.LocalizableString{
		// 																																											{
		// 																																												LocalizedValue: to.Ptr("LUN"),
		// 																																												Value: to.Ptr("LUN"),
		// 																																											},
		// 																																											{
		// 																																												LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																												Value: to.Ptr("Microsoft.ResourceId"),
		// 																																											},
		// 																																											{
		// 																																												LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																												Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																										}},
		// 																																										DisplayDescription: to.Ptr("Premium Data Disk Cache Read Hit"),
		// 																																										ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Premium Data Disk Cache Read Hit"),
		// 																																										IsDimensionRequired: to.Ptr(false),
		// 																																										MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																											{
		// 																																												Retention: to.Ptr("P93D"),
		// 																																												TimeGrain: to.Ptr("PT1M"),
		// 																																											},
		// 																																											{
		// 																																												Retention: to.Ptr("P93D"),
		// 																																												TimeGrain: to.Ptr("PT5M"),
		// 																																											},
		// 																																											{
		// 																																												Retention: to.Ptr("P93D"),
		// 																																												TimeGrain: to.Ptr("PT15M"),
		// 																																											},
		// 																																											{
		// 																																												Retention: to.Ptr("P93D"),
		// 																																												TimeGrain: to.Ptr("PT30M"),
		// 																																											},
		// 																																											{
		// 																																												Retention: to.Ptr("P93D"),
		// 																																												TimeGrain: to.Ptr("PT1H"),
		// 																																											},
		// 																																											{
		// 																																												Retention: to.Ptr("P93D"),
		// 																																												TimeGrain: to.Ptr("PT6H"),
		// 																																											},
		// 																																											{
		// 																																												Retention: to.Ptr("P93D"),
		// 																																												TimeGrain: to.Ptr("PT12H"),
		// 																																											},
		// 																																											{
		// 																																												Retention: to.Ptr("P93D"),
		// 																																												TimeGrain: to.Ptr("P1D"),
		// 																																										}},
		// 																																										Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																										PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																										ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																										SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																											to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																											to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																											to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																											to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																											to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																											to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																											Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																																										},
		// 																																										{
		// 																																											Name: &armmonitor.LocalizableString{
		// 																																												LocalizedValue: to.Ptr("Premium Data Disk Cache Read Miss"),
		// 																																												Value: to.Ptr("Premium Data Disk Cache Read Miss"),
		// 																																											},
		// 																																											Dimensions: []*armmonitor.LocalizableString{
		// 																																												{
		// 																																													LocalizedValue: to.Ptr("LUN"),
		// 																																													Value: to.Ptr("LUN"),
		// 																																												},
		// 																																												{
		// 																																													LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																													Value: to.Ptr("Microsoft.ResourceId"),
		// 																																												},
		// 																																												{
		// 																																													LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																													Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																											}},
		// 																																											DisplayDescription: to.Ptr("Premium Data Disk Cache Read Miss"),
		// 																																											ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Premium Data Disk Cache Read Miss"),
		// 																																											IsDimensionRequired: to.Ptr(false),
		// 																																											MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																												{
		// 																																													Retention: to.Ptr("P93D"),
		// 																																													TimeGrain: to.Ptr("PT1M"),
		// 																																												},
		// 																																												{
		// 																																													Retention: to.Ptr("P93D"),
		// 																																													TimeGrain: to.Ptr("PT5M"),
		// 																																												},
		// 																																												{
		// 																																													Retention: to.Ptr("P93D"),
		// 																																													TimeGrain: to.Ptr("PT15M"),
		// 																																												},
		// 																																												{
		// 																																													Retention: to.Ptr("P93D"),
		// 																																													TimeGrain: to.Ptr("PT30M"),
		// 																																												},
		// 																																												{
		// 																																													Retention: to.Ptr("P93D"),
		// 																																													TimeGrain: to.Ptr("PT1H"),
		// 																																												},
		// 																																												{
		// 																																													Retention: to.Ptr("P93D"),
		// 																																													TimeGrain: to.Ptr("PT6H"),
		// 																																												},
		// 																																												{
		// 																																													Retention: to.Ptr("P93D"),
		// 																																													TimeGrain: to.Ptr("PT12H"),
		// 																																												},
		// 																																												{
		// 																																													Retention: to.Ptr("P93D"),
		// 																																													TimeGrain: to.Ptr("P1D"),
		// 																																											}},
		// 																																											Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																											PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																											ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																											SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																												to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																												to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																												to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																												to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																												to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																												to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																												Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																																											},
		// 																																											{
		// 																																												Name: &armmonitor.LocalizableString{
		// 																																													LocalizedValue: to.Ptr("Premium OS Disk Cache Read Hit"),
		// 																																													Value: to.Ptr("Premium OS Disk Cache Read Hit"),
		// 																																												},
		// 																																												Dimensions: []*armmonitor.LocalizableString{
		// 																																													{
		// 																																														LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																														Value: to.Ptr("Microsoft.ResourceId"),
		// 																																													},
		// 																																													{
		// 																																														LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																														Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																												}},
		// 																																												DisplayDescription: to.Ptr("Premium OS Disk Cache Read Hit"),
		// 																																												ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Premium OS Disk Cache Read Hit"),
		// 																																												IsDimensionRequired: to.Ptr(false),
		// 																																												MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																													{
		// 																																														Retention: to.Ptr("P93D"),
		// 																																														TimeGrain: to.Ptr("PT1M"),
		// 																																													},
		// 																																													{
		// 																																														Retention: to.Ptr("P93D"),
		// 																																														TimeGrain: to.Ptr("PT5M"),
		// 																																													},
		// 																																													{
		// 																																														Retention: to.Ptr("P93D"),
		// 																																														TimeGrain: to.Ptr("PT15M"),
		// 																																													},
		// 																																													{
		// 																																														Retention: to.Ptr("P93D"),
		// 																																														TimeGrain: to.Ptr("PT30M"),
		// 																																													},
		// 																																													{
		// 																																														Retention: to.Ptr("P93D"),
		// 																																														TimeGrain: to.Ptr("PT1H"),
		// 																																													},
		// 																																													{
		// 																																														Retention: to.Ptr("P93D"),
		// 																																														TimeGrain: to.Ptr("PT6H"),
		// 																																													},
		// 																																													{
		// 																																														Retention: to.Ptr("P93D"),
		// 																																														TimeGrain: to.Ptr("PT12H"),
		// 																																													},
		// 																																													{
		// 																																														Retention: to.Ptr("P93D"),
		// 																																														TimeGrain: to.Ptr("P1D"),
		// 																																												}},
		// 																																												Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																												PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																												ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																												SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																													to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																													to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																													to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																													to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																													to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																													to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																													Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																																												},
		// 																																												{
		// 																																													Name: &armmonitor.LocalizableString{
		// 																																														LocalizedValue: to.Ptr("Premium OS Disk Cache Read Miss"),
		// 																																														Value: to.Ptr("Premium OS Disk Cache Read Miss"),
		// 																																													},
		// 																																													Dimensions: []*armmonitor.LocalizableString{
		// 																																														{
		// 																																															LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																															Value: to.Ptr("Microsoft.ResourceId"),
		// 																																														},
		// 																																														{
		// 																																															LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																															Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																													}},
		// 																																													DisplayDescription: to.Ptr("Premium OS Disk Cache Read Miss"),
		// 																																													ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Premium OS Disk Cache Read Miss"),
		// 																																													IsDimensionRequired: to.Ptr(false),
		// 																																													MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																														{
		// 																																															Retention: to.Ptr("P93D"),
		// 																																															TimeGrain: to.Ptr("PT1M"),
		// 																																														},
		// 																																														{
		// 																																															Retention: to.Ptr("P93D"),
		// 																																															TimeGrain: to.Ptr("PT5M"),
		// 																																														},
		// 																																														{
		// 																																															Retention: to.Ptr("P93D"),
		// 																																															TimeGrain: to.Ptr("PT15M"),
		// 																																														},
		// 																																														{
		// 																																															Retention: to.Ptr("P93D"),
		// 																																															TimeGrain: to.Ptr("PT30M"),
		// 																																														},
		// 																																														{
		// 																																															Retention: to.Ptr("P93D"),
		// 																																															TimeGrain: to.Ptr("PT1H"),
		// 																																														},
		// 																																														{
		// 																																															Retention: to.Ptr("P93D"),
		// 																																															TimeGrain: to.Ptr("PT6H"),
		// 																																														},
		// 																																														{
		// 																																															Retention: to.Ptr("P93D"),
		// 																																															TimeGrain: to.Ptr("PT12H"),
		// 																																														},
		// 																																														{
		// 																																															Retention: to.Ptr("P93D"),
		// 																																															TimeGrain: to.Ptr("P1D"),
		// 																																													}},
		// 																																													Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																													PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																													ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																													SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																														to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																														to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																														to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																														to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																														to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																														to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																														Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																																													},
		// 																																													{
		// 																																														Name: &armmonitor.LocalizableString{
		// 																																															LocalizedValue: to.Ptr("VM Cached Bandwidth Consumed Percentage"),
		// 																																															Value: to.Ptr("VM Cached Bandwidth Consumed Percentage"),
		// 																																														},
		// 																																														Dimensions: []*armmonitor.LocalizableString{
		// 																																															{
		// 																																																LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																																Value: to.Ptr("Microsoft.ResourceId"),
		// 																																															},
		// 																																															{
		// 																																																LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																																Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																														}},
		// 																																														DisplayDescription: to.Ptr("Percentage of cached disk bandwidth consumed by the VM"),
		// 																																														ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/VM Cached Bandwidth Consumed Percentage"),
		// 																																														IsDimensionRequired: to.Ptr(false),
		// 																																														MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																															{
		// 																																																Retention: to.Ptr("P93D"),
		// 																																																TimeGrain: to.Ptr("PT1M"),
		// 																																															},
		// 																																															{
		// 																																																Retention: to.Ptr("P93D"),
		// 																																																TimeGrain: to.Ptr("PT5M"),
		// 																																															},
		// 																																															{
		// 																																																Retention: to.Ptr("P93D"),
		// 																																																TimeGrain: to.Ptr("PT15M"),
		// 																																															},
		// 																																															{
		// 																																																Retention: to.Ptr("P93D"),
		// 																																																TimeGrain: to.Ptr("PT30M"),
		// 																																															},
		// 																																															{
		// 																																																Retention: to.Ptr("P93D"),
		// 																																																TimeGrain: to.Ptr("PT1H"),
		// 																																															},
		// 																																															{
		// 																																																Retention: to.Ptr("P93D"),
		// 																																																TimeGrain: to.Ptr("PT6H"),
		// 																																															},
		// 																																															{
		// 																																																Retention: to.Ptr("P93D"),
		// 																																																TimeGrain: to.Ptr("PT12H"),
		// 																																															},
		// 																																															{
		// 																																																Retention: to.Ptr("P93D"),
		// 																																																TimeGrain: to.Ptr("P1D"),
		// 																																														}},
		// 																																														Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																														PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																														ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																														SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																															to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																															to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																															to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																															to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																															to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																															to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																															Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																																														},
		// 																																														{
		// 																																															Name: &armmonitor.LocalizableString{
		// 																																																LocalizedValue: to.Ptr("VM Cached IOPS Consumed Percentage"),
		// 																																																Value: to.Ptr("VM Cached IOPS Consumed Percentage"),
		// 																																															},
		// 																																															Dimensions: []*armmonitor.LocalizableString{
		// 																																																{
		// 																																																	LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																																	Value: to.Ptr("Microsoft.ResourceId"),
		// 																																																},
		// 																																																{
		// 																																																	LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																																	Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																															}},
		// 																																															DisplayDescription: to.Ptr("Percentage of cached disk IOPS consumed by the VM"),
		// 																																															ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/VM Cached IOPS Consumed Percentage"),
		// 																																															IsDimensionRequired: to.Ptr(false),
		// 																																															MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																																{
		// 																																																	Retention: to.Ptr("P93D"),
		// 																																																	TimeGrain: to.Ptr("PT1M"),
		// 																																																},
		// 																																																{
		// 																																																	Retention: to.Ptr("P93D"),
		// 																																																	TimeGrain: to.Ptr("PT5M"),
		// 																																																},
		// 																																																{
		// 																																																	Retention: to.Ptr("P93D"),
		// 																																																	TimeGrain: to.Ptr("PT15M"),
		// 																																																},
		// 																																																{
		// 																																																	Retention: to.Ptr("P93D"),
		// 																																																	TimeGrain: to.Ptr("PT30M"),
		// 																																																},
		// 																																																{
		// 																																																	Retention: to.Ptr("P93D"),
		// 																																																	TimeGrain: to.Ptr("PT1H"),
		// 																																																},
		// 																																																{
		// 																																																	Retention: to.Ptr("P93D"),
		// 																																																	TimeGrain: to.Ptr("PT6H"),
		// 																																																},
		// 																																																{
		// 																																																	Retention: to.Ptr("P93D"),
		// 																																																	TimeGrain: to.Ptr("PT12H"),
		// 																																																},
		// 																																																{
		// 																																																	Retention: to.Ptr("P93D"),
		// 																																																	TimeGrain: to.Ptr("P1D"),
		// 																																															}},
		// 																																															Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																															PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																															ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																															SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																																to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																																to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																																to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																																to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																																to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																																to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																																Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																																															},
		// 																																															{
		// 																																																Name: &armmonitor.LocalizableString{
		// 																																																	LocalizedValue: to.Ptr("VM Uncached Bandwidth Consumed Percentage"),
		// 																																																	Value: to.Ptr("VM Uncached Bandwidth Consumed Percentage"),
		// 																																																},
		// 																																																Dimensions: []*armmonitor.LocalizableString{
		// 																																																	{
		// 																																																		LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																																		Value: to.Ptr("Microsoft.ResourceId"),
		// 																																																	},
		// 																																																	{
		// 																																																		LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																																		Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																																}},
		// 																																																DisplayDescription: to.Ptr("Percentage of uncached disk bandwidth consumed by the VM"),
		// 																																																ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/VM Uncached Bandwidth Consumed Percentage"),
		// 																																																IsDimensionRequired: to.Ptr(false),
		// 																																																MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																																	{
		// 																																																		Retention: to.Ptr("P93D"),
		// 																																																		TimeGrain: to.Ptr("PT1M"),
		// 																																																	},
		// 																																																	{
		// 																																																		Retention: to.Ptr("P93D"),
		// 																																																		TimeGrain: to.Ptr("PT5M"),
		// 																																																	},
		// 																																																	{
		// 																																																		Retention: to.Ptr("P93D"),
		// 																																																		TimeGrain: to.Ptr("PT15M"),
		// 																																																	},
		// 																																																	{
		// 																																																		Retention: to.Ptr("P93D"),
		// 																																																		TimeGrain: to.Ptr("PT30M"),
		// 																																																	},
		// 																																																	{
		// 																																																		Retention: to.Ptr("P93D"),
		// 																																																		TimeGrain: to.Ptr("PT1H"),
		// 																																																	},
		// 																																																	{
		// 																																																		Retention: to.Ptr("P93D"),
		// 																																																		TimeGrain: to.Ptr("PT6H"),
		// 																																																	},
		// 																																																	{
		// 																																																		Retention: to.Ptr("P93D"),
		// 																																																		TimeGrain: to.Ptr("PT12H"),
		// 																																																	},
		// 																																																	{
		// 																																																		Retention: to.Ptr("P93D"),
		// 																																																		TimeGrain: to.Ptr("P1D"),
		// 																																																}},
		// 																																																Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																																PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																																ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																																SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																																	to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																																	to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																																	to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																																	to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																																	to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																																	to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																																	Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																																																},
		// 																																																{
		// 																																																	Name: &armmonitor.LocalizableString{
		// 																																																		LocalizedValue: to.Ptr("VM Uncached IOPS Consumed Percentage"),
		// 																																																		Value: to.Ptr("VM Uncached IOPS Consumed Percentage"),
		// 																																																	},
		// 																																																	Dimensions: []*armmonitor.LocalizableString{
		// 																																																		{
		// 																																																			LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																																			Value: to.Ptr("Microsoft.ResourceId"),
		// 																																																		},
		// 																																																		{
		// 																																																			LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																																			Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																																	}},
		// 																																																	DisplayDescription: to.Ptr("Percentage of uncached disk IOPS consumed by the VM"),
		// 																																																	ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/VM Uncached IOPS Consumed Percentage"),
		// 																																																	IsDimensionRequired: to.Ptr(false),
		// 																																																	MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																																		{
		// 																																																			Retention: to.Ptr("P93D"),
		// 																																																			TimeGrain: to.Ptr("PT1M"),
		// 																																																		},
		// 																																																		{
		// 																																																			Retention: to.Ptr("P93D"),
		// 																																																			TimeGrain: to.Ptr("PT5M"),
		// 																																																		},
		// 																																																		{
		// 																																																			Retention: to.Ptr("P93D"),
		// 																																																			TimeGrain: to.Ptr("PT15M"),
		// 																																																		},
		// 																																																		{
		// 																																																			Retention: to.Ptr("P93D"),
		// 																																																			TimeGrain: to.Ptr("PT30M"),
		// 																																																		},
		// 																																																		{
		// 																																																			Retention: to.Ptr("P93D"),
		// 																																																			TimeGrain: to.Ptr("PT1H"),
		// 																																																		},
		// 																																																		{
		// 																																																			Retention: to.Ptr("P93D"),
		// 																																																			TimeGrain: to.Ptr("PT6H"),
		// 																																																		},
		// 																																																		{
		// 																																																			Retention: to.Ptr("P93D"),
		// 																																																			TimeGrain: to.Ptr("PT12H"),
		// 																																																		},
		// 																																																		{
		// 																																																			Retention: to.Ptr("P93D"),
		// 																																																			TimeGrain: to.Ptr("P1D"),
		// 																																																	}},
		// 																																																	Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																																	PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																																	ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																																	SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																																		to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																																		to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																																		to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																																		to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																																		to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																																		to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																																		Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																																																	},
		// 																																																	{
		// 																																																		Name: &armmonitor.LocalizableString{
		// 																																																			LocalizedValue: to.Ptr("Network In Total"),
		// 																																																			Value: to.Ptr("Network In Total"),
		// 																																																		},
		// 																																																		Dimensions: []*armmonitor.LocalizableString{
		// 																																																			{
		// 																																																				LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																																				Value: to.Ptr("Microsoft.ResourceId"),
		// 																																																			},
		// 																																																			{
		// 																																																				LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																																				Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																																		}},
		// 																																																		DisplayDescription: to.Ptr("The number of bytes received on all network interfaces by the Virtual Machine(s) (Incoming Traffic)"),
		// 																																																		ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Network In Total"),
		// 																																																		IsDimensionRequired: to.Ptr(false),
		// 																																																		MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																																			{
		// 																																																				Retention: to.Ptr("P93D"),
		// 																																																				TimeGrain: to.Ptr("PT1M"),
		// 																																																			},
		// 																																																			{
		// 																																																				Retention: to.Ptr("P93D"),
		// 																																																				TimeGrain: to.Ptr("PT5M"),
		// 																																																			},
		// 																																																			{
		// 																																																				Retention: to.Ptr("P93D"),
		// 																																																				TimeGrain: to.Ptr("PT15M"),
		// 																																																			},
		// 																																																			{
		// 																																																				Retention: to.Ptr("P93D"),
		// 																																																				TimeGrain: to.Ptr("PT30M"),
		// 																																																			},
		// 																																																			{
		// 																																																				Retention: to.Ptr("P93D"),
		// 																																																				TimeGrain: to.Ptr("PT1H"),
		// 																																																			},
		// 																																																			{
		// 																																																				Retention: to.Ptr("P93D"),
		// 																																																				TimeGrain: to.Ptr("PT6H"),
		// 																																																			},
		// 																																																			{
		// 																																																				Retention: to.Ptr("P93D"),
		// 																																																				TimeGrain: to.Ptr("PT12H"),
		// 																																																			},
		// 																																																			{
		// 																																																				Retention: to.Ptr("P93D"),
		// 																																																				TimeGrain: to.Ptr("P1D"),
		// 																																																		}},
		// 																																																		Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																																		PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																																		ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																																		SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																																			to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																																			to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																																			to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																																			to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																																			to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																																			to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																																			Unit: to.Ptr(armmonitor.MetricUnitBytes),
		// 																																																		},
		// 																																																		{
		// 																																																			Name: &armmonitor.LocalizableString{
		// 																																																				LocalizedValue: to.Ptr("Network Out Total"),
		// 																																																				Value: to.Ptr("Network Out Total"),
		// 																																																			},
		// 																																																			Dimensions: []*armmonitor.LocalizableString{
		// 																																																				{
		// 																																																					LocalizedValue: to.Ptr("Microsoft.ResourceId"),
		// 																																																					Value: to.Ptr("Microsoft.ResourceId"),
		// 																																																				},
		// 																																																				{
		// 																																																					LocalizedValue: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																																					Value: to.Ptr("Microsoft.ResourceGroupName"),
		// 																																																			}},
		// 																																																			DisplayDescription: to.Ptr("The number of bytes out on all network interfaces by the Virtual Machine(s) (Outgoing Traffic)"),
		// 																																																			ID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5/providers/microsoft.insights/metricdefinitions/Network Out Total"),
		// 																																																			IsDimensionRequired: to.Ptr(false),
		// 																																																			MetricAvailabilities: []*armmonitor.MetricAvailability{
		// 																																																				{
		// 																																																					Retention: to.Ptr("P93D"),
		// 																																																					TimeGrain: to.Ptr("PT1M"),
		// 																																																				},
		// 																																																				{
		// 																																																					Retention: to.Ptr("P93D"),
		// 																																																					TimeGrain: to.Ptr("PT5M"),
		// 																																																				},
		// 																																																				{
		// 																																																					Retention: to.Ptr("P93D"),
		// 																																																					TimeGrain: to.Ptr("PT15M"),
		// 																																																				},
		// 																																																				{
		// 																																																					Retention: to.Ptr("P93D"),
		// 																																																					TimeGrain: to.Ptr("PT30M"),
		// 																																																				},
		// 																																																				{
		// 																																																					Retention: to.Ptr("P93D"),
		// 																																																					TimeGrain: to.Ptr("PT1H"),
		// 																																																				},
		// 																																																				{
		// 																																																					Retention: to.Ptr("P93D"),
		// 																																																					TimeGrain: to.Ptr("PT6H"),
		// 																																																				},
		// 																																																				{
		// 																																																					Retention: to.Ptr("P93D"),
		// 																																																					TimeGrain: to.Ptr("PT12H"),
		// 																																																				},
		// 																																																				{
		// 																																																					Retention: to.Ptr("P93D"),
		// 																																																					TimeGrain: to.Ptr("P1D"),
		// 																																																			}},
		// 																																																			Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 																																																			PrimaryAggregationType: to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																																			ResourceID: to.Ptr("subscriptions/92d2a2d8-b514-432d-8cc9-a5f9272630d5"),
		// 																																																			SupportedAggregationTypes: []*armmonitor.MetricAggregationType{
		// 																																																				to.Ptr(armmonitor.MetricAggregationTypeNone),
		// 																																																				to.Ptr(armmonitor.MetricAggregationTypeAverage),
		// 																																																				to.Ptr(armmonitor.MetricAggregationTypeMinimum),
		// 																																																				to.Ptr(armmonitor.MetricAggregationTypeMaximum),
		// 																																																				to.Ptr(armmonitor.MetricAggregationTypeTotal),
		// 																																																				to.Ptr(armmonitor.MetricAggregationTypeCount)},
		// 																																																				Unit: to.Ptr(armmonitor.MetricUnitBytes),
		// 																																																		}},
		// 																																																	}
	}
}
