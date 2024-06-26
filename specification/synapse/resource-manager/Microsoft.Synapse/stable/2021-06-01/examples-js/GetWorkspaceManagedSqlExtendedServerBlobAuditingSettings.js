const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a workspace SQL server's extended blob auditing policy.
 *
 * @summary Get a workspace SQL server's extended blob auditing policy.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlExtendedServerBlobAuditingSettings.json
 */
async function getWorkspaceManagedSqlServersExtendedBlobAuditingSettings() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "wsg-7398";
  const workspaceName = "testWorkspace";
  const blobAuditingPolicyName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.workspaceManagedSqlServerExtendedBlobAuditingPolicies.get(
    resourceGroupName,
    workspaceName,
    blobAuditingPolicyName
  );
  console.log(result);
}
