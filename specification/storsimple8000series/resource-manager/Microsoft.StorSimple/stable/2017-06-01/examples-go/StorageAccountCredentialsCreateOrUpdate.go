package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/StorageAccountCredentialsCreateOrUpdate.json
func ExampleStorageAccountCredentialsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageAccountCredentialsClient().BeginCreateOrUpdate(ctx, "SACForTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", armstorsimple8000series.StorageAccountCredential{
		Properties: &armstorsimple8000series.StorageAccountCredentialProperties{
			AccessKey: &armstorsimple8000series.AsymmetricEncryptedSecret{
				EncryptionAlgorithm:      to.Ptr(armstorsimple8000series.EncryptionAlgorithmRSAESPKCS1V15),
				EncryptionCertThumbprint: to.Ptr("A872A2DF196AC7682EE24791E7DE2E2A360F5926"),
				Value:                    to.Ptr("ATuJSkmrFk4h8r1jrZ4nd3nthLSddcguEO5QLO/NECUtTuB9kL4dNv3/jC4WOvFkeVr3x1UvfhlIeMmJBF1SMr6hR1JzD0xNU/TtQqUeXN7V3jk7I+2l67P9StuHWR6OMd3XOLwvznxOEQtEWpweDiobZU1ZiY03WafcGZFpV5j6tEoHeopoZ1J/GhPtkYmx+TqxzUN6qnir5rP3NSYiZciImP/qu8U9yUV/xpVRv39KvFc2Yr5SpKpMMRUj55XW10UnPer63M6KovF8X9Wi/fNnrZAs1Esl5XddZETGrW/e5B++VMJ6w0Q/uvPR+UBwrOU0804l0SzwdIe3qVVd0Q=="),
			},
			EndPoint:  to.Ptr("blob.core.windows.net"),
			SSLStatus: to.Ptr(armstorsimple8000series.SSLStatusEnabled),
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
	// res.StorageAccountCredential = armstorsimple8000series.StorageAccountCredential{
	// 	Name: to.Ptr("SACForTest"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/storageAccountCredentials"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/storageAccountCredentials/SACForTest"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.StorageAccountCredentialProperties{
	// 		EndPoint: to.Ptr("blob.core.windows.net"),
	// 		SSLStatus: to.Ptr(armstorsimple8000series.SSLStatusEnabled),
	// 		VolumesCount: to.Ptr[int32](0),
	// 	},
	// }
}
