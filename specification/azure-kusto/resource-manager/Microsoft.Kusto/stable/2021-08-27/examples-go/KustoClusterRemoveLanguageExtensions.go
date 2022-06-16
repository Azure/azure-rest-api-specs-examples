package armkusto_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClusterRemoveLanguageExtensions.json
func ExampleClustersClient_BeginRemoveLanguageExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRemoveLanguageExtensions(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armkusto.LanguageExtensionsList{
			Value: []*armkusto.LanguageExtension{
				{
					LanguageExtensionName: armkusto.LanguageExtensionName("PYTHON").ToPtr(),
				},
				{
					LanguageExtensionName: armkusto.LanguageExtensionName("R").ToPtr(),
				}},
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
