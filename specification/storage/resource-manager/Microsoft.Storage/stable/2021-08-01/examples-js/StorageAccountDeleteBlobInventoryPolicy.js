const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountDeleteBlobInventoryPolicy() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const accountName = "sto2527";
  const blobInventoryPolicyName = "default";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobInventoryPolicies.delete(
    resourceGroupName,
    accountName,
    blobInventoryPolicyName
  );
  console.log(result);
}

storageAccountDeleteBlobInventoryPolicy().catch(console.error);
