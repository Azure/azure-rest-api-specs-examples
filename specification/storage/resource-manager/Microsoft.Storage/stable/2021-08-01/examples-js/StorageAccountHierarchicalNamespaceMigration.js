const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountHierarchicalNamespaceMigration() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4228";
  const accountName = "sto2434";
  const requestType = "HnsOnValidationRequest";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.beginHierarchicalNamespaceMigrationAndWait(
    resourceGroupName,
    accountName,
    requestType
  );
  console.log(result);
}

storageAccountHierarchicalNamespaceMigration().catch(console.error);
