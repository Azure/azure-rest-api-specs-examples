const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a server's DevOps audit settings.
 *
 * @summary Creates or updates a server's DevOps audit settings.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ServerDevOpsAuditCreateMin.json
 */
async function updateAServerDevOpsAuditSettingsWithMinimalInput() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "devAuditTestRG";
  const serverName = "devOpsAuditTestSvr";
  const devOpsAuditingSettingsName = "Default";
  const parameters = {
    state: "Enabled",
    storageAccountAccessKey: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
    storageEndpoint: "https://mystorage.blob.core.windows.net",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.serverDevOpsAuditSettings.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    devOpsAuditingSettingsName,
    parameters,
  );
  console.log(result);
}
