package armnetworkanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkanalytics/armnetworkanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21a8d55d74e4425e96d76e5835f52cfc9eb95a22/specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_RemoveUserRole_MinimumSet_Gen.json
func ExampleDataProductsClient_RemoveUserRole_dataProductsRemoveUserRoleMaximumSetGenGeneratedByMinimumSetRuleMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDataProductsClient().RemoveUserRole(ctx, "aoiresourceGroupName", "dataproduct01", armnetworkanalytics.RoleAssignmentDetail{
		DataTypeScope: []*string{
			to.Ptr("scope")},
		PrincipalID:      to.Ptr("00000000-0000-0000-0000-00000000000"),
		PrincipalType:    to.Ptr("User"),
		Role:             to.Ptr(armnetworkanalytics.DataProductUserRoleReader),
		RoleAssignmentID: to.Ptr("00000000-0000-0000-0000-00000000000"),
		RoleID:           to.Ptr("00000000-0000-0000-0000-00000000000"),
		UserName:         to.Ptr("UserName"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
