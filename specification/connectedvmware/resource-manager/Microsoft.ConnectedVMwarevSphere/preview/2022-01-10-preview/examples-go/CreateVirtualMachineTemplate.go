package armconnectedvmware_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/CreateVirtualMachineTemplate.json
func ExampleVirtualMachineTemplatesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineTemplatesClient().BeginCreate(ctx, "testrg", "WebFrontEndTemplate", armconnectedvmware.VirtualMachineTemplate{
		ExtendedLocation: &armconnectedvmware.ExtendedLocation{
			Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
			Type: to.Ptr("customLocation"),
		},
		Location: to.Ptr("East US"),
		Properties: &armconnectedvmware.VirtualMachineTemplateProperties{
			MoRefID:   to.Ptr("aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
			VCenterID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter"),
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
	// res.VirtualMachineTemplate = armconnectedvmware.VirtualMachineTemplate{
	// 	Name: to.Ptr("WebFrontEndTemplate"),
	// 	Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VirtualMachineTemplates"),
	// 	ExtendedLocation: &armconnectedvmware.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VirtualMachineTemplates/WebFrontEndTemplate"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armconnectedvmware.VirtualMachineTemplateProperties{
	// 		MoRefID: to.Ptr("aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		VCenterID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter"),
	// 	},
	// }
}
