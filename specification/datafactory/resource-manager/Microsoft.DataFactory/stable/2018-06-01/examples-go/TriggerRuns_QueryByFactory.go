package armdatafactory_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v10"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/TriggerRuns_QueryByFactory.json
func ExampleTriggerRunsClient_QueryByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTriggerRunsClient().QueryByFactory(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.RunFilterParameters{
		Filters: []*armdatafactory.RunQueryFilter{
			{
				Operand:  to.Ptr(armdatafactory.RunQueryFilterOperandTriggerName),
				Operator: to.Ptr(armdatafactory.RunQueryFilterOperatorEquals),
				Values: []*string{
					to.Ptr("exampleTrigger")},
			}},
		LastUpdatedAfter:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:36:44.334Z"); return t }()),
		LastUpdatedBefore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:49:48.368Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TriggerRunsQueryResponse = armdatafactory.TriggerRunsQueryResponse{
	// 	Value: []*armdatafactory.TriggerRun{
	// 		{
	// 			Message: to.Ptr(""),
	// 			Properties: map[string]*string{
	// 				"ScheduleTime": to.Ptr("6/16/2018 12:43:14 AM"),
	// 				"TriggerTime": to.Ptr("6/16/2018 12:43:15 AM"),
	// 			},
	// 			Status: to.Ptr(armdatafactory.TriggerRunStatusSucceeded),
	// 			TriggerName: to.Ptr("exampleTrigger"),
	// 			TriggerRunID: to.Ptr("08586724970898148904457116912CU27"),
	// 			TriggerRunTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:43:15.660Z"); return t}()),
	// 			TriggerType: to.Ptr("ScheduleTrigger"),
	// 			TriggeredPipelines: map[string]*string{
	// 				"examplePipeline": to.Ptr("9f3ce8b3-37d7-43eb-96ac-a656c0476283"),
	// 			},
	// 	}},
	// }
}
