const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified flow log resource.
 *
 * @summary Deletes the specified flow log resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkWatcherFlowLogDelete.json
 */
async function deleteFlowLog() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const flowLogName = "fl";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.flowLogs.beginDeleteAndWait(
    resourceGroupName,
    networkWatcherName,
    flowLogName
  );
  console.log(result);
}

deleteFlowLog().catch(console.error);
