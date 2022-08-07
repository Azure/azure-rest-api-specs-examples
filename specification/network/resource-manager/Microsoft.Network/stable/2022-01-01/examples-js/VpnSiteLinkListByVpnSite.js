const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the vpnSiteLinks in a resource group for a vpn site.
 *
 * @summary Lists all the vpnSiteLinks in a resource group for a vpn site.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VpnSiteLinkListByVpnSite.json
 */
async function vpnSiteLinkListByVpnSite() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const vpnSiteName = "vpnSite1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.vpnSiteLinks.listByVpnSite(resourceGroupName, vpnSiteName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

vpnSiteLinkListByVpnSite().catch(console.error);
