const { StorageManagementClient } = require("@azure/arm-storage-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Revoke user delegation keys.
 *
 * @summary Revoke user delegation keys.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2019-06-01/examples/StorageAccountRevokeUserDelegationKeys.json
 */
async function storageAccountRevokeUserDelegationKeys() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res4167";
  const accountName = "sto3539";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.revokeUserDelegationKeys(
    resourceGroupName,
    accountName
  );
  console.log(result);
}
