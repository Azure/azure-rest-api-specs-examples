package armvisualstudio_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/visualstudio/armvisualstudio"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/GetOperations.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvisualstudio.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().List(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationListResult = armvisualstudio.OperationListResult{
	// 	Value: []*armvisualstudio.Operation{
	// 		{
	// 			Name: to.Ptr("Microsoft.VisualStudio/Account/Write"),
	// 			Display: &armvisualstudio.OperationProperties{
	// 				Description: to.Ptr("Set Account"),
	// 				Operation: to.Ptr("Creates or updates the Account"),
	// 				Provider: to.Ptr("Visual Studio"),
	// 				Resource: to.Ptr("Account"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.VisualStudio/Account/Delete"),
	// 			Display: &armvisualstudio.OperationProperties{
	// 				Description: to.Ptr("Delete Account"),
	// 				Operation: to.Ptr("Deletes the Account"),
	// 				Provider: to.Ptr("Visual Studio"),
	// 				Resource: to.Ptr("Account"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.VisualStudio/Account/Read"),
	// 			Display: &armvisualstudio.OperationProperties{
	// 				Description: to.Ptr("Read Account"),
	// 				Operation: to.Ptr("Reads the Account"),
	// 				Provider: to.Ptr("Visual Studio"),
	// 				Resource: to.Ptr("Account"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.VisualStudio/Project/Write"),
	// 			Display: &armvisualstudio.OperationProperties{
	// 				Description: to.Ptr("Set Project"),
	// 				Operation: to.Ptr("Creates or updates the Project"),
	// 				Provider: to.Ptr("Visual Studio"),
	// 				Resource: to.Ptr("Project"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.VisualStudio/Project/Delete"),
	// 			Display: &armvisualstudio.OperationProperties{
	// 				Description: to.Ptr("Delete Project"),
	// 				Operation: to.Ptr("Deletes the Project"),
	// 				Provider: to.Ptr("Visual Studio"),
	// 				Resource: to.Ptr("Project"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.VisualStudio/Project/Read"),
	// 			Display: &armvisualstudio.OperationProperties{
	// 				Description: to.Ptr("Read Project"),
	// 				Operation: to.Ptr("Reads the Project"),
	// 				Provider: to.Ptr("Visual Studio"),
	// 				Resource: to.Ptr("Project"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.VisualStudio/Extension/Write"),
	// 			Display: &armvisualstudio.OperationProperties{
	// 				Description: to.Ptr("Set Extension"),
	// 				Operation: to.Ptr("Creates or updates the Extension"),
	// 				Provider: to.Ptr("Visual Studio"),
	// 				Resource: to.Ptr("Extension"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.VisualStudio/Extension/Delete"),
	// 			Display: &armvisualstudio.OperationProperties{
	// 				Description: to.Ptr("Delete Extension"),
	// 				Operation: to.Ptr("Deletes the Extension"),
	// 				Provider: to.Ptr("Visual Studio"),
	// 				Resource: to.Ptr("Extension"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.VisualStudio/Extension/Read"),
	// 			Display: &armvisualstudio.OperationProperties{
	// 				Description: to.Ptr("Read Extension"),
	// 				Operation: to.Ptr("Reads the Extension"),
	// 				Provider: to.Ptr("Visual Studio"),
	// 				Resource: to.Ptr("Extension"),
	// 			},
	// 	}},
	// }
}
