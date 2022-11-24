package armdatafactory_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_QueryByFactory.json
func ExamplePipelineRunsClient_QueryByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelineRunsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.QueryByFactory(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.RunFilterParameters{
		Filters: []*armdatafactory.RunQueryFilter{
			{
				Operand:  to.Ptr(armdatafactory.RunQueryFilterOperandPipelineName),
				Operator: to.Ptr(armdatafactory.RunQueryFilterOperatorEquals),
				Values: []*string{
					to.Ptr("examplePipeline")},
			}},
		LastUpdatedAfter:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:36:44.3345758Z"); return t }()),
		LastUpdatedBefore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:49:48.3686473Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
