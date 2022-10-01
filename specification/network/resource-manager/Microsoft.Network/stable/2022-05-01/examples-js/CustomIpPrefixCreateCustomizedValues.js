const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a custom IP prefix.
 *
 * @summary Creates or updates a custom IP prefix.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/CustomIpPrefixCreateCustomizedValues.json
 */
async function createCustomIPPrefixAllocationMethod() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const customIpPrefixName = "test-customipprefix";
  const parameters = { cidr: "0.0.0.0/24", location: "westus" };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.customIPPrefixes.beginCreateOrUpdateAndWait(
    resourceGroupName,
    customIpPrefixName,
    parameters
  );
  console.log(result);
}

createCustomIPPrefixAllocationMethod().catch(console.error);
