package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry"
)

// Generated from example definition: 2024-11-01/Get_BillingContainer.json
func ExampleBillingContainersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBillingContainersClient().Get(ctx, "my-billingContainer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdeviceregistry.BillingContainersClientGetResponse{
	// 	BillingContainer: &armdeviceregistry.BillingContainer{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DeviceRegistry/billingContainers/my-billingContainer"),
	// 		Name: to.Ptr("my-billingContainer"),
	// 		Type: to.Ptr("Microsoft.DeviceRegistry/billingcontainers"),
	// 		Etag: to.Ptr("\"00001300-0000-0100-0000-6671f0170000\""),
	// 		SystemData: &armdeviceregistry.SystemData{
	// 			CreatedBy: to.Ptr("0e1b4448-67b9-46a1-8158-e3dade4c008e"),
	// 			CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-20T21:51:21.169954Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("0e1b4448-67b9-46a1-8158-e3dade4c008e"),
	// 			LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-18T20:37:41.9371846Z"); return t}()),
	// 		},
	// 		Properties: &armdeviceregistry.BillingContainerProperties{
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
