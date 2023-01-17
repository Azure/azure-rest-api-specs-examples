const { FluidRelayManagementClient } = require("@azure/arm-fluidrelay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all Fluid Relay servers in a subscription.
 *
 * @summary List all Fluid Relay servers in a subscription.
 * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_ListBySubscription.json
 */
async function listAllFluidRelayServersInASubscription() {
  const subscriptionId = process.env["FLUIDRELAY_SUBSCRIPTION_ID"] || "xxxx-xxxx-xxxx-xxxx";
  const credential = new DefaultAzureCredential();
  const client = new FluidRelayManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fluidRelayServers.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
