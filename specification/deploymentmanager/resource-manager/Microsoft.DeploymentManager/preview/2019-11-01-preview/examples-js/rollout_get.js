const { AzureDeploymentManager } = require("@azure/arm-deploymentmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets detailed information of a rollout.
 *
 * @summary Gets detailed information of a rollout.
 * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollout_get.json
 */
async function getRollout() {
  const subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
  const resourceGroupName = "myResourceGroup";
  const rolloutName = "myRollout";
  const credential = new DefaultAzureCredential();
  const client = new AzureDeploymentManager(credential, subscriptionId);
  const result = await client.rollouts.get(resourceGroupName, rolloutName);
  console.log(result);
}

getRollout().catch(console.error);
