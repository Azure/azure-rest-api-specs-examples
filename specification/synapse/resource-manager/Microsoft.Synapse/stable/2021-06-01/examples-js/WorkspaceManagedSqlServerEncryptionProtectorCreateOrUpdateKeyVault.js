const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates workspace managed sql server's encryption protector.
 *
 * @summary Updates workspace managed sql server's encryption protector.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateKeyVault.json
 */
async function updateTheEncryptionProtectorToKeyVault() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "wsg-7398";
  const workspaceName = "testWorkspace";
  const encryptionProtectorName = "current";
  const parameters = {
    serverKeyName: "someVault_someKey_01234567890123456789012345678901",
    serverKeyType: "AzureKeyVault",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result =
    await client.workspaceManagedSqlServerEncryptionProtector.beginCreateOrUpdateAndWait(
      resourceGroupName,
      workspaceName,
      encryptionProtectorName,
      parameters
    );
  console.log(result);
}
