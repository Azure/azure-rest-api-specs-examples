const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list the operations for the provider
 *
 * @summary list the operations for the provider
 * x-ms-original-file: 2026-03-01/computeRPCommonExamples/Operations_List_MinimumSet_Gen.json
 */
async function operationsListMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.operations.list()) {
    resArray.push(item);
  }

  console.log(resArray);
}
