const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified Virtual Appliance Site.
 *
 * @summary Gets the specified Virtual Appliance Site.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/NetworkVirtualApplianceSiteGet.json
 */
async function getNetworkVirtualApplianceSite() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkVirtualApplianceName = "nva";
  const siteName = "site1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualApplianceSites.get(
    resourceGroupName,
    networkVirtualApplianceName,
    siteName,
  );
  console.log(result);
}
