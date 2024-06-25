const { DeploymentStacksClient } = require("@azure/arm-resourcesdeploymentstacks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Exports the template used to create the Deployment stack at Management Group scope.
 *
 * @summary Exports the template used to create the Deployment stack at Management Group scope.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/DeploymentStackManagementGroupExportTemplate.json
 */
async function deploymentStacksManagementGroupExportTemplate() {
  const managementGroupId = "myMg";
  const deploymentStackName = "simpleDeploymentStack";
  const credential = new DefaultAzureCredential();
  const client = new DeploymentStacksClient(credential);
  const result = await client.deploymentStacks.exportTemplateAtManagementGroup(
    managementGroupId,
    deploymentStackName,
  );
  console.log(result);
}
