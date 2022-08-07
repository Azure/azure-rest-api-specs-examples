const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the specified static member.
 *
 * @summary Lists the specified static member.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerStaticMemberList.json
 */
async function staticMembersList() {
  const subscriptionId = "subscriptionC";
  const resourceGroupName = "rg1";
  const networkManagerName = "testNetworkManager";
  const networkGroupName = "testNetworkGroup";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.staticMembers.list(
    resourceGroupName,
    networkManagerName,
    networkGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

staticMembersList().catch(console.error);
