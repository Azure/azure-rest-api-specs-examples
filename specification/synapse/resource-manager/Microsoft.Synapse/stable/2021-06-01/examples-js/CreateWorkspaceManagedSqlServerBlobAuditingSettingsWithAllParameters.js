const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or Update a workspace managed sql server's blob auditing policy.
 *
 * @summary Create or Update a workspace managed sql server's blob auditing policy.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateWorkspaceManagedSqlServerBlobAuditingSettingsWithAllParameters.json
 */
async function createOrUpdateBlobAuditingPolicyOfWorkspaceSqlServerWithAllParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "wsg-7398";
  const workspaceName = "testWorkspace";
  const blobAuditingPolicyName = "default";
  const parameters = {
    auditActionsAndGroups: [
      "SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP",
      "FAILED_DATABASE_AUTHENTICATION_GROUP",
      "BATCH_COMPLETED_GROUP",
    ],
    isAzureMonitorTargetEnabled: true,
    isStorageSecondaryKeyInUse: false,
    queueDelayMs: 4000,
    retentionDays: 6,
    state: "Enabled",
    storageAccountAccessKey:
      "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==",
    storageAccountSubscriptionId: "00000000-1234-0000-5678-000000000000",
    storageEndpoint: "https://mystorage.blob.core.windows.net",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result =
    await client.workspaceManagedSqlServerBlobAuditingPolicies.beginCreateOrUpdateAndWait(
      resourceGroupName,
      workspaceName,
      blobAuditingPolicyName,
      parameters
    );
  console.log(result);
}

createOrUpdateBlobAuditingPolicyOfWorkspaceSqlServerWithAllParameters().catch(console.error);
