const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Calculate the hash of the given template.
 *
 * @summary Calculate the hash of the given template.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/CalculateTemplateHash.json
 */
async function calculateTemplateHash() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const template = {
    $schema:
      "http://schemas.management.azure.com/deploymentTemplate?api-version=2014-04-01-preview",
    contentVersion: "1.0.0.0",
    outputs: { string: { type: "string", value: "myvalue" } },
    parameters: { string: { type: "string" } },
    resources: [],
    variables: {
      array: [1, 2, 3, 4],
      bool: true,
      int: 42,
      object: { object: { location: "West US", vmSize: "Large" } },
      string: "string",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.deployments.calculateTemplateHash(template);
  console.log(result);
}

calculateTemplateHash().catch(console.error);
