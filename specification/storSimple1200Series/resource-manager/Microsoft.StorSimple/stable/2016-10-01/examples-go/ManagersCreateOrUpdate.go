package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ManagersCreateOrUpdate.json
func ExampleManagersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagersClient().CreateOrUpdate(ctx, "ResourceGroupForSDKTest", "hManagerForSDKTest", armstorsimple1200series.Manager{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Manager = armstorsimple1200series.Manager{
	// 	Name: to.Ptr("hManagerForSDKTest"),
	// 	Type: to.Ptr("Microsoft.StorSimple/Managers"),
	// 	ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/Managers/hManagerForSDKTest"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"TagName": to.Ptr("Demo manager for SDK test"),
	// 	},
	// 	Etag: to.Ptr("W/\"datetime'2018-08-12T15%3A10%3A31.6040125Z'\""),
	// 	Properties: &armstorsimple1200series.ManagerProperties{
	// 		CisIntrinsicSettings: &armstorsimple1200series.ManagerIntrinsicSettings{
	// 			Type: to.Ptr(armstorsimple1200series.ManagerTypeHelsinkiV1),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SKU: &armstorsimple1200series.ManagerSKU{
	// 			Name: to.Ptr("Standard"),
	// 		},
	// 	},
	// }
}
