package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2027-01-01/SolutionTypes_ListByResourceGroup_MaximumSet_Gen.json
func ExampleSolutionTypesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSolutionTypesClient("5ACC4579-DB34-4C2F-8F8C-25061168F342").NewListByResourceGroupPager("rghybridconnectivity", nil)
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
		// page = armhybridconnectivity.SolutionTypesClientListByResourceGroupResponse{
		// 	SolutionTypeResourceListResult: armhybridconnectivity.SolutionTypeResourceListResult{
		// 		Value: []*armhybridconnectivity.SolutionTypeResource{
		// 			{
		// 				Properties: &armhybridconnectivity.SolutionTypeProperties{
		// 					SolutionType: to.Ptr("dembhpcydwoiyszmjtniletpy"),
		// 					SupportedAzureRegions: []*string{
		// 						to.Ptr("ujawbfint"),
		// 					},
		// 					SolutionSettings: []*armhybridconnectivity.SolutionTypeSettingsProperties{
		// 						{
		// 							Name: to.Ptr("i"),
		// 							DisplayName: to.Ptr("forzmqskffaub"),
		// 							Type: to.Ptr("d"),
		// 							AllowedValues: []*string{
		// 								to.Ptr("uecqnmmssdeusxejcxrtkskfugvl"),
		// 							},
		// 							DefaultValue: to.Ptr("uzwiymoxrummkoowwvzjhyazeavzr"),
		// 							Description: to.Ptr("zrvionyqmsb"),
		// 							HostTypes: []*armhybridconnectivity.HostType{
		// 								to.Ptr(armhybridconnectivity.HostTypeAWS),
		// 							},
		// 						},
		// 					},
		// 					Description: to.Ptr("fhgkvkmrwozqlhpzycasojtlbrbx"),
		// 					HostTypes: []*armhybridconnectivity.HostType{
		// 						to.Ptr(armhybridconnectivity.HostTypeAWS),
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/testSubcrptions/resourceGroups/testResourceGroup/providers/Microsoft.HybridConnectivity/solutionTypes/i"),
		// 				Name: to.Ptr("jtlxwihbuftmaobxfmfjojalhpwrv"),
		// 				Type: to.Ptr("zditfautattfhnffvjw"),
		// 				SystemData: &armhybridconnectivity.SystemData{
		// 					CreatedBy: to.Ptr("rpxzkcrobprrdvuoqxz"),
		// 					CreatedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2024, time.January, 18, 22, 52, 7, 890000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("jidegyskxi"),
		// 					LastModifiedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2024, time.January, 18, 22, 52, 7, 890000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
