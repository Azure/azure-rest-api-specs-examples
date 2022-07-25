const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates workspace managed sql server's encryption protector.
 *
 * @summary Updates workspace managed sql server's encryption protector.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateServiceManaged.json
 */
async function updateTheEncryptionProtectorToServiceManaged() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "wsg-7398";
  const workspaceName = "testWorkspace";
  const encryptionProtectorName = "current";
  const parameters = {
    serverKeyName: "ServiceManaged",
    serverKeyType: "ServiceManaged",
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

updateTheEncryptionProtectorToServiceManaged().catch(console.error);
