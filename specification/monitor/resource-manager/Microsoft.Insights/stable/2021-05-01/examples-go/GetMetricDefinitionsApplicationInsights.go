package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2021-05-01/examples/GetMetricDefinitionsApplicationInsights.json
func ExampleMetricDefinitionsClient_NewListPager_getApplicationInsightsMetricDefinitionsWithoutFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMetricDefinitionsClient().NewListPager("subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions", &armmonitor.MetricDefinitionsClientListOptions{Metricnamespace: to.Ptr("microsoft.insights/components")})
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
		// 				LocalizedValue: to.Ptr("Availability"),
		// 				Value: to.Ptr("availabilityResults/availabilityPercentage"),
		// 			},
		// 			Category: to.Ptr("Availability"),
		// 			Dimensions: []*armmonitor.LocalizableString{
		// 				{
		// 					LocalizedValue: to.Ptr("Test name"),
		// 					Value: to.Ptr("availabilityResult/name"),
		// 				},
		// 				{
		// 					LocalizedValue: to.Ptr("Run location"),
		// 					Value: to.Ptr("availabilityResult/location"),
		// 			}},
		// 			DisplayDescription: to.Ptr("Percentage of successfully completed availability tests"),
		// 			ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/availabilityResults/availabilityPercentage"),
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
		// 			Namespace: to.Ptr("microsoft.insights/components"),
		// 			PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 			ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 			SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 				to.Ptr(armmonitor.AggregationTypeAverage)},
		// 				Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 			},
		// 			{
		// 				Name: &armmonitor.LocalizableString{
		// 					LocalizedValue: to.Ptr("Availability tests"),
		// 					Value: to.Ptr("availabilityResults/count"),
		// 				},
		// 				Category: to.Ptr("Availability"),
		// 				Dimensions: []*armmonitor.LocalizableString{
		// 					{
		// 						LocalizedValue: to.Ptr("Test name"),
		// 						Value: to.Ptr("availabilityResult/name"),
		// 					},
		// 					{
		// 						LocalizedValue: to.Ptr("Run location"),
		// 						Value: to.Ptr("availabilityResult/location"),
		// 					},
		// 					{
		// 						LocalizedValue: to.Ptr("Test result"),
		// 						Value: to.Ptr("availabilityResult/success"),
		// 				}},
		// 				DisplayDescription: to.Ptr("Count of availability tests"),
		// 				ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/availabilityResults/count"),
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
		// 				Namespace: to.Ptr("microsoft.insights/components"),
		// 				PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeCount),
		// 				ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 				SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 					to.Ptr(armmonitor.AggregationTypeCount)},
		// 					Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 				},
		// 				{
		// 					Name: &armmonitor.LocalizableString{
		// 						LocalizedValue: to.Ptr("Availability test duration"),
		// 						Value: to.Ptr("availabilityResults/duration"),
		// 					},
		// 					Category: to.Ptr("Availability"),
		// 					Dimensions: []*armmonitor.LocalizableString{
		// 						{
		// 							LocalizedValue: to.Ptr("Test name"),
		// 							Value: to.Ptr("availabilityResult/name"),
		// 						},
		// 						{
		// 							LocalizedValue: to.Ptr("Run location"),
		// 							Value: to.Ptr("availabilityResult/location"),
		// 						},
		// 						{
		// 							LocalizedValue: to.Ptr("Test result"),
		// 							Value: to.Ptr("availabilityResult/success"),
		// 					}},
		// 					DisplayDescription: to.Ptr("Availability test duration"),
		// 					ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/availabilityResults/duration"),
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
		// 					Namespace: to.Ptr("microsoft.insights/components"),
		// 					PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 					ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 					SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 						to.Ptr(armmonitor.AggregationTypeAverage),
		// 						to.Ptr(armmonitor.AggregationTypeMaximum),
		// 						to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 						Unit: to.Ptr(armmonitor.MetricUnitMilliSeconds),
		// 					},
		// 					{
		// 						Name: &armmonitor.LocalizableString{
		// 							LocalizedValue: to.Ptr("Page load network connect time"),
		// 							Value: to.Ptr("browserTimings/networkDuration"),
		// 						},
		// 						Category: to.Ptr("Browser"),
		// 						DisplayDescription: to.Ptr("Time between user request and network connection. Includes DNS lookup and transport connection."),
		// 						ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/browserTimings/networkDuration"),
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
		// 						Namespace: to.Ptr("microsoft.insights/components"),
		// 						PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 						ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 						SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 							to.Ptr(armmonitor.AggregationTypeAverage),
		// 							to.Ptr(armmonitor.AggregationTypeMaximum),
		// 							to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 							Unit: to.Ptr(armmonitor.MetricUnitMilliSeconds),
		// 						},
		// 						{
		// 							Name: &armmonitor.LocalizableString{
		// 								LocalizedValue: to.Ptr("Client processing time"),
		// 								Value: to.Ptr("browserTimings/processingDuration"),
		// 							},
		// 							Category: to.Ptr("Browser"),
		// 							DisplayDescription: to.Ptr("Time between receiving the last byte of a document until the DOM is loaded. Async requests may still be processing."),
		// 							ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/browserTimings/processingDuration"),
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
		// 							Namespace: to.Ptr("microsoft.insights/components"),
		// 							PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 							ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 							SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 								to.Ptr(armmonitor.AggregationTypeAverage),
		// 								to.Ptr(armmonitor.AggregationTypeMaximum),
		// 								to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 								Unit: to.Ptr(armmonitor.MetricUnitMilliSeconds),
		// 							},
		// 							{
		// 								Name: &armmonitor.LocalizableString{
		// 									LocalizedValue: to.Ptr("Receiving response time"),
		// 									Value: to.Ptr("browserTimings/receiveDuration"),
		// 								},
		// 								Category: to.Ptr("Browser"),
		// 								DisplayDescription: to.Ptr("Time between the first and last bytes, or until disconnection."),
		// 								ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/browserTimings/receiveDuration"),
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
		// 								Namespace: to.Ptr("microsoft.insights/components"),
		// 								PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 								ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 								SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 									to.Ptr(armmonitor.AggregationTypeAverage),
		// 									to.Ptr(armmonitor.AggregationTypeMaximum),
		// 									to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 									Unit: to.Ptr(armmonitor.MetricUnitMilliSeconds),
		// 								},
		// 								{
		// 									Name: &armmonitor.LocalizableString{
		// 										LocalizedValue: to.Ptr("Send request time"),
		// 										Value: to.Ptr("browserTimings/sendDuration"),
		// 									},
		// 									Category: to.Ptr("Browser"),
		// 									DisplayDescription: to.Ptr("Time between network connection and receiving the first byte."),
		// 									ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/browserTimings/sendDuration"),
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
		// 									Namespace: to.Ptr("microsoft.insights/components"),
		// 									PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 									ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 									SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 										to.Ptr(armmonitor.AggregationTypeAverage),
		// 										to.Ptr(armmonitor.AggregationTypeMaximum),
		// 										to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 										Unit: to.Ptr(armmonitor.MetricUnitMilliSeconds),
		// 									},
		// 									{
		// 										Name: &armmonitor.LocalizableString{
		// 											LocalizedValue: to.Ptr("Browser page load time"),
		// 											Value: to.Ptr("browserTimings/totalDuration"),
		// 										},
		// 										Category: to.Ptr("Browser"),
		// 										DisplayDescription: to.Ptr("Time from user request until DOM, stylesheets, scripts and images are loaded."),
		// 										ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/browserTimings/totalDuration"),
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
		// 										Namespace: to.Ptr("microsoft.insights/components"),
		// 										PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 										ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 										SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 											to.Ptr(armmonitor.AggregationTypeAverage),
		// 											to.Ptr(armmonitor.AggregationTypeMaximum),
		// 											to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 											Unit: to.Ptr(armmonitor.MetricUnitMilliSeconds),
		// 										},
		// 										{
		// 											Name: &armmonitor.LocalizableString{
		// 												LocalizedValue: to.Ptr("Dependency calls"),
		// 												Value: to.Ptr("dependencies/count"),
		// 											},
		// 											Category: to.Ptr("Server"),
		// 											Dimensions: []*armmonitor.LocalizableString{
		// 												{
		// 													LocalizedValue: to.Ptr("Dependency type"),
		// 													Value: to.Ptr("dependency/type"),
		// 												},
		// 												{
		// 													LocalizedValue: to.Ptr("Dependency performance"),
		// 													Value: to.Ptr("dependency/performanceBucket"),
		// 												},
		// 												{
		// 													LocalizedValue: to.Ptr("Successful call"),
		// 													Value: to.Ptr("dependency/success"),
		// 												},
		// 												{
		// 													LocalizedValue: to.Ptr("Target of a dependency call"),
		// 													Value: to.Ptr("dependency/target"),
		// 												},
		// 												{
		// 													LocalizedValue: to.Ptr("Result code"),
		// 													Value: to.Ptr("dependency/resultCode"),
		// 												},
		// 												{
		// 													LocalizedValue: to.Ptr("Is traffic synthetic"),
		// 													Value: to.Ptr("operation/synthetic"),
		// 												},
		// 												{
		// 													LocalizedValue: to.Ptr("Cloud role instance"),
		// 													Value: to.Ptr("cloud/roleInstance"),
		// 												},
		// 												{
		// 													LocalizedValue: to.Ptr("Cloud role name"),
		// 													Value: to.Ptr("cloud/roleName"),
		// 											}},
		// 											DisplayDescription: to.Ptr("Count of calls made by the application to external resources."),
		// 											ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/dependencies/count"),
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
		// 											Namespace: to.Ptr("microsoft.insights/components"),
		// 											PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeCount),
		// 											ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 											SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 												to.Ptr(armmonitor.AggregationTypeCount)},
		// 												Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 											},
		// 											{
		// 												Name: &armmonitor.LocalizableString{
		// 													LocalizedValue: to.Ptr("Dependency duration"),
		// 													Value: to.Ptr("dependencies/duration"),
		// 												},
		// 												Category: to.Ptr("Server"),
		// 												Dimensions: []*armmonitor.LocalizableString{
		// 													{
		// 														LocalizedValue: to.Ptr("Dependency type"),
		// 														Value: to.Ptr("dependency/type"),
		// 													},
		// 													{
		// 														LocalizedValue: to.Ptr("Dependency performance"),
		// 														Value: to.Ptr("dependency/performanceBucket"),
		// 													},
		// 													{
		// 														LocalizedValue: to.Ptr("Successful call"),
		// 														Value: to.Ptr("dependency/success"),
		// 													},
		// 													{
		// 														LocalizedValue: to.Ptr("Target of a dependency call"),
		// 														Value: to.Ptr("dependency/target"),
		// 													},
		// 													{
		// 														LocalizedValue: to.Ptr("Result code"),
		// 														Value: to.Ptr("dependency/resultCode"),
		// 													},
		// 													{
		// 														LocalizedValue: to.Ptr("Is traffic synthetic"),
		// 														Value: to.Ptr("operation/synthetic"),
		// 													},
		// 													{
		// 														LocalizedValue: to.Ptr("Cloud role instance"),
		// 														Value: to.Ptr("cloud/roleInstance"),
		// 													},
		// 													{
		// 														LocalizedValue: to.Ptr("Cloud role name"),
		// 														Value: to.Ptr("cloud/roleName"),
		// 												}},
		// 												DisplayDescription: to.Ptr("Duration of calls made by the application to external resources."),
		// 												ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/dependencies/duration"),
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
		// 												Namespace: to.Ptr("microsoft.insights/components"),
		// 												PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 												ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 												SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 													to.Ptr(armmonitor.AggregationTypeAverage),
		// 													to.Ptr(armmonitor.AggregationTypeMaximum),
		// 													to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 													Unit: to.Ptr(armmonitor.MetricUnitMilliSeconds),
		// 												},
		// 												{
		// 													Name: &armmonitor.LocalizableString{
		// 														LocalizedValue: to.Ptr("Dependency call failures"),
		// 														Value: to.Ptr("dependencies/failed"),
		// 													},
		// 													Category: to.Ptr("Failures"),
		// 													Dimensions: []*armmonitor.LocalizableString{
		// 														{
		// 															LocalizedValue: to.Ptr("Dependency type"),
		// 															Value: to.Ptr("dependency/type"),
		// 														},
		// 														{
		// 															LocalizedValue: to.Ptr("Dependency performance"),
		// 															Value: to.Ptr("dependency/performanceBucket"),
		// 														},
		// 														{
		// 															LocalizedValue: to.Ptr("Target of a dependency call"),
		// 															Value: to.Ptr("dependency/target"),
		// 														},
		// 														{
		// 															LocalizedValue: to.Ptr("Result code"),
		// 															Value: to.Ptr("dependency/resultCode"),
		// 														},
		// 														{
		// 															LocalizedValue: to.Ptr("Is traffic synthetic"),
		// 															Value: to.Ptr("operation/synthetic"),
		// 														},
		// 														{
		// 															LocalizedValue: to.Ptr("Cloud role instance"),
		// 															Value: to.Ptr("cloud/roleInstance"),
		// 														},
		// 														{
		// 															LocalizedValue: to.Ptr("Cloud role name"),
		// 															Value: to.Ptr("cloud/roleName"),
		// 													}},
		// 													DisplayDescription: to.Ptr("Count of failed dependency calls made by the application to external resources."),
		// 													ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/dependencies/failed"),
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
		// 													Namespace: to.Ptr("microsoft.insights/components"),
		// 													PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeCount),
		// 													ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 													SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 														to.Ptr(armmonitor.AggregationTypeCount)},
		// 														Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 													},
		// 													{
		// 														Name: &armmonitor.LocalizableString{
		// 															LocalizedValue: to.Ptr("Page views"),
		// 															Value: to.Ptr("pageViews/count"),
		// 														},
		// 														Category: to.Ptr("Usage"),
		// 														Dimensions: []*armmonitor.LocalizableString{
		// 															{
		// 																LocalizedValue: to.Ptr("Is traffic synthetic"),
		// 																Value: to.Ptr("operation/synthetic"),
		// 															},
		// 															{
		// 																LocalizedValue: to.Ptr("Cloud role name"),
		// 																Value: to.Ptr("cloud/roleName"),
		// 														}},
		// 														DisplayDescription: to.Ptr("Count of page views."),
		// 														ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/pageViews/count"),
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
		// 														Namespace: to.Ptr("microsoft.insights/components"),
		// 														PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeCount),
		// 														ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 														SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 															to.Ptr(armmonitor.AggregationTypeCount)},
		// 															Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 														},
		// 														{
		// 															Name: &armmonitor.LocalizableString{
		// 																LocalizedValue: to.Ptr("Page view load time"),
		// 																Value: to.Ptr("pageViews/duration"),
		// 															},
		// 															Category: to.Ptr("Usage"),
		// 															Dimensions: []*armmonitor.LocalizableString{
		// 																{
		// 																	LocalizedValue: to.Ptr("Is traffic synthetic"),
		// 																	Value: to.Ptr("operation/synthetic"),
		// 																},
		// 																{
		// 																	LocalizedValue: to.Ptr("Cloud role name"),
		// 																	Value: to.Ptr("cloud/roleName"),
		// 															}},
		// 															DisplayDescription: to.Ptr("Page view load time"),
		// 															ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/pageViews/duration"),
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
		// 															Namespace: to.Ptr("microsoft.insights/components"),
		// 															PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 															ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 															SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																to.Ptr(armmonitor.AggregationTypeAverage),
		// 																to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 																Unit: to.Ptr(armmonitor.MetricUnitMilliSeconds),
		// 															},
		// 															{
		// 																Name: &armmonitor.LocalizableString{
		// 																	LocalizedValue: to.Ptr("HTTP request execution time"),
		// 																	Value: to.Ptr("performanceCounters/requestExecutionTime"),
		// 																},
		// 																Category: to.Ptr("Performance counters"),
		// 																Dimensions: []*armmonitor.LocalizableString{
		// 																	{
		// 																		LocalizedValue: to.Ptr("Cloud role instance"),
		// 																		Value: to.Ptr("cloud/roleInstance"),
		// 																}},
		// 																DisplayDescription: to.Ptr("Execution time of the most recent request."),
		// 																ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/performanceCounters/requestExecutionTime"),
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
		// 																Namespace: to.Ptr("microsoft.insights/components"),
		// 																PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																	to.Ptr(armmonitor.AggregationTypeAverage),
		// 																	to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																	to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 																	Unit: to.Ptr(armmonitor.MetricUnitMilliSeconds),
		// 																},
		// 																{
		// 																	Name: &armmonitor.LocalizableString{
		// 																		LocalizedValue: to.Ptr("HTTP requests in application queue"),
		// 																		Value: to.Ptr("performanceCounters/requestsInQueue"),
		// 																	},
		// 																	Category: to.Ptr("Performance counters"),
		// 																	Dimensions: []*armmonitor.LocalizableString{
		// 																		{
		// 																			LocalizedValue: to.Ptr("Cloud role instance"),
		// 																			Value: to.Ptr("cloud/roleInstance"),
		// 																	}},
		// 																	DisplayDescription: to.Ptr("Length of the application request queue."),
		// 																	ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/performanceCounters/requestsInQueue"),
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
		// 																	Namespace: to.Ptr("microsoft.insights/components"),
		// 																	PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																	ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																	SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																		to.Ptr(armmonitor.AggregationTypeAverage),
		// 																		to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																		to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 																		Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																	},
		// 																	{
		// 																		Name: &armmonitor.LocalizableString{
		// 																			LocalizedValue: to.Ptr("HTTP request rate"),
		// 																			Value: to.Ptr("performanceCounters/requestsPerSecond"),
		// 																		},
		// 																		Category: to.Ptr("Performance counters"),
		// 																		Dimensions: []*armmonitor.LocalizableString{
		// 																			{
		// 																				LocalizedValue: to.Ptr("Cloud role instance"),
		// 																				Value: to.Ptr("cloud/roleInstance"),
		// 																		}},
		// 																		DisplayDescription: to.Ptr("Rate of all requests to the application per second from ASP.NET."),
		// 																		ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/performanceCounters/requestsPerSecond"),
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
		// 																		Namespace: to.Ptr("microsoft.insights/components"),
		// 																		PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																		ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																		SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																			to.Ptr(armmonitor.AggregationTypeAverage),
		// 																			to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																			to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 																			Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 																		},
		// 																		{
		// 																			Name: &armmonitor.LocalizableString{
		// 																				LocalizedValue: to.Ptr("Exception rate"),
		// 																				Value: to.Ptr("performanceCounters/exceptionsPerSecond"),
		// 																			},
		// 																			Category: to.Ptr("Performance counters"),
		// 																			Dimensions: []*armmonitor.LocalizableString{
		// 																				{
		// 																					LocalizedValue: to.Ptr("Cloud role instance"),
		// 																					Value: to.Ptr("cloud/roleInstance"),
		// 																			}},
		// 																			DisplayDescription: to.Ptr("Count of handled and unhandled exceptions reported to windows, including .NET exceptions and unmanaged exceptions that are converted into .NET exceptions."),
		// 																			ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/performanceCounters/exceptionsPerSecond"),
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
		// 																			Namespace: to.Ptr("microsoft.insights/components"),
		// 																			PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																			ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																			SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																				to.Ptr(armmonitor.AggregationTypeAverage),
		// 																				to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																				to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 																				Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 																			},
		// 																			{
		// 																				Name: &armmonitor.LocalizableString{
		// 																					LocalizedValue: to.Ptr("Process IO rate"),
		// 																					Value: to.Ptr("performanceCounters/processIOBytesPerSecond"),
		// 																				},
		// 																				Category: to.Ptr("Performance counters"),
		// 																				Dimensions: []*armmonitor.LocalizableString{
		// 																					{
		// 																						LocalizedValue: to.Ptr("Cloud role instance"),
		// 																						Value: to.Ptr("cloud/roleInstance"),
		// 																				}},
		// 																				DisplayDescription: to.Ptr("Total bytes per second read and written to files, network and devices."),
		// 																				ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/performanceCounters/processIOBytesPerSecond"),
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
		// 																				Namespace: to.Ptr("microsoft.insights/components"),
		// 																				PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																				ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																				SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																					to.Ptr(armmonitor.AggregationTypeAverage),
		// 																					to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																					to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 																					Unit: to.Ptr(armmonitor.MetricUnitBytesPerSecond),
		// 																				},
		// 																				{
		// 																					Name: &armmonitor.LocalizableString{
		// 																						LocalizedValue: to.Ptr("Process CPU"),
		// 																						Value: to.Ptr("performanceCounters/processCpuPercentage"),
		// 																					},
		// 																					Category: to.Ptr("Performance counters"),
		// 																					Dimensions: []*armmonitor.LocalizableString{
		// 																						{
		// 																							LocalizedValue: to.Ptr("Cloud role instance"),
		// 																							Value: to.Ptr("cloud/roleInstance"),
		// 																					}},
		// 																					DisplayDescription: to.Ptr("The percentage of elapsed time that all process threads used the processor to execute instructions. This can vary between 0 to 100. This metric indicates the performance of w3wp process alone."),
		// 																					ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/performanceCounters/processCpuPercentage"),
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
		// 																					Namespace: to.Ptr("microsoft.insights/components"),
		// 																					PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																					ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																					SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																						to.Ptr(armmonitor.AggregationTypeAverage),
		// 																						to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																						to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 																						Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																					},
		// 																					{
		// 																						Name: &armmonitor.LocalizableString{
		// 																							LocalizedValue: to.Ptr("Processor time"),
		// 																							Value: to.Ptr("performanceCounters/processorCpuPercentage"),
		// 																						},
		// 																						Category: to.Ptr("Performance counters"),
		// 																						Dimensions: []*armmonitor.LocalizableString{
		// 																							{
		// 																								LocalizedValue: to.Ptr("Cloud role instance"),
		// 																								Value: to.Ptr("cloud/roleInstance"),
		// 																						}},
		// 																						DisplayDescription: to.Ptr("The percentage of time that the processor spends in non-idle threads."),
		// 																						ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/performanceCounters/processorCpuPercentage"),
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
		// 																						Namespace: to.Ptr("microsoft.insights/components"),
		// 																						PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																						ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																						SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																							to.Ptr(armmonitor.AggregationTypeAverage),
		// 																							to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																							to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 																							Unit: to.Ptr(armmonitor.MetricUnitPercent),
		// 																						},
		// 																						{
		// 																							Name: &armmonitor.LocalizableString{
		// 																								LocalizedValue: to.Ptr("Available memory"),
		// 																								Value: to.Ptr("performanceCounters/memoryAvailableBytes"),
		// 																							},
		// 																							Category: to.Ptr("Performance counters"),
		// 																							Dimensions: []*armmonitor.LocalizableString{
		// 																								{
		// 																									LocalizedValue: to.Ptr("Cloud role instance"),
		// 																									Value: to.Ptr("cloud/roleInstance"),
		// 																							}},
		// 																							DisplayDescription: to.Ptr("Physical memory immediately available for allocation to a process or for system use."),
		// 																							ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/performanceCounters/memoryAvailableBytes"),
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
		// 																							Namespace: to.Ptr("microsoft.insights/components"),
		// 																							PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																							ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																							SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																								to.Ptr(armmonitor.AggregationTypeAverage),
		// 																								to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																								to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 																								Unit: to.Ptr(armmonitor.MetricUnitBytes),
		// 																							},
		// 																							{
		// 																								Name: &armmonitor.LocalizableString{
		// 																									LocalizedValue: to.Ptr("Process private bytes"),
		// 																									Value: to.Ptr("performanceCounters/processPrivateBytes"),
		// 																								},
		// 																								Category: to.Ptr("Performance counters"),
		// 																								Dimensions: []*armmonitor.LocalizableString{
		// 																									{
		// 																										LocalizedValue: to.Ptr("Cloud role instance"),
		// 																										Value: to.Ptr("cloud/roleInstance"),
		// 																								}},
		// 																								DisplayDescription: to.Ptr("Memory exclusively assigned to the monitored application's processes."),
		// 																								ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/performanceCounters/processPrivateBytes"),
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
		// 																								Namespace: to.Ptr("microsoft.insights/components"),
		// 																								PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																								ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																								SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																									to.Ptr(armmonitor.AggregationTypeAverage),
		// 																									to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																									to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 																									Unit: to.Ptr(armmonitor.MetricUnitBytes),
		// 																								},
		// 																								{
		// 																									Name: &armmonitor.LocalizableString{
		// 																										LocalizedValue: to.Ptr("Server response time"),
		// 																										Value: to.Ptr("requests/duration"),
		// 																									},
		// 																									Category: to.Ptr("Server"),
		// 																									Dimensions: []*armmonitor.LocalizableString{
		// 																										{
		// 																											LocalizedValue: to.Ptr("Request performance"),
		// 																											Value: to.Ptr("request/performanceBucket"),
		// 																										},
		// 																										{
		// 																											LocalizedValue: to.Ptr("Result code"),
		// 																											Value: to.Ptr("request/resultCode"),
		// 																										},
		// 																										{
		// 																											LocalizedValue: to.Ptr("Is traffic synthetic"),
		// 																											Value: to.Ptr("operation/synthetic"),
		// 																										},
		// 																										{
		// 																											LocalizedValue: to.Ptr("Cloud role instance"),
		// 																											Value: to.Ptr("cloud/roleInstance"),
		// 																										},
		// 																										{
		// 																											LocalizedValue: to.Ptr("Successful request"),
		// 																											Value: to.Ptr("request/success"),
		// 																										},
		// 																										{
		// 																											LocalizedValue: to.Ptr("Cloud role name"),
		// 																											Value: to.Ptr("cloud/roleName"),
		// 																									}},
		// 																									DisplayDescription: to.Ptr("Time between receiving an HTTP request and finishing sending the response."),
		// 																									ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/requests/duration"),
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
		// 																									Namespace: to.Ptr("microsoft.insights/components"),
		// 																									PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																									ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																									SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																										to.Ptr(armmonitor.AggregationTypeAverage),
		// 																										to.Ptr(armmonitor.AggregationTypeMaximum),
		// 																										to.Ptr(armmonitor.AggregationTypeMinimum)},
		// 																										Unit: to.Ptr(armmonitor.MetricUnitMilliSeconds),
		// 																									},
		// 																									{
		// 																										Name: &armmonitor.LocalizableString{
		// 																											LocalizedValue: to.Ptr("Server requests"),
		// 																											Value: to.Ptr("requests/count"),
		// 																										},
		// 																										Category: to.Ptr("Server"),
		// 																										Dimensions: []*armmonitor.LocalizableString{
		// 																											{
		// 																												LocalizedValue: to.Ptr("Request performance"),
		// 																												Value: to.Ptr("request/performanceBucket"),
		// 																											},
		// 																											{
		// 																												LocalizedValue: to.Ptr("Result code"),
		// 																												Value: to.Ptr("request/resultCode"),
		// 																											},
		// 																											{
		// 																												LocalizedValue: to.Ptr("Is traffic synthetic"),
		// 																												Value: to.Ptr("operation/synthetic"),
		// 																											},
		// 																											{
		// 																												LocalizedValue: to.Ptr("Cloud role instance"),
		// 																												Value: to.Ptr("cloud/roleInstance"),
		// 																											},
		// 																											{
		// 																												LocalizedValue: to.Ptr("Successful request"),
		// 																												Value: to.Ptr("request/success"),
		// 																											},
		// 																											{
		// 																												LocalizedValue: to.Ptr("Cloud role name"),
		// 																												Value: to.Ptr("cloud/roleName"),
		// 																										}},
		// 																										DisplayDescription: to.Ptr("Count of HTTP requests completed."),
		// 																										ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/requests/count"),
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
		// 																										Namespace: to.Ptr("microsoft.insights/components"),
		// 																										PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeCount),
		// 																										ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																										SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																											to.Ptr(armmonitor.AggregationTypeCount)},
		// 																											Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																										},
		// 																										{
		// 																											Name: &armmonitor.LocalizableString{
		// 																												LocalizedValue: to.Ptr("Failed requests"),
		// 																												Value: to.Ptr("requests/failed"),
		// 																											},
		// 																											Category: to.Ptr("Failures"),
		// 																											Dimensions: []*armmonitor.LocalizableString{
		// 																												{
		// 																													LocalizedValue: to.Ptr("Request performance"),
		// 																													Value: to.Ptr("request/performanceBucket"),
		// 																												},
		// 																												{
		// 																													LocalizedValue: to.Ptr("Result code"),
		// 																													Value: to.Ptr("request/resultCode"),
		// 																												},
		// 																												{
		// 																													LocalizedValue: to.Ptr("Is traffic synthetic"),
		// 																													Value: to.Ptr("operation/synthetic"),
		// 																												},
		// 																												{
		// 																													LocalizedValue: to.Ptr("Cloud role instance"),
		// 																													Value: to.Ptr("cloud/roleInstance"),
		// 																												},
		// 																												{
		// 																													LocalizedValue: to.Ptr("Cloud role name"),
		// 																													Value: to.Ptr("cloud/roleName"),
		// 																											}},
		// 																											DisplayDescription: to.Ptr("Count of HTTP requests marked as failed. In most cases these are requests with a response code >= 400 and not equal to 401."),
		// 																											ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/requests/failed"),
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
		// 																											Namespace: to.Ptr("microsoft.insights/components"),
		// 																											PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeCount),
		// 																											ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																											SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																												to.Ptr(armmonitor.AggregationTypeCount)},
		// 																												Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																											},
		// 																											{
		// 																												Name: &armmonitor.LocalizableString{
		// 																													LocalizedValue: to.Ptr("Server request rate"),
		// 																													Value: to.Ptr("requests/rate"),
		// 																												},
		// 																												Category: to.Ptr("Server"),
		// 																												Dimensions: []*armmonitor.LocalizableString{
		// 																													{
		// 																														LocalizedValue: to.Ptr("Request performance"),
		// 																														Value: to.Ptr("request/performanceBucket"),
		// 																													},
		// 																													{
		// 																														LocalizedValue: to.Ptr("Result code"),
		// 																														Value: to.Ptr("request/resultCode"),
		// 																													},
		// 																													{
		// 																														LocalizedValue: to.Ptr("Is traffic synthetic"),
		// 																														Value: to.Ptr("operation/synthetic"),
		// 																													},
		// 																													{
		// 																														LocalizedValue: to.Ptr("Cloud role instance"),
		// 																														Value: to.Ptr("cloud/roleInstance"),
		// 																													},
		// 																													{
		// 																														LocalizedValue: to.Ptr("Successful request"),
		// 																														Value: to.Ptr("request/success"),
		// 																													},
		// 																													{
		// 																														LocalizedValue: to.Ptr("Cloud role name"),
		// 																														Value: to.Ptr("cloud/roleName"),
		// 																												}},
		// 																												DisplayDescription: to.Ptr("Rate of server requests per second"),
		// 																												ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/requests/rate"),
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
		// 																												Namespace: to.Ptr("microsoft.insights/components"),
		// 																												PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeAverage),
		// 																												ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																												SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																													to.Ptr(armmonitor.AggregationTypeAverage)},
		// 																													Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
		// 																												},
		// 																												{
		// 																													Name: &armmonitor.LocalizableString{
		// 																														LocalizedValue: to.Ptr("Exceptions"),
		// 																														Value: to.Ptr("exceptions/count"),
		// 																													},
		// 																													Category: to.Ptr("Failures"),
		// 																													Dimensions: []*armmonitor.LocalizableString{
		// 																														{
		// 																															LocalizedValue: to.Ptr("Cloud role name"),
		// 																															Value: to.Ptr("cloud/roleName"),
		// 																														},
		// 																														{
		// 																															LocalizedValue: to.Ptr("Cloud role instance"),
		// 																															Value: to.Ptr("cloud/roleInstance"),
		// 																														},
		// 																														{
		// 																															LocalizedValue: to.Ptr("Device type"),
		// 																															Value: to.Ptr("client/type"),
		// 																													}},
		// 																													DisplayDescription: to.Ptr("Combined count of all uncaught exceptions."),
		// 																													ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/exceptions/count"),
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
		// 																													Namespace: to.Ptr("microsoft.insights/components"),
		// 																													PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeCount),
		// 																													ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																													SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																														to.Ptr(armmonitor.AggregationTypeCount)},
		// 																														Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																													},
		// 																													{
		// 																														Name: &armmonitor.LocalizableString{
		// 																															LocalizedValue: to.Ptr("Browser exceptions"),
		// 																															Value: to.Ptr("exceptions/browser"),
		// 																														},
		// 																														Category: to.Ptr("Failures"),
		// 																														Dimensions: []*armmonitor.LocalizableString{
		// 																															{
		// 																																LocalizedValue: to.Ptr("Cloud role name"),
		// 																																Value: to.Ptr("cloud/roleName"),
		// 																														}},
		// 																														DisplayDescription: to.Ptr("Count of uncaught exceptions thrown in the browser."),
		// 																														ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/exceptions/browser"),
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
		// 																														Namespace: to.Ptr("microsoft.insights/components"),
		// 																														PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeCount),
		// 																														ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																														SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																															to.Ptr(armmonitor.AggregationTypeCount)},
		// 																															Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																														},
		// 																														{
		// 																															Name: &armmonitor.LocalizableString{
		// 																																LocalizedValue: to.Ptr("Server exceptions"),
		// 																																Value: to.Ptr("exceptions/server"),
		// 																															},
		// 																															Category: to.Ptr("Failures"),
		// 																															Dimensions: []*armmonitor.LocalizableString{
		// 																																{
		// 																																	LocalizedValue: to.Ptr("Cloud role name"),
		// 																																	Value: to.Ptr("cloud/roleName"),
		// 																																},
		// 																																{
		// 																																	LocalizedValue: to.Ptr("Cloud role instance"),
		// 																																	Value: to.Ptr("cloud/roleInstance"),
		// 																															}},
		// 																															DisplayDescription: to.Ptr("Count of uncaught exceptions thrown in the server application."),
		// 																															ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/exceptions/server"),
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
		// 																															Namespace: to.Ptr("microsoft.insights/components"),
		// 																															PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeCount),
		// 																															ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																															SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																																to.Ptr(armmonitor.AggregationTypeCount)},
		// 																																Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																															},
		// 																															{
		// 																																Name: &armmonitor.LocalizableString{
		// 																																	LocalizedValue: to.Ptr("Traces"),
		// 																																	Value: to.Ptr("traces/count"),
		// 																																},
		// 																																Category: to.Ptr("Usage"),
		// 																																Dimensions: []*armmonitor.LocalizableString{
		// 																																	{
		// 																																		LocalizedValue: to.Ptr("Severity level"),
		// 																																		Value: to.Ptr("trace/severityLevel"),
		// 																																	},
		// 																																	{
		// 																																		LocalizedValue: to.Ptr("Is traffic synthetic"),
		// 																																		Value: to.Ptr("operation/synthetic"),
		// 																																	},
		// 																																	{
		// 																																		LocalizedValue: to.Ptr("Cloud role name"),
		// 																																		Value: to.Ptr("cloud/roleName"),
		// 																																	},
		// 																																	{
		// 																																		LocalizedValue: to.Ptr("Cloud role instance"),
		// 																																		Value: to.Ptr("cloud/roleInstance"),
		// 																																}},
		// 																																DisplayDescription: to.Ptr("Trace document count"),
		// 																																ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions/traces/count"),
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
		// 																																Namespace: to.Ptr("microsoft.insights/components"),
		// 																																PrimaryAggregationType: to.Ptr(armmonitor.AggregationTypeCount),
		// 																																ResourceID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill"),
		// 																																SupportedAggregationTypes: []*armmonitor.AggregationType{
		// 																																	to.Ptr(armmonitor.AggregationTypeCount)},
		// 																																	Unit: to.Ptr(armmonitor.MetricUnitCount),
		// 																															}},
		// 																														}
	}
}
