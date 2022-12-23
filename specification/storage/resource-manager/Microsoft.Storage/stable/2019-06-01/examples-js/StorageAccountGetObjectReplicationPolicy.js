const { StorageManagementClient } = require("@azure/arm-storage-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the object replication policy of the storage account by policy ID.
 *
 * @summary Get the object replication policy of the storage account by policy ID.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2019-06-01/examples/StorageAccountGetObjectReplicationPolicy.json
 */
async function storageAccountGetObjectReplicationPolicies() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res6977";
  const accountName = "sto2527";
  const objectReplicationPolicyId = "{objectReplicationPolicy-Id}";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.objectReplicationPoliciesOperations.get(
    resourceGroupName,
    accountName,
    objectReplicationPolicyId
  );
  console.log(result);
}
