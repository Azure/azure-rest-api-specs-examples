const { WorkloadsClient } = require("@azure/arm-migrationdiscoverysap");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the operations for the provider
 *
 * @summary List the operations for the provider
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/operations/preview/2023-10-01-preview/examples/Operations_List.json
 */
async function listTheOperationsForTheProvider() {
  const subscriptionId =
    process.env["MIGRATIONDISCOVERY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
