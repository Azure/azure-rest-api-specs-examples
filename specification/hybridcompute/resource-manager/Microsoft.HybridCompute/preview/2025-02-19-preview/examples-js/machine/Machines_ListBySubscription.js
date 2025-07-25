const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all the hybrid machines in the specified subscription. Use the nextLink property in the response to get the next page of hybrid machines.
 *
 * @summary Lists all the hybrid machines in the specified subscription. Use the nextLink property in the response to get the next page of hybrid machines.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/machine/Machines_ListBySubscription.json
 */
async function listMachinesByResourceGroup() {
  const subscriptionId = process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.machines.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
