const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists extended auditing settings of a server.
 *
 * @summary Lists extended auditing settings of a server.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerExtendedAuditingSettingsList.json
 */
async function listExtendedAuditingSettingsOfAServer() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "blobauditingtest-4799";
  const serverName = "blobauditingtest-6440";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.extendedServerBlobAuditingPolicies.listByServer(
    resourceGroupName,
    serverName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listExtendedAuditingSettingsOfAServer().catch(console.error);
