const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a private endpoint connection.
 *
 * @summary Deletes a private endpoint connection.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/private-endpoint-connection-delete.json
 */
async function deletePrivateEndpointConnection() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contososports";
  const name = "connectionName1";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.privateEndpointConnections.delete(
    resourceGroupName,
    accountName,
    name
  );
  console.log(result);
}

deletePrivateEndpointConnection().catch(console.error);
