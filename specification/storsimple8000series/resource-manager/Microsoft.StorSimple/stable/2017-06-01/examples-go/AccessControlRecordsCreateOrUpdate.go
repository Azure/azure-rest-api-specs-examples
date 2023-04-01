package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/AccessControlRecordsCreateOrUpdate.json
func ExampleAccessControlRecordsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccessControlRecordsClient().BeginCreateOrUpdate(ctx, "ACRForTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", armstorsimple8000series.AccessControlRecord{
		Properties: &armstorsimple8000series.AccessControlRecordProperties{
			InitiatorName: to.Ptr("iqn.2017-06.com.contoso:ForTest"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessControlRecord = armstorsimple8000series.AccessControlRecord{
	// 	Name: to.Ptr("ACRForTest"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/accessControlRecords"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/accessControlRecords/ACRForTest"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.AccessControlRecordProperties{
	// 		InitiatorName: to.Ptr("iqn.2017-06.com.contoso:ForTest"),
	// 		VolumeCount: to.Ptr[int32](0),
	// 	},
	// }
}
