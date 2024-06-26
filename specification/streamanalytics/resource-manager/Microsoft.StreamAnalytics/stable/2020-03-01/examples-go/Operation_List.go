package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Operation_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armstreamanalytics.OperationListResult{
		// 	Value: []*armstreamanalytics.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/locations/quotas/Read"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Read Stream Analytics Subscription Quota"),
		// 				Operation: to.Ptr("Read Stream Analytics Subscription Quota"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Subscription Quota"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/operations/Read"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Read Stream Analytics Operations"),
		// 				Operation: to.Ptr("Read Stream Analytics Operations"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Operations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/Register/action"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Register subscription with Stream Analytics Resource Provider"),
		// 				Operation: to.Ptr("Register subscription with Stream Analytics Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/Delete"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Delete Stream Analytics Job"),
		// 				Operation: to.Ptr("Delete Stream Analytics Job"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/functions/Delete"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Delete Stream Analytics Job Function"),
		// 				Operation: to.Ptr("Delete Stream Analytics Job Function"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Function"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/functions/operationresults/Read"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Read operation results for Stream Analytics Job Function"),
		// 				Operation: to.Ptr("Read operation results for Stream Analytics Job Function"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Function"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/functions/Read"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Read Stream Analytics Job Function"),
		// 				Operation: to.Ptr("Read Stream Analytics Job Function"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Function"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/functions/RetrieveDefaultDefinition/action"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Retrieve Default Definition of a Stream Analytics Job Function"),
		// 				Operation: to.Ptr("Retrieve Default Definition of a Stream Analytics Job Function"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Function"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/functions/Test/action"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Test Stream Analytics Job Function"),
		// 				Operation: to.Ptr("Test Stream Analytics Job Function"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Function"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/functions/Write"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Write Stream Analytics Job Function"),
		// 				Operation: to.Ptr("Write Stream Analytics Job Function"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Function"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs/Delete"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Delete Stream Analytics Job Input"),
		// 				Operation: to.Ptr("Delete Stream Analytics Job Input"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Input"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs/operationresults/Read"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Read operation results for Stream Analytics Job Input"),
		// 				Operation: to.Ptr("Read operation results for Stream Analytics Job Input"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Input"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs/Read"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Read Stream Analytics Job Input"),
		// 				Operation: to.Ptr("Read Stream Analytics Job Input"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Input"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs/Sample/action"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Sample Stream Analytics Job Input"),
		// 				Operation: to.Ptr("Sample Stream Analytics Job Input"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Input"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs/Test/action"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Test Stream Analytics Job Input"),
		// 				Operation: to.Ptr("Test Stream Analytics Job Input"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Input"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs/Write"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Write Stream Analytics Job Input"),
		// 				Operation: to.Ptr("Write Stream Analytics Job Input"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Input"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/metricdefinitions/Read"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Read Metric Definitions"),
		// 				Operation: to.Ptr("Read Metric Definitions"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Metric Definitions"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/operationresults/Read"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Read operation results for Stream Analytics Job"),
		// 				Operation: to.Ptr("Read operation results for Stream Analytics Job"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs/Delete"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Delete Stream Analytics Job Output"),
		// 				Operation: to.Ptr("Delete Stream Analytics Job Output"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Output"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs/operationresults/Read"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Read operation results for Stream Analytics Job Output"),
		// 				Operation: to.Ptr("Read operation results for Stream Analytics Job Output"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Output"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs/Read"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Read Stream Analytics Job Output"),
		// 				Operation: to.Ptr("Read Stream Analytics Job Output"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Output"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs/Test/action"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Test Stream Analytics Job Output"),
		// 				Operation: to.Ptr("Test Stream Analytics Job Output"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Output"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs/Write"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Write Stream Analytics Job Output"),
		// 				Operation: to.Ptr("Write Stream Analytics Job Output"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Output"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Read diagnostic setting."),
		// 				Operation: to.Ptr("Read diagnostic setting."),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("streamingjobs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Write diagnostic setting."),
		// 				Operation: to.Ptr("Write diagnostic setting."),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("streamingjobs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/PublishEdgePackage/action"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Publish edge package for Stream Analytics Job"),
		// 				Operation: to.Ptr("Publish edge package for Stream Analytics Job"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/Read"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Read Stream Analytics Job"),
		// 				Operation: to.Ptr("Read Stream Analytics Job"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/Scale/action"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Scale Stream Analytics Job"),
		// 				Operation: to.Ptr("Scale Stream Analytics Job"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/Start/action"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Start Stream Analytics Job"),
		// 				Operation: to.Ptr("Start Stream Analytics Job"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/Stop/action"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Stop Stream Analytics Job"),
		// 				Operation: to.Ptr("Stop Stream Analytics Job"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/transformations/Delete"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Delete Stream Analytics Job Transformation"),
		// 				Operation: to.Ptr("Delete Stream Analytics Job Transformation"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Transformation"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/transformations/Read"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Read Stream Analytics Job Transformation"),
		// 				Operation: to.Ptr("Read Stream Analytics Job Transformation"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Transformation"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/transformations/Write"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Write Stream Analytics Job Transformation"),
		// 				Operation: to.Ptr("Write Stream Analytics Job Transformation"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job Transformation"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/Write"),
		// 			Display: &armstreamanalytics.OperationDisplay{
		// 				Description: to.Ptr("Write Stream Analytics Job"),
		// 				Operation: to.Ptr("Write Stream Analytics Job"),
		// 				Provider: to.Ptr("Microsoft Azure Stream Analytics"),
		// 				Resource: to.Ptr("Stream Analytics Job"),
		// 			},
		// 	}},
		// }
	}
}
