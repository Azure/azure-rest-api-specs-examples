const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Cancel a running Experiment resource.
 *
 * @summary Cancel a running Experiment resource.
 * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2022-07-01-preview/examples/CancelAExperiment.json
 */
async function cancelARunningExperiment() {
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const resourceGroupName = "exampleRG";
  const experimentName = "exampleExperiment";
  const credential = new DefaultAzureCredential();
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.experiments.cancel(resourceGroupName, experimentName);
  console.log(result);
}

cancelARunningExperiment().catch(console.error);
