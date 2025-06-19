package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: 2025-03-01/FlexComponents_Get_MaximumSet_Gen.json
func ExampleFlexComponentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFlexComponentsClient().Get(ctx, "eastus", "flexComponent", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armoracledatabase.FlexComponentsClientGetResponse{
	// 	FlexComponent: &armoracledatabase.FlexComponent{
	// 		Properties: &armoracledatabase.FlexComponentProperties{
	// 			MinimumCoreCount: to.Ptr[int32](16),
	// 			AvailableCoreCount: to.Ptr[int32](11),
	// 			AvailableDbStorageInGbs: to.Ptr[int32](8),
	// 			RuntimeMinimumCoreCount: to.Ptr[int32](13),
	// 			Shape: to.Ptr("Exadata.X11M"),
	// 			AvailableMemoryInGbs: to.Ptr[int32](15),
	// 			AvailableLocalStorageInGbs: to.Ptr[int32](13),
	// 			ComputeModel: to.Ptr("ECPU"),
	// 			HardwareType: to.Ptr(armoracledatabase.HardwareTypeCOMPUTE),
	// 			DescriptionSummary: to.Ptr("The description summary for this Flex Component"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/flexComponents/name"),
	// 		Name: to.Ptr("fc"),
	// 		Type: to.Ptr("Oracle.Database/flexComponents"),
	// 		SystemData: &armoracledatabase.SystemData{
	// 			CreatedBy: to.Ptr("ilrpjodjmvzhybazxipoplnql"),
	// 			CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-09T21:02:12.592Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("lhjbxchqkaia"),
	// 			LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-09T21:02:12.592Z"); return t}()),
	// 		},
	// 	},
	// }
}
