const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a machine pool.
 *
 * @summary creates or updates a machine pool.
 * x-ms-original-file: 2026-01-01-preview/Pools_PutWithManagedNetwork.json
 */
async function poolsCreateOrUpdateWithManagedNetwork() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58fffff";
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.pools.createOrUpdate("rg1", "DevProject", "DevPool", {
    location: "centralus",
    devBoxDefinitionName: "WebDevBox",
    displayName: "Developer Pool",
    licenseType: "Windows_Client",
    localAdministrator: "Enabled",
    managedVirtualNetworkRegions: ["centralus"],
    networkConnectionName: "managedNetwork",
    singleSignOnStatus: "Disabled",
    stopOnDisconnect: { gracePeriodMinutes: 60, status: "Enabled" },
    stopOnNoConnect: { gracePeriodMinutes: 120, status: "Enabled" },
    virtualNetworkType: "Managed",
  });
  console.log(result);
}
