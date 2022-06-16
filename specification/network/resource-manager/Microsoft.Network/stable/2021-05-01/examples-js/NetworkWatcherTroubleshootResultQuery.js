const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the last completed troubleshooting result on a specified resource.
 *
 * @summary Get the last completed troubleshooting result on a specified resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherTroubleshootResultQuery.json
 */
async function getTroubleshootResult() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const parameters = {
    targetResourceId:
      "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkWatchers.beginGetTroubleshootingResultAndWait(
    resourceGroupName,
    networkWatcherName,
    parameters
  );
  console.log(result);
}

getTroubleshootResult().catch(console.error);
