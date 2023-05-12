package armextendedlocation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fb9c8e2ca33e9723c2b2fc849f627329067feb54/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/CustomLocationsListOperations.json
func ExampleCustomLocationsClient_NewListOperationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomLocationsClient().NewListOperationsPager(nil)
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
		// page.CustomLocationOperationsList = armextendedlocation.CustomLocationOperationsList{
		// 	Value: []*armextendedlocation.CustomLocationOperation{
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/operations/read"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Gets list of Available Operations for Custom Locations"),
		// 				Operation: to.Ptr("List Available Operations for Custom Locations"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Operations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/register/action"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Registers the subscription for Custom Location resource provider and enables the creation of Custom Location."),
		// 				Operation: to.Ptr("Registers the Custom Location Resource Provider"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Locations Resource Provider"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/unregister/action"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("UnRegisters the subscription for Custom Location resource provider and disables the creation of Custom Location."),
		// 				Operation: to.Ptr("UnRegisters the Custom Location Resource Provider"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Locations Resource Provider"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/customLocations/read"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Gets an Custom Location resource"),
		// 				Operation: to.Ptr("Get Custom Location"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Locations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/customLocations/write"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Creates or Updates Custom Location resource"),
		// 				Operation: to.Ptr("Create or Update Custom Location"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Locations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/customLocations/deploy/action"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Deploy permissions to a Custom Location resource"),
		// 				Operation: to.Ptr("Deploy permissions to Custom Location"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Locations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/customLocations/delete"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Deletes Custom Location resource"),
		// 				Operation: to.Ptr("Delete Custom Location"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Locations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/customLocations/enabledresourcetypes/read"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Gets EnabledResourceTypes for a Custom Location resource"),
		// 				Operation: to.Ptr("Get EnabledResourceTypes for Custom Location"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Locations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/locations/operationsstatus/read"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Get result of Custom Location operation"),
		// 				Operation: to.Ptr("Get status of Custom Location operation"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Location Operation Status"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/locations/operationresults/read"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Get result of Custom Location operation"),
		// 				Operation: to.Ptr("Get the status of Custom Location operation"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Location Operation Result"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/customLocations/resourceSyncRules/read"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Gets a Resource Sync Rule resource"),
		// 				Operation: to.Ptr("Get Resource Sync Rule"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Resource Sync Rules"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/customLocations/resourceSyncRules/write"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Creates or Updates a Resource Sync Rule resource"),
		// 				Operation: to.Ptr("Create or Update Resource Sync Rule"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Resource Sync Rules"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/customLocations/resourceSyncRules/delete"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Deletes Resource Sync Rule resource"),
		// 				Operation: to.Ptr("Delete Resource Sync Rule"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Resource Sync Rules"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/customLocations/findTargetResourceGroup/action"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Evaluate Labels Against Resource Sync Rules to Get Target Resource Group"),
		// 				Operation: to.Ptr("Find Target Resource Group Action"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Resource Sync Rules"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}
