package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/ExtendedUeInfo5GGet.json
func ExampleExtendedUeInformationClient_Get_getUeInformation5G() {
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
	// 	Properties: &armmobilenetwork.UeInfo5G{
	// 		LastReadAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		RatType: to.Ptr(armmobilenetwork.RatTypeFiveG),
	// 		Info: &armmobilenetwork.UeInfo5GProperties{
	// 			ConnectionInfo: &armmobilenetwork.UeConnectionInfo5G{
	// 				AllowedNssai: []*armmobilenetwork.Snssai{
	// 					{
	// 						Sst: to.Ptr[int32](1),
	// 					},
	// 					{
	// 						Sd: to.Ptr("abcdef"),
	// 						Sst: to.Ptr[int32](2),
	// 				}},
	// 				AmfUeNgapID: to.Ptr[int64](549755813888),
	// 				GlobalRanNodeID: &armmobilenetwork.GlobalRanNodeID{
	// 					GNbID: &armmobilenetwork.GNbID{
	// 						BitLength: to.Ptr[int32](32),
	// 						GNBValue: to.Ptr("01234567"),
	// 					},
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
	// 				PerUeTnla: to.Ptr("00101.0.macroEnbId-10.232.228.84:36412"),
	// 				RanUeNgapID: to.Ptr[int32](12345678),
	// 				RrcEstablishmentCause: to.Ptr(armmobilenetwork.RrcEstablishmentCauseEmergency),
	// 				UeState: to.Ptr(armmobilenetwork.UeStateConnected),
	// 				UeUsageSetting: to.Ptr(armmobilenetwork.UeUsageSettingDataCentric),
	// 			},
	// 			FivegGuti: &armmobilenetwork.Guti5G{
	// 				AmfID: &armmobilenetwork.AmfID{
	// 					Pointer: to.Ptr[int32](1),
	// 					RegionID: to.Ptr[int32](1),
	// 					SetID: to.Ptr[int32](1),
	// 				},
	// 				FivegTmsi: to.Ptr[int32](907),
	// 				Plmn: &armmobilenetwork.PlmnID{
	// 					Mcc: to.Ptr("001"),
	// 					Mnc: to.Ptr("01"),
	// 				},
	// 			},
	// 			Pei: to.Ptr("imei-123456789012345"),
	// 			SessionInfo: []*armmobilenetwork.UeSessionInfo5G{
	// 				{
	// 					Ambr: &armmobilenetwork.Ambr{
	// 						Downlink: to.Ptr("2 Gbps"),
	// 						Uplink: to.Ptr("2 Gbps"),
	// 					},
	// 					Dnn: to.Ptr("internet"),
	// 					PdnType: to.Ptr(armmobilenetwork.PdnTypeIPV4),
	// 					PduSessionID: to.Ptr[int32](15),
	// 					QosFlow: []*armmobilenetwork.UeQOSFlow{
	// 						{
	// 							Fiveqi: to.Ptr[int32](9),
	// 							Qfi: to.Ptr[int32](9),
	// 						},
	// 						{
	// 							Fiveqi: to.Ptr[int32](1),
	// 							Gbr: &armmobilenetwork.Ambr{
	// 								Downlink: to.Ptr("10.0 Mbps"),
	// 								Uplink: to.Ptr("100.0 Mbps"),
	// 							},
	// 							Mbr: &armmobilenetwork.Ambr{
	// 								Downlink: to.Ptr("10.0 Kbps"),
	// 								Uplink: to.Ptr("100.0 Kbps"),
	// 							},
	// 							Qfi: to.Ptr[int32](1),
	// 					}},
	// 					Snssai: &armmobilenetwork.Snssai{
	// 						Sd: to.Ptr("abcdef"),
	// 						Sst: to.Ptr[int32](1),
	// 					},
	// 					UeIPAddress: &armmobilenetwork.UeIPAddress{
	// 						IPV4Addr: to.Ptr("10.10.0.1"),
	// 					},
	// 			}},
	// 			Supi: to.Ptr("imsi-84449105622"),
	// 		},
	// 	},
	// }
}
