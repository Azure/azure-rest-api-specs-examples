package armlabservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4c2cdccf6ca3281dd50ed8788ce1de2e0d480973/specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Images/putImage.json
func ExampleImagesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlabservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewImagesClient().CreateOrUpdate(ctx, "testrg123", "testlabplan", "image1", armlabservices.Image{
		Properties: &armlabservices.ImageProperties{
			EnabledState: to.Ptr(armlabservices.EnableStateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Image = armlabservices.Image{
	// 	Name: to.Ptr("image1"),
	// 	Type: to.Ptr("Microsoft.LabServices/Image"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.LabServices/labPlans/testlabplan/images/image1"),
	// 	Properties: &armlabservices.ImageProperties{
	// 		EnabledState: to.Ptr(armlabservices.EnableStateEnabled),
	// 		Description: to.Ptr("A description of the image"),
	// 		DisplayName: to.Ptr("Windows Server 2019 Datacenter"),
	// 		Offer: to.Ptr("WindowsServer"),
	// 		OSType: to.Ptr(armlabservices.OsTypeWindows),
	// 		ProvisioningState: to.Ptr(armlabservices.ProvisioningStateSucceeded),
	// 		Publisher: to.Ptr("Microsoft"),
	// 		SKU: to.Ptr("2019-Datacenter"),
	// 		Version: to.Ptr("2019.0.20190410"),
	// 	},
	// 	SystemData: &armlabservices.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T10:00:00.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("identity123"),
	// 		CreatedByType: to.Ptr(armlabservices.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T09:12:28.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identity123"),
	// 		LastModifiedByType: to.Ptr(armlabservices.CreatedByTypeUser),
	// 	},
	// }
}
