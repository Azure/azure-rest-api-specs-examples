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

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/entities/timeline/PostTimelineEntity.json
func ExampleEntitiesGetTimelineClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewEntitiesGetTimelineClient("<subscription-id>", cred, nil)
	res, err := client.List(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<entity-id>",
		armsecurityinsight.EntityTimelineParameters{
			EndTime:        to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-01T00:00:00.000Z"); return t }()),
			NumberOfBucket: to.Int32Ptr(4),
			StartTime:      to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T00:00:00.000Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EntitiesGetTimelineClientListResult)
}
```
