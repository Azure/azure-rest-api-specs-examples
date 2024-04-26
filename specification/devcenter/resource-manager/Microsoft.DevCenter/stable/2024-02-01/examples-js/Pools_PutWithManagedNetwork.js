const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a machine pool
 *
 * @summary Creates or updates a machine pool
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/Pools_PutWithManagedNetwork.json
 */
async function poolsCreateOrUpdateWithManagedNetwork() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
  const projectName = "DevProject";
  const poolName = "DevPool";
  const body = {
    devBoxDefinitionName: "WebDevBox",
    displayName: "Developer Pool",
    licenseType: "Windows_Client",
    localAdministrator: "Enabled",
    location: "centralus",
    managedVirtualNetworkRegions: ["centralus"],
    networkConnectionName: "managedNetwork",
    singleSignOnStatus: "Disabled",
    stopOnDisconnect: { gracePeriodMinutes: 60, status: "Enabled" },
    virtualNetworkType: "Managed",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.pools.beginCreateOrUpdateAndWait(
    resourceGroupName,
    projectName,
    poolName,
    body,
  );
  console.log(result);
}
