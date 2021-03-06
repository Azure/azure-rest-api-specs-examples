const { PostgreSQLManagementClient } = require("@azure/arm-postgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing server. The request body can contain one to many of the properties present in the normal server definition.
 *
 * @summary Updates an existing server. The request body can contain one to many of the properties present in the normal server definition.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerUpdate.json
 */
async function serverUpdate() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testrg";
  const serverName = "pgtestsvc4";
  const parameters = {
    administratorLoginPassword: "<administratorLoginPassword>",
    minimalTlsVersion: "TLS1_2",
    sslEnforcement: "Enabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementClient(credential, subscriptionId);
  const result = await client.servers.beginUpdateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}

serverUpdate().catch(console.error);
