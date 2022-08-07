const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates VpnSite tags.
 *
 * @summary Updates VpnSite tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VpnSiteUpdateTags.json
 */
async function vpnSiteUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const vpnSiteName = "vpnSite1";
  const vpnSiteParameters = {
    tags: { key1: "value1", key2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnSites.updateTags(
    resourceGroupName,
    vpnSiteName,
    vpnSiteParameters
  );
  console.log(result);
}

vpnSiteUpdate().catch(console.error);
