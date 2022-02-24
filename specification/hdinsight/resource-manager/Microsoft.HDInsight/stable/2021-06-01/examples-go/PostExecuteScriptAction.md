Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhdinsight%2Farmhdinsight%2Fv0.2.1/sdk/resourcemanager/hdinsight/armhdinsight/README.md) on how to add the SDK to your project and authenticate.

```go
package armhdinsight_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/PostExecuteScriptAction.json
func ExampleClustersClient_BeginExecuteScriptActions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginExecuteScriptActions(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armhdinsight.ExecuteScriptActionParameters{
			PersistOnSuccess: to.BoolPtr(false),
			ScriptActions: []*armhdinsight.RuntimeScriptAction{
				{
					Name:       to.StringPtr("<name>"),
					Parameters: to.StringPtr("<parameters>"),
					Roles: []*string{
						to.StringPtr("headnode"),
						to.StringPtr("workernode")},
					URI: to.StringPtr("<uri>"),
				}},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
