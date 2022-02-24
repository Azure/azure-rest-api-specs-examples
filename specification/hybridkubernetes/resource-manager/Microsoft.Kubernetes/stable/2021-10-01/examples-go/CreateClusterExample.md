Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridkubernetes%2Farmhybridkubernetes%2Fv0.2.1/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybridkubernetes_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes"
)

// x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/stable/2021-10-01/examples/CreateClusterExample.json
func ExampleConnectedClusterClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridkubernetes.NewConnectedClusterClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armhybridkubernetes.ConnectedCluster{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Identity: &armhybridkubernetes.ConnectedClusterIdentity{
				Type: armhybridkubernetes.ResourceIdentityTypeSystemAssigned.ToPtr(),
			},
			Properties: &armhybridkubernetes.ConnectedClusterProperties{
				AgentPublicKeyCertificate: to.StringPtr("<agent-public-key-certificate>"),
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
	log.Printf("Response result: %#v\n", res.ConnectedClusterClientCreateResult)
}
```
