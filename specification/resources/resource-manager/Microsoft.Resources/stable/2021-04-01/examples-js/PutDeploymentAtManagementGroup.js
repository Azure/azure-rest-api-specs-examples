const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to You can provide the template and parameters directly in the request or link to JSON files.
 *
 * @summary You can provide the template and parameters directly in the request or link to JSON files.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PutDeploymentAtManagementGroup.json
 */
async function createDeploymentAtManagementGroupScope() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const groupId = "my-management-group-id";
  const deploymentName = "my-deployment";
  const parameters = {
    location: "eastus",
    properties: {
      mode: "Incremental",
      parameters: {},
      templateLink: { uri: "https://example.com/exampleTemplate.json" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginCreateOrUpdateAtManagementGroupScopeAndWait(
    groupId,
    deploymentName,
    parameters
  );
  console.log(result);
}

createDeploymentAtManagementGroupScope().catch(console.error);
