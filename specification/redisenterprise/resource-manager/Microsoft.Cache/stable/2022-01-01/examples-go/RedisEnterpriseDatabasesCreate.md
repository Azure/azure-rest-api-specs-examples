Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fredisenterprise%2Farmredisenterprise%2Fv0.4.0/sdk/resourcemanager/redisenterprise/armredisenterprise/README.md) on how to add the SDK to your project and authenticate.

```go
package armredisenterprise_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDatabasesCreate.json
func ExampleDatabasesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armredisenterprise.NewDatabasesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<database-name>",
		armredisenterprise.Database{
			Properties: &armredisenterprise.DatabaseProperties{
				ClientProtocol:   to.Ptr(armredisenterprise.ProtocolEncrypted),
				ClusteringPolicy: to.Ptr(armredisenterprise.ClusteringPolicyEnterpriseCluster),
				EvictionPolicy:   to.Ptr(armredisenterprise.EvictionPolicyAllKeysLRU),
				Modules: []*armredisenterprise.Module{
					{
						Name: to.Ptr("<name>"),
						Args: to.Ptr("<args>"),
					},
					{
						Name: to.Ptr("<name>"),
						Args: to.Ptr("<args>"),
					},
					{
						Name: to.Ptr("<name>"),
					}},
				Persistence: &armredisenterprise.Persistence{
					AofEnabled:   to.Ptr(true),
					AofFrequency: to.Ptr(armredisenterprise.AofFrequencyOneS),
				},
				Port: to.Ptr[int32](10000),
			},
		},
		&armredisenterprise.DatabasesClientBeginCreateOptions{ResumeToken: ""})
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
