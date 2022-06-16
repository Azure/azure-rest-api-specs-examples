const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the managementpolicy associated with the specified storage account.
 *
 * @summary Gets the managementpolicy associated with the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountGetManagementPolicy.json
 */
async function storageAccountGetManagementPolicies() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const accountName = "sto2527";
  const managementPolicyName = "default";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.managementPolicies.get(
    resourceGroupName,
    accountName,
    managementPolicyName
  );
  console.log(result);
}

storageAccountGetManagementPolicies().catch(console.error);
