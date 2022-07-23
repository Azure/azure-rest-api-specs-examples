const { FluidRelayManagementClient } = require("@azure/arm-fluidrelay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all Fluid Relay servers in a resource group.
 *
 * @summary List all Fluid Relay servers in a resource group.
 * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_ListByResourceGroup.json
 */
async function listAllFluidRelayServersInAResourceGroup() {
  const subscriptionId = "xxxx-xxxx-xxxx-xxxx";
  const resourceGroup = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new FluidRelayManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fluidRelayServers.listByResourceGroup(resourceGroup)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllFluidRelayServersInAResourceGroup().catch(console.error);
