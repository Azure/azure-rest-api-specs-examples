const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all private endpoint connections.
 *
 * @summary List all private endpoint connections.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/private-endpoint-connection-list.json
 */
async function getAllPrivateEndpointConnections() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contososports";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.privateEndpointConnections.list(resourceGroupName, accountName);
  console.log(result);
}

getAllPrivateEndpointConnections().catch(console.error);
