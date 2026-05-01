const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Create_WithApplicationProfile.json
 */
async function createAVmWithApplicationProfile() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "myVM", {
    location: "westus",
    hardwareProfile: { vmSize: "Standard_D1_v2" },
    storageProfile: {
      imageReference: {
        sku: "{image_sku}",
        publisher: "{image_publisher}",
        version: "latest",
        offer: "{image_offer}",
      },
      osDisk: {
        caching: "ReadWrite",
        managedDisk: { storageAccountType: "Standard_LRS" },
        name: "myVMosdisk",
        createOption: "FromImage",
      },
    },
    applicationProfile: {
      galleryApplications: [
        {
          tags: "myTag1",
          order: 1,
          packageReferenceId:
            "/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdb/resourceGroups/myresourceGroupName2/providers/Microsoft.Compute/galleries/myGallery1/applications/MyApplication1/versions/1.0",
          configurationReference:
            "https://mystorageaccount.blob.core.windows.net/configurations/settings.config",
          treatFailureAsDeploymentFailure: false,
          enableAutomaticUpgrade: false,
        },
        {
          packageReferenceId:
            "/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdg/resourceGroups/myresourceGroupName3/providers/Microsoft.Compute/galleries/myGallery2/applications/MyApplication2/versions/1.1",
        },
      ],
    },
    osProfile: {
      adminUsername: "{your-username}",
      computerName: "myVM",
      adminPassword: "{your-password}",
    },
    networkProfile: {
      networkInterfaces: [
        {
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
          primary: true,
        },
      ],
    },
  });
  console.log(result);
}
