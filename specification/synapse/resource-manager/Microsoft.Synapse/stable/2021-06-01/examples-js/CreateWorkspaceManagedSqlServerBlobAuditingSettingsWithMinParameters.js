const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or Update a workspace managed sql server's blob auditing policy.
 *
 * @summary Create or Update a workspace managed sql server's blob auditing policy.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateWorkspaceManagedSqlServerBlobAuditingSettingsWithMinParameters.json
 */
async function createOrUpdateBlobAuditingPolicyOfWorkspaceManagedSqlServerWithMinimalParameters() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "wsg-7398";
  const workspaceName = "testWorkspace";
  const blobAuditingPolicyName = "default";
  const parameters = {
    state: "Enabled",
    storageAccountAccessKey:
      "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==",
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
