package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/ExtendedUeInfo4GGet.json
func ExampleExtendedUeInformationClient_Get_getUeInformation4G() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExtendedUeInformationClient().Get(ctx, "rg1", "TestPacketCoreCP", "84449105622", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExtendedUeInfo = armmobilenetwork.ExtendedUeInfo{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/packetCoreControlPlanes/ues/extendedInformation"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/ues/84449105622/extendedInformation/default"),
	// 	Properties: &armmobilenetwork.UeInfo4G{
	// 		LastReadAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		RatType: to.Ptr(armmobilenetwork.RatTypeFourG),
	// 		Info: &armmobilenetwork.UeInfo4GProperties{
	// 			ConnectionInfo: &armmobilenetwork.UeConnectionInfo4G{
	// 				EnbS1ApID: to.Ptr[int32](12345678),
	// 				GlobalRanNodeID: &armmobilenetwork.GlobalRanNodeID{
	// 					ENbID: to.Ptr("MacroeNB-ABCDE"),
	// 					PlmnID: &armmobilenetwork.PlmnID{
	// 						Mcc: to.Ptr("001"),
	// 						Mnc: to.Ptr("01"),
	// 					},
	// 				},
	// 				LastActivityTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 				LastVisitedTai: to.Ptr("00101 000001"),
	// 				LocationInfo: &armmobilenetwork.UeLocationInfo{
	// 					LocationType: to.Ptr("EUTRA"),
	// 					Plmn: &armmobilenetwork.PlmnID{
	// 						Mcc: to.Ptr("001"),
	// 						Mnc: to.Ptr("01"),
	// 					},
	// 					Tac: to.Ptr("000001"),
	// 				},
	// 				MmeS1ApID: to.Ptr[int32](12345678),
	// 				PerUeTnla: to.Ptr("00101.0.macroEnbId-10.232.228.84:36412"),
	// 				RrcEstablishmentCause: to.Ptr(armmobilenetwork.RrcEstablishmentCauseEmergency),
	// 				UeState: to.Ptr(armmobilenetwork.UeStateConnected),
	// 				UeUsageSetting: to.Ptr(armmobilenetwork.UeUsageSettingDataCentric),
	// 			},
	// 			Guti: &armmobilenetwork.Guti4G{
	// 				MTmsi: to.Ptr[int32](907),
	// 				MmeID: &armmobilenetwork.MmeID{
	// 					Code: to.Ptr[int32](1),
	// 					GroupID: to.Ptr[int32](1),
	// 				},
	// 				Plmn: &armmobilenetwork.PlmnID{
	// 					Mcc: to.Ptr("001"),
	// 					Mnc: to.Ptr("01"),
	// 				},
	// 			},
	// 			Imei: to.Ptr("123456789012345"),
	// 			Imeisv: to.Ptr("2993972087439794"),
	// 			Imsi: to.Ptr("84449105622"),
	// 			SessionInfo: []*armmobilenetwork.UeSessionInfo4G{
	// 				{
	// 					Apn: to.Ptr("internet"),
	// 					Ebi: to.Ptr[int32](15),
	// 					PdnType: to.Ptr(armmobilenetwork.PdnTypeIPV4),
	// 					UeIPAddress: &armmobilenetwork.UeIPAddress{
	// 						IPV4Addr: to.Ptr("10.10.0.1"),
	// 					},
	// 			}},
	// 		},
	// 	},
	// }
}
