const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets, for the specified location, the current compute resource usage information as well as the limits for compute resources under the subscription.
 *
 * @summary Gets, for the specified location, the current compute resource usage information as well as the limits for compute resources under the subscription.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/computeRPCommonExamples/Usage_List_MinimumSet_Gen.json
 */
async function usageListMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const location = "_--";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.usageOperations.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
