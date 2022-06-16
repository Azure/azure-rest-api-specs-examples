const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new server or updates an existing server. The update action will overwrite the existing server.
 *
 * @summary Creates a new server or updates an existing server. The update action will overwrite the existing server.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerCreate.json
 */
async function createANewServer() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testrg";
  const serverName = "mysqltestsvc4";
  const parameters = {
    location: "westus",
    properties: {
      administratorLogin: "cloudsa",
      administratorLoginPassword: "<administratorLoginPassword>",
      createMode: "Default",
      sslEnforcement: "Enabled",
      storageProfile: {
        backupRetentionDays: 7,
        geoRedundantBackup: "Enabled",
        storageMB: 128000,
      },
    },
    sku: {
      name: "GP_Gen5_2",
      capacity: 2,
      family: "Gen5",
      tier: "GeneralPurpose",
    },
    tags: { elasticServer: "1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.servers.beginCreateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}

createANewServer().catch(console.error);
