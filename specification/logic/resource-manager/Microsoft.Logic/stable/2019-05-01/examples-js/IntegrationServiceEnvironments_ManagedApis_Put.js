const { LogicManagementClient } = require("@azure/arm-logic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Puts the integration service environment managed Api.
 *
 * @summary Puts the integration service environment managed Api.
 * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ManagedApis_Put.json
 */
async function getsTheIntegrationServiceEnvironmentManagedApis() {
  const subscriptionId =
    process.env["LOGIC_SUBSCRIPTION_ID"] || "f34b22a3-2202-4fb1-b040-1332bd928c84";
  const resourceGroup = "testResourceGroup";
  const integrationServiceEnvironmentName = "testIntegrationServiceEnvironment";
  const apiName = "servicebus";
  const integrationServiceEnvironmentManagedApi = {
    deploymentParameters: {},
    location: "brazilsouth",
  };
  const credential = new DefaultAzureCredential();
  const client = new LogicManagementClient(credential, subscriptionId);
  const result = await client.integrationServiceEnvironmentManagedApis.beginPutAndWait(
    resourceGroup,
    integrationServiceEnvironmentName,
    apiName,
    integrationServiceEnvironmentManagedApi
  );
  console.log(result);
}
