package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDatabasesCreate.json
func ExampleDatabasesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredisenterprise.NewDatabasesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"rg1",
		"cache1",
		"default",
		armredisenterprise.Database{
			Properties: &armredisenterprise.DatabaseProperties{
				ClientProtocol:   to.Ptr(armredisenterprise.ProtocolEncrypted),
				ClusteringPolicy: to.Ptr(armredisenterprise.ClusteringPolicyEnterpriseCluster),
				EvictionPolicy:   to.Ptr(armredisenterprise.EvictionPolicyAllKeysLRU),
				Modules: []*armredisenterprise.Module{
					{
						Name: to.Ptr("RedisBloom"),
						Args: to.Ptr("ERROR_RATE 0.00 INITIAL_SIZE 400"),
					},
					{
						Name: to.Ptr("RedisTimeSeries"),
						Args: to.Ptr("RETENTION_POLICY 20"),
					},
					{
						Name: to.Ptr("RediSearch"),
					}},
				Persistence: &armredisenterprise.Persistence{
					AofEnabled:   to.Ptr(true),
					AofFrequency: to.Ptr(armredisenterprise.AofFrequencyOneS),
				},
				Port: to.Ptr[int32](10000),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
