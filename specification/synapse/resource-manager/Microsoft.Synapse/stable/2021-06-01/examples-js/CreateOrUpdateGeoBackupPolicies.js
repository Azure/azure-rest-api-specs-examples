const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a SQL Pool geo backup policy.
 *
 * @summary Updates a SQL Pool geo backup policy.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateGeoBackupPolicies.json
 */
async function createGeoBackupPolicy() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "testrg";
  const workspaceName = "testws";
  const sqlPoolName = "testdw";
  const geoBackupPolicyName = "Default";
  const parameters = { state: "Enabled" };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolGeoBackupPolicies.createOrUpdate(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    geoBackupPolicyName,
    parameters
  );
  console.log(result);
}
