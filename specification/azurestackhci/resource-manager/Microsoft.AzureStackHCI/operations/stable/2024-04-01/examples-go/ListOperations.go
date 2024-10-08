package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c9b146c38df5f76e2d34a3ef771979805ff4ff73/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/operations/stable/2024-04-01/examples/ListOperations.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.OperationListResult = armazurestackhci.OperationListResult{
	// 	Value: []*armazurestackhci.Operation{
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Register/Action"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Registers the subscription for the Azure Stack HCI resource provider and enables the creation of Azure Stack HCI resources."),
	// 				Operation: to.Ptr("Registers the Azure Stack HCI Resource Provider"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Register"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Unregister/Action"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Unregisters the subscription for the Azure Stack HCI resource provider."),
	// 				Operation: to.Ptr("Unregisters the Azure Stack HCI Resource Provider"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Unregister"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Operations/Read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Gets operations"),
	// 				Operation: to.Ptr("Gets/List operations resources"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Operations"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/Read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Gets clusters"),
	// 				Operation: to.Ptr("Gets/List cluster resources"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/Write"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Creates or updates a cluster"),
	// 				Operation: to.Ptr("Create/update cluster resources"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/Delete"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Deletes cluster resource"),
	// 				Operation: to.Ptr("Deletes cluster resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/ArcSettings/Read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Gets arc resource of HCI cluster"),
	// 				Operation: to.Ptr("Gets/List arc resources"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/ArcSettings"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/ArcSettings/Write"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Create or updates arc resource of HCI cluster"),
	// 				Operation: to.Ptr("Create/Update arc resources"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/ArcSettings"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/ArcSettings/Delete"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Delete arc resource of HCI cluster"),
	// 				Operation: to.Ptr("Delete arc resources"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/ArcSettings"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/ArcSettings/Extensions/Read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Gets extension resource of HCI cluster"),
	// 				Operation: to.Ptr("Gets/List extension resources of HCI cluster"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/ArcSettings/Extensions"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/ArcSettings/Extensions/Write"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Create or update extension resource of HCI cluster"),
	// 				Operation: to.Ptr("Create/Update extension resources of HCI cluster"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/ArcSettings/Extensions"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/ArcSettings/Extensions/Delete"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Delete extension resources of HCI cluster"),
	// 				Operation: to.Ptr("Delete extension resources of HCI cluster"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/ArcSettings/Extensions"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualMachines/Restart/Action"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Restarts virtual machine resource"),
	// 				Operation: to.Ptr("Restarts virtual machine resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualMachines"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualMachines/Start/Action"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Starts virtual machine resource"),
	// 				Operation: to.Ptr("Starts virtual machine resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualMachines"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualMachines/Stop/Action"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Stops virtual machine resource"),
	// 				Operation: to.Ptr("Stops virtual machine resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualMachines"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualMachines/Delete"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Deletes virtual machine resource"),
	// 				Operation: to.Ptr("Deletes virtual machine resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualMachines"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualMachines/Write"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Creates/Updates virtual machine resource"),
	// 				Operation: to.Ptr("Creates/Updates virtual machine resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualMachines"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualMachines/Read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Gets/Lists virtual machine resource"),
	// 				Operation: to.Ptr("Gets/Lists virtual machine resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualMachines"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualNetworks/Delete"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Deletes virtual networks resource"),
	// 				Operation: to.Ptr("Deletes virtual networks resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualNetworks"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualNetworks/Write"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Creates/Updates virtual networks resource"),
	// 				Operation: to.Ptr("Creates/Updates virtual networks resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualNetworks"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualNetworks/Read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Gets/Lists virtual networks resource"),
	// 				Operation: to.Ptr("Gets/Lists virtual networks resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualNetworks"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualHardDisks/Delete"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Deletes virtual hard disk resource"),
	// 				Operation: to.Ptr("Deletes virtual hard disk resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualHardDisks"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualHardDisks/Write"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Creates/Updates virtual hard disk resource"),
	// 				Operation: to.Ptr("Creates/Updates virtual hard disk resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualHardDisks"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualHardDisks/Read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Gets/Lists virtual hard disk resource"),
	// 				Operation: to.Ptr("Gets/Lists virtual hard disk resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualHardDisks"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/NetworkInterfaces/Delete"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Deletes network interfaces resource"),
	// 				Operation: to.Ptr("Deletes network interfaces resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("NetworkInterfaces"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/NetworkInterfaces/Write"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Creates/Updates network interfaces resource"),
	// 				Operation: to.Ptr("Creates/Updates network interfaces resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("NetworkInterfaces"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/NetworkInterfaces/Read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Gets/Lists network interfaces resource"),
	// 				Operation: to.Ptr("Gets/Lists network interfaces resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("NetworkInterfaces"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/GalleryImages/Delete"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Deletes gallery images resource"),
	// 				Operation: to.Ptr("Deletes gallery images resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("GalleryImages"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/GalleryImages/Write"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Creates/Updates gallery images resource"),
	// 				Operation: to.Ptr("Creates/Updates gallery images resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("GalleryImages"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/GalleryImages/Read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Gets/Lists gallery images resource"),
	// 				Operation: to.Ptr("Gets/Lists gallery images resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("GalleryImages"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualMachines/HybridIdentityMetadata/Read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Gets/Lists virtual machine hybrid identity metadata proxy resource"),
	// 				Operation: to.Ptr("Gets/Lists virtual machine hybrid identity metadata proxy resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualMachines/HybridIdentityMetadata"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualMachines/Extensions/Read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Gets/Lists virtual machine extensions resource"),
	// 				Operation: to.Ptr("Gets/Lists virtual machine extensions resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualMachines/Extensions"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualMachines/Extensions/Write"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Creates/Updates virtual machine extensions resource"),
	// 				Operation: to.Ptr("Creates/Updates virtual machine extensions resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualMachines/Extensions"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/VirtualMachines/Extensions/Delete"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Deletes virtual machine extensions resource"),
	// 				Operation: to.Ptr("Deletes virtual machine extensions resource"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("VirtualMachines/Extensions"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/RegisteredSubscriptions/read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Reads registered subscriptions"),
	// 				Operation: to.Ptr("Gets/Lists registered subscriptions"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("RegisteredSubscriptions"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/Updates/Read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Gets available updates for HCI cluster"),
	// 				Operation: to.Ptr("Gets/List available updates for HCI cluster"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/Updates"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/Updates/Write"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Create or update updates resource of HCI cluster"),
	// 				Operation: to.Ptr("Create/Update updates resource of HCI cluster"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/Updates"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/Updates/Delete"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Delete updates resources of HCI cluster"),
	// 				Operation: to.Ptr("Delete updates resources of HCI cluster"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/Updates"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/UpdateSummaries/Read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Gets available update summaries for HCI cluster"),
	// 				Operation: to.Ptr("Gets/List available update summaries for HCI cluster"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/UpdateSummaries"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/UpdateSummaries/Write"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Create or update update summaries resource of HCI cluster"),
	// 				Operation: to.Ptr("Create/Update update summaries resource of HCI cluster"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/UpdateSummaries"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/UpdateSummaries/Delete"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Delete update summaries resources of HCI cluster"),
	// 				Operation: to.Ptr("Delete updates resource summaries of HCI cluster"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/UpdateSummaries"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/Updates/UpdateRuns/Read"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Gets available update runs for HCI cluster"),
	// 				Operation: to.Ptr("Gets/List available update runs for HCI cluster"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/Updates/UpdateRuns"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/Updates/UpdateRuns/Write"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Create or update update runs resource of HCI cluster"),
	// 				Operation: to.Ptr("Create/Update update runs resource of HCI cluster"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/Updates/UpdateRuns"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.AzureStackHCI/Clusters/Updates/UpdateRuns/Delete"),
	// 			Display: &armazurestackhci.OperationDisplay{
	// 				Description: to.Ptr("Delete update runs resources of HCI cluster"),
	// 				Operation: to.Ptr("Delete updates resource runs of HCI cluster"),
	// 				Provider: to.Ptr("Microsoft.AzureStackHCI"),
	// 				Resource: to.Ptr("Clusters/Updates/UpdateRuns"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 	}},
	// }
}
