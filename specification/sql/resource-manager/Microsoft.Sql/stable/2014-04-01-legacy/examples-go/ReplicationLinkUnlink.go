package armsql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01-legacy/examples/ReplicationLinkUnlink.json
func ExampleReplicationLinksClient_BeginUnlink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewReplicationLinksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUnlink(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		"<link-id>",
		armsql.UnlinkParameters{
			ForcedTermination: to.BoolPtr(true),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
