const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the configured and effective security group rules on the specified VM.
 *
 * @summary Gets the configured and effective security group rules on the specified VM.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherSecurityGroupViewGet.json
 */
async function getSecurityGroupView() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const parameters = {
    targetResourceId:
      "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkWatchers.beginGetVMSecurityRulesAndWait(
    resourceGroupName,
    networkWatcherName,
    parameters
  );
  console.log(result);
}

getSecurityGroupView().catch(console.error);
