Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.2.1/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolWorkloadClassifierMax.json
func ExampleSQLPoolWorkloadClassifierClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewSQLPoolWorkloadClassifierClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<sql-pool-name>",
		"<workload-group-name>",
		"<workload-classifier-name>",
		armsynapse.WorkloadClassifier{
			Properties: &armsynapse.WorkloadClassifierProperties{
				Context:    to.StringPtr("<context>"),
				EndTime:    to.StringPtr("<end-time>"),
				Importance: to.StringPtr("<importance>"),
				Label:      to.StringPtr("<label>"),
				MemberName: to.StringPtr("<member-name>"),
				StartTime:  to.StringPtr("<start-time>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SQLPoolWorkloadClassifierClientCreateOrUpdateResult)
}
```
