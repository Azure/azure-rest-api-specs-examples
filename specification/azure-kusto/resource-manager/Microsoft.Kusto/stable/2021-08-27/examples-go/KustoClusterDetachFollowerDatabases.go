package armkusto_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClusterDetachFollowerDatabases.json
func ExampleClustersClient_BeginDetachFollowerDatabases() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDetachFollowerDatabases(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armkusto.FollowerDatabaseDefinition{
			AttachedDatabaseConfigurationName: to.StringPtr("<attached-database-configuration-name>"),
			ClusterResourceID:                 to.StringPtr("<cluster-resource-id>"),
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
