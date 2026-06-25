const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Create_CustomImageVmFromAnUnmanagedGeneralizedOsImage.json
 */
async function createACustomImageVmFromAnUnmanagedGeneralizedOsImage() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "{vm-name}", {
    location: "westus",
    hardwareProfile: { vmSize: "Standard_D1_v2" },
    storageProfile: {
      osDisk: {
        name: "myVMosdisk",
        image: {
          uri: "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/{existing-generalized-os-image-blob-name}.vhd",
        },
        osType: "Windows",
        createOption: "FromImage",
        caching: "ReadWrite",
        vhd: {
          uri: "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd",
        },
      },
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
