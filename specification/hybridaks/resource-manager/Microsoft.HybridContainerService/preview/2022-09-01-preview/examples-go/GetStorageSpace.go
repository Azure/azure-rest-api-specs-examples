package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/GetStorageSpace.json
func ExampleStorageSpacesClient_Retrieve() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStorageSpacesClient().Retrieve(ctx, "test-arcappliance-resgrp", "test-storage", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageSpaces = armhybridcontainerservice.StorageSpaces{
	// 	Name: to.Ptr("test-storage"),
	// 	Type: to.Ptr("microsoft.hybridcontainerservice/storageSpaces"),
	// 	ID: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.HybridContainerService/storageSpaces/test-storage"),
	// 	Location: to.Ptr("westus"),
	// 	ExtendedLocation: &armhybridcontainerservice.StorageSpacesExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armhybridcontainerservice.StorageSpacesProperties{
	// 		HciStorageProfile: &armhybridcontainerservice.StorageSpacesPropertiesHciStorageProfile{
	// 			MocGroup: to.Ptr("target-group"),
	// 			MocLocation: to.Ptr("MocLocation"),
	// 			MocStorageContainer: to.Ptr("WssdStorageContainer"),
	// 		},
	// 		ProvisioningState: to.Ptr(armhybridcontainerservice.ProvisioningStateSucceeded),
	// 	},
	// }
}
