package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/ManagersGetExtendedInfo.json
func ExampleManagersClient_GetExtendedInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagersClient().GetExtendedInfo(ctx, "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagerExtendedInfo = armstorsimple8000series.ManagerExtendedInfo{
	// 	Name: to.Ptr("vaultExtendedInfo"),
	// 	Type: to.Ptr("Microsoft.StorSimple/Managers/extendedInformation"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/Managers/ManagerForSDKTest1extendedInformation/vaultExtendedInfo"),
	// 	Etag: to.Ptr("bcb55c1e-29b2-42b3-a507-f2a0cd27a963"),
	// 	Properties: &armstorsimple8000series.ManagerExtendedInfoProperties{
	// 		Algorithm: to.Ptr("None"),
	// 		EncryptionKey: to.Ptr("MIIDQjCCAiqgAwIBAgIQQmzC+SfPepJF20zGMFBbJjANBgkqhkiG9w0BAQUFADBdMVswWQYDVQQDHlIAQwBCAF8AMgA0ADMANQAyADQAMwAzADYANwAzADkAOAA0ADcAMQA5ADEANQBfADYAMwA2ADMAMgAzADQAMgA2ADEAOQAxADMANgAwADgAOAAxMB4XDTE3MDYwMTA3MDAwMFoXDTIwMDYwNjEwNDMzOVowXTFbMFkGA1UEAx5SAEMAQgBfADIANAAzADUAMgA0ADMAMwA2ADcAMwA5ADgANAA3ADEAOQAxADUAXwA2ADMANgAzADIAMwA0ADIANgAxADkAMQAzADYAMAA4ADgAMTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAJvO+6ytUY0Z/x2GVE2grF8L7XvuJ63mN3ebHKncQVlJiKjyfuG7iDzQUfW0HZZ/TtgxineJLHPNg67NIxjwaqhr+c0kbR6VUtyShqBxfqFAQt+Lx1bySxZ5S/G/ZSKc++evMwhYMoCkU8k1uQWOcco4N4+7HrB3XPkZee1TckhXnsjS5d9OawNCfVxt9PsCEwhX8Ezj58u1FN21OOosxKdSBTaRI1Hh+kIAP2ZPeldRvgjJ+HVRzR7hqm6OLVg/azaSziGMkP+UmvvwsyG8EI2xfhkF3K11rEFRng0fB06V01LN11hpVEL1uvNbGbfkmlIjEWmiwujiFGMdVK0eYfUCAwEAATANBgkqhkiG9w0BAQUFAAOCAQEAcztBrLgiL3fqykBXQ/3GjJXZ1PX7z9CFjoLkHdv6ZG+mkAkAkjaOxYvzCueIhX61u2zOgg9AJYW+J3BKvSGe5W5O7W8OLL0TRCmzSASH4Bap9oZmb6exs8D45aPgGRW/Mmhm2W+jGk9NIg6W8GFPksqP2XY3DlKSfAztzIU86fPRfXlNHknvpp7rt+gv7WjbSt8smSTJ7PJwUh9s/oYL+k5GPtYxHkB8A7YCgVj0rk6v9uHvhmOXm3Cm+1wOJA9ebk0U1CUUJ8maMi89JvFuxhps8kBObnp6M52AzVErMwqn9zH0gbDRlh9fIQPwvXsXm7DosxccD5KuG9/oX4eqfQ=="),
	// 		EncryptionKeyThumbprint: to.Ptr("A872A2DF196AC7682EE24791E7DE2E2A360F5926"),
	// 		IntegrityKey: to.Ptr("x3V1JcHquNW1P87Vgz10Pw=="),
	// 	},
	// }
}
