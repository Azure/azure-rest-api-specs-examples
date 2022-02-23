Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhdinsight%2Farmhdinsight%2Fv0.2.1/sdk/resourcemanager/hdinsight/armhdinsight/README.md) on how to add the SDK to your project and authenticate.

```go
package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/PatchLinuxHadoopCluster.json
func ExampleClustersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewClustersClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armhdinsight.ClusterPatchParameters{
			Tags: map[string]*string{
				"key1": to.StringPtr("val1"),
				"key2": to.StringPtr("val2"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClustersClientUpdateResult)
}
```
