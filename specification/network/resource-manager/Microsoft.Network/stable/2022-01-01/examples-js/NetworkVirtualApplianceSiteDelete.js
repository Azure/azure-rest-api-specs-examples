const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified site from a Virtual Appliance.
 *
 * @summary Deletes the specified site from a Virtual Appliance.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkVirtualApplianceSiteDelete.json
 */
async function deleteNetworkVirtualApplianceSite() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkVirtualApplianceName = "nva";
  const siteName = "site1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualApplianceSites.beginDeleteAndWait(
    resourceGroupName,
    networkVirtualApplianceName,
    siteName
  );
  console.log(result);
}

deleteNetworkVirtualApplianceSite().catch(console.error);
