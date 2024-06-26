package armscvmm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.ResourceProviderOperationList = armscvmm.ResourceProviderOperationList{
		// 	Value: []*armscvmm.ResourceProviderOperation{
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/VmmServers/Read"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Gets/List the VmmServer resource data."),
		// 				Operation: to.Ptr("Gets/List VmmServer resources."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("VmmServers"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/VmmServers/Write"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Read VmmServer."),
		// 				Operation: to.Ptr("Create or Update VmmServer resource."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("VmmServers"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/VmmServers/Delete"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the VmmServer resource."),
		// 				Operation: to.Ptr("Deletes the VmmServer resource."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("VmmServers"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/Clouds/Read"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Gets/List the Cloud resource data."),
		// 				Operation: to.Ptr("Gets/List Cloud resources."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("Clouds"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/Clouds/Write"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Create or Update Cloud resource data."),
		// 				Operation: to.Ptr("Create or Update Cloud resource."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("Clouds"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/Clouds/Delete"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the Cloud resource."),
		// 				Operation: to.Ptr("Deletes the Cloud resource."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("Clouds"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/VirtualNetworks/Read"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Gets/List the VirtualNetwork resource data."),
		// 				Operation: to.Ptr("Gets/List VirtualNetwork resources."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("VirtualNetworks"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/VirtualNetworks/Write"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Create or Update VirtualNetwork resource data."),
		// 				Operation: to.Ptr("Create or Update VirtualNetwork resource."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("VirtualNetworks"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/VirtualNetworks/Delete"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the VirtualNetwork resource."),
		// 				Operation: to.Ptr("Deletes the VirtualNetwork resource."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("VirtualNetworks"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/VirtualMachineTemplates/Read"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Gets/List the VirtualMachineTemplate resource data."),
		// 				Operation: to.Ptr("Gets/List VirtualMachineTemplate resources."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("VirtualMachineTemplates"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/VirtualMachineTemplates/Write"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Create or Update VirtualMachineTemplate resource data."),
		// 				Operation: to.Ptr("Create or Update VirtualMachineTemplate resource."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("VirtualMachineTemplates"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/VirtualMachineTemplates/Delete"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the VirtualMachineTemplate resource."),
		// 				Operation: to.Ptr("Deletes the VirtualMachineTemplate resource."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("VirtualMachineTemplates"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/VirtualMachines/Read"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Gets/List the VirtualMachine resource data."),
		// 				Operation: to.Ptr("Gets/List VirtualMachine resources."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("VirtualMachines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/VirtualMachines/Write"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Create or Update VirtualMachine resource data."),
		// 				Operation: to.Ptr("Create or Update VirtualMachine resource."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("VirtualMachines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ScVmm/VirtualMachines/Delete"),
		// 			Display: &armscvmm.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the VirtualMachine resource."),
		// 				Operation: to.Ptr("Deletes the VirtualMachine resource."),
		// 				Provider: to.Ptr("Microsoft.ScVmm resource provider"),
		// 				Resource: to.Ptr("VirtualMachines"),
		// 			},
		// 	}},
		// }
	}
}
