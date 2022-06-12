```go
package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/Assessments_Get.json
func ExampleAssessmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmigrate.NewAssessmentsClient("6393a73f-8d55-47ef-b6dd-179b3e0c7910", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"abgoyal-westEurope",
		"abgoyalWEselfhostb72bproject",
		"Test1",
		"assessment_5_9_2019_16_22_14",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmigrate%2Farmmigrate%2Fv1.0.0/sdk/resourcemanager/migrate/armmigrate/README.md) on how to add the SDK to your project and authenticate.
