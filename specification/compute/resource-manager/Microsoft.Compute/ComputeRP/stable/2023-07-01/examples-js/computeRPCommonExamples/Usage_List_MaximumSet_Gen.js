const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets, for the specified location, the current compute resource usage information as well as the limits for compute resources under the subscription.
 *
 * @summary Gets, for the specified location, the current compute resource usage information as well as the limits for compute resources under the subscription.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/computeRPCommonExamples/Usage_List_MaximumSet_Gen.json
 */
async function usageListMaximumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const location = "4_.";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.usageOperations.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
