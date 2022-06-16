const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the blob inventory policy associated with the specified storage account.
 *
 * @summary Deletes the blob inventory policy associated with the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountDeleteBlobInventoryPolicy.json
 */
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
