const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to The operation to get all gateways of a non-Azure machine
 *
 * @summary The operation to get all gateways of a non-Azure machine
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/gateway/Gateway_ListBySubscription.json
 */
async function listGatewaysBySubscription() {
  const subscriptionId =
    process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "ffd506c8-3415-42d3-9612-fdb423fb17df";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.gateways.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
