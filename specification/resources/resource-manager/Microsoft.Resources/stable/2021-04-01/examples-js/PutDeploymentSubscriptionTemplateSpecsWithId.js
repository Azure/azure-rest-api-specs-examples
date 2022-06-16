const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to You can provide the template and parameters directly in the request or link to JSON files.
 *
 * @summary You can provide the template and parameters directly in the request or link to JSON files.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PutDeploymentSubscriptionTemplateSpecsWithId.json
 */
async function createADeploymentThatWillDeployATemplateSpecWithTheGivenResourceId() {
  const subscriptionId = "00000000-0000-0000-0000-000000000001";
  const deploymentName = "my-deployment";
  const parameters = {
    location: "eastus",
    properties: {
      mode: "Incremental",
      parameters: {},
      templateLink: {
        id: "/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group/providers/Microsoft.Resources/TemplateSpecs/TemplateSpec-Name/versions/v1",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginCreateOrUpdateAtSubscriptionScopeAndWait(
    deploymentName,
    parameters
  );
  console.log(result);
}

createADeploymentThatWillDeployATemplateSpecWithTheGivenResourceId().catch(console.error);
