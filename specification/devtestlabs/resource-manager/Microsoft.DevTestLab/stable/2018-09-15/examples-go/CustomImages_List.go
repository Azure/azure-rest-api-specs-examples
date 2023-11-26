package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/CustomImages_List.json
func ExampleCustomImagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomImagesClient().NewListPager("resourceGroupName", "{labName}", &armdevtestlabs.CustomImagesClientListOptions{Expand: nil,
		Filter:  nil,
		Top:     nil,
		Orderby: nil,
	})
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
		// page.CustomImageList = armdevtestlabs.CustomImageList{
		// 	Value: []*armdevtestlabs.CustomImage{
		// 		{
		// 			Name: to.Ptr("{customImageName}"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs/customImages"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/customimages/{customImageName}"),
		// 			Location: to.Ptr("{location}"),
		// 			Properties: &armdevtestlabs.CustomImageProperties{
		// 				Description: to.Ptr("My Custom Image"),
		// 				Author: to.Ptr("{authorName}"),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-10T09:59:28.798Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
		// 				VM: &armdevtestlabs.CustomImagePropertiesFromVM{
		// 					LinuxOsInfo: &armdevtestlabs.LinuxOsInfo{
		// 						LinuxOsState: to.Ptr(armdevtestlabs.LinuxOsStateNonDeprovisioned),
		// 					},
		// 					SourceVMID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualmachines/{vmName}"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
