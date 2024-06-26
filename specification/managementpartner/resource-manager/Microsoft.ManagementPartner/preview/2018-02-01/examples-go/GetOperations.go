package armmanagementpartner_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementpartner/armmanagementpartner"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementpartner/resource-manager/Microsoft.ManagementPartner/preview/2018-02-01/examples/GetOperations.json
func ExampleOperationClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementpartner.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationClient().NewListPager(nil)
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
		// page.OperationList = armmanagementpartner.OperationList{
		// 	Value: []*armmanagementpartner.OperationResponse{
		// 		{
		// 			Name: to.Ptr("Microsoft.ManagementPartner/partners/read"),
		// 			Display: &armmanagementpartner.OperationDisplay{
		// 				Description: to.Ptr("Read All ManagementPartner"),
		// 				Operation: to.Ptr("Get ManagementPartner"),
		// 				Provider: to.Ptr("Microsoft ManagementPartner"),
		// 				Resource: to.Ptr("ManagementPartner"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ManagementPartner/partners/write"),
		// 			Display: &armmanagementpartner.OperationDisplay{
		// 				Description: to.Ptr("Create any ManagementPartner"),
		// 				Operation: to.Ptr("Create ManagementPartner"),
		// 				Provider: to.Ptr("Microsoft ManagementPartner"),
		// 				Resource: to.Ptr("ManagementPartner"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ManagementPartner/partners/action"),
		// 			Display: &armmanagementpartner.OperationDisplay{
		// 				Description: to.Ptr("Update any ManagementPartner"),
		// 				Operation: to.Ptr("Update ManagementPartner"),
		// 				Provider: to.Ptr("Microsoft ManagementPartner"),
		// 				Resource: to.Ptr("ManagementPartner"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ManagementPartner/partners/delete"),
		// 			Display: &armmanagementpartner.OperationDisplay{
		// 				Description: to.Ptr("Delete any ManagementPartner"),
		// 				Operation: to.Ptr("Delete ManagementPartner"),
		// 				Provider: to.Ptr("Microsoft ManagementPartner"),
		// 				Resource: to.Ptr("ManagementPartner"),
		// 			},
		// 	}},
		// }
	}
}
