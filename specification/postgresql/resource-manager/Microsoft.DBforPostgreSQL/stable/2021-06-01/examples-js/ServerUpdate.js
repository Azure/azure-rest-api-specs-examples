const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing server. The request body can contain one to many of the properties present in the normal server definition.
 *
 * @summary Updates an existing server. The request body can contain one to many of the properties present in the normal server definition.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/ServerUpdate.json
 */
async function serverUpdate() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "TestGroup";
  const serverName = "pgtestsvc4";
  const parameters = {
    administratorLoginPassword: "newpassword",
    backup: { backupRetentionDays: 20 },
    createMode: "Update",
    location: "westus",
    sku: { name: "Standard_D8s_v3", tier: "GeneralPurpose" },
    storage: { storageSizeGB: 1024 },
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.beginUpdateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}

serverUpdate().catch(console.error);
