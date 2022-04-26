Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmariadb%2Farmmariadb%2Fv0.4.0/sdk/resourcemanager/mariadb/armmariadb/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerCreatePointInTimeRestore.json
func ExampleServersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmariadb.NewServersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armmariadb.ServerForCreate{
			Location: to.Ptr("<location>"),
			Properties: &armmariadb.ServerPropertiesForRestore{
				CreateMode:         to.Ptr(armmariadb.CreateModePointInTimeRestore),
				RestorePointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-14T00:00:37.467Z"); return t }()),
				SourceServerID:     to.Ptr("<source-server-id>"),
			},
			SKU: &armmariadb.SKU{
				Name:     to.Ptr("<name>"),
				Capacity: to.Ptr[int32](2),
				Family:   to.Ptr("<family>"),
				Tier:     to.Ptr(armmariadb.SKUTierGeneralPurpose),
			},
			Tags: map[string]*string{
				"ElasticServer": to.Ptr("1"),
			},
		},
		&armmariadb.ServersClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
