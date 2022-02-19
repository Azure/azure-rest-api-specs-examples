Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsight%2Farmsecurityinsight%2Fv0.2.1/sdk/resourcemanager/securityinsight/armsecurityinsight/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurityinsight_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/entities/expand/PostExpandEntity.json
func ExampleEntitiesClient_Expand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewEntitiesClient("<subscription-id>", cred, nil)
	res, err := client.Expand(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<entity-id>",
		armsecurityinsight.EntityExpandParameters{
			EndTime:     to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-26T00:00:00.000Z"); return t }()),
			ExpansionID: to.StringPtr("<expansion-id>"),
			StartTime:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-25T00:00:00.000Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EntitiesClientExpandResult)
}
```
