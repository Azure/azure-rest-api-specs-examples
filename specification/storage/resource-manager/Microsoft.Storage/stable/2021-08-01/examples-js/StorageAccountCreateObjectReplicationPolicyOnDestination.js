const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountCreateObjectReplicationPolicyOnDestination() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res7687";
  const accountName = "dst112";
  const objectReplicationPolicyId = "default";
  const properties = {
    destinationAccount: "dst112",
    rules: [
      {
        destinationContainer: "dcont139",
        filters: { prefixMatch: ["blobA", "blobB"] },
        sourceContainer: "scont139",
      },
    ],
    sourceAccount: "src1122",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.objectReplicationPoliciesOperations.createOrUpdate(
    resourceGroupName,
    accountName,
    objectReplicationPolicyId,
    properties
  );
  console.log(result);
}

storageAccountCreateObjectReplicationPolicyOnDestination().catch(console.error);
