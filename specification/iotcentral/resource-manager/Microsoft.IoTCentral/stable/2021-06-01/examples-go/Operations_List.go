package armiotcentral_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armiotcentral.OperationListResult{
		// 	Value: []*armiotcentral.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTCentral/IoTApps/read"),
		// 			Display: &armiotcentral.OperationDisplay{
		// 				Description: to.Ptr("Gets a single IoT Central Application"),
		// 				Operation: to.Ptr("Get IoT Central Application"),
		// 				Provider: to.Ptr("Azure IoT Central"),
		// 				Resource: to.Ptr("IoTApps"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTCentral/IoTApps/write"),
		// 			Display: &armiotcentral.OperationDisplay{
		// 				Description: to.Ptr("Creates or Updates an IoT Central Applications"),
		// 				Operation: to.Ptr("Create or Update IoT Central Application"),
		// 				Provider: to.Ptr("Azure IoT Central"),
		// 				Resource: to.Ptr("IoTApps"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTCentral/IoTApps/delete"),
		// 			Display: &armiotcentral.OperationDisplay{
		// 				Description: to.Ptr("Deletes an IoT Central Applications"),
		// 				Operation: to.Ptr("Delete IoT Central Application"),
		// 				Provider: to.Ptr("Azure IoT Central"),
		// 				Resource: to.Ptr("IoTApps"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTCentral/checkNameAvailability/action"),
		// 			Display: &armiotcentral.OperationDisplay{
		// 				Description: to.Ptr("Checks if an IoT Central Application name is available"),
		// 				Operation: to.Ptr("Check resource name availability"),
		// 				Provider: to.Ptr("Azure IoT Central"),
		// 				Resource: to.Ptr(""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTCentral/checkSubdomainAvailability/action"),
		// 			Display: &armiotcentral.OperationDisplay{
		// 				Description: to.Ptr("Checks if an IoT Central Application subdomain is available"),
		// 				Operation: to.Ptr("Check resource subdomain availability"),
		// 				Provider: to.Ptr("Azure IoT Central"),
		// 				Resource: to.Ptr(""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTCentral/operations/read"),
		// 			Display: &armiotcentral.OperationDisplay{
		// 				Description: to.Ptr("Gets all the available operations on IoT Central Applications"),
		// 				Operation: to.Ptr("Get all the available operations"),
		// 				Provider: to.Ptr("Azure IoT Central"),
		// 				Resource: to.Ptr(""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTCentral/appTemplates/action"),
		// 			Display: &armiotcentral.OperationDisplay{
		// 				Description: to.Ptr("Gets all the available application templates on Azure IoT Central"),
		// 				Operation: to.Ptr("Get all available application templates"),
		// 				Provider: to.Ptr("Azure IoT Central"),
		// 				Resource: to.Ptr("Azure IoT Central Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTCentral/register/action"),
		// 			Display: &armiotcentral.OperationDisplay{
		// 				Description: to.Ptr("Register the subscription for Azure IoT Central resource provider"),
		// 				Operation: to.Ptr("Register Azure IoT Central resource provider"),
		// 				Provider: to.Ptr("Azure IoT Central"),
		// 				Resource: to.Ptr("Azure IoT Central Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTCentral/IoTApps/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armiotcentral.OperationDisplay{
		// 				Description: to.Ptr("Gets all the available Metrics definitions on Azure IoT Central"),
		// 				Operation: to.Ptr("Get all available Metrics definitions"),
		// 				Provider: to.Ptr("Azure IoT Central"),
		// 				Resource: to.Ptr("IoTApps"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 			Properties: map[string]any{
		// 				"serviceSpecification":map[string]any{
		// 					"metricSpecifications":[]any{
		// 						map[string]any{
		// 							"name": "connectedDeviceCount",
		// 							"aggregationType": "Average",
		// 							"displayDescription": "Number of devices connected to IoT Central",
		// 							"displayName": "Total Connected Devices",
		// 							"lockAggregationType": "Total",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 							},
		// 							"supportedTimeGrainTypes":[]any{
		// 								"PT1M",
		// 								"PT5M",
		// 								"PT15M",
		// 								"PT30M",
		// 								"PT1H",
		// 								"PT6H",
		// 								"PT12H",
		// 								"P1D",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "c2d.property.read.success",
		// 							"aggregationType": "Total",
		// 							"displayDescription": "The count of all successful property reads initiated from IoT Central",
		// 							"displayName": "Successful Device Property Reads from IoT Central",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"supportedTimeGrainTypes":[]any{
		// 								"PT1M",
		// 								"PT5M",
		// 								"PT15M",
		// 								"PT30M",
		// 								"PT1H",
		// 								"PT6H",
		// 								"PT12H",
		// 								"P1D",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "c2d.property.read.failure",
		// 							"aggregationType": "Total",
		// 							"displayDescription": "The count of all failed property reads initiated from IoT Central",
		// 							"displayName": "Failed Device Property Reads from IoT Central",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"supportedTimeGrainTypes":[]any{
		// 								"PT1M",
		// 								"PT5M",
		// 								"PT15M",
		// 								"PT30M",
		// 								"PT1H",
		// 								"PT6H",
		// 								"PT12H",
		// 								"P1D",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "d2c.property.read.success",
		// 							"aggregationType": "Total",
		// 							"displayDescription": "The count of all successful property reads initiated from devices",
		// 							"displayName": "Successful Device Property Reads from Devices",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"supportedTimeGrainTypes":[]any{
		// 								"PT1M",
		// 								"PT5M",
		// 								"PT15M",
		// 								"PT30M",
		// 								"PT1H",
		// 								"PT6H",
		// 								"PT12H",
		// 								"P1D",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "d2c.property.read.failure",
		// 							"aggregationType": "Total",
		// 							"displayDescription": "The count of all failed property reads initiated from devices",
		// 							"displayName": "Failed Device Property Reads from Devices",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"supportedTimeGrainTypes":[]any{
		// 								"PT1M",
		// 								"PT5M",
		// 								"PT15M",
		// 								"PT30M",
		// 								"PT1H",
		// 								"PT6H",
		// 								"PT12H",
		// 								"P1D",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "c2d.property.update.success",
		// 							"aggregationType": "Total",
		// 							"displayDescription": "The count of all successful property updates initiated from IoT Central",
		// 							"displayName": "Successful Device Property Updates from IoT Central",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"supportedTimeGrainTypes":[]any{
		// 								"PT1M",
		// 								"PT5M",
		// 								"PT15M",
		// 								"PT30M",
		// 								"PT1H",
		// 								"PT6H",
		// 								"PT12H",
		// 								"P1D",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "c2d.property.update.failure",
		// 							"aggregationType": "Total",
		// 							"displayDescription": "The count of all failed property updates initiated from IoT Central",
		// 							"displayName": "Failed Device Property Updates from IoT Central",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"supportedTimeGrainTypes":[]any{
		// 								"PT1M",
		// 								"PT5M",
		// 								"PT15M",
		// 								"PT30M",
		// 								"PT1H",
		// 								"PT6H",
		// 								"PT12H",
		// 								"P1D",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "d2c.property.update.success",
		// 							"aggregationType": "Total",
		// 							"displayDescription": "The count of all successful property updates initiated from devices",
		// 							"displayName": "Successful Device Property Updates from Devices",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"supportedTimeGrainTypes":[]any{
		// 								"PT1M",
		// 								"PT5M",
		// 								"PT15M",
		// 								"PT30M",
		// 								"PT1H",
		// 								"PT6H",
		// 								"PT12H",
		// 								"P1D",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "d2c.property.update.failure",
		// 							"aggregationType": "Total",
		// 							"displayDescription": "The count of all failed property updates initiated from devices",
		// 							"displayName": "Failed Device Property Updates from Devices",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"supportedTimeGrainTypes":[]any{
		// 								"PT1M",
		// 								"PT5M",
		// 								"PT15M",
		// 								"PT30M",
		// 								"PT1H",
		// 								"PT6H",
		// 								"PT12H",
		// 								"P1D",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTCentral/IoTApps/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armiotcentral.OperationDisplay{
		// 				Description: to.Ptr("Gets the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Read diagnostic setting"),
		// 				Provider: to.Ptr("Azure IoT Central"),
		// 				Resource: to.Ptr("IoTApps"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTCentral/IoTApps/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armiotcentral.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Write diagnostic setting"),
		// 				Provider: to.Ptr("Azure IoT Central"),
		// 				Resource: to.Ptr("IoTApps"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 	}},
		// }
	}
}
