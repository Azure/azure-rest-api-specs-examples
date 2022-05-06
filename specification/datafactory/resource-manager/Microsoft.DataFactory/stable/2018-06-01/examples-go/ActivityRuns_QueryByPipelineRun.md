Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatafactory%2Farmdatafactory%2Fv0.5.0/sdk/resourcemanager/datafactory/armdatafactory/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ActivityRuns_QueryByPipelineRun.json
func ExampleActivityRunsClient_QueryByPipelineRun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatafactory.NewActivityRunsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.QueryByPipelineRun(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<run-id>",
		armdatafactory.RunFilterParameters{
			LastUpdatedAfter:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:36:44.3345758Z"); return t }()),
			LastUpdatedBefore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:49:48.3686473Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
