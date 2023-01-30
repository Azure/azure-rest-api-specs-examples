const { LogicManagementClient } = require("@azure/arm-logic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an integration service environment.
 *
 * @summary Updates an integration service environment.
 * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Patch.json
 */
async function patchAnIntegrationServiceEnvironment() {
  const subscriptionId =
    process.env["LOGIC_SUBSCRIPTION_ID"] || "f34b22a3-2202-4fb1-b040-1332bd928c84";
  const resourceGroup = "testResourceGroup";
  const integrationServiceEnvironmentName = "testIntegrationServiceEnvironment";
  const integrationServiceEnvironment = {
    sku: { name: "Developer", capacity: 0 },
    tags: { tag1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new LogicManagementClient(credential, subscriptionId);
  const result = await client.integrationServiceEnvironments.beginUpdateAndWait(
    resourceGroup,
    integrationServiceEnvironmentName,
    integrationServiceEnvironment
  );
  console.log(result);
}
