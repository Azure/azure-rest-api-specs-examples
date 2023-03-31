package armoperationsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationsmanagement/armoperationsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/ManagementAssociationCreate.json
func ExampleManagementAssociationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementAssociationsClient().CreateOrUpdate(ctx, "rg1", "providerName", "resourceType", "resourceName", "managementAssociation1", armoperationsmanagement.ManagementAssociation{
		Location: to.Ptr("East US"),
		Properties: &armoperationsmanagement.ManagementAssociationProperties{
			ApplicationID: to.Ptr("/subscriptions/sub1/resourcegroups/rg1/providers/Microsoft.Appliance/Appliances/appliance1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagementAssociation = armoperationsmanagement.ManagementAssociation{
	// 	Name: to.Ptr("managementAssociation1"),
	// 	Type: to.Ptr("Microsoft.OperationsManagement/ManagementAssociations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.OperationalInsights/workspaces/ws1/Microsoft.OperationsManagement/ManagementAssociations/managementAssociation1"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armoperationsmanagement.ManagementAssociationProperties{
	// 		ApplicationID: to.Ptr("/subscriptions/sub1/resourcegroups/rg1/providers/Microsoft.Appliance/Appliances/appliance1"),
	// 	},
	// }
}
