const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of execution details of a Experiment resource.
 *
 * @summary Get a list of execution details of a Experiment resource.
 * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2022-07-01-preview/examples/ListExperimentExecutionsDetails.json
 */
async function listExperimentExecutionsDetails() {
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const resourceGroupName = "exampleRG";
  const experimentName = "exampleExperiment";
  const credential = new DefaultAzureCredential();
  const client = new ChaosManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.experiments.listExecutionDetails(
    resourceGroupName,
    experimentName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listExperimentExecutionsDetails().catch(console.error);
