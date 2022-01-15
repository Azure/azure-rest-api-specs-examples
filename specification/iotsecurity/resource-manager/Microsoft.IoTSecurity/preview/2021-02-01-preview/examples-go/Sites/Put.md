Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fiotsecurity%2Farmiotsecurity%2Fv0.2.0/sdk/resourcemanager/iotsecurity/armiotsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// x-ms-original-file: specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/Sites/Put.json
func ExampleSitesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotsecurity.NewSitesClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<scope>",
		armiotsecurity.SiteModel{
			Properties: &armiotsecurity.SiteProperties{
				DisplayName: to.StringPtr("<display-name>"),
				Tags: map[string]*string{
					"key1": to.StringPtr("value1"),
					"key2": to.StringPtr("value2"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SitesClientCreateOrUpdateResult)
}
```
