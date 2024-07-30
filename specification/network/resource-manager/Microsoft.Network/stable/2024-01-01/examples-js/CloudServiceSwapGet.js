const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the SwapResource which identifies the slot type for the specified cloud service. The slot type on a cloud service can either be Staging or Production
 *
 * @summary Gets the SwapResource which identifies the slot type for the specified cloud service. The slot type on a cloud service can either be Staging or Production
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/CloudServiceSwapGet.json
 */
async function getSwapResource() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const groupName = "rg1";
  const resourceName = "testCloudService";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vipSwap.get(groupName, resourceName);
  console.log(result);
}
