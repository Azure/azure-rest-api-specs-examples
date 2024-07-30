const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of currently active sessions on the Bastion.
 *
 * @summary Returns the list of currently active sessions on the Bastion.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/BastionSessionDelete.json
 */
async function deletesTheSpecifiedActiveSession() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const bastionHostName = "bastionhosttenant";
  const sessionIds = {};
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.listDisconnectActiveSessions(
    resourceGroupName,
    bastionHostName,
    sessionIds,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
