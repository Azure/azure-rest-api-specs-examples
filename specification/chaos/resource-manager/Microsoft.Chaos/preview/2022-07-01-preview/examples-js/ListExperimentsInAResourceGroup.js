const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of Experiment resources in a resource group.
 *
 * @summary Get a list of Experiment resources in a resource group.
 * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2022-07-01-preview/examples/ListExperimentsInAResourceGroup.json
 */
async function listAllExperimentsInAResourceGroup() {
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const resourceGroupName = "exampleRG";
  const continuationToken = undefined;
  const options = { continuationToken };
  const credential = new DefaultAzureCredential();
  const client = new ChaosManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.experiments.list(resourceGroupName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllExperimentsInAResourceGroup().catch(console.error);
