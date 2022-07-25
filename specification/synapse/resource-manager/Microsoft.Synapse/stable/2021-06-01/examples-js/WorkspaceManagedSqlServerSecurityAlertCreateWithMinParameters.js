const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or Update a workspace managed sql server's threat detection policy.
 *
 * @summary Create or Update a workspace managed sql server's threat detection policy.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/WorkspaceManagedSqlServerSecurityAlertCreateWithMinParameters.json
 */
async function updateAWorkspaceManagedSqlServerThreatDetectionPolicyWithMinimalParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "wsg-7398";
  const workspaceName = "testWorkspace";
  const securityAlertPolicyName = "Default";
  const parameters = {
    emailAccountAdmins: true,
    state: "Disabled",
    storageAccountAccessKey:
      "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==",
    storageEndpoint: "https://mystorage.blob.core.windows.net",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result =
    await client.workspaceManagedSqlServerSecurityAlertPolicy.beginCreateOrUpdateAndWait(
      resourceGroupName,
      workspaceName,
      securityAlertPolicyName,
      parameters
    );
  console.log(result);
}

updateAWorkspaceManagedSqlServerThreatDetectionPolicyWithMinimalParameters().catch(console.error);
