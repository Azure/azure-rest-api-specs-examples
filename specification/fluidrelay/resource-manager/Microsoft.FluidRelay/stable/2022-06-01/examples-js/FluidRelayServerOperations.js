const { FluidRelayManagementClient } = require("@azure/arm-fluidrelay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all operations provided by Microsoft.FluidRelay.
 *
 * @summary List all operations provided by Microsoft.FluidRelay.
 * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServerOperations.json
 */
async function listFluidRelayServerOperations() {
  const subscriptionId =
    process.env["FLUIDRELAY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new FluidRelayManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fluidRelayOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
