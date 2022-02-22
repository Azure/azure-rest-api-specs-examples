Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.3.1/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateInstancePoolMax.json
func ExampleInstancePoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewInstancePoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<instance-pool-name>",
		armsql.InstancePool{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"a": to.StringPtr("b"),
			},
			Properties: &armsql.InstancePoolProperties{
				LicenseType: armsql.InstancePoolLicenseType("LicenseIncluded").ToPtr(),
				SubnetID:    to.StringPtr("<subnet-id>"),
				VCores:      to.Int32Ptr(8),
			},
			SKU: &armsql.SKU{
				Name:   to.StringPtr("<name>"),
				Family: to.StringPtr("<family>"),
				Tier:   to.StringPtr("<tier>"),
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
	log.Printf("Response result: %#v\n", res.InstancePoolsClientCreateOrUpdateResult)
}
```
