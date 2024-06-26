package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Event_Get.json
func ExampleEventClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEventClient().Get(ctx, "rgrecoveryservicesdatareplication", "4", "231CIG", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EventModel = armrecoveryservicesdatareplication.EventModel{
	// 	Name: to.Ptr("j"),
	// 	Type: to.Ptr("lgk"),
	// 	ID: to.Ptr("fgbppsytxctsrsxlfugyohhu"),
	// 	Properties: &armrecoveryservicesdatareplication.EventModelProperties{
	// 		Description: to.Ptr("hkeiogebluvfpdgxogwesjmtbbok"),
	// 		CorrelationID: to.Ptr("lwfsxforxnhvpmheujutjicflmxv"),
	// 		CustomProperties: &armrecoveryservicesdatareplication.EventModelCustomProperties{
	// 			InstanceType: to.Ptr("EventModelCustomProperties"),
	// 		},
	// 		EventName: to.Ptr("s"),
	// 		EventType: to.Ptr("npumqmvspm"),
	// 		HealthErrors: []*armrecoveryservicesdatareplication.HealthErrorModel{
	// 			{
	// 				AffectedResourceCorrelationIDs: []*string{
	// 					to.Ptr("fope")},
	// 					AffectedResourceType: to.Ptr("scfniv"),
	// 					Category: to.Ptr("leigw"),
	// 					Causes: to.Ptr("xznphqrrmsdzm"),
	// 					ChildErrors: []*armrecoveryservicesdatareplication.InnerHealthErrorModel{
	// 						{
	// 							Category: to.Ptr("lcsdxrqxquke"),
	// 							Causes: to.Ptr("kefaugkpxjkpulimjthjnl"),
	// 							Code: to.Ptr("yuxxpblihirpedwkigywgwjjrlzq"),
	// 							CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:52.128Z"); return t}()),
	// 							HealthCategory: to.Ptr("mhdgfjqwbikhxmhtomkl"),
	// 							IsCustomerResolvable: to.Ptr(true),
	// 							Message: to.Ptr("sskcei"),
	// 							Recommendation: to.Ptr("kqybwaesqumywtjepi"),
	// 							Severity: to.Ptr("wqxxiuaqjyagq"),
	// 							Source: to.Ptr("wevvftugwydzzw"),
	// 							Summary: to.Ptr("djsmgrltruljo"),
	// 					}},
	// 					Code: to.Ptr("dgxkefzmeukd"),
	// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:52.128Z"); return t}()),
	// 					HealthCategory: to.Ptr("itc"),
	// 					IsCustomerResolvable: to.Ptr(true),
	// 					Message: to.Ptr("lbywtdprdqdekl"),
	// 					Recommendation: to.Ptr("gmssteizlhjtclyeoo"),
	// 					Severity: to.Ptr("vvdajssdcypewdyechilxjmuijvdd"),
	// 					Source: to.Ptr("iy"),
	// 					Summary: to.Ptr("jtooblbvaxxrvcwgscbobq"),
	// 			}},
	// 			ResourceName: to.Ptr("yhpkowkbvtqnbiklnjzc"),
	// 			ResourceType: to.Ptr("surgdzezskgregozynvlinfutyh"),
	// 			Severity: to.Ptr("sjous"),
	// 			TimeOfOccurrence: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.432Z"); return t}()),
	// 		},
	// 		SystemData: &armrecoveryservicesdatareplication.EventModelSystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.432Z"); return t}()),
	// 			CreatedBy: to.Ptr("uske"),
	// 			CreatedByType: to.Ptr("luzowppyxjalugkef"),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.432Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ufrixpmhben"),
	// 			LastModifiedByType: to.Ptr("aubgraubkuaeipwzvbcgnlpseobx"),
	// 		},
	// 	}
}
