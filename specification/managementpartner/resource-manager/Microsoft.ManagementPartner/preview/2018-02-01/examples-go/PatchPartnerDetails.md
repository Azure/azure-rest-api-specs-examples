Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmanagementpartner%2Farmmanagementpartner%2Fv0.1.0/sdk/resourcemanager/managementpartner/armmanagementpartner/README.md) on how to add the SDK to your project and authenticate.

```go
package armmanagementpartner_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementpartner/armmanagementpartner"
)

// x-ms-original-file: specification/managementpartner/resource-manager/Microsoft.ManagementPartner/preview/2018-02-01/examples/PatchPartnerDetails.json
func ExamplePartnerClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagementpartner.NewPartnerClient(cred, nil)
	res, err := client.Update(ctx,
		"<partner-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PartnerResponse.ID: %s\n", *res.ID)
}
```
