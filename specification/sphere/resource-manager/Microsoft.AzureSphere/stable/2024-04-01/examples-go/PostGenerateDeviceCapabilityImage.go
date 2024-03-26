package armsphere_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sphere/armsphere"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/09c37754dac91874ff689ed1e60effb4268c8669/specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/PostGenerateDeviceCapabilityImage.json
func ExampleDevicesClient_BeginGenerateCapabilityImage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDevicesClient().BeginGenerateCapabilityImage(ctx, "MyResourceGroup1", "MyCatalog1", "MyProduct1", "myDeviceGroup1", "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", armsphere.GenerateCapabilityImageRequest{
		Capabilities: []*armsphere.CapabilityType{
			to.Ptr(armsphere.CapabilityTypeApplicationDevelopment)},
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
	// res.SignedCapabilityImageResponse = armsphere.SignedCapabilityImageResponse{
	// 	Image: to.Ptr("TheDeviceCapabilityImage"),
	// }
}
