const { ComputeClient } = require("@azure/arm-compute-bulkactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to bulkVdiFlexCreate: Bulk create  operation for a batch of virtual machines, this operation supports flex properties to give options on Sku and zone selection.
 *
 * @summary bulkVdiFlexCreate: Bulk create  operation for a batch of virtual machines, this operation supports flex properties to give options on Sku and zone selection.
 * x-ms-original-file: 2026-08-06-preview/VirtualMachineBulkOperations_BulkVdiFlexCreate_MinimumSet_Gen.json
 */
async function virtualMachineBulkOperationsBulkVdiFlexCreateExampleGeneratedByMinimumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "1FBA3C66-5C9C-4391-B72F-9F52735FC9F2";
  const client = new ComputeClient(credential, subscriptionId);
  const result = await client.virtualMachineBulkOperations.bulkVdiFlexCreateOperation(
    "rgBulkactions",
    "useast2euap",
    {
      resourceConfigParameters: {
        resourceCount: 10,
        flexProperties: {
          vmSizeProfiles: [{ name: "Standard_D2ads_v5", rank: 7 }],
          osType: "Windows",
          priorityProfile: {},
        },
      },
      executionParameters: {},
    },
  );
  console.log(result);
}
