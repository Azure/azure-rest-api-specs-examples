const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an extended Sql pool's blob auditing policy.
 *
 * @summary Gets an extended Sql pool's blob auditing policy.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ExtendedSqlPoolBlobAuditingGet.json
 */
async function getAnExtendedDatabaseBlobAuditingPolicy() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "blobauditingtest-6852";
  const workspaceName = "blobauditingtest-2080";
  const sqlPoolName = "testdb";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.extendedSqlPoolBlobAuditingPolicies.get(
    resourceGroupName,
    workspaceName,
    sqlPoolName
  );
  console.log(result);
}
