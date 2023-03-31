package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ChapSettingsCreateOrUpdate.json
func ExampleChapSettingsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewChapSettingsClient().BeginCreateOrUpdate(ctx, "HSDK-WSJQERQW3F", "ChapSettingForSDK", "ResourceGroupForSDKTest", "hAzureSDKOperations", armstorsimple1200series.ChapSettings{
		Name: to.Ptr("ChapSettingForSDK"),
		Properties: &armstorsimple1200series.ChapProperties{
			Password: &armstorsimple1200series.AsymmetricEncryptedSecret{
				EncryptionAlgorithm:             to.Ptr(armstorsimple1200series.EncryptionAlgorithmRSAESPKCS1V15),
				EncryptionCertificateThumbprint: to.Ptr("D73DB57C4CDD6761E159F8D1E8A7D759424983FD"),
				Value:                           to.Ptr("<value>"),
			},
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
	// res.ChapSettings = armstorsimple1200series.ChapSettings{
	// 	Name: to.Ptr("ChapSettingForSDK"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/chapSettings"),
	// 	ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/chapSettings/ChapSettingForSDK"),
	// 	Properties: &armstorsimple1200series.ChapProperties{
	// 		Password: &armstorsimple1200series.AsymmetricEncryptedSecret{
	// 			EncryptionAlgorithm: to.Ptr(armstorsimple1200series.EncryptionAlgorithmRSAESPKCS1V15),
	// 			EncryptionCertificateThumbprint: to.Ptr("D73DB57C4CDD6761E159F8D1E8A7D759424983FD"),
	// 			Value: to.Ptr("<value>"),
	// 		},
	// 	},
	// }
}
