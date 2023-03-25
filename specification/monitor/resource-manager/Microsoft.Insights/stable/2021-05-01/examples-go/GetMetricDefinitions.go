package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2021-05-01/examples/GetMetricDefinitions.json
func ExampleMetricDefinitionsClient_NewListPager_getMetricDefinitionsWithoutFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMetricDefinitionsClient().NewListPager("subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricDefinitions", &armmonitor.MetricDefinitionsClientListOptions{Metricnamespace: to.Ptr("Microsoft.Web/sites")})
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
		// page.MetricDefinitionCollection = armmonitor.MetricDefinitionCollection{
		// 	Value: []*armmonitor.MetricDefinition{
		// 		{
		// 			Name: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("CPU Time"),
		// 				Value: to.Ptr("CpuTime"),
		// 			},
		// 			Dimensions: []*armmonitor.LocalizableString{
		// 				{
		// 					LocalizedValue: to.Ptr("Instance"),
		// 					Value: to.Ptr("Instance"),
		// 			}},
		// 			ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/CpuTime"),
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
		// 			Namespace: to.Ptr("Microsoft.Web/sites"),
		// 			PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeTotal),
		// 			ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 			SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 				to.Ptr(armmonitor.AggregationTypeNone),
		// 				to.Ptr(armmonitor.AggregationTypeAverage),
		// 				to.Ptr(armmonitor.AggregationTypeMinimum),
		// 				to.Ptr(armmonitor.AggregationTypeMaximum),
		// 				to.Ptr(armmonitor.AggregationTypeTotal),
		// 				to.Ptr(armmonitor.AggregationTypeCount)},
		// 				Unit: to.Ptr(armmonitor.MetricUnitSeconds),
		// 			},
		// 			{
		// 				Name: &armmonitor.LocalizableString{
		// 					LocalizedValue: to.Ptr("Requests"),
		// 					Value: to.Ptr("Requests"),
		// 				},
		// 				Dimensions: []*armmonitor.LocalizableString{
		// 					{
		// 						LocalizedValue: to.Ptr("Instance"),
		// 						Value: to.Ptr("Instance"),
		// 				}},
		// 				ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/Requests"),
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
		// 				Namespace: to.Ptr("Microsoft.Web/sites"),
		// 				PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeTotal),
		// 				ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 				SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 					to.Ptr(armmonitor.AggregationTypeNone),
		// 					to.Ptr(armmonitor.AggregationTypeAverage),
		// 					to.Ptr(armmonitor.AggregationTypeMinimum),
		// 					to.Ptr(armmonitor.AggregationTypeMaximum),
		// 					to.Ptr(armmonitor.AggregationTypeTotal),
		// 					to.Ptr(armmonitor.AggregationTypeCount)},
		// 					Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 				},
		// 				{
		// 					Name: &armmonitor.LocalizableString{
		// 						LocalizedValue: to.Ptr("Data In"),
		// 						Value: to.Ptr("BytesReceived"),
		// 					},
		// 					Dimensions: []*armmonitor.LocalizableString{
		// 						{
		// 							LocalizedValue: to.Ptr("Instance"),
		// 							Value: to.Ptr("Instance"),
		// 					}},
		// 					ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/BytesReceived"),
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
		// 					Namespace: to.Ptr("Microsoft.Web/sites"),
		// 					PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeTotal),
		// 					ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 					SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 						to.Ptr(armmonitor.AggregationTypeNone),
		// 						to.Ptr(armmonitor.AggregationTypeAverage),
		// 						to.Ptr(armmonitor.AggregationTypeMinimum),
		// 						to.Ptr(armmonitor.AggregationTypeMaximum),
		// 						to.Ptr(armmonitor.AggregationTypeTotal),
		// 						to.Ptr(armmonitor.AggregationTypeCount)},
		// 						Unit: to.Ptr(armmonitor.MetricUnitBytes),
		// 					},
		// 					{
		// 						Name: &armmonitor.LocalizableString{
		// 							LocalizedValue: to.Ptr("Data Out"),
		// 							Value: to.Ptr("BytesSent"),
		// 						},
		// 						Dimensions: []*armmonitor.LocalizableString{
		// 							{
		// 								LocalizedValue: to.Ptr("Instance"),
		// 								Value: to.Ptr("Instance"),
		// 						}},
		// 						ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/BytesSent"),
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
		// 						Namespace: to.Ptr("Microsoft.Web/sites"),
		// 						PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeTotal),
		// 						ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 						SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 							to.Ptr(armmonitor.AggregationTypeNone),
		// 							to.Ptr(armmonitor.AggregationTypeAverage),
		// 							to.Ptr(armmonitor.AggregationTypeMinimum),
		// 							to.Ptr(armmonitor.AggregationTypeMaximum),
		// 							to.Ptr(armmonitor.AggregationTypeTotal),
		// 							to.Ptr(armmonitor.AggregationTypeCount)},
		// 							Unit: to.Ptr(armmonitor.MetricUnitBytes),
		// 						},
		// 						{
		// 							Name: &armmonitor.LocalizableString{
		// 								LocalizedValue: to.Ptr("Http 101"),
		// 								Value: to.Ptr("Http101"),
		// 							},
		// 							Dimensions: []*armmonitor.LocalizableString{
		// 								{
		// 									LocalizedValue: to.Ptr("Instance"),
		// 									Value: to.Ptr("Instance"),
		// 							}},
		// 							ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/Http101"),
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
		// 							Namespace: to.Ptr("Microsoft.Web/sites"),
		// 							PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeTotal),
		// 							ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 							SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 								to.Ptr(armmonitor.AggregationTypeNone),
		// 								to.Ptr(armmonitor.AggregationTypeAverage),
		// 								to.Ptr(armmonitor.AggregationTypeMinimum),
		// 								to.Ptr(armmonitor.AggregationTypeMaximum),
		// 								to.Ptr(armmonitor.AggregationTypeTotal),
		// 								to.Ptr(armmonitor.AggregationTypeCount)},
		// 								Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 							},
		// 							{
		// 								Name: &armmonitor.LocalizableString{
		// 									LocalizedValue: to.Ptr("Http 2xx"),
		// 									Value: to.Ptr("Http2xx"),
		// 								},
		// 								Dimensions: []*armmonitor.LocalizableString{
		// 									{
		// 										LocalizedValue: to.Ptr("Instance"),
		// 										Value: to.Ptr("Instance"),
		// 								}},
		// 								ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/Http2xx"),
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
		// 								Namespace: to.Ptr("Microsoft.Web/sites"),
		// 								PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeTotal),
		// 								ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 								SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 									to.Ptr(armmonitor.AggregationTypeNone),
		// 									to.Ptr(armmonitor.AggregationTypeAverage),
		// 									to.Ptr(armmonitor.AggregationTypeMinimum),
		// 									to.Ptr(armmonitor.AggregationTypeMaximum),
		// 									to.Ptr(armmonitor.AggregationTypeTotal),
		// 									to.Ptr(armmonitor.AggregationTypeCount)},
		// 									Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 								},
		// 								{
		// 									Name: &armmonitor.LocalizableString{
		// 										LocalizedValue: to.Ptr("Http 3xx"),
		// 										Value: to.Ptr("Http3xx"),
		// 									},
		// 									Dimensions: []*armmonitor.LocalizableString{
		// 										{
		// 											LocalizedValue: to.Ptr("Instance"),
		// 											Value: to.Ptr("Instance"),
		// 									}},
		// 									ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/Http3xx"),
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
		// 									Namespace: to.Ptr("Microsoft.Web/sites"),
		// 									PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeTotal),
		// 									ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 									SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 										to.Ptr(armmonitor.AggregationTypeNone),
		// 										to.Ptr(armmonitor.AggregationTypeAverage),
		// 										to.Ptr(armmonitor.AggregationTypeMinimum),
		// 										to.Ptr(armmonitor.AggregationTypeMaximum),
		// 										to.Ptr(armmonitor.AggregationTypeTotal),
		// 										to.Ptr(armmonitor.AggregationTypeCount)},
		// 										Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 									},
		// 									{
		// 										Name: &armmonitor.LocalizableString{
		// 											LocalizedValue: to.Ptr("Http 401"),
		// 											Value: to.Ptr("Http401"),
		// 										},
		// 										Dimensions: []*armmonitor.LocalizableString{
		// 											{
		// 												LocalizedValue: to.Ptr("Instance"),
		// 												Value: to.Ptr("Instance"),
		// 										}},
		// 										ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/Http401"),
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
		// 										Namespace: to.Ptr("Microsoft.Web/sites"),
		// 										PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeTotal),
		// 										ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 										SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 											to.Ptr(armmonitor.AggregationTypeNone),
		// 											to.Ptr(armmonitor.AggregationTypeAverage),
		// 											to.Ptr(armmonitor.AggregationTypeMinimum),
		// 											to.Ptr(armmonitor.AggregationTypeMaximum),
		// 											to.Ptr(armmonitor.AggregationTypeTotal),
		// 											to.Ptr(armmonitor.AggregationTypeCount)},
		// 											Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 										},
		// 										{
		// 											Name: &armmonitor.LocalizableString{
		// 												LocalizedValue: to.Ptr("Http 403"),
		// 												Value: to.Ptr("Http403"),
		// 											},
		// 											Dimensions: []*armmonitor.LocalizableString{
		// 												{
		// 													LocalizedValue: to.Ptr("Instance"),
		// 													Value: to.Ptr("Instance"),
		// 											}},
		// 											ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/Http403"),
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
		// 											Namespace: to.Ptr("Microsoft.Web/sites"),
		// 											PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeTotal),
		// 											ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 											SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 												to.Ptr(armmonitor.AggregationTypeNone),
		// 												to.Ptr(armmonitor.AggregationTypeAverage),
		// 												to.Ptr(armmonitor.AggregationTypeMinimum),
		// 												to.Ptr(armmonitor.AggregationTypeMaximum),
		// 												to.Ptr(armmonitor.AggregationTypeTotal),
		// 												to.Ptr(armmonitor.AggregationTypeCount)},
		// 												Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 											},
		// 											{
		// 												Name: &armmonitor.LocalizableString{
		// 													LocalizedValue: to.Ptr("Http 404"),
		// 													Value: to.Ptr("Http404"),
		// 												},
		// 												Dimensions: []*armmonitor.LocalizableString{
		// 													{
		// 														LocalizedValue: to.Ptr("Instance"),
		// 														Value: to.Ptr("Instance"),
		// 												}},
		// 												ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/Http404"),
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
		// 												Namespace: to.Ptr("Microsoft.Web/sites"),
		// 												PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeTotal),
		// 												ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 												SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 													to.Ptr(armmonitor.AggregationTypeNone),
		// 													to.Ptr(armmonitor.AggregationTypeAverage),
		// 													to.Ptr(armmonitor.AggregationTypeMinimum),
		// 													to.Ptr(armmonitor.AggregationTypeMaximum),
		// 													to.Ptr(armmonitor.AggregationTypeTotal),
		// 													to.Ptr(armmonitor.AggregationTypeCount)},
		// 													Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 												},
		// 												{
		// 													Name: &armmonitor.LocalizableString{
		// 														LocalizedValue: to.Ptr("Http 406"),
		// 														Value: to.Ptr("Http406"),
		// 													},
		// 													Dimensions: []*armmonitor.LocalizableString{
		// 														{
		// 															LocalizedValue: to.Ptr("Instance"),
		// 															Value: to.Ptr("Instance"),
		// 													}},
		// 													ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/Http406"),
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
		// 													Namespace: to.Ptr("Microsoft.Web/sites"),
		// 													PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeTotal),
		// 													ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 													SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 														to.Ptr(armmonitor.AggregationTypeNone),
		// 														to.Ptr(armmonitor.AggregationTypeAverage),
		// 														to.Ptr(armmonitor.AggregationTypeMinimum),
		// 														to.Ptr(armmonitor.AggregationTypeMaximum),
		// 														to.Ptr(armmonitor.AggregationTypeTotal),
		// 														to.Ptr(armmonitor.AggregationTypeCount)},
		// 														Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 													},
		// 													{
		// 														Name: &armmonitor.LocalizableString{
		// 															LocalizedValue: to.Ptr("Http 4xx"),
		// 															Value: to.Ptr("Http4xx"),
		// 														},
		// 														Dimensions: []*armmonitor.LocalizableString{
		// 															{
		// 																LocalizedValue: to.Ptr("Instance"),
		// 																Value: to.Ptr("Instance"),
		// 														}},
		// 														ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/Http4xx"),
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
		// 														Namespace: to.Ptr("Microsoft.Web/sites"),
		// 														PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeTotal),
		// 														ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 														SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 															to.Ptr(armmonitor.AggregationTypeNone),
		// 															to.Ptr(armmonitor.AggregationTypeAverage),
		// 															to.Ptr(armmonitor.AggregationTypeMinimum),
		// 															to.Ptr(armmonitor.AggregationTypeMaximum),
		// 															to.Ptr(armmonitor.AggregationTypeTotal),
		// 															to.Ptr(armmonitor.AggregationTypeCount)},
		// 															Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 														},
		// 														{
		// 															Name: &armmonitor.LocalizableString{
		// 																LocalizedValue: to.Ptr("Http Server Errors"),
		// 																Value: to.Ptr("Http5xx"),
		// 															},
		// 															Dimensions: []*armmonitor.LocalizableString{
		// 																{
		// 																	LocalizedValue: to.Ptr("Instance"),
		// 																	Value: to.Ptr("Instance"),
		// 															}},
		// 															ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/Http5xx"),
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
		// 															Namespace: to.Ptr("Microsoft.Web/sites"),
		// 															PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeTotal),
		// 															ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 															SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																to.Ptr(armmonitor.AggregationTypeNone),
		// 																to.Ptr(armmonitor.AggregationTypeAverage),
		// 																to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																to.Ptr(armmonitor.AggregationTypeTotal),
		// 																to.Ptr(armmonitor.AggregationTypeCount)},
		// 																Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 															},
		// 															{
		// 																Name: &armmonitor.LocalizableString{
		// 																	LocalizedValue: to.Ptr("Memory working set"),
		// 																	Value: to.Ptr("MemoryWorkingSet"),
		// 																},
		// 																Dimensions: []*armmonitor.LocalizableString{
		// 																	{
		// 																		LocalizedValue: to.Ptr("Instance"),
		// 																		Value: to.Ptr("Instance"),
		// 																}},
		// 																ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/MemoryWorkingSet"),
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
		// 																Namespace: to.Ptr("Microsoft.Web/sites"),
		// 																PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 																SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																	to.Ptr(armmonitor.AggregationTypeNone),
		// 																	to.Ptr(armmonitor.AggregationTypeAverage),
		// 																	to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																	to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																	to.Ptr(armmonitor.AggregationTypeTotal),
		// 																	to.Ptr(armmonitor.AggregationTypeCount)},
		// 																	Unit: to.Ptr(armmonitor.MetricUnitBytes),
		// 																},
		// 																{
		// 																	Name: &armmonitor.LocalizableString{
		// 																		LocalizedValue: to.Ptr("Average memory working set"),
		// 																		Value: to.Ptr("AverageMemoryWorkingSet"),
		// 																	},
		// 																	Dimensions: []*armmonitor.LocalizableString{
		// 																		{
		// 																			LocalizedValue: to.Ptr("Instance"),
		// 																			Value: to.Ptr("Instance"),
		// 																	}},
		// 																	ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/AverageMemoryWorkingSet"),
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
		// 																	Namespace: to.Ptr("Microsoft.Web/sites"),
		// 																	PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																	ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 																	SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																		to.Ptr(armmonitor.AggregationTypeNone),
		// 																		to.Ptr(armmonitor.AggregationTypeAverage),
		// 																		to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																		to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																		to.Ptr(armmonitor.AggregationTypeTotal),
		// 																		to.Ptr(armmonitor.AggregationTypeCount)},
		// 																		Unit: to.Ptr(armmonitor.MetricUnitBytes),
		// 																	},
		// 																	{
		// 																		Name: &armmonitor.LocalizableString{
		// 																			LocalizedValue: to.Ptr("Average Response Time"),
		// 																			Value: to.Ptr("AverageResponseTime"),
		// 																		},
		// 																		Dimensions: []*armmonitor.LocalizableString{
		// 																			{
		// 																				LocalizedValue: to.Ptr("Instance"),
		// 																				Value: to.Ptr("Instance"),
		// 																		}},
		// 																		ID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricdefinitions/AverageResponseTime"),
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
		// 																		Namespace: to.Ptr("Microsoft.Web/sites"),
		// 																		PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																		ResourceID: to.Ptr("/subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest"),
		// 																		SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																			to.Ptr(armmonitor.AggregationTypeNone),
		// 																			to.Ptr(armmonitor.AggregationTypeAverage),
		// 																			to.Ptr(armmonitor.AggregationTypeMinimum),
		// 																			to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																			to.Ptr(armmonitor.AggregationTypeTotal),
		// 																			to.Ptr(armmonitor.AggregationTypeCount)},
		// 																			Unit: to.Ptr(armmonitor.MetricUnitSeconds),
		// 																	}},
		// 																}
	}
}
