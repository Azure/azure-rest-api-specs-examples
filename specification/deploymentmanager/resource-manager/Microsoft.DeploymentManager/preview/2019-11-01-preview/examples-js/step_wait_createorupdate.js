const { AzureDeploymentManager } = require("@azure/arm-deploymentmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Synchronously creates a new step or updates an existing step.
 *
 * @summary Synchronously creates a new step or updates an existing step.
 * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/step_wait_createorupdate.json
 */
async function createWaitStep() {
  const subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
  const resourceGroupName = "myResourceGroup";
  const stepName = "waitStep";
  const stepInfo = {
    location: "centralus",
    properties: { attributes: { duration: "PT20M" }, stepType: "Wait" },
    tags: {},
  };
  const options = { stepInfo };
  const credential = new DefaultAzureCredential();
  const client = new AzureDeploymentManager(credential, subscriptionId);
  const result = await client.steps.createOrUpdate(resourceGroupName, stepName, options);
  console.log(result);
}

createWaitStep().catch(console.error);
