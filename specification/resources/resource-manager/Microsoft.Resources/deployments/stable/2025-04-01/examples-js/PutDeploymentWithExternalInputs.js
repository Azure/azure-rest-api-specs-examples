const { DeploymentsClient } = require("@azure/arm-resourcesdeployments");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to you can provide the template and parameters directly in the request or link to JSON files.
 *
 * @summary you can provide the template and parameters directly in the request or link to JSON files.
 * x-ms-original-file: 2025-04-01/PutDeploymentWithExternalInputs.json
 */
async function createDeploymentUsingExternalInputs() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000001";
  const client = new DeploymentsClient(credential, subscriptionId);
  const result = await client.deployments.createOrUpdate("my-resource-group", "my-deployment", {
    properties: {
      externalInputDefinitions: { fooValue: { config: "FOO_VALUE", kind: "sys.envVar" } },
      externalInputs: { fooValue: { value: "baz" } },
      mode: "Incremental",
      parameters: { inputObj: { expression: "[createObject('foo', externalInputs('fooValue'))]" } },
      template: {
        $schema: "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
        contentVersion: "1.0.0.0",
        outputs: { inputObj: { type: "object", value: "[parameters('inputObj')]" } },
        parameters: { inputObj: { type: "object" } },
        resources: [],
      },
    },
  });
  console.log(result);
}
