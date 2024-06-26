const { DeploymentStacksClient } = require("@azure/arm-resourcesdeploymentstacks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Deployment stack at Subscription scope.
 *
 * @summary Creates or updates a Deployment stack at Subscription scope.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/DeploymentStackSubscriptionCreate.json
 */
async function deploymentStacksSubscriptionCreateOrUpdate() {
  const subscriptionId =
    process.env["RESOURCESDEPLOYMENTSTACKS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const deploymentStackName = "simpleDeploymentStack";
  const deploymentStack = {
    location: "eastus",
    properties: {
      actionOnUnmanage: {
        managementGroups: "detach",
        resourceGroups: "delete",
        resources: "delete",
      },
      denySettings: {
        applyToChildScopes: false,
        excludedActions: ["action"],
        excludedPrincipals: ["principal"],
        mode: "denyDelete",
      },
      parameters: { parameter1: { value: "a string" } },
    },
    tags: { tagkey: "tagVal" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DeploymentStacksClient(credential, subscriptionId);
  const result = await client.deploymentStacks.beginCreateOrUpdateAtSubscriptionAndWait(
    deploymentStackName,
    deploymentStack,
  );
  console.log(result);
}
