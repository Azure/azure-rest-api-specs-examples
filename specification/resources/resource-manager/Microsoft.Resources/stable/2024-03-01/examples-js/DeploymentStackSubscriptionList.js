const { DeploymentStacksClient } = require("@azure/arm-resourcesdeploymentstacks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the Deployment stacks within the specified Subscription.
 *
 * @summary Lists all the Deployment stacks within the specified Subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/DeploymentStackSubscriptionList.json
 */
async function deploymentStacksSubscriptionList() {
  const subscriptionId =
    process.env["RESOURCESDEPLOYMENTSTACKS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new DeploymentStacksClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.deploymentStacks.listAtSubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
