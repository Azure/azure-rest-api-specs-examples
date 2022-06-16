const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a static or dynamic public IP address.
 *
 * @summary Creates or updates a static or dynamic public IP address.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/PublicIpAddressCreateDns.json
 */
async function createPublicIPAddressDns() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const publicIpAddressName = "test-ip";
  const parameters = {
    dnsSettings: { domainNameLabel: "dnslbl" },
    location: "eastus",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.publicIPAddresses.beginCreateOrUpdateAndWait(
    resourceGroupName,
    publicIpAddressName,
    parameters
  );
  console.log(result);
}

createPublicIPAddressDns().catch(console.error);
