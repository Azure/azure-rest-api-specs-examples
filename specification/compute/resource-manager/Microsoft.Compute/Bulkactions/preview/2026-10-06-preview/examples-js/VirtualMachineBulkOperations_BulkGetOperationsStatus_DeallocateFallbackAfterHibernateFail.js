const { ComputeClient } = require("@azure/arm-compute-bulkactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get the current status of one or more operations identified by their Bulk Action Operation Ids.
 *
 * @summary get the current status of one or more operations identified by their Bulk Action Operation Ids.
 * x-ms-original-file: 2026-10-06-preview/VirtualMachineBulkOperations_BulkGetOperationsStatus_DeallocateFallbackAfterHibernateFail.json
 */
async function _04ResponseWithSuccessfulDeallocationFallbackAfterHibernationFails() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ComputeClient(credential, subscriptionId);
  const result = await client.virtualMachineBulkOperations.bulkGetOperationsStatus(
    "example-rg",
    "eastus",
    { operationIds: ["ffffffff-ffff-ffff-ffff-ffffffffffff"] },
  );
  console.log(result);
}
