Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpostgresql%2Farmpostgresql%2Fv0.5.0/sdk/resourcemanager/postgresql/armpostgresql/README.md) on how to add the SDK to your project and authenticate.

```go
package armpostgresql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerCreatePointInTimeRestore.json
func ExampleServersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpostgresql.NewServersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armpostgresql.ServerForCreate{
			Location: to.Ptr("<location>"),
			Properties: &armpostgresql.ServerPropertiesForRestore{
				CreateMode:         to.Ptr(armpostgresql.CreateModePointInTimeRestore),
				RestorePointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-14T00:00:37.467Z"); return t }()),
				SourceServerID:     to.Ptr("<source-server-id>"),
			},
			SKU: &armpostgresql.SKU{
				Name:     to.Ptr("<name>"),
				Capacity: to.Ptr[int32](2),
				Family:   to.Ptr("<family>"),
				Tier:     to.Ptr(armpostgresql.SKUTierBasic),
			},
			Tags: map[string]*string{
				"ElasticServer": to.Ptr("1"),
			},
		},
		&armpostgresql.ServersClientBeginCreateOptions{ResumeToken: ""})
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
