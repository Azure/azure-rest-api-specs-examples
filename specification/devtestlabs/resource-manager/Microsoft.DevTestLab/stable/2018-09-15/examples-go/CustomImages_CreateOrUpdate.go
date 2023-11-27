package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/CustomImages_CreateOrUpdate.json
func ExampleCustomImagesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomImagesClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "{labName}", "{customImageName}", armdevtestlabs.CustomImage{
		Tags: map[string]*string{
			"tagName1": to.Ptr("tagValue1"),
		},
		Properties: &armdevtestlabs.CustomImageProperties{
			Description: to.Ptr("My Custom Image"),
			VM: &armdevtestlabs.CustomImagePropertiesFromVM{
				LinuxOsInfo: &armdevtestlabs.LinuxOsInfo{
					LinuxOsState: to.Ptr(armdevtestlabs.LinuxOsStateNonDeprovisioned),
				},
				SourceVMID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualmachines/{vmName}"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomImage = armdevtestlabs.CustomImage{
	// 	Name: to.Ptr("{customImageName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/customImages"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/customimages/{customImageName}"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.CustomImageProperties{
	// 		Description: to.Ptr("My Custom Image"),
	// 		Author: to.Ptr("{authorName}"),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-10T09:59:28.798Z"); return t}()),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 		VM: &armdevtestlabs.CustomImagePropertiesFromVM{
	// 			LinuxOsInfo: &armdevtestlabs.LinuxOsInfo{
	// 				LinuxOsState: to.Ptr(armdevtestlabs.LinuxOsStateNonDeprovisioned),
	// 			},
	// 			SourceVMID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualmachines/{vmName}"),
	// 		},
	// 	},
	// }
}
