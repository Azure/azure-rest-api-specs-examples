package armprogrammableconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/programmableconnectivity/armprogrammableconnectivity"
)

// Generated from example definition: 2024-01-15-preview/OperatorApiPlans_Get_MaximumSet_Gen.json
func ExampleOperatorAPIPlansClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprogrammableconnectivity.NewClientFactory("B976474B-99FA-4C25-A3BD-8B05C3C3D07A", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperatorAPIPlansClient().Get(ctx, "etzfxkqegslnxdhdmzvbtzxahnyq", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armprogrammableconnectivity.OperatorAPIPlansClientGetResponse{
	// 	OperatorAPIPlan: &armprogrammableconnectivity.OperatorAPIPlan{
	// 		Properties: &armprogrammableconnectivity.OperatorAPIPlanProperties{
	// 			OperatorName: to.Ptr("csisc"),
	// 			CamaraAPIName: to.Ptr("vfugujiismnjtcwrzkdvxaj"),
	// 			SupportedLocations: []*string{
	// 				to.Ptr("m"),
	// 			},
	// 			OperatorRegions: []*string{
	// 				to.Ptr("licmuodzevtmowxdzzbstssmghl"),
	// 			},
	// 			Markets: []*string{
	// 				to.Ptr("yncxinbbsipilayubuq"),
	// 			},
	// 			Limits: to.Ptr("tdcuvjvnyctdlcxebcrcwquq"),
	// 			MarketplaceProperties: &armprogrammableconnectivity.MarketplaceProperties{
	// 				OfferID: to.Ptr("zlhnteqphbk"),
	// 				LegacyOfferID: to.Ptr("x"),
	// 				PublisherID: to.Ptr("zxlhmlhmkhuqggwmzyxg"),
	// 				PlanID: to.Ptr("ukdqvgtelddrssctjfbqjospcshfcg"),
	// 				TermID: to.Ptr("uuwqguxsrfnl"),
	// 			},
	// 			ProvisioningState: to.Ptr(armprogrammableconnectivity.ProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/providers/Microsoft.ProgrammableConnectivity/operatorApiPlans/livmzrh"),
	// 		Name: to.Ptr("bfrnaknndvhdkmaqgfwcmnxjsjtg"),
	// 		Type: to.Ptr("hclykzxrxyojkpfxnckjkcndlf"),
	// 		SystemData: &armprogrammableconnectivity.SystemData{
	// 			CreatedBy: to.Ptr("kuprrapuolhnvju"),
	// 			CreatedByType: to.Ptr(armprogrammableconnectivity.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T16:41:38.838Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("lsmrhxnvkpmrxncylgqpkr"),
	// 			LastModifiedByType: to.Ptr(armprogrammableconnectivity.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T16:41:38.838Z"); return t}()),
	// 		},
	// 	},
	// }
}
