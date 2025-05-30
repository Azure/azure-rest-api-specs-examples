const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Reimages one VM belonging to the specified Network Virtual Appliance.
 *
 * @summary Reimages one VM belonging to the specified Network Virtual Appliance.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkVirtualApplianceSpecificReimage.json
 */
async function reimagesSpecificNetworkVirtualApplianceVMSInVMScaleSet() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkVirtualApplianceName = "nva";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkVirtualAppliances.beginReimageAndWait(
    resourceGroupName,
    networkVirtualApplianceName,
  );
  console.log(result);
}
