package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetTXTDelete.json
func ExampleRecordSetsClient_Delete_deletePrivateDnsZoneTxtRecordSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatedns.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewRecordSetsClient().Delete(ctx, "resourceGroup1", "privatezone1.com", armprivatedns.RecordTypeTXT, "recordTXT", &armprivatedns.RecordSetsClientDeleteOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
