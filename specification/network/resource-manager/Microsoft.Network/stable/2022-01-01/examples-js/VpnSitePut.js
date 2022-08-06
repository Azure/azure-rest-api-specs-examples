const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a VpnSite resource if it doesn't exist else updates the existing VpnSite.
 *
 * @summary Creates a VpnSite resource if it doesn't exist else updates the existing VpnSite.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VpnSitePut.json
 */
async function vpnSiteCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const vpnSiteName = "vpnSite1";
  const vpnSiteParameters = {
    addressSpace: { addressPrefixes: ["10.0.0.0/16"] },
    isSecuritySite: false,
    location: "West US",
    o365Policy: {
      breakOutCategories: { default: false, allow: true, optimize: true },
    },
    tags: { key1: "value1" },
    virtualWan: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWANs/wan1",
    },
    vpnSiteLinks: [
      {
        name: "vpnSiteLink1",
        bgpProperties: { asn: 1234, bgpPeeringAddress: "192.168.0.0" },
        fqdn: "link1.vpnsite1.contoso.com",
        ipAddress: "50.50.50.56",
        linkProperties: { linkProviderName: "vendor1", linkSpeedInMbps: 0 },
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnSites.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vpnSiteName,
    vpnSiteParameters
  );
  console.log(result);
}

vpnSiteCreate().catch(console.error);
