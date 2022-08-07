const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates public IP address tags.
 *
 * @summary Updates public IP address tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/PublicIpAddressUpdateTags.json
 */
async function updatePublicIPAddressTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const publicIpAddressName = "test-ip";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.publicIPAddresses.updateTags(
    resourceGroupName,
    publicIpAddressName,
    parameters
  );
  console.log(result);
}

updatePublicIPAddressTags().catch(console.error);
