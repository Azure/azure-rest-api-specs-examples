const { FluidRelayManagementClient } = require("@azure/arm-fluidrelay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or Update a Fluid Relay server.
 *
 * @summary Create or Update a Fluid Relay server.
 * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_CreateWithAmi.json
 */
async function createAFluidRelayServerWithAmi() {
  const subscriptionId = process.env["FLUIDRELAY_SUBSCRIPTION_ID"] || "xxxx-xxxx-xxxx-xxxx";
  const resourceGroup = "myResourceGroup";
  const fluidRelayServerName = "myFluidRelayServer";
  const resource = {
    identity: {
      type: "SystemAssigned, UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/xxxxXxxxXxxxXxxx/resourceGroups/myResourceGroup/providers/MicrosoftManagedIdentity/userAssignedIdentities/id1":
          {},
        "/subscriptions/xxxxXxxxXxxxXxxx/resourceGroups/myResourceGroup/providers/MicrosoftManagedIdentity/userAssignedIdentities/id2":
          {},
      },
    },
    location: "west-us",
    storagesku: "basic",
    tags: { category: "sales" },
  };
  const credential = new DefaultAzureCredential();
  const client = new FluidRelayManagementClient(credential, subscriptionId);
  const result = await client.fluidRelayServers.createOrUpdate(
    resourceGroup,
    fluidRelayServerName,
    resource
  );
  console.log(result);
}
