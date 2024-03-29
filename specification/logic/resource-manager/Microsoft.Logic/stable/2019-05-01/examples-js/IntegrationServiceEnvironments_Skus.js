const { LogicManagementClient } = require("@azure/arm-logic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of integration service environment Skus.
 *
 * @summary Gets a list of integration service environment Skus.
 * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Skus.json
 */
async function listIntegrationServiceEnvironmentSkus() {
  const subscriptionId =
    process.env["LOGIC_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroup = "testResourceGroup";
  const integrationServiceEnvironmentName = "testIntegrationServiceEnvironment";
  const credential = new DefaultAzureCredential();
  const client = new LogicManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.integrationServiceEnvironmentSkus.list(
    resourceGroup,
    integrationServiceEnvironmentName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
