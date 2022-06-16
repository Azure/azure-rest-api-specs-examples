const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountGetEncryptionScope() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "resource-group-name";
  const accountName = "{storage-account-name}";
  const encryptionScopeName = "{encryption-scope-name}";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.encryptionScopes.get(
    resourceGroupName,
    accountName,
    encryptionScopeName
  );
  console.log(result);
}

storageAccountGetEncryptionScope().catch(console.error);
