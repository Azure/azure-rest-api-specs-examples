const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns changes that will be made by the deployment if executed at the scope of the resource group.
 *
 * @summary Returns changes that will be made by the deployment if executed at the scope of the resource group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PostDeploymentWhatIfOnResourceGroup.json
 */
async function predictTemplateChangesAtResourceGroupScope() {
  const subscriptionId =
    process.env["RESOURCES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000001";
  const resourceGroupName = process.env["RESOURCES_RESOURCE_GROUP"] || "my-resource-group";
  const deploymentName = "my-deployment";
  const parameters = {
    properties: { mode: "Incremental", parameters: {}, templateLink: {} },
  };
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginWhatIfAndWait(
    resourceGroupName,
    deploymentName,
    parameters
  );
  console.log(result);
}
