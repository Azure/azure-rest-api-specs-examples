const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Revoke user delegation keys.
 *
 * @summary Revoke user delegation keys.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountRevokeUserDelegationKeys.json
 */
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
