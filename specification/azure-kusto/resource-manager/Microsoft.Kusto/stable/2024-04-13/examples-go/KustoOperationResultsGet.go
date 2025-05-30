package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8adce17dc500f338f86f18af30aac61dcb71c5f/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoOperationResultsGet.json
func ExampleOperationsResultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsResultsClient().Get(ctx, "westus", "30972f1b-b61d-4fd8-bd34-3dcfa24670f3", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationResult = armkusto.OperationResult{
	// 	Name: to.Ptr("30972f1b-b61d-4fd8-bd34-3dcfa24670f3"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-26T13:45:42.702Z"); return t}()),
	// 	Error: &armkusto.OperationResultErrorProperties{
	// 		Code: to.Ptr("CannotAlterFollowerDatabase"),
	// 		Message: to.Ptr("[BadRequest] Cannot alter leader cluster 'test' for resource name 'adc'."),
	// 	},
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/providers/Microsoft.Kusto/locations/westus/operationresults/30972f1b-b61d-4fd8-bd34-3dcfa24670f3"),
	// 	PercentComplete: to.Ptr[float64](1),
	// 	Properties: &armkusto.OperationResultProperties{
	// 		OperationKind: to.Ptr("FollowerDatabaseCreate"),
	// 		OperationState: to.Ptr("BadInput"),
	// 	},
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-26T13:45:39.655Z"); return t}()),
	// 	Status: to.Ptr(armkusto.StatusFailed),
	// }
}
