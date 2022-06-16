package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/ManagersGetDevicePublicEncryptionKey.json
func ExampleManagersClient_GetDevicePublicEncryptionKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple8000series.NewManagersClient("d3ebfe71-b7a9-4c57-92b9-68a2afde4de5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetDevicePublicEncryptionKey(ctx,
		"sca01forsdktest",
		"ResourceGroupForSDKTest",
		"ManagerForSDKTest1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
