const { ResourceConnectorManagementClient } = require("@azure/arm-resourceconnector");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an Appliance in the specified Subscription and Resource Group.
 *
 * @summary creates or updates an Appliance in the specified Subscription and Resource Group.
 * x-ms-original-file: 2025-03-01-preview/AppliancesUpdateProxy.json
 */
async function updateApplianceProxyConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "11111111-2222-3333-4444-555555555555";
  const client = new ResourceConnectorManagementClient(credential, subscriptionId);
  const result = await client.appliances.createOrUpdate("testresourcegroup", "appliance01", {
    location: "West US",
    distro: "AKSEdge",
    publicKey: "xxxxxxxx",
    infrastructureConfig: { provider: "VMWare" },
    networkProfile: { proxyConfiguration: { version: "latest" } },
  });
  console.log(result);
}
