const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates private endpoint connection with the specified tags.
 *
 * @summary Updates private endpoint connection with the specified tags.
 * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/PrivateEndpointConnectionUpdateTags.json
 */
async function updatePrivateEndpointConnectionTags() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const serverName = "test-svr";
  const privateEndpointConnectionName = "private-endpoint-connection-name";
  const parameters = { tags: { key1: "val1", key2: "val2" } };
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginUpdateTagsAndWait(
    resourceGroupName,
    serverName,
    privateEndpointConnectionName,
    parameters
  );
  console.log(result);
}

updatePrivateEndpointConnectionTags().catch(console.error);
