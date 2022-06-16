package armdns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
)

// x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListRecordSetsByZone.json
func ExampleRecordSetsClient_ListAllByDNSZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdns.NewRecordSetsClient("<subscription-id>", cred, nil)
	pager := client.ListAllByDNSZone("<resource-group-name>",
		"<zone-name>",
		&armdns.RecordSetsListAllByDNSZoneOptions{Top: nil,
			RecordSetNameSuffix: nil,
		})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("RecordSet.ID: %s\n", *v.ID)
		}
	}
}
