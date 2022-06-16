const { PostgreSQLManagementClient } = require("@azure/arm-postgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new server, or will overwrite an existing server.
 *
 * @summary Creates a new server, or will overwrite an existing server.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerCreate.json
 */
async function createANewServer() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "TestGroup";
  const serverName = "pgtestsvc4";
  const parameters = {
    location: "westus",
    properties: {
      administratorLogin: "cloudsa",
      administratorLoginPassword: "<administratorLoginPassword>",
      createMode: "Default",
      minimalTlsVersion: "TLS1_2",
      sslEnforcement: "Enabled",
      storageProfile: {
        backupRetentionDays: 7,
        geoRedundantBackup: "Disabled",
        storageMB: 128000,
      },
    },
    sku: { name: "B_Gen5_2", capacity: 2, family: "Gen5", tier: "Basic" },
    tags: { elasticServer: "1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementClient(credential, subscriptionId);
  const result = await client.servers.beginCreateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}

createANewServer().catch(console.error);
