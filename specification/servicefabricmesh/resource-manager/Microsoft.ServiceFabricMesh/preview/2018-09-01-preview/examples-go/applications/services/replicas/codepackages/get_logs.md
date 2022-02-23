Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicefabricmesh%2Farmservicefabricmesh%2Fv0.2.1/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicefabricmesh_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
)

// x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/applications/services/replicas/codepackages/get_logs.json
func ExampleCodePackageClient_GetContainerLogs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabricmesh.NewCodePackageClient("<subscription-id>", cred, nil)
	res, err := client.GetContainerLogs(ctx,
		"<resource-group-name>",
		"<application-resource-name>",
		"<service-resource-name>",
		"<replica-name>",
		"<code-package-name>",
		&armservicefabricmesh.CodePackageClientGetContainerLogsOptions{Tail: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CodePackageClientGetContainerLogsResult)
}
```
