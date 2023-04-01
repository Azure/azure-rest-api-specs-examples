package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsUpdateSecuritySettings.json
func ExampleDeviceSettingsClient_BeginUpdateSecuritySettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeviceSettingsClient().BeginUpdateSecuritySettings(ctx, "Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", armstorsimple8000series.SecuritySettingsPatch{
		Properties: &armstorsimple8000series.SecuritySettingsPatchProperties{
			ChapSettings: &armstorsimple8000series.ChapSettings{
				InitiatorSecret: &armstorsimple8000series.AsymmetricEncryptedSecret{
					EncryptionAlgorithm:      to.Ptr(armstorsimple8000series.EncryptionAlgorithmRSAESPKCS1V15),
					EncryptionCertThumbprint: to.Ptr("A872A2DF196AC7682EE24791E7DE2E2A360F5926"),
					Value:                    to.Ptr("V/uVfWk5OcXfMC0HvUV89o9+cmF636jBnqhFM1pD/zHhmh8Z1KB5/LhVV3T53uGjIlKL3wjhwg+9NIQrIbYuKhl/r8jSftSSH+WqUnQHTRDWazjPAeMu6ozrL5RYzP1h5mgw7XtidZPaaV9ae/uF1KQPkK6TIARaOTdr8I/BLWUg7WdDrfARNYHnW6ezXek1M9Qhv1sL9fZY+JrGB58LF6D2aC2Xjed4K4Jk6v2T1ieneNV27uIdnt21TajuM7w90UlRiVZJZtq/KdEUfqI28C7VoUdcXluAwzR95Ho8hmyIJDqeW3/Wxymdjv+Rctwqtmcka9i2G85Hj8SVV3g4kA=="),
				},
				InitiatorUser: to.Ptr("test-initiator-user"),
				TargetSecret: &armstorsimple8000series.AsymmetricEncryptedSecret{
					EncryptionAlgorithm:      to.Ptr(armstorsimple8000series.EncryptionAlgorithmRSAESPKCS1V15),
					EncryptionCertThumbprint: to.Ptr("A872A2DF196AC7682EE24791E7DE2E2A360F5926"),
					Value:                    to.Ptr("OTR4uwVpy+pf0zthnCIAUXurC8NdSh8RpRG5GWL9TSv4WtkVmpeU/U2A4vjkrchfQOzI1x+uooWikWW9txwwQOM+/N3NG44+/dlHoaEe7AxjmItCKhNj8K2RM6D1mb45wicbF/M4uanuXnGXuT+JmZ+1Lcy2k1GXsk67ejplz2K08h37B+oIW85qMUHLdKuuQlAA/fFS+q6qMti3j2Q8Fr+Sh4U76/2AZVkKRtFeqPB1QhC12dFx6TFoZJkMFzdQz4WNvWVelIK2McKNnOiH0/Z5lAXC7164uzReAoTEfqoNU7qqqRrHhsdwWPu6jbeUn8BQnr7A/X6NWvgeax+HGA=="),
				},
				TargetUser: to.Ptr("test-target-user"),
			},
			DeviceAdminPassword: &armstorsimple8000series.AsymmetricEncryptedSecret{
				EncryptionAlgorithm:      to.Ptr(armstorsimple8000series.EncryptionAlgorithmRSAESPKCS1V15),
				EncryptionCertThumbprint: to.Ptr("A872A2DF196AC7682EE24791E7DE2E2A360F5926"),
				Value:                    to.Ptr("<value>"),
			},
			RemoteManagementSettings: &armstorsimple8000series.RemoteManagementSettingsPatch{
				RemoteManagementMode: to.Ptr(armstorsimple8000series.RemoteManagementModeConfigurationHTTPSAndHTTPEnabled),
			},
			SnapshotPassword: &armstorsimple8000series.AsymmetricEncryptedSecret{
				EncryptionAlgorithm:      to.Ptr(armstorsimple8000series.EncryptionAlgorithmRSAESPKCS1V15),
				EncryptionCertThumbprint: to.Ptr("A872A2DF196AC7682EE24791E7DE2E2A360F5926"),
				Value:                    to.Ptr("<value>"),
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
	// res.SecuritySettings = armstorsimple8000series.SecuritySettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/securitySettings"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/securitySettings/default"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.SecuritySettingsProperties{
	// 		ChapSettings: &armstorsimple8000series.ChapSettings{
	// 			InitiatorUser: to.Ptr("test-initiator-user"),
	// 			TargetUser: to.Ptr("test-target-user"),
	// 		},
	// 		RemoteManagementSettings: &armstorsimple8000series.RemoteManagementSettings{
	// 			RemoteManagementCertificate: to.Ptr(""),
	// 			RemoteManagementMode: to.Ptr(armstorsimple8000series.RemoteManagementModeConfigurationHTTPSAndHTTPEnabled),
	// 		},
	// 	},
	// }
}
