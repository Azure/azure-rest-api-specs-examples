const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the specified ExpressRouteCrossConnection.
 *
 * @summary Update the specified ExpressRouteCrossConnection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/ExpressRouteCrossConnectionUpdate.json
 */
async function updateExpressRouteCrossConnection() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName =
    process.env["NETWORK_RESOURCE_GROUP"] || "CrossConnection-SiliconValley";
  const crossConnectionName = "<circuitServiceKey>";
  const parameters = {
    serviceProviderProvisioningState: "NotProvisioned",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCrossConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    crossConnectionName,
    parameters,
  );
  console.log(result);
}
