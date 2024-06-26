package armtimeseriesinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/Operation_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtimeseriesinsights.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armtimeseriesinsights.OperationListResult{
		// 	Value: []*armtimeseriesinsights.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/register/action"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Registers the subscription for the Time Series Insights resource provider and enables the creation of Time Series Insights environments."),
		// 				Operation: to.Ptr("Registers the Time Series Insights Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Time Series Insights Resource Provider"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets the available metrics for environments"),
		// 				Operation: to.Ptr("Read environments metric definitions"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("The metrics definition of environments"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 			OperationProperties: &armtimeseriesinsights.OperationProperties{
		// 				ServiceSpecification: &armtimeseriesinsights.ServiceSpecification{
		// 					MetricSpecifications: []*armtimeseriesinsights.MetricSpecification{
		// 						{
		// 							Name: to.Ptr("IngressReceivedMessages"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Count of messages read from all Event hub or IoT hub event sources"),
		// 							DisplayName: to.Ptr("Ingress Received Messages"),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("IngressReceivedInvalidMessages"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Count of invalid messages read from all Event hub or IoT hub event sources"),
		// 							DisplayName: to.Ptr("Ingress Received Invalid Messages"),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("IngressReceivedBytes"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Count of bytes read from all event sources"),
		// 							DisplayName: to.Ptr("Ingress Received Bytes"),
		// 							Unit: to.Ptr("Bytes"),
		// 						},
		// 						{
		// 							Name: to.Ptr("IngressStoredBytes"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Total size of events successfully processed and available for query"),
		// 							DisplayName: to.Ptr("Ingress Stored Bytes"),
		// 							Unit: to.Ptr("Bytes"),
		// 						},
		// 						{
		// 							Name: to.Ptr("IngressStoredEvents"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Count of flattened events successfully processed and available for query"),
		// 							DisplayName: to.Ptr("Ingress Stored Events"),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("IngressReceivedMessagesTimeLag"),
		// 							AggregationType: to.Ptr("Maximum"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Difference between the time that the message is enqueued in the event source and the time it is processed in Ingress"),
		// 							DisplayName: to.Ptr("Ingress Received Messages Time Lag"),
		// 							Unit: to.Ptr("Seconds"),
		// 						},
		// 						{
		// 							Name: to.Ptr("IngressReceivedMessagesCountLag"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Difference between the sequence number of last enqueued message in the event source partition and sequence number of messages being processed in Ingress"),
		// 							DisplayName: to.Ptr("Ingress Received Messages Count Lag"),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("WarmStorageMaxProperties"),
		// 							AggregationType: to.Ptr("Maximum"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Maximum number of properties used allowed by the environment for S1/S2 SKU and maximum number of properties allowed by Warm Store for PAYG SKU"),
		// 							DisplayName: to.Ptr("Warm Storage Max Properties"),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("WarmStorageUsedProperties"),
		// 							AggregationType: to.Ptr("Maximum"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Number of properties used by the environment for S1/S2 SKU and number of properties used by Warm Store for PAYG SKU"),
		// 							DisplayName: to.Ptr("Warm Storage Used Properties "),
		// 							Unit: to.Ptr("Count"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Read diagnostic setting."),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("environments"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Write diagnostic setting."),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("environments"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/providers/Microsoft.Insights/logDefinitions/read"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets the available logs for environments"),
		// 				Operation: to.Ptr("Read environments log definitions"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("environments"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 			OperationProperties: &armtimeseriesinsights.OperationProperties{
		// 				ServiceSpecification: &armtimeseriesinsights.ServiceSpecification{
		// 					LogSpecifications: []*armtimeseriesinsights.LogSpecification{
		// 						{
		// 							Name: to.Ptr("Ingress"),
		// 							DisplayName: to.Ptr("Ingress"),
		// 						},
		// 						{
		// 							Name: to.Ptr("Management"),
		// 							DisplayName: to.Ptr("Management"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/eventsources/providers/Microsoft.Insights/logDefinitions/read"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets the available logs for the event source"),
		// 				Operation: to.Ptr("Read event source log definitions"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Event Source"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 			OperationProperties: &armtimeseriesinsights.OperationProperties{
		// 				ServiceSpecification: &armtimeseriesinsights.ServiceSpecification{
		// 					LogSpecifications: []*armtimeseriesinsights.LogSpecification{
		// 						{
		// 							Name: to.Ptr("Ingress"),
		// 							DisplayName: to.Ptr("Ingress"),
		// 						},
		// 						{
		// 							Name: to.Ptr("Management"),
		// 							DisplayName: to.Ptr("Management"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/eventsources/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets the available metrics for eventsources"),
		// 				Operation: to.Ptr("Read eventsources metric definitions"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("The metrics definition of environments/eventsources"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 			OperationProperties: &armtimeseriesinsights.OperationProperties{
		// 				ServiceSpecification: &armtimeseriesinsights.ServiceSpecification{
		// 					MetricSpecifications: []*armtimeseriesinsights.MetricSpecification{
		// 						{
		// 							Name: to.Ptr("IngressReceivedMessages"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Count of messages read from the event source"),
		// 							DisplayName: to.Ptr("Ingress Received Messages"),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("IngressReceivedInvalidMessages"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Count of invalid messages read from the event source"),
		// 							DisplayName: to.Ptr("Ingress Received Invalid Messages"),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("IngressReceivedBytes"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Count of bytes read from the event source"),
		// 							DisplayName: to.Ptr("Ingress Received Bytes"),
		// 							Unit: to.Ptr("Bytes"),
		// 						},
		// 						{
		// 							Name: to.Ptr("IngressStoredBytes"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Total size of events successfully processed and available for query"),
		// 							DisplayName: to.Ptr("Ingress Stored Bytes"),
		// 							Unit: to.Ptr("Bytes"),
		// 						},
		// 						{
		// 							Name: to.Ptr("IngressStoredEvents"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Count of flattened events successfully processed and available for query"),
		// 							DisplayName: to.Ptr("Ingress Stored Events"),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("IngressReceivedMessagesTimeLag"),
		// 							AggregationType: to.Ptr("Maximum"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Difference between the time that the message is enqueued in the event source and the time it is processed in Ingress"),
		// 							DisplayName: to.Ptr("Ingress Received Messages Time Lag"),
		// 							Unit: to.Ptr("Seconds"),
		// 						},
		// 						{
		// 							Name: to.Ptr("IngressReceivedMessagesCountLag"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Difference between the sequence number of last enqueued message in the event source partition and sequence number of messages being processed in Ingress"),
		// 							DisplayName: to.Ptr("Ingress Received Messages Count Lag"),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("WarmStorageMaxProperties"),
		// 							AggregationType: to.Ptr("Maximum"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Maximum number of properties used allowed by the environment for S1/S2 SKU and maximum number of properties allowed by Warm Store for PAYG SKU"),
		// 							DisplayName: to.Ptr("Warm Storage Max Properties"),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("WarmStorageUsedProperties"),
		// 							AggregationType: to.Ptr("Maximum"),
		// 							Availabilities: []*armtimeseriesinsights.MetricAvailability{
		// 								{
		// 									BlobDuration: to.Ptr("PT1H"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Number of properties used by the environment for S1/S2 SKU and number of properties used by Warm Store for PAYG SKU"),
		// 							DisplayName: to.Ptr("Warm Storage Used Properties "),
		// 							Unit: to.Ptr("Count"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/eventsources/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Read diagnostic setting."),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("environments/eventsources"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/eventsources/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Write diagnostic setting."),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("environments/eventsources"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/read"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Get the properties of an environment."),
		// 				Operation: to.Ptr("Read Environment"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Environment"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/write"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Creates a new environment, or updates an existing environment."),
		// 				Operation: to.Ptr("Create or Update Environment"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Environment"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/delete"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes the environment."),
		// 				Operation: to.Ptr("Delete Environment"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Environment"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/status/read"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Get the status of the environment, state of its associated operations like ingress."),
		// 				Operation: to.Ptr("Read Environment status"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Environment"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/eventsources/read"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Get the properties of an event source."),
		// 				Operation: to.Ptr("Read Event Source"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Event Source"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/eventsources/write"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Creates a new event source for an environment, or updates an existing event source."),
		// 				Operation: to.Ptr("Create or Update Event Source"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Event Source"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/eventsources/delete"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes the event source."),
		// 				Operation: to.Ptr("Delete Event Source"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Event Source"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/referencedatasets/read"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Get the properties of a reference data set."),
		// 				Operation: to.Ptr("Read Reference Data Set"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Reference Data Set"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/referencedatasets/write"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Creates a new reference data set for an environment, or updates an existing reference data set."),
		// 				Operation: to.Ptr("Create or Update Reference Data Set"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Reference Data Set"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/referencedatasets/delete"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes the reference data set."),
		// 				Operation: to.Ptr("Delete Reference Data Set"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Reference Data Set"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/accesspolicies/read"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Get the properties of an access policy."),
		// 				Operation: to.Ptr("Read Access Policy"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Access Policy"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/accesspolicies/write"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Creates a new access policy for an environment, or updates an existing access policy."),
		// 				Operation: to.Ptr("Create or Update Access Policy"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Access Policy"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/accesspolicies/delete"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes the access policy."),
		// 				Operation: to.Ptr("Delete Access Policy"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Access Policy"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/privateEndpointConnectionProxies/read"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Get the properties of a private endpoint connection proxy."),
		// 				Operation: to.Ptr("Read private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Private endpoint connection proxy"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/privateEndpointConnectionProxies/write"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Creates a new private endpoint connection proxy for an environment, or updates an existing connection proxy."),
		// 				Operation: to.Ptr("Create or Update private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Private endpoint connection proxy"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/privateEndpointConnectionProxies/delete"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes the private endpoint connection proxy."),
		// 				Operation: to.Ptr("Delete the private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Private endpoint connection proxy"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/privateEndpointConnectionProxies/validate/action"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Validate the private endpoint connection proxy object before creation."),
		// 				Operation: to.Ptr("Validate the private endpoint connection proxy."),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Private endpoint connection proxy"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/privateEndpointConnectionProxies/operationresults/read"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Validate the private endpoint connection proxy operation status."),
		// 				Operation: to.Ptr("Get private endpoint connection proxy operation status."),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Private endpoint connection proxy"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/privateendpointConnections/read"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Get the properties of a private endpoint connection."),
		// 				Operation: to.Ptr("Read private endpoint connection."),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Private endpoint connection."),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/privateendpointConnections/write"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Creates a new private endpoint connection for an environment, or updates an existing connection."),
		// 				Operation: to.Ptr("Create or Update private endpoint connection."),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Private endpoint connection."),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TimeSeriesInsights/environments/privateendpointConnections/delete"),
		// 			Display: &armtimeseriesinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes the private endpoint connection."),
		// 				Operation: to.Ptr("Delete the private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft Time Series Insights"),
		// 				Resource: to.Ptr("Private endpoint connection"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}
