const { ResourceManagementClient } = require("@azure/arm-resources-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to You can provide the template and parameters directly in the request or link to JSON files.
 *
 * @summary You can provide the template and parameters directly in the request or link to JSON files.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2019-10-01/examples/PutDeploymentAtScope.json
 */
async function createDeploymentAtAGivenScope() {
  const subscriptionId =
    process.env["RESOURCES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope = "providers/Microsoft.Management/managementGroups/tiano-group1";
  const deploymentName = "mg-dep01";
  const parameters = {
    location: "eastus",
    properties: {
      mode: "Incremental",
      parameters: {},
      templateLink: { uri: "{templateUri}" },
    },
    tags: { tagKey1: "tagValue1", tagKey2: "tagValue2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginCreateOrUpdateAtScopeAndWait(
    scope,
    deploymentName,
    parameters
  );
  console.log(result);
}
