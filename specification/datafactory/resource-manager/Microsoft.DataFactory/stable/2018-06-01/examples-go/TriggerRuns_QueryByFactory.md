Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatafactory%2Farmdatafactory%2Fv0.3.0/sdk/resourcemanager/datafactory/armdatafactory/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatafactory_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/TriggerRuns_QueryByFactory.json
func ExampleTriggerRunsClient_QueryByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewTriggerRunsClient("<subscription-id>", cred, nil)
	res, err := client.QueryByFactory(ctx,
		"<resource-group-name>",
		"<factory-name>",
		armdatafactory.RunFilterParameters{
			Filters: []*armdatafactory.RunQueryFilter{
				{
					Operand:  armdatafactory.RunQueryFilterOperand("TriggerName").ToPtr(),
					Operator: armdatafactory.RunQueryFilterOperator("Equals").ToPtr(),
					Values: []*string{
						to.StringPtr("exampleTrigger")},
				}},
			LastUpdatedAfter:  to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:36:44.3345758Z"); return t }()),
			LastUpdatedBefore: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:49:48.3686473Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TriggerRunsClientQueryByFactoryResult)
}
```
