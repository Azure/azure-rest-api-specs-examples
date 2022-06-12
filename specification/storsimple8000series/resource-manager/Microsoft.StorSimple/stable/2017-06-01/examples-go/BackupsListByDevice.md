```go
package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupsListByDevice.json
func ExampleBackupsClient_NewListByDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple8000series.NewBackupsClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByDevicePager("Device05ForSDKTest",
		"ResourceGroupForSDKTest",
		"ManagerForSDKTest1",
		&armstorsimple8000series.BackupsClientListByDeviceOptions{Filter: to.Ptr("createdTime%20ge%20'2017-06-22T18:30:00Z'%20and%20backupPolicyId%20eq%20'%2Fsubscriptions%2F4385cf00-2d3a-425a-832f-f4285b1c9dce%2FresourceGroups%2FResourceGroupForSDKTest%2Fproviders%2FMicrosoft.StorSimple%2Fmanagers%2FManagerForSDKTest1%2Fdevices%2FDevice05ForSDKTest%2FbackupPolicies%2FBkUpPolicy01ForSDKTest'")})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorsimple8000series%2Farmstorsimple8000series%2Fv1.0.0/sdk/resourcemanager/storsimple8000series/armstorsimple8000series/README.md) on how to add the SDK to your project and authenticate.
