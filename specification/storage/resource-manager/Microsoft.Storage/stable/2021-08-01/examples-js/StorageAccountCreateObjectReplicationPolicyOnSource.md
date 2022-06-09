```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountCreateObjectReplicationPolicyOnSource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res7687";
  const accountName = "src1122";
  const objectReplicationPolicyId = "2a20bb73-5717-4635-985a-5d4cf777438f";
  const properties = {
    destinationAccount: "dst112",
    rules: [
      {
        destinationContainer: "dcont139",
        filters: {
          minCreationTime: "2020-02-19T16:05:00Z",
          prefixMatch: ["blobA", "blobB"],
        },
        ruleId: "d5d18a48-8801-4554-aeaa-74faf65f5ef9",
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

storageAccountCreateObjectReplicationPolicyOnSource().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.
