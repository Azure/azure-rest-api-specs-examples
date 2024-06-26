package armconnectedvmware_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/ListOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationsList = armconnectedvmware.OperationsList{
		// 	Value: []*armconnectedvmware.Operation{
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/vcenters/Read"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Read vcenters"),
		// 				Operation: to.Ptr("Gets/List vcenters resources"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("vcenters"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/vcenters/Write"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Writes vcenters"),
		// 				Operation: to.Ptr("Create/update vcenters resources"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("vcenters"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/vcenters/Delete"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Deletes vcenters"),
		// 				Operation: to.Ptr("Deletes vcenters resource"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("vcenters"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/resourcepools/Read"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Read resourcepools"),
		// 				Operation: to.Ptr("Gets/List resourcepools resources"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("resourcepools"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/resourcepools/Write"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Writes resourcepools"),
		// 				Operation: to.Ptr("Create/update resourcepools resources"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("resourcepools"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/resourcepools/Delete"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Deletes resourcepools"),
		// 				Operation: to.Ptr("Deletes resourcepools resource"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("resourcepools"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/virtualmachines/Read"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Read virtualmachines"),
		// 				Operation: to.Ptr("Gets/List virtualmachines resources"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("virtualmachines"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/virtualmachines/Write"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Writes virtualmachines"),
		// 				Operation: to.Ptr("Create/update virtualmachines resources"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("virtualmachines"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/virtualmachines/Delete"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Deletes virtualmachines"),
		// 				Operation: to.Ptr("Deletes virtualmachines resource"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("virtualmachines"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/virtualmachinetemplates/Read"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Read virtualmachinetemplates"),
		// 				Operation: to.Ptr("Gets/List virtualmachinetemplates resources"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("virtualmachinetemplates"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/virtualmachinetemplates/Write"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Writes virtualmachinetemplates"),
		// 				Operation: to.Ptr("Create/update virtualmachinetemplates resources"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("virtualmachinetemplates"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/virtualmachinetemplates/Delete"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Deletes virtualmachinetemplates"),
		// 				Operation: to.Ptr("Deletes virtualmachinetemplates resource"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("virtualmachinetemplates"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/virtualnetworks/Read"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Read virtualnetworks"),
		// 				Operation: to.Ptr("Gets/List virtualnetworks resources"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("virtualnetworks"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/virtualnetworks/Write"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Writes virtualnetworks"),
		// 				Operation: to.Ptr("Create/update virtualnetworks resources"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("virtualnetworks"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/virtualnetworks/Delete"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Deletes virtualnetworks"),
		// 				Operation: to.Ptr("Deletes virtualnetworks resource"),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("virtualnetworks"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/arczones/deploy/action"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Deploy on arczone."),
		// 				Operation: to.Ptr("Deploy on arczone."),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("arczones"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/virtualmachinetemplates/clone/action"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Clone on virtual machine templates."),
		// 				Operation: to.Ptr("Clone on virtual machine templates."),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("virtualmachinetemplates"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/resourcepools/deploy/action"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Deploy on resource pool."),
		// 				Operation: to.Ptr("Deploy on resource pool."),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("resourcepools"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.vmware/virtualnetworks/join/action"),
		// 			Display: &armconnectedvmware.OperationDisplay{
		// 				Description: to.Ptr("Join virtual network."),
		// 				Operation: to.Ptr("Join virtual network."),
		// 				Provider: to.Ptr("microsoft.vmware"),
		// 				Resource: to.Ptr("virtualnetworks"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
