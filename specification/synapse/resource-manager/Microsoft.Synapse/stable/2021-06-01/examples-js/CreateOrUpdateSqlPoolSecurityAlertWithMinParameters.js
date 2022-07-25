const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Sql pool's security alert policy.
 *
 * @summary Create or update a Sql pool's security alert policy.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolSecurityAlertWithMinParameters.json
 */
async function updateASqlPoolThreatDetectionPolicyWithMinimalParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "securityalert-4799";
  const workspaceName = "securityalert-6440";
  const sqlPoolName = "testdb";
  const securityAlertPolicyName = "default";
  const parameters = { state: "Enabled" };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolSecurityAlertPolicies.createOrUpdate(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    securityAlertPolicyName,
    parameters
  );
  console.log(result);
}

updateASqlPoolThreatDetectionPolicyWithMinimalParameters().catch(console.error);
