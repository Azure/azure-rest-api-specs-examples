const { AzureDeploymentManager } = require("@azure/arm-deploymentmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the artifact sources in a resource group.
 *
 * @summary Lists the artifact sources in a resource group.
 * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/artifactsources_list.json
 */
async function listSteps() {
  const subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
  const resourceGroupName = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new AzureDeploymentManager(credential, subscriptionId);
  const result = await client.artifactSources.list(resourceGroupName);
  console.log(result);
}

listSteps().catch(console.error);
