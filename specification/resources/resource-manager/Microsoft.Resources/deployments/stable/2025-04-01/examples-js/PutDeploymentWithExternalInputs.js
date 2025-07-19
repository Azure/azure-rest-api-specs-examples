const { DeploymentsClient } = require("@azure/arm-resourcesdeployments");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to You can provide the template and parameters directly in the request or link to JSON files.
 *
 * @summary You can provide the template and parameters directly in the request or link to JSON files.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/deployments/stable/2025-04-01/examples/PutDeploymentWithExternalInputs.json
 */
async function createDeploymentUsingExternalInputs() {
  const subscriptionId =
    process.env["RESOURCES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000001";
  const resourceGroupName = process.env["RESOURCES_RESOURCE_GROUP"] || "my-resource-group";
  const deploymentName = "my-deployment";
  const parameters = {
    properties: {
      externalInputDefinitions: {
        fooValue: { config: "FOO_VALUE", kind: "sys.envVar" },
      },
      externalInputs: { fooValue: { value: "baz" } },
      mode: "Incremental",
      parameters: {
        inputObj: {
          expression: "[createObject('foo', externalInputs('fooValue'))]",
        },
      },
      template: {
        $schema: "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
        contentVersion: "1.0.0.0",
        outputs: {
          inputObj: { type: "object", value: "[parameters('inputObj')]" },
        },
        parameters: { inputObj: { type: "object" } },
        resources: [],
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DeploymentsClient(credential, subscriptionId);
  const result = await client.deployments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    deploymentName,
    parameters,
  );
  console.log(result);
}
