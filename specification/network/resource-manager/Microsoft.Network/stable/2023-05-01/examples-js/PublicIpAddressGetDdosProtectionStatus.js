const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Ddos Protection Status of a Public IP Address
 *
 * @summary Gets the Ddos Protection Status of a Public IP Address
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/PublicIpAddressGetDdosProtectionStatus.json
 */
async function getDdosProtectionStatusOfAPublicIPAddress() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const publicIpAddressName = "test-pip";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.publicIPAddresses.beginDdosProtectionStatusAndWait(
    resourceGroupName,
    publicIpAddressName
  );
  console.log(result);
}
