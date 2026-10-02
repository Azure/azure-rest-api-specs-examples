const { ComputeClient } = require("@azure/arm-compute-bulkactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to acknowledge errors for specified operations in a resource group.
 *
 * @summary acknowledge errors for specified operations in a resource group.
 * x-ms-original-file: 2026-10-06-preview/VirtualMachineBulkOperations_BulkAcknowledgeOperationErrors_OperationNotFoundError.json
 */
async function _03AcknowledgeOperationErrorsWithAnUnknownOperationIdResultingInNotFound() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ComputeClient(credential, subscriptionId);
  const result = await client.virtualMachineBulkOperations.bulkAcknowledgeOperationErrors(
    "example-rg",
    "eastus",
    { operationIds: ["dddddddd-dddd-dddd-dddd-dddddddddddd"] },
  );
  console.log(result);
}
