Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Favs%2Farmavs%2Fv0.1.0/sdk/resourcemanager/avs/armavs/README.md) on how to add the SDK to your project and authenticate.

```go
package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/HcxEnterpriseSites_CreateOrUpdate.json
func ExampleHcxEnterpriseSitesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewHcxEnterpriseSitesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<hcx-enterprise-site-name>",
		armavs.HcxEnterpriseSite{},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("HcxEnterpriseSite.ID: %s\n", *res.ID)
}
```
