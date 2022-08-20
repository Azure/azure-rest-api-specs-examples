const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the blob inventory policy associated with the specified storage account.
 *
 * @summary Gets the blob inventory policy associated with the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/StorageAccountGetBlobInventoryPolicy.json
 */
async function storageAccountGetBlobInventoryPolicy() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res7687";
  const accountName = "sto9699";
  const blobInventoryPolicyName = "default";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobInventoryPolicies.get(
    resourceGroupName,
    accountName,
    blobInventoryPolicyName
  );
  console.log(result);
}

storageAccountGetBlobInventoryPolicy().catch(console.error);
