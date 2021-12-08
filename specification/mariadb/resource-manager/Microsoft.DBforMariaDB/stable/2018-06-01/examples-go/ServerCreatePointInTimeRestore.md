Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmariadb%2Farmmariadb%2Fv0.1.0/sdk/resourcemanager/mariadb/armmariadb/README.md) on how to add the SDK to your project and authenticate.

```go
package armmariadb_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerCreatePointInTimeRestore.json
func ExampleServersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmariadb.NewServersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armmariadb.ServerForCreate{
			Location: to.StringPtr("<location>"),
			Properties: &armmariadb.ServerPropertiesForRestore{
				ServerPropertiesForCreate: armmariadb.ServerPropertiesForCreate{
					CreateMode: armmariadb.CreateModePointInTimeRestore.ToPtr(),
				},
				RestorePointInTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-14T00:00:37.467Z"); return t }()),
				SourceServerID:     to.StringPtr("<source-server-id>"),
			},
			SKU: &armmariadb.SKU{
				Name:     to.StringPtr("<name>"),
				Capacity: to.Int32Ptr(2),
				Family:   to.StringPtr("<family>"),
				Tier:     armmariadb.SKUTierGeneralPurpose.ToPtr(),
			},
			Tags: map[string]*string{
				"ElasticServer": to.StringPtr("1"),
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
	log.Printf("Server.ID: %s\n", *res.ID)
}
```
