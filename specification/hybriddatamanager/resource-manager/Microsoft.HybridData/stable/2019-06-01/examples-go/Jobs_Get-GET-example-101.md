Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybriddatamanager%2Farmhybriddatamanager%2Fv0.1.0/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/Jobs_Get-GET-example-101.json
func ExampleJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewJobsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<data-service-name>",
		"<job-definition-name>",
		"<job-id>",
		"<resource-group-name>",
		"<data-manager-name>",
		&armhybriddatamanager.JobsGetOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Job.ID: %s\n", *res.ID)
}
```
