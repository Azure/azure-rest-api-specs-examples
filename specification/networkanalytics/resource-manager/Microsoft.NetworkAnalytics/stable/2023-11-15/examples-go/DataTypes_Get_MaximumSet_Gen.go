package armnetworkanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkanalytics/armnetworkanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21a8d55d74e4425e96d76e5835f52cfc9eb95a22/specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataTypes_Get_MaximumSet_Gen.json
func ExampleDataTypesClient_Get_dataTypesGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataTypesClient().Get(ctx, "aoiresourceGroupName", "dataproduct01", "datatypename", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataType = armnetworkanalytics.DataType{
	// 	Name: to.Ptr("datatypename"),
	// 	Type: to.Ptr("Microsoft.NetworkAnalytics/DataProducts/DataTypes"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/aoiresourceGroupName/providers/Microsoft.NetworkAnalytics/dataProducts/dataproduct01/dataTypes/datatypename"),
	// 	SystemData: &armnetworkanalytics.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-04T08:26:27.150Z"); return t}()),
	// 		CreatedBy: to.Ptr("abc@micros.com"),
	// 		CreatedByType: to.Ptr(armnetworkanalytics.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-04T08:26:27.150Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("abc@micros.com"),
	// 		LastModifiedByType: to.Ptr(armnetworkanalytics.CreatedByTypeUser),
	// 	},
	// 	Properties: &armnetworkanalytics.DataTypeProperties{
	// 		DatabaseCacheRetention: to.Ptr[int32](23),
	// 		DatabaseRetention: to.Ptr[int32](6),
	// 		ProvisioningState: to.Ptr(armnetworkanalytics.ProvisioningStateSucceeded),
	// 		State: to.Ptr(armnetworkanalytics.DataTypeState("STARTED")),
	// 		StateReason: to.Ptr("state Reason"),
	// 		StorageOutputRetention: to.Ptr[int32](27),
	// 		VisualizationURL: to.Ptr("visualizationUrl"),
	// 	},
	// }
}
