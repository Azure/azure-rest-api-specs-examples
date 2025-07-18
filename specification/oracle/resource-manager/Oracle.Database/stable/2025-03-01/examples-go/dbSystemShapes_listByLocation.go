package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: 2025-03-01/dbSystemShapes_listByLocation.json
func ExampleDbSystemShapesClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDbSystemShapesClient().NewListByLocationPager("eastus", nil)
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
		// page = armoracledatabase.DbSystemShapesClientListByLocationResponse{
		// 	DbSystemShapeListResult: armoracledatabase.DbSystemShapeListResult{
		// 		Value: []*armoracledatabase.DbSystemShape{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/dbSystemShapes/EXADATA.X9M"),
		// 				Type: to.Ptr("Oracle.Database/locations/dbSystemShapes"),
		// 				Properties: &armoracledatabase.DbSystemShapeProperties{
		// 					ShapeFamily: to.Ptr("EXADATA"),
		// 					ShapeName: to.Ptr("EXADATA.X9M"),
		// 					AvailableCoreCount: to.Ptr[int32](100),
		// 					MinimumCoreCount: to.Ptr[int32](1),
		// 					RuntimeMinimumCoreCount: to.Ptr[int32](1),
		// 					CoreCountIncrement: to.Ptr[int32](1),
		// 					MinStorageCount: to.Ptr[int32](0),
		// 					MaxStorageCount: to.Ptr[int32](100),
		// 					AvailableDataStoragePerServerInTbs: to.Ptr[float64](100),
		// 					AvailableMemoryPerNodeInGbs: to.Ptr[int32](10),
		// 					AvailableDbNodePerNodeInGbs: to.Ptr[int32](10),
		// 					MinCoreCountPerNode: to.Ptr[int32](0),
		// 					AvailableMemoryInGbs: to.Ptr[int32](10),
		// 					MinMemoryPerNodeInGbs: to.Ptr[int32](0),
		// 					AvailableDbNodeStorageInGbs: to.Ptr[int32](10),
		// 					MinDbNodeStoragePerNodeInGbs: to.Ptr[int32](0),
		// 					AvailableDataStorageInTbs: to.Ptr[int32](10),
		// 					MinDataStorageInTbs: to.Ptr[int32](0),
		// 					MinimumNodeCount: to.Ptr[int32](0),
		// 					MaximumNodeCount: to.Ptr[int32](1000),
		// 					AvailableCoreCountPerNode: to.Ptr[int32](1000),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
