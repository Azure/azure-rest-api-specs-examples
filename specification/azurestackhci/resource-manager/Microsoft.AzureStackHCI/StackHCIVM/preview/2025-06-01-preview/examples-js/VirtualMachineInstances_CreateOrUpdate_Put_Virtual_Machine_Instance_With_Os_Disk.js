const { AzureStackHCIVMManagementClient } = require("@azure/arm-azurestackhcivm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine instance. Please note some properties can be set only during virtual machine instance creation.
 *
 * @summary the operation to create or update a virtual machine instance. Please note some properties can be set only during virtual machine instance creation.
 * x-ms-original-file: 2025-06-01-preview/VirtualMachineInstances_CreateOrUpdate_Put_Virtual_Machine_Instance_With_Os_Disk.json
 */
async function putVirtualMachineInstanceWithOsDisk() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const client = new AzureStackHCIVMManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineInstances.createOrUpdate(
    "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM",
    {
      extendedLocation: {
        name: "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location",
        type: "CustomLocation",
      },
      properties: {
        hardwareProfile: { vmSize: "Default" },
        networkProfile: {
          networkInterfaces: [
            {
              id: "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/networkInterfaces/test-nic",
            },
          ],
        },
        securityProfile: {
          enableTPM: true,
          uefiSettings: { secureBootEnabled: true },
        },
        storageProfile: {
          osDisk: {
            id: "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/virtualHardDisks/test-vhd",
          },
          vmConfigStoragePathId:
            "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-container",
        },
      },
    },
  );
  console.log(result);
}
