package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/getRequiredAmlFSSubnetsSize.json
func ExampleManagementClient_GetRequiredAmlFSSubnetsSize() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().GetRequiredAmlFSSubnetsSize(ctx, &armstoragecache.ManagementClientGetRequiredAmlFSSubnetsSizeOptions{RequiredAMLFilesystemSubnetsSizeInfo: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RequiredAmlFilesystemSubnetsSize = armstoragecache.RequiredAmlFilesystemSubnetsSize{
	// 	FilesystemSubnetSize: to.Ptr[int32](24),
	// }
}
