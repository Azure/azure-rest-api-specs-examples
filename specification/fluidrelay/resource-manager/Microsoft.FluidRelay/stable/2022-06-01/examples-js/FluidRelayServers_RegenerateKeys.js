const { FluidRelayManagementClient } = require("@azure/arm-fluidrelay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate the primary or secondary key for this server.
 *
 * @summary Regenerate the primary or secondary key for this server.
 * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_RegenerateKeys.json
 */
async function regenerateKeysForAFluidRelayServer() {
  const subscriptionId = "xxxx-xxxx-xxxx-xxxx";
  const resourceGroup = "myResourceGroup";
  const fluidRelayServerName = "myFluidRelayServer";
  const parameters = { keyName: "key1" };
  const credential = new DefaultAzureCredential();
  const client = new FluidRelayManagementClient(credential, subscriptionId);
  const result = await client.fluidRelayServers.regenerateKey(
    resourceGroup,
    fluidRelayServerName,
    parameters
  );
  console.log(result);
}

regenerateKeysForAFluidRelayServer().catch(console.error);
