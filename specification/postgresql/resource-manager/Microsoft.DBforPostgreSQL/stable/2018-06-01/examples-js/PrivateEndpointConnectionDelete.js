const { PostgreSQLManagementClient } = require("@azure/arm-postgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a private endpoint connection with a given name.
 *
 * @summary Deletes a private endpoint connection with a given name.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/PrivateEndpointConnectionDelete.json
 */
async function deletesAPrivateEndpointConnectionWithAGivenName() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const serverName = "test-svr";
  const privateEndpointConnectionName = "private-endpoint-connection-name";
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginDeleteAndWait(
    resourceGroupName,
    serverName,
    privateEndpointConnectionName
  );
  console.log(result);
}

deletesAPrivateEndpointConnectionWithAGivenName().catch(console.error);
