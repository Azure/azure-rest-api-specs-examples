const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountRevokeUserDelegationKeys() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4167";
  const accountName = "sto3539";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.revokeUserDelegationKeys(
    resourceGroupName,
    accountName
  );
  console.log(result);
}

storageAccountRevokeUserDelegationKeys().catch(console.error);
