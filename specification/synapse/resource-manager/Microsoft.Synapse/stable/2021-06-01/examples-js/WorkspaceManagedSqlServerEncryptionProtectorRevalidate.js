const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Revalidates workspace managed sql server's existing encryption protector.
 *
 * @summary Revalidates workspace managed sql server's existing encryption protector.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/WorkspaceManagedSqlServerEncryptionProtectorRevalidate.json
 */
async function revalidatesTheEncryptionProtector() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "wsg-7398";
  const workspaceName = "testWorkspace";
  const encryptionProtectorName = "current";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.workspaceManagedSqlServerEncryptionProtector.beginRevalidateAndWait(
    resourceGroupName,
    workspaceName,
    encryptionProtectorName
  );
  console.log(result);
}

revalidatesTheEncryptionProtector().catch(console.error);
