const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the specified Network Virtual Appliance Site.
 *
 * @summary Creates or updates the specified Network Virtual Appliance Site.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkVirtualApplianceSitePut.json
 */
async function createNetworkVirtualApplianceSite() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkVirtualApplianceName = "nva";
  const siteName = "site1";
  const parameters = {
    addressPrefix: "192.168.1.0/24",
    o365Policy: {
      breakOutCategories: { default: true, allow: true, optimize: true },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualApplianceSites.beginCreateOrUpdateAndWait(
    resourceGroupName,
    networkVirtualApplianceName,
    siteName,
    parameters
  );
  console.log(result);
}

createNetworkVirtualApplianceSite().catch(console.error);
