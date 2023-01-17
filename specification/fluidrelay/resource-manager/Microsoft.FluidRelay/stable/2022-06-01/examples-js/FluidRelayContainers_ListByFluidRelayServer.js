const { FluidRelayManagementClient } = require("@azure/arm-fluidrelay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all Fluid Relay containers which are children of a given Fluid Relay server.
 *
 * @summary List all Fluid Relay containers which are children of a given Fluid Relay server.
 * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayContainers_ListByFluidRelayServer.json
 */
async function listAllFluidRelayContainersInAFluidRelayServer() {
  const subscriptionId = process.env["FLUIDRELAY_SUBSCRIPTION_ID"] || "xxxx-xxxx-xxxx-xxxx";
  const resourceGroup = "myResourceGroup";
  const fluidRelayServerName = "myFluidRelayServer";
  const credential = new DefaultAzureCredential();
  const client = new FluidRelayManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fluidRelayContainers.listByFluidRelayServers(
    resourceGroup,
    fluidRelayServerName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
