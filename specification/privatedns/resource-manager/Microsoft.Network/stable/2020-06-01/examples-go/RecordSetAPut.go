package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetAPut.json
func ExampleRecordSetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armprivatedns.NewRecordSetsClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"resourceGroup1",
		"privatezone1.com",
		armprivatedns.RecordTypeA,
		"recordA",
		armprivatedns.RecordSet{
			Properties: &armprivatedns.RecordSetProperties{
				ARecords: []*armprivatedns.ARecord{
					{
						IPv4Address: to.Ptr("1.2.3.4"),
					}},
				Metadata: map[string]*string{
					"key1": to.Ptr("value1"),
				},
				TTL: to.Ptr[int64](3600),
			},
		},
		&armprivatedns.RecordSetsClientCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
