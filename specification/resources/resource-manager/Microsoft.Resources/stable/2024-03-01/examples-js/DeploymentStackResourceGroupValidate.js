const { DeploymentStacksClient } = require("@azure/arm-resourcesdeploymentstacks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Runs preflight validation on the Resource Group scoped Deployment stack template to verify its acceptance to Azure Resource Manager.
 *
 * @summary Runs preflight validation on the Resource Group scoped Deployment stack template to verify its acceptance to Azure Resource Manager.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/DeploymentStackResourceGroupValidate.json
 */
async function deploymentStacksResourceGroupValidate() {
  const subscriptionId =
    process.env["RESOURCESDEPLOYMENTSTACKS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName =
    process.env["RESOURCESDEPLOYMENTSTACKS_RESOURCE_GROUP"] || "deploymentStacksRG";
  const deploymentStackName = "simpleDeploymentStack";
  const deploymentStack = {
    properties: {
      actionOnUnmanage: {
        managementGroups: "delete",
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
      templateLink: { uri: "https://example.com/exampleTemplate.json" },
    },
    tags: { tagkey: "tagVal" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DeploymentStacksClient(credential, subscriptionId);
  const result = await client.deploymentStacks.beginValidateStackAtResourceGroupAndWait(
    resourceGroupName,
    deploymentStackName,
    deploymentStack,
  );
  console.log(result);
}
