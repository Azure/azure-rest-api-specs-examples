const { ResourceManagementClient } = require("@azure/arm-resources-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to You can provide the template and parameters directly in the request or link to JSON files.
 *
 * @summary You can provide the template and parameters directly in the request or link to JSON files.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2019-10-01/examples/PutDeploymentWithOnErrorDeploymentSpecificDeployment.json
 */
async function createADeploymentThatWillRedeployAnotherDeploymentOnFailure() {
  const subscriptionId = process.env["RESOURCES_SUBSCRIPTION_ID"] || "{subscriptionId}";
  const resourceGroupName = process.env["RESOURCES_RESOURCE_GROUP"] || "myResourceGroup";
  const deploymentName = "exampleDeploymentName";
  const parameters = {
    properties: {
      mode: "Complete",
      onErrorDeployment: {
        type: "SpecificDeployment",
        deploymentName: "{nameOfDeploymentToUse}",
      },
      parameters: {},
      templateLink: { uri: "{templateUri}" },
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
