const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns changes that will be made by the deployment if executed at the scope of the management group.
 *
 * @summary Returns changes that will be made by the deployment if executed at the scope of the management group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PostDeploymentWhatIfOnManagementGroup.json
 */
async function predictTemplateChangesAtManagementGroupScope() {
  const subscriptionId =
    process.env["RESOURCES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const groupId = "myManagementGruop";
  const deploymentName = "exampleDeploymentName";
  const parameters = {
    location: "eastus",
    properties: { mode: "Incremental", parameters: {}, templateLink: {} },
  };
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginWhatIfAtManagementGroupScopeAndWait(
    groupId,
    deploymentName,
    parameters
  );
  console.log(result);
}
