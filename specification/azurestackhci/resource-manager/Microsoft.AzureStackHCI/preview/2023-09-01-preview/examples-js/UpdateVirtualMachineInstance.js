const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update a virtual machine instance.
 *
 * @summary The operation to update a virtual machine instance.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/UpdateVirtualMachineInstance.json
 */
async function updateVirtualMachine() {
  const resourceUri =
    "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM";
  const virtualMachineInstance = {
    properties: {
      storageProfile: {
        dataDisks: [
          {
            id: "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/virtualHardDisks/test-vhd",
          },
        ],
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential);
  const result = await client.virtualMachineInstances.beginUpdateAndWait(
    resourceUri,
    virtualMachineInstance
  );
  console.log(result);
}
