Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmysql%2Farmmysql%2Fv0.3.1/sdk/resourcemanager/mysql/armmysql/README.md) on how to add the SDK to your project and authenticate.

```go
package armmysql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerCreatePointInTimeRestore.json
func ExampleServersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmysql.NewServersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armmysql.ServerForCreate{
			Location: to.StringPtr("<location>"),
			Properties: &armmysql.ServerPropertiesForRestore{
				CreateMode:         armmysql.CreateMode("PointInTimeRestore").ToPtr(),
				RestorePointInTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-14T00:00:37.467Z"); return t }()),
				SourceServerID:     to.StringPtr("<source-server-id>"),
			},
			SKU: &armmysql.SKU{
				Name:     to.StringPtr("<name>"),
				Capacity: to.Int32Ptr(2),
				Family:   to.StringPtr("<family>"),
				Tier:     armmysql.SKUTier("GeneralPurpose").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.ServersClientCreateResult)
}
```
