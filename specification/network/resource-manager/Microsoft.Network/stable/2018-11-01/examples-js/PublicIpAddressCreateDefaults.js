const { NetworkManagementClient } = require("@azure/arm-network-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a static or dynamic public IP address.
 *
 * @summary Creates or updates a static or dynamic public IP address.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2018-11-01/examples/PublicIpAddressCreateDefaults.json
 */
async function createPublicIPAddressDefaults() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const publicIpAddressName = "test-ip";
  const parameters = {};
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.publicIPAddresses.beginCreateOrUpdateAndWait(
    resourceGroupName,
    publicIpAddressName,
    parameters
  );
  console.log(result);
}
