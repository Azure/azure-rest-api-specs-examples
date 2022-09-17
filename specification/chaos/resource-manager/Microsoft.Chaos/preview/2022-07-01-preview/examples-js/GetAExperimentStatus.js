const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a status of a Experiment resource.
 *
 * @summary Get a status of a Experiment resource.
 * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2022-07-01-preview/examples/GetAExperimentStatus.json
 */
async function getTheStatusOfAExperiment() {
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const resourceGroupName = "exampleRG";
  const experimentName = "exampleExperiment";
  const statusId = "50734542-2e64-4e08-814c-cc0e7475f7e4";
  const credential = new DefaultAzureCredential();
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.experiments.getStatus(resourceGroupName, experimentName, statusId);
  console.log(result);
}

getTheStatusOfAExperiment().catch(console.error);
