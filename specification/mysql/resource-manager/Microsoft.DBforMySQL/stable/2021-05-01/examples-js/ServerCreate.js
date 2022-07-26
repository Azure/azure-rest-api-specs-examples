const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new server or updates an existing server. The update action will overwrite the existing server.
 *
 * @summary Creates a new server or updates an existing server. The update action will overwrite the existing server.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ServerCreate.json
 */
async function createANewServer() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testrg";
  const serverName = "mysqltestserver";
  const parameters = {
    administratorLogin: "cloudsa",
    administratorLoginPassword: "your_password",
    availabilityZone: "1",
    backup: { backupRetentionDays: 7, geoRedundantBackup: "Disabled" },
    createMode: "Default",
    highAvailability: { mode: "ZoneRedundant", standbyAvailabilityZone: "3" },
    location: "southeastasia",
    sku: { name: "Standard_D2ds_v4", tier: "GeneralPurpose" },
    storage: { autoGrow: "Disabled", iops: 600, storageSizeGB: 100 },
    tags: { num: "1" },
    version: "5.7",
  };
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.beginCreateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}

createANewServer().catch(console.error);
