package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/065033d1c4087a2b009e71c0b3f0666718354ebd/specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheExport.json
func ExampleClient_BeginExportData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginExportData(ctx, "rg1", "cache1", armredis.ExportRDBParameters{
		Format:    to.Ptr("RDB"),
		Container: to.Ptr("https://contosostorage.blob.core.window.net/urltoBlobContainer?sasKeyParameters"),
		Prefix:    to.Ptr("datadump1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
