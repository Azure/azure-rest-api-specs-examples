const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Performs vip swap operation on swappable cloud services.
 *
 * @summary Performs vip swap operation on swappable cloud services.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/CloudServiceSwapPut.json
 */
async function putVipSwapOperation() {
  const subscriptionId = "subid";
  const groupName = "rg1";
  const resourceName = "testCloudService";
  const parameters = { properties: { slotType: "Production" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vipSwap.beginCreateAndWait(groupName, resourceName, parameters);
  console.log(result);
}

putVipSwapOperation().catch(console.error);
