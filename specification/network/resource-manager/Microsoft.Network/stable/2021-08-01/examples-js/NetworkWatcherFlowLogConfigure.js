const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Configures flow log and traffic analytics (optional) on a specified resource.
 *
 * @summary Configures flow log and traffic analytics (optional) on a specified resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkWatcherFlowLogConfigure.json
 */
async function configureFlowLog() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const parameters = {
    enabled: true,
    storageId:
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/st1",
    targetResourceId:
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkWatchers.beginSetFlowLogConfigurationAndWait(
    resourceGroupName,
    networkWatcherName,
    parameters
  );
  console.log(result);
}

configureFlowLog().catch(console.error);
