const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a flow log for the specified network security group.
 *
 * @summary Create or update a flow log for the specified network security group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkWatcherFlowLogCreate.json
 */
async function createOrUpdateFlowLog() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const flowLogName = "fl";
  const parameters = {
    format: { type: "JSON", version: 1 },
    enabled: true,
    location: "centraluseuap",
    storageId:
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/nwtest1mgvbfmqsigdxe",
    targetResourceId:
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/desmondcentral-nsg",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.flowLogs.beginCreateOrUpdateAndWait(
    resourceGroupName,
    networkWatcherName,
    flowLogName,
    parameters
  );
  console.log(result);
}

createOrUpdateFlowLog().catch(console.error);
