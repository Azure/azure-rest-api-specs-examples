const { ComputeClient } = require("@azure/arm-compute-bulkactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to bulkCreate: Execute create operation for a batch of virtual machines, this operation is triggered as soon as Computeschedule receives it.
 *
 * @summary bulkCreate: Execute create operation for a batch of virtual machines, this operation is triggered as soon as Computeschedule receives it.
 * x-ms-original-file: 2026-07-06-preview/VirtualMachineBulkOperations_BulkCreate_MinimumSet_Gen.json
 */
async function virtualMachineBulkOperationsBulkCreateExampleGeneratedByMinimumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "1FBA3C66-5C9C-4391-B72F-9F52735FC9F2";
  const client = new ComputeClient(credential, subscriptionId);
  const result = await client.virtualMachineBulkOperations.bulkCreateOperation(
    "rgBulkactions",
    "useast2euap",
    { resourceConfigParameters: { resourceCount: 23 }, executionParameters: {} },
  );
  console.log(result);
}
