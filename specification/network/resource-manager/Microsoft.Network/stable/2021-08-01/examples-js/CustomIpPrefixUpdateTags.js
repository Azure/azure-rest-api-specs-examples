const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates custom IP prefix tags.
 *
 * @summary Updates custom IP prefix tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/CustomIpPrefixUpdateTags.json
 */
async function updatePublicIPAddressTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const customIpPrefixName = "test-customipprefix";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.customIPPrefixes.updateTags(
    resourceGroupName,
    customIpPrefixName,
    parameters
  );
  console.log(result);
}

updatePublicIPAddressTags().catch(console.error);
