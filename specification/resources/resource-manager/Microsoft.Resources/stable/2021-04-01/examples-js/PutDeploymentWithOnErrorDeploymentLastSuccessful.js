const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to You can provide the template and parameters directly in the request or link to JSON files.
 *
 * @summary You can provide the template and parameters directly in the request or link to JSON files.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PutDeploymentWithOnErrorDeploymentLastSuccessful.json
 */
async function createADeploymentThatWillRedeployTheLastSuccessfulDeploymentOnFailure() {
  const subscriptionId =
    process.env["RESOURCES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["RESOURCES_RESOURCE_GROUP"] || "my-resource-group";
  const deploymentName = "my-deployment";
  const parameters = {
    properties: {
      mode: "Complete",
      onErrorDeployment: { type: "LastSuccessful" },
      parameters: {},
      templateLink: { uri: "https://example.com/exampleTemplate.json" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    deploymentName,
    parameters
  );
  console.log(result);
}
