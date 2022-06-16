const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a static or dynamic public IP prefix.
 *
 * @summary Creates or updates a static or dynamic public IP prefix.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PublicIpPrefixCreateDefaults.json
 */
async function createPublicIPPrefixDefaults() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const publicIpPrefixName = "test-ipprefix";
  const parameters = {
    location: "westus",
    prefixLength: 30,
    sku: { name: "Standard" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.publicIPPrefixes.beginCreateOrUpdateAndWait(
    resourceGroupName,
    publicIpPrefixName,
    parameters
  );
  console.log(result);
}

createPublicIPPrefixDefaults().catch(console.error);
