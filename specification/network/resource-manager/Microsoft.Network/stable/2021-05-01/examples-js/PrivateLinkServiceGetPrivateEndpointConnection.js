const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specific private end point connection by specific private link service in the resource group.
 *
 * @summary Get the specific private end point connection by specific private link service in the resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateLinkServiceGetPrivateEndpointConnection.json
 */
async function getPrivateEndPointConnection() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const serviceName = "testPls";
  const peConnectionName = "testPlePeConnection";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.privateLinkServices.getPrivateEndpointConnection(
    resourceGroupName,
    serviceName,
    peConnectionName
  );
  console.log(result);
}

getPrivateEndPointConnection().catch(console.error);
