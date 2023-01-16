const { FluidRelayManagementClient } = require("@azure/arm-fluidrelay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get primary and secondary key for this server.
 *
 * @summary Get primary and secondary key for this server.
 * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_ListKeys.json
 */
async function getKeysForAFluidRelayServer() {
  const subscriptionId = process.env["FLUIDRELAY_SUBSCRIPTION_ID"] || "xxxx-xxxx-xxxx-xxxx";
  const resourceGroup = "myResourceGroup";
  const fluidRelayServerName = "myFluidRelayServer";
  const credential = new DefaultAzureCredential();
  const client = new FluidRelayManagementClient(credential, subscriptionId);
  const result = await client.fluidRelayServers.listKeys(resourceGroup, fluidRelayServerName);
  console.log(result);
}
