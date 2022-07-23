const { FluidRelayManagementClient } = require("@azure/arm-fluidrelay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a Fluid Relay container.
 *
 * @summary Delete a Fluid Relay container.
 * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayContainers_Delete.json
 */
async function deleteAFluidRelayContainer() {
  const subscriptionId = "xxxx-xxxx-xxxx-xxxx";
  const resourceGroup = "myResourceGroup";
  const fluidRelayServerName = "myFluidRelayServer";
  const fluidRelayContainerName = "myFluidRelayContainer";
  const credential = new DefaultAzureCredential();
  const client = new FluidRelayManagementClient(credential, subscriptionId);
  const result = await client.fluidRelayContainers.delete(
    resourceGroup,
    fluidRelayServerName,
    fluidRelayContainerName
  );
  console.log(result);
}

deleteAFluidRelayContainer().catch(console.error);
