Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fredhatopenshift%2Farmredhatopenshift%2Fv0.2.1/sdk/resourcemanager/redhatopenshift/armredhatopenshift/README.md) on how to add the SDK to your project and authenticate.

```go
package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// x-ms-original-file: specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2020-04-30/examples/OpenShiftClusters_ListCredentials.json
func ExampleOpenShiftClustersClient_ListCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armredhatopenshift.NewOpenShiftClustersClient("<subscription-id>", cred, nil)
	res, err := client.ListCredentials(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OpenShiftClustersClientListCredentialsResult)
}
```
