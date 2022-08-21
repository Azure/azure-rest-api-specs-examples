const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the managementpolicy associated with the specified storage account.
 *
 * @summary Deletes the managementpolicy associated with the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/StorageAccountDeleteManagementPolicy.json
 */
async function storageAccountDeleteManagementPolicies() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const accountName = "sto2527";
  const managementPolicyName = "default";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.managementPolicies.delete(
    resourceGroupName,
    accountName,
    managementPolicyName
  );
  console.log(result);
}

storageAccountDeleteManagementPolicies().catch(console.error);
