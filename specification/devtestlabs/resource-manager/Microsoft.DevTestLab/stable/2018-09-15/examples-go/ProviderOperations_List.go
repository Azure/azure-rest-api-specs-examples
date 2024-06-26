package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ProviderOperations_List.json
func ExampleProviderOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProviderOperationsClient().NewListPager(nil)
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
		// page.ProviderOperationResult = armdevtestlabs.ProviderOperationResult{
		// 	Value: []*armdevtestlabs.OperationMetadata{
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/register/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Registers the subscription"),
		// 				Operation: to.Ptr("Register Subscription"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("register"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/policySets/EvaluatePolicies/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Evaluates lab policy."),
		// 				Operation: to.Ptr("Evaluate policy"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("policy sets"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete labs."),
		// 				Operation: to.Ptr("Delete labs."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Labs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read labs."),
		// 				Operation: to.Ptr("Read labs."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Labs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify labs."),
		// 				Operation: to.Ptr("Add or modify labs."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Labs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/ListVhds/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("List disk images available for custom image creation."),
		// 				Operation: to.Ptr("List VHDs"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Labs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/GenerateUploadUri/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Generate a URI for uploading custom disk images to a Lab."),
		// 				Operation: to.Ptr("Generate image upload URI"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Labs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/CreateEnvironment/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Create virtual machines in a lab."),
		// 				Operation: to.Ptr("Create a virtual machine"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Labs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/ClaimAnyVm/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Claim a random claimable virtual machine in the lab."),
		// 				Operation: to.Ptr("Claim Any Virtual Machine"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Labs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/ExportResourceUsage/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Exports the lab resource usage into a storage account"),
		// 				Operation: to.Ptr("Exports the lab resource usage into a storage account"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Labs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/ImportVirtualMachine/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Import a virtual machine into a different lab."),
		// 				Operation: to.Ptr("Import a virtual machine"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Labs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/policySets/policies/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete policies."),
		// 				Operation: to.Ptr("Delete policies."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("policies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/policySets/policies/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read policies."),
		// 				Operation: to.Ptr("Read policies."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("policies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/policySets/policies/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify policies."),
		// 				Operation: to.Ptr("Add or modify policies."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("policies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/schedules/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete schedules."),
		// 				Operation: to.Ptr("Delete schedules."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/schedules/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read schedules."),
		// 				Operation: to.Ptr("Read schedules."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/schedules/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify schedules."),
		// 				Operation: to.Ptr("Add or modify schedules."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/schedules/Execute/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Execute a schedule."),
		// 				Operation: to.Ptr("Execute schedule"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/schedules/ListApplicable/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Lists all applicable schedules"),
		// 				Operation: to.Ptr("List all applicable schedules"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/schedules/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete schedules."),
		// 				Operation: to.Ptr("Delete schedules."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/schedules/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read schedules."),
		// 				Operation: to.Ptr("Read schedules."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/schedules/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify schedules."),
		// 				Operation: to.Ptr("Add or modify schedules."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/schedules/Execute/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Execute a schedule."),
		// 				Operation: to.Ptr("Execute schedule"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/schedules/Retarget/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Updates a schedule's target resource Id."),
		// 				Operation: to.Ptr("Retarget schedule."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/schedules/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete schedules."),
		// 				Operation: to.Ptr("Delete schedules."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/schedules/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read schedules."),
		// 				Operation: to.Ptr("Read schedules."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/schedules/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify schedules."),
		// 				Operation: to.Ptr("Add or modify schedules."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/schedules/Execute/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Execute a schedule."),
		// 				Operation: to.Ptr("Execute schedule"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/serviceFabrics/schedules/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete schedules."),
		// 				Operation: to.Ptr("Delete schedules."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/serviceFabrics/schedules/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read schedules."),
		// 				Operation: to.Ptr("Read schedules."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/serviceFabrics/schedules/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify schedules."),
		// 				Operation: to.Ptr("Add or modify schedules."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/serviceFabrics/schedules/Execute/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Execute a schedule."),
		// 				Operation: to.Ptr("Execute schedule"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("schedules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/artifactSources/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete artifact sources."),
		// 				Operation: to.Ptr("Delete artifact sources."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Artifact sources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/artifactSources/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read artifact sources."),
		// 				Operation: to.Ptr("Read artifact sources."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Artifact sources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/artifactSources/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify artifact sources."),
		// 				Operation: to.Ptr("Add or modify artifact sources."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Artifact sources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/artifactSources/artifacts/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read artifacts."),
		// 				Operation: to.Ptr("Read artifacts."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Artifacts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/artifactSources/artifacts/GenerateArmTemplate/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Generates an ARM template for the given artifact, uploads the required files to a storage account, and validates the generated artifact."),
		// 				Operation: to.Ptr("Generates an ARM template for the given artifact"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Artifacts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/galleryImages/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read gallery images."),
		// 				Operation: to.Ptr("Read gallery images."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("gallery images"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/customImages/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete custom images."),
		// 				Operation: to.Ptr("Delete custom images."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("custom images"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/customImages/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read custom images."),
		// 				Operation: to.Ptr("Read custom images."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("custom images"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/customImages/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify custom images."),
		// 				Operation: to.Ptr("Add or modify custom images."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("custom images"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualNetworks/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete virtual networks."),
		// 				Operation: to.Ptr("Delete virtual networks."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("virtual networks"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualNetworks/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read virtual networks."),
		// 				Operation: to.Ptr("Read virtual networks."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("virtual networks"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualNetworks/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify virtual networks."),
		// 				Operation: to.Ptr("Add or modify virtual networks."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("virtual networks"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete virtual machines."),
		// 				Operation: to.Ptr("Delete virtual machines."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read virtual machines."),
		// 				Operation: to.Ptr("Read virtual machines."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify virtual machines."),
		// 				Operation: to.Ptr("Add or modify virtual machines."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/Start/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Start a virtual machine."),
		// 				Operation: to.Ptr("Start virtual machine"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/Stop/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Stop a virtual machine"),
		// 				Operation: to.Ptr("Stop virtual machine"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/Restart/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Restart a virtual machine."),
		// 				Operation: to.Ptr("Restart virtual machine"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/Redeploy/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Redeploy a virtual machine"),
		// 				Operation: to.Ptr("Redeploy a virtual machine"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/Resize/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Resize Virtual Machine."),
		// 				Operation: to.Ptr("Resize Virtual Machine."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/ApplyArtifacts/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Apply artifacts to virtual machine."),
		// 				Operation: to.Ptr("Apply artifacts to virtual machine"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/AddDataDisk/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Attach a new or existing data disk to virtual machine."),
		// 				Operation: to.Ptr("Add or attach a data disk"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/DetachDataDisk/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Detach the specified disk from the virtual machine."),
		// 				Operation: to.Ptr("Detach the specified disk from the virtual machine"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/Claim/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Take ownership of an existing virtual machine"),
		// 				Operation: to.Ptr("Claim a virtual machine"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/UnClaim/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Release ownership of an existing virtual machine"),
		// 				Operation: to.Ptr("Unclaim a virtual machine"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/TransferDisks/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Transfer ownership of virtual machine data disks to yourself"),
		// 				Operation: to.Ptr("Transfer data disks to yourself"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/ListApplicableSchedules/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Lists all applicable schedules"),
		// 				Operation: to.Ptr("List all applicable schedules"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/GetRdpFileContents/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Gets a string that represents the contents of the RDP file for the virtual machine"),
		// 				Operation: to.Ptr("Get RDP file contents"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Virtual machines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/formulas/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete formulas."),
		// 				Operation: to.Ptr("Delete formulas."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Formulas"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/formulas/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read formulas."),
		// 				Operation: to.Ptr("Read formulas."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Formulas"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/formulas/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify formulas."),
		// 				Operation: to.Ptr("Add or modify formulas."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Formulas"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/costs/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read costs."),
		// 				Operation: to.Ptr("Read costs."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("costs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/costs/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify costs."),
		// 				Operation: to.Ptr("Add or modify costs."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("costs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/disks/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete disks."),
		// 				Operation: to.Ptr("Delete disks."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("disks"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/disks/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read disks."),
		// 				Operation: to.Ptr("Read disks."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("disks"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/disks/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify disks."),
		// 				Operation: to.Ptr("Add or modify disks."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("disks"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/disks/Attach/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Attach and create the lease of the disk to the virtual machine."),
		// 				Operation: to.Ptr("Attach disk"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("disks"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/disks/Detach/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Detach and break the lease of the disk attached to the virtual machine."),
		// 				Operation: to.Ptr("Detach and break the lease of the disk attached to the virtual machine"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("disks"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete user profiles."),
		// 				Operation: to.Ptr("Delete user profiles."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("user profiles"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read user profiles."),
		// 				Operation: to.Ptr("Read user profiles."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("user profiles"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify user profiles."),
		// 				Operation: to.Ptr("Add or modify user profiles."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("user profiles"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/notificationChannels/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete notification channels."),
		// 				Operation: to.Ptr("Delete notification channels."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("notificationChannels"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/notificationChannels/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read notification channels."),
		// 				Operation: to.Ptr("Read notification channels."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("notificationChannels"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/notificationChannels/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify notification channels."),
		// 				Operation: to.Ptr("Add or modify notification channels."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("notificationChannels"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/notificationChannels/Notify/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Send notification to provided channel."),
		// 				Operation: to.Ptr("Notify"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("notificationChannels"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/secrets/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete secrets."),
		// 				Operation: to.Ptr("Delete secrets."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("secrets"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/secrets/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read secrets."),
		// 				Operation: to.Ptr("Read secrets."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("secrets"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/secrets/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify secrets."),
		// 				Operation: to.Ptr("Add or modify secrets."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("secrets"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/locations/operations/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read operations."),
		// 				Operation: to.Ptr("Read operations."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("operations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/artifactSources/armTemplates/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read azure resource manager templates."),
		// 				Operation: to.Ptr("Read azure resource manager templates."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Azure Resource Manager templates"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/environments/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete environments."),
		// 				Operation: to.Ptr("Delete environments."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("environments"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/environments/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read environments."),
		// 				Operation: to.Ptr("Read environments."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("environments"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/environments/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify environments."),
		// 				Operation: to.Ptr("Add or modify environments."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("environments"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/serviceRunners/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete service runners."),
		// 				Operation: to.Ptr("Delete service runners."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Service runners"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/serviceRunners/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read service runners."),
		// 				Operation: to.Ptr("Read service runners."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Service runners"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/serviceRunners/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify service runners."),
		// 				Operation: to.Ptr("Add or modify service runners."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Service runners"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/serviceFabrics/delete"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Delete service fabrics."),
		// 				Operation: to.Ptr("Delete service fabrics."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Service Fabrics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/serviceFabrics/read"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Read service fabrics."),
		// 				Operation: to.Ptr("Read service fabrics."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Service Fabrics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/serviceFabrics/write"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Add or modify service fabrics."),
		// 				Operation: to.Ptr("Add or modify service fabrics."),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Service Fabrics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/serviceFabrics/Start/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Start a service fabric."),
		// 				Operation: to.Ptr("Start service fabric"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Service Fabrics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/serviceFabrics/Stop/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Stop a service fabric"),
		// 				Operation: to.Ptr("Stop service fabric"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Service Fabrics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevTestLab/labs/users/serviceFabrics/ListApplicableSchedules/action"),
		// 			Display: &armdevtestlabs.OperationMetadataDisplay{
		// 				Description: to.Ptr("Lists all applicable schedules"),
		// 				Operation: to.Ptr("List all applicable schedules"),
		// 				Provider: to.Ptr("Microsoft DevTest Labs"),
		// 				Resource: to.Ptr("Service Fabrics"),
		// 			},
		// 	}},
		// }
	}
}
