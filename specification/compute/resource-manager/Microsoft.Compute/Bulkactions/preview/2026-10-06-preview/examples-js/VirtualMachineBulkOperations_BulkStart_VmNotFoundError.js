const { ComputeClient } = require("@azure/arm-compute-bulkactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to start one or more virtual machines. Bulk Actions begins processing the request immediately and returns a Bulk Action Operation Id for each virtual machine. Use the returned IDs to get operation status updates.
 *
 * @summary start one or more virtual machines. Bulk Actions begins processing the request immediately and returns a Bulk Action Operation Id for each virtual machine. Use the returned IDs to get operation status updates.
 * x-ms-original-file: 2026-10-06-preview/VirtualMachineBulkOperations_BulkStart_VmNotFoundError.json
 */
async function _04ResponseWhenAVirtualMachineDoesNotExist() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ComputeClient(credential, subscriptionId);
  const result = await client.virtualMachineBulkOperations.bulkStartOperation(
    "example-rg",
    "eastus",
    {
      executionParameters: {},
      resources: {
        ids: [
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01",
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/missing-vm",
        ],
      },
    },
  );
  console.log(result);
}
