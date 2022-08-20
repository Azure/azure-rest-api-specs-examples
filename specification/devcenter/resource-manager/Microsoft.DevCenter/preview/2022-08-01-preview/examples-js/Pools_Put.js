const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a machine pool
 *
 * @summary Creates or updates a machine pool
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Pools_Put.json
 */
async function poolsCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const projectName = "{projectName}";
  const poolName = "{poolName}";
  const body = {
    devBoxDefinitionName: "WebDevBox",
    licenseType: "Windows_Client",
    localAdministrator: "Enabled",
    location: "centralus",
    networkConnectionName: "Network1-westus2",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.pools.beginCreateOrUpdateAndWait(
    resourceGroupName,
    projectName,
    poolName,
    body
  );
  console.log(result);
}

poolsCreateOrUpdate().catch(console.error);
