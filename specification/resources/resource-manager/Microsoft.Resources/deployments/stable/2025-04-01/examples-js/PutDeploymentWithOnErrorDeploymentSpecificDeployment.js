const { DeploymentsClient } = require("@azure/arm-resourcesdeployments");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to you can provide the template and parameters directly in the request or link to JSON files.
 *
 * @summary you can provide the template and parameters directly in the request or link to JSON files.
 * x-ms-original-file: 2025-04-01/PutDeploymentWithOnErrorDeploymentSpecificDeployment.json
 */
async function createADeploymentThatWillRedeployAnotherDeploymentOnFailure() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DeploymentsClient(credential, subscriptionId);
  const result = await client.deployments.createOrUpdate("my-resource-group", "my-deployment", {
    properties: {
      mode: "Complete",
      onErrorDeployment: {
        type: "SpecificDeployment",
        deploymentName: "name-of-deployment-to-use",
      },
      parameters: {},
      templateLink: { uri: "https://example.com/exampleTemplate.json" },
    },
  });
  console.log(result);
}
