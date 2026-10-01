const { ComputeClient } = require("@azure/arm-compute-bulkactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get the current status of one or more operations identified by their Bulk Action Operation Ids.
 *
 * @summary get the current status of one or more operations identified by their Bulk Action Operation Ids.
 * x-ms-original-file: 2026-10-06-preview/VirtualMachineBulkOperations_BulkGetOperationsStatus_FailedOperation.json
 */
async function _02GetTheStatusOfAFailedOperation() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ComputeClient(credential, subscriptionId);
  const result = await client.virtualMachineBulkOperations.bulkGetOperationsStatus(
    "example-rg",
    "eastus",
    { operationIds: ["e69c80d2-4f31-46ac-9e35-c6a7cb63fe12"] },
  );
  console.log(result);
}
