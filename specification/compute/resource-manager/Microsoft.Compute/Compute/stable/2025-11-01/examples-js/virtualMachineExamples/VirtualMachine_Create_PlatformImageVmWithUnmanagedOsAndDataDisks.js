const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Create_PlatformImageVmWithUnmanagedOsAndDataDisks.json
 */
async function createAPlatformImageVmWithUnmanagedOsAndDataDisks() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "{vm-name}", {
    location: "westus",
    hardwareProfile: { vmSize: "Standard_D2_v2" },
    storageProfile: {
      imageReference: {
        sku: "2016-Datacenter",
        publisher: "MicrosoftWindowsServer",
        version: "latest",
        offer: "WindowsServer",
      },
      osDisk: {
        caching: "ReadWrite",
        vhd: {
          uri: "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd",
        },
        createOption: "FromImage",
        name: "myVMosdisk",
      },
      dataDisks: [
        {
          diskSizeGB: 1023,
          createOption: "Empty",
          lun: 0,
          vhd: {
            uri: "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk0.vhd",
          },
        },
        {
          diskSizeGB: 1023,
          createOption: "Empty",
          lun: 1,
          vhd: {
            uri: "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk1.vhd",
          },
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
