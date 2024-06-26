const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Sql pool's security alert policy.
 *
 * @summary Get a Sql pool's security alert policy.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolSecurityAlert.json
 */
async function getASecurityAlertOfASqlAnalyticsPool() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "securityalert-6852";
  const workspaceName = "securityalert-2080";
  const sqlPoolName = "testdb";
  const securityAlertPolicyName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolSecurityAlertPolicies.get(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    securityAlertPolicyName
  );
  console.log(result);
}
