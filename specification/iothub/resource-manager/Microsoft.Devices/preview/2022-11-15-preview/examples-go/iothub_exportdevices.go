package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/iothub/resource-manager/Microsoft.Devices/preview/2022-11-15-preview/examples/iothub_exportdevices.json
func ExampleResourceClient_ExportDevices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceClient().ExportDevices(ctx, "myResourceGroup", "testHub", armiothub.ExportDevicesRequest{
		AuthenticationType:     to.Ptr(armiothub.AuthenticationTypeIdentityBased),
		ExcludeKeys:            to.Ptr(true),
		ExportBlobContainerURI: to.Ptr("testBlob"),
		Identity: &armiothub.ManagedIdentity{
			UserAssignedIdentity: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobResponse = armiothub.JobResponse{
	// 	Type: to.Ptr(armiothub.JobTypeUnknown),
	// 	EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 15 Jun 2017 19:20:58 GMT"); return t}()),
	// 	JobID: to.Ptr("test"),
	// 	StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 15 Jun 2017 19:20:58 GMT"); return t}()),
	// 	Status: to.Ptr(armiothub.JobStatusUnknown),
	// }
}
