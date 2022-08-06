const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an existing private endpoint connection.
 *
 * @summary Update an existing private endpoint connection.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/private-endpoint-connection-put.json
 */
async function updatePrivateEndpointConnection() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contososports";
  const name = "connectionName1";
  const parameters = {
    privateLinkServiceConnectionState: {
      description: "Test description.",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.privateEndpointConnections.createOrUpdate(
    resourceGroupName,
    accountName,
    name,
    parameters
  );
  console.log(result);
}

updatePrivateEndpointConnection().catch(console.error);
