const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an extended server's blob auditing policy.
 *
 * @summary Gets an extended server's blob auditing policy.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ExtendedServerBlobAuditingGet.json
 */
async function getAServerBlobExtendedAuditingPolicy() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "blobauditingtest-4799";
  const serverName = "blobauditingtest-6440";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.extendedServerBlobAuditingPolicies.get(resourceGroupName, serverName);
  console.log(result);
}

getAServerBlobExtendedAuditingPolicy().catch(console.error);
