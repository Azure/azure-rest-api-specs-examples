package armnetworkanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkanalytics/armnetworkanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21a8d55d74e4425e96d76e5835f52cfc9eb95a22/specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_ListRolesAssignments_MaximumSet_Gen.json
func ExampleDataProductsClient_ListRolesAssignments_dataProductsListRolesAssignmentsMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataProductsClient().ListRolesAssignments(ctx, "aoiresourceGroupName", "dataproduct01", map[string]any{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListRoleAssignments = armnetworkanalytics.ListRoleAssignments{
	// 	Count: to.Ptr[int32](1),
	// 	RoleAssignmentResponse: []*armnetworkanalytics.RoleAssignmentDetail{
	// 		{
	// 			DataTypeScope: []*string{
	// 				to.Ptr("scope")},
	// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-00000000000"),
	// 				PrincipalType: to.Ptr("User"),
	// 				Role: to.Ptr(armnetworkanalytics.DataProductUserRoleReader),
	// 				RoleAssignmentID: to.Ptr("00000000-0000-0000-0000-00000000000"),
	// 				RoleID: to.Ptr("00000000-0000-0000-0000-00000000000"),
	// 				UserName: to.Ptr("UserName"),
	// 		}},
	// 	}
}
