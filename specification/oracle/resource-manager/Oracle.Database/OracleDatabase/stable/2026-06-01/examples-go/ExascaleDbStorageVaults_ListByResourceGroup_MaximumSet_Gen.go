package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v3"
)

// Generated from example definition: 2026-06-01/ExascaleDbStorageVaults_ListByResourceGroup_MaximumSet_Gen.json
func ExampleExascaleDbStorageVaultsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExascaleDbStorageVaultsClient().NewListByResourceGroupPager("rgopenapi", nil)
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
		// page = armoracledatabase.ExascaleDbStorageVaultsClientListByResourceGroupResponse{
		// 	ExascaleDbStorageVaultListResult: armoracledatabase.ExascaleDbStorageVaultListResult{
		// 		Value: []*armoracledatabase.ExascaleDbStorageVault{
		// 			{
		// 				Properties: &armoracledatabase.ExascaleDbStorageVaultProperties{
		// 					AdditionalFlashCacheInPercent: to.Ptr[int32](0),
		// 					Description: to.Ptr("example"),
		// 					DisplayName: to.Ptr("resource1"),
		// 					HighCapacityDatabaseStorageInput: &armoracledatabase.ExascaleDbStorageInputDetails{
		// 						TotalSizeInGbs: to.Ptr[int32](24),
		// 					},
		// 					HighCapacityDatabaseStorage: &armoracledatabase.ExascaleDbStorageDetails{
		// 						AvailableSizeInGbs: to.Ptr[int32](26),
		// 						TotalSizeInGbs: to.Ptr[int32](18),
		// 					},
		// 					TimeZone: to.Ptr("2026-06-01T00:00:00Z"),
		// 					ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
		// 					LifecycleState: to.Ptr(armoracledatabase.ExascaleDbStorageVaultLifecycleStateProvisioning),
		// 					LifecycleDetails: to.Ptr("tuedx"),
		// 					VMClusterCount: to.Ptr[int32](9),
		// 					Ocid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					OciURL: to.Ptr("https://microsoft.com/aak"),
		// 					ExadataInfrastructureID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
		// 					AttachedShapeAttributes: []*armoracledatabase.ShapeAttribute{
		// 						to.Ptr(armoracledatabase.ShapeAttributeSMARTSTORAGE),
		// 						to.Ptr(armoracledatabase.ShapeAttributeSMARTSTORAGE),
		// 					},
		// 					IsAutoscaleEnabled: to.Ptr(true),
		// 					AutoscaleLimitInGbs: to.Ptr[int32](10),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("zsw"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key4308": to.Ptr("example"),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/resources/resource1"),
		// 				Name: to.Ptr("nnbi"),
		// 				Type: to.Ptr("Oracle.Database/resource"),
		// 				SystemData: &armoracledatabase.SystemData{
		// 					CreatedBy: to.Ptr("ns"),
		// 					CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2026, time.July, 28, 21, 47, 8, 141000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("example"),
		// 					LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2026, time.July, 28, 21, 47, 8, 141000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
