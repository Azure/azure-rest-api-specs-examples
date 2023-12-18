package armbaremetalinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/baremetalinfrastructure/armbaremetalinfrastructure/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalStorageInstances_Create.json
func ExampleAzureBareMetalStorageInstancesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbaremetalinfrastructure.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAzureBareMetalStorageInstancesClient().Create(ctx, "myResourceGroup", "myAzureBareMetalStorageInstance", armbaremetalinfrastructure.AzureBareMetalStorageInstance{
		Location: to.Ptr("westus2"),
		Tags: map[string]*string{
			"key": to.Ptr("value"),
		},
		Properties: &armbaremetalinfrastructure.AzureBareMetalStorageInstanceProperties{
			AzureBareMetalStorageInstanceUniqueIdentifier: to.Ptr("23415635-4d7e-41dc-9598-8194f22c24e9"),
			StorageProperties: &armbaremetalinfrastructure.StorageProperties{
				Generation:        to.Ptr("Gen4"),
				HardwareType:      to.Ptr("NetApp"),
				OfferingType:      to.Ptr("EPIC"),
				ProvisioningState: to.Ptr(armbaremetalinfrastructure.ProvisioningStateSucceeded),
				StorageBillingProperties: &armbaremetalinfrastructure.StorageBillingProperties{
					AzureBareMetalStorageInstanceSize: to.Ptr(""),
					BillingMode:                       to.Ptr("PAYG"),
				},
				StorageType:  to.Ptr("FC"),
				WorkloadType: to.Ptr("ODB"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureBareMetalStorageInstance = armbaremetalinfrastructure.AzureBareMetalStorageInstance{
	// 	Name: to.Ptr("myAzureBareMetalStorageInstance"),
	// 	Type: to.Ptr("Microsoft.BareMetalInfrastructure/bareMetalStorageInstances"),
	// 	ID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.BareMetalInfrastructure/bareMetalStorageInstances/myAzureBareMetalStorageInstance"),
	// 	SystemData: &armbaremetalinfrastructure.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-20T23:10:22.682Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@microsoft.com"),
	// 		CreatedByType: to.Ptr(armbaremetalinfrastructure.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armbaremetalinfrastructure.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westus2"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armbaremetalinfrastructure.AzureBareMetalStorageInstanceProperties{
	// 		AzureBareMetalStorageInstanceUniqueIdentifier: to.Ptr("23415635-4d7e-41dc-9598-8194f22c24e9"),
	// 		StorageProperties: &armbaremetalinfrastructure.StorageProperties{
	// 			Generation: to.Ptr("Gen4"),
	// 			HardwareType: to.Ptr("NetApp"),
	// 			OfferingType: to.Ptr("EPIC"),
	// 			ProvisioningState: to.Ptr(armbaremetalinfrastructure.ProvisioningStateSucceeded),
	// 			StorageBillingProperties: &armbaremetalinfrastructure.StorageBillingProperties{
	// 				AzureBareMetalStorageInstanceSize: to.Ptr(""),
	// 				BillingMode: to.Ptr("PAYG"),
	// 			},
	// 			StorageType: to.Ptr("FC"),
	// 			WorkloadType: to.Ptr("ODB"),
	// 		},
	// 	},
	// }
}
