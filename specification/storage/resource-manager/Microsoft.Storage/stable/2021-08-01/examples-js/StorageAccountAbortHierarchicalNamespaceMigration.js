const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountAbortHierarchicalNamespaceMigration() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4228";
  const accountName = "sto2434";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.beginAbortHierarchicalNamespaceMigrationAndWait(
    resourceGroupName,
    accountName
  );
  console.log(result);
}

storageAccountAbortHierarchicalNamespaceMigration().catch(console.error);
