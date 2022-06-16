package armredisenterprise_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
)

// x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2021-08-01/examples/RedisEnterpriseDatabasesCreate.json
func ExampleDatabasesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armredisenterprise.NewDatabasesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<database-name>",
		armredisenterprise.Database{
			Properties: &armredisenterprise.DatabaseProperties{
				ClientProtocol:   armredisenterprise.Protocol("Encrypted").ToPtr(),
				ClusteringPolicy: armredisenterprise.ClusteringPolicy("EnterpriseCluster").ToPtr(),
				EvictionPolicy:   armredisenterprise.EvictionPolicy("AllKeysLRU").ToPtr(),
				Modules: []*armredisenterprise.Module{
					{
						Name: to.StringPtr("<name>"),
						Args: to.StringPtr("<args>"),
					},
					{
						Name: to.StringPtr("<name>"),
						Args: to.StringPtr("<args>"),
					},
					{
						Name: to.StringPtr("<name>"),
					}},
				Persistence: &armredisenterprise.Persistence{
					AofEnabled:   to.BoolPtr(true),
					AofFrequency: armredisenterprise.AofFrequency("1s").ToPtr(),
				},
				Port: to.Int32Ptr(10000),
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
	log.Printf("Response result: %#v\n", res.DatabasesClientCreateResult)
}
