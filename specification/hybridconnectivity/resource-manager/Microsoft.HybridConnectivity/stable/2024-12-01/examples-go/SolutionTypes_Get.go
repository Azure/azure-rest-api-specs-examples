package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2024-12-01/SolutionTypes_Get.json
func ExampleSolutionTypesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSolutionTypesClient("5ACC4579-DB34-4C2F-8F8C-25061168F342").Get(ctx, "rgpublicCloud", "lulzqllpu", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.SolutionTypesClientGetResponse{
	// 	SolutionTypeResource: &armhybridconnectivity.SolutionTypeResource{
	// 		Properties: &armhybridconnectivity.SolutionTypeProperties{
	// 			SolutionType: to.Ptr("tjtoeycxhyqxtgd"),
	// 			Description: to.Ptr("wxyxcvtzuxgodtlanjevedwfdwnznc"),
	// 			SupportedAzureRegions: []*string{
	// 				to.Ptr("cimocdh"),
	// 			},
	// 			SolutionSettings: []*armhybridconnectivity.SolutionTypeSettingsProperties{
	// 				{
	// 					Name: to.Ptr("tepghdgbefujhnnue"),
	// 					DisplayName: to.Ptr("mwlzepoin"),
	// 					Type: to.Ptr("je"),
	// 					Description: to.Ptr("soq"),
	// 					AllowedValues: []*string{
	// 						to.Ptr("pwizyngpkpxsllpluffjspx"),
	// 					},
	// 					DefaultValue: to.Ptr("laekyetgapdpxyqervqaqfscfwagek"),
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/testSubcrptions/resourceGroups/testResourceGroup/providers/Microsoft.HybridConnectivity/solutionTypes/j"),
	// 		Name: to.Ptr("xczyyxuphhacyyj"),
	// 		Type: to.Ptr("mf"),
	// 		SystemData: &armhybridconnectivity.SystemData{
	// 			CreatedBy: to.Ptr("rpxzkcrobprrdvuoqxz"),
	// 			CreatedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("jidegyskxi"),
	// 			LastModifiedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
	// 		},
	// 	},
	// }
}
