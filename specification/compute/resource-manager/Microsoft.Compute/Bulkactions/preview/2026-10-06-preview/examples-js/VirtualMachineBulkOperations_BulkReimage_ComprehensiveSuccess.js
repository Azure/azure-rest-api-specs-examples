const { ComputeClient } = require("@azure/arm-compute-bulkactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to this feature is currently in preview.
 *
 * Reimage one or more virtual machines. Reimaging is destructive and can replace operating system disk contents. Bulk Actions begins processing the request immediately and returns a Bulk Action Operation Id for each virtual machine. Use the returned IDs to get operation status updates.
 *
 * @summary this feature is currently in preview.
 *
 * Reimage one or more virtual machines. Reimaging is destructive and can replace operating system disk contents. Bulk Actions begins processing the request immediately and returns a Bulk Action Operation Id for each virtual machine. Use the returned IDs to get operation status updates.
 * x-ms-original-file: 2026-10-06-preview/VirtualMachineBulkOperations_BulkReimage_ComprehensiveSuccess.json
 */
async function _02ReimageVirtualMachinesWithSharedSettingsAndAPerVMOverride() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ComputeClient(credential, subscriptionId);
  const result = await client.virtualMachineBulkOperations.bulkReimageOperation(
    "example-rg",
    "eastus",
    {
      executionParameters: { retryPolicy: { retryWindowInMinutes: 30 } },
      resources: {
        ids: [
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01",
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-02",
        ],
      },
      reimageParameters: {
        baseProfile: {
          tempDisk: false,
          exactVersion: "1.0.0",
          osProfile: { customData: "I2Nsb3VkLWNvbmZpZwpwYWNrYWdlX3VwZ3JhZGU6IHRydWUK" },
        },
        resourceOverrides: [
          {
            resourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-02",
            profile: { tempDisk: false, exactVersion: "1.1.0" },
          },
        ],
      },
    },
  );
  console.log(result);
}
