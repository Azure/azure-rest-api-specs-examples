```go
package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ManagersCreateOrUpdate.json
func ExampleManagersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewManagersClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"ResourceGroupForSDKTest",
		"hManagerForSDKTest",
		armstorsimple1200series.Manager{
			Name:     to.Ptr("hManagerForSDKTest"),
			Location: to.Ptr("westus"),
			Properties: &armstorsimple1200series.ManagerProperties{
				CisIntrinsicSettings: &armstorsimple1200series.ManagerIntrinsicSettings{
					Type: to.Ptr(armstorsimple1200series.ManagerTypeHelsinkiV1),
				},
				SKU: &armstorsimple1200series.ManagerSKU{
					Name: to.Ptr("Standard"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorsimple1200series%2Farmstorsimple1200series%2Fv1.0.0/sdk/resourcemanager/storsimple1200series/armstorsimple1200series/README.md) on how to add the SDK to your project and authenticate.
