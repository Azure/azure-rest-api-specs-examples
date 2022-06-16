package armdns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/CreateOrUpdateARecordset.json
func ExampleRecordSetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdns.NewRecordSetsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg1",
		"zone1",
		"record1",
		armdns.RecordTypeA,
		armdns.RecordSet{
			Properties: &armdns.RecordSetProperties{
				ARecords: []*armdns.ARecord{
					{
						IPv4Address: to.Ptr("127.0.0.1"),
					}},
				TTL: to.Ptr[int64](3600),
				Metadata: map[string]*string{
					"key1": to.Ptr("value1"),
				},
			},
		},
		&armdns.RecordSetsClientCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
