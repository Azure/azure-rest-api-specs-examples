const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a HubVirtualNetworkConnection.
 *
 * @summary Deletes a HubVirtualNetworkConnection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/HubVirtualNetworkConnectionDelete.json
 */
async function hubVirtualNetworkConnectionDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualHubName = "virtualHub1";
  const connectionName = "connection1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.hubVirtualNetworkConnections.beginDeleteAndWait(
    resourceGroupName,
    virtualHubName,
    connectionName
  );
  console.log(result);
}

hubVirtualNetworkConnectionDelete().catch(console.error);
