package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/CloudAppliancesListSupportedConfigurations.json
func ExampleCloudAppliancesClient_NewListSupportedConfigurationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudAppliancesClient().NewListSupportedConfigurationsPager("ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.CloudApplianceConfigurationList = armstorsimple8000series.CloudApplianceConfigurationList{
		// 	Value: []*armstorsimple8000series.CloudApplianceConfiguration{
		// 		{
		// 			Name: to.Ptr("8010"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/cloudApplianceConfigurations"),
		// 			ID: to.Ptr("/subscriptions/d3ebfe71-b7a9-4c57-92b9-68a2afde4de5/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/managers/cloudApplianceConfigurations/8010"),
		// 			Kind: to.Ptr("Series8000"),
		// 			Properties: &armstorsimple8000series.CloudApplianceConfigurationProperties{
		// 				AcsConfiguration: &armstorsimple8000series.AcsConfiguration{
		// 					Namespace: to.Ptr("wuspod01rp1users"),
		// 					Realm: to.Ptr("http://windowscloudbackup/m3"),
		// 					ServiceURL: to.Ptr("accesscontrol.windows.net"),
		// 				},
		// 				CloudPlatform: to.Ptr("public"),
		// 				ModelNumber: to.Ptr("8010"),
		// 				SupportedRegions: []*string{
		// 					to.Ptr("All")},
		// 					SupportedStorageAccountTypes: []*string{
		// 						to.Ptr("Standard_LRS"),
		// 						to.Ptr("Standard_GRS"),
		// 						to.Ptr("Standard_RAGRS")},
		// 						SupportedVMImages: []*armstorsimple8000series.VMImage{
		// 							{
		// 								Name: to.Ptr("StorSimple 8000 Series Update 4"),
		// 								Offer: to.Ptr("StorSimple"),
		// 								Publisher: to.Ptr("MicrosoftHybridCloudStorage"),
		// 								SKU: to.Ptr("StorSimple-Garda-8000-Series-BBUpdate"),
		// 								Version: to.Ptr("9600.17820.170208"),
		// 						}},
		// 						SupportedVMTypes: []*string{
		// 							to.Ptr("Standard_A3")},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("8020"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/cloudApplianceConfigurations"),
		// 						ID: to.Ptr("/subscriptions/d3ebfe71-b7a9-4c57-92b9-68a2afde4de5/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/managers/cloudApplianceConfigurations/8020"),
		// 						Kind: to.Ptr("Series8000"),
		// 						Properties: &armstorsimple8000series.CloudApplianceConfigurationProperties{
		// 							AcsConfiguration: &armstorsimple8000series.AcsConfiguration{
		// 								Namespace: to.Ptr("wuspod01rp1users"),
		// 								Realm: to.Ptr("http://windowscloudbackup/m3"),
		// 								ServiceURL: to.Ptr("accesscontrol.windows.net"),
		// 							},
		// 							CloudPlatform: to.Ptr("public"),
		// 							ModelNumber: to.Ptr("8020"),
		// 							SupportedRegions: []*string{
		// 								to.Ptr("Central US"),
		// 								to.Ptr("East US"),
		// 								to.Ptr("East US 2"),
		// 								to.Ptr("South Central US"),
		// 								to.Ptr("West US"),
		// 								to.Ptr("North Europe"),
		// 								to.Ptr("West Europe"),
		// 								to.Ptr("East Asia"),
		// 								to.Ptr("Southeast Asia"),
		// 								to.Ptr("Japan East"),
		// 								to.Ptr("Japan West"),
		// 								to.Ptr("Australia East"),
		// 								to.Ptr("Australia Southeast")},
		// 								SupportedStorageAccountTypes: []*string{
		// 									to.Ptr("Premium_LRS")},
		// 									SupportedVMImages: []*armstorsimple8000series.VMImage{
		// 										{
		// 											Name: to.Ptr("StorSimple 8000 Series Update 4"),
		// 											Offer: to.Ptr("StorSimple"),
		// 											Publisher: to.Ptr("MicrosoftHybridCloudStorage"),
		// 											SKU: to.Ptr("StorSimple-Garda-8000-Series-BBUpdate"),
		// 											Version: to.Ptr("9600.17820.170208"),
		// 									}},
		// 									SupportedVMTypes: []*string{
		// 										to.Ptr("Standard_DS3")},
		// 									},
		// 							}},
		// 						}
	}
}
