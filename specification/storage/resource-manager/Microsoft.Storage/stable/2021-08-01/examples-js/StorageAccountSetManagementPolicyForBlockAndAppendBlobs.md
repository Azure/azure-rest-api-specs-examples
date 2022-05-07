Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountSetManagementPolicyForBlockAndAppendBlobs() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res7687";
  const accountName = "sto9699";
  const managementPolicyName = "default";
  const properties = {
    policy: {
      rules: [
        {
          name: "olcmtest1",
          type: "Lifecycle",
          definition: {
            actions: {
              baseBlob: { delete: { daysAfterModificationGreaterThan: 90 } },
              snapshot: { delete: { daysAfterCreationGreaterThan: 90 } },
              version: { delete: { daysAfterCreationGreaterThan: 90 } },
            },
            filters: {
              blobTypes: ["blockBlob", "appendBlob"],
              prefixMatch: ["olcmtestcontainer1"],
            },
          },
          enabled: true,
        },
      ],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.managementPolicies.createOrUpdate(
    resourceGroupName,
    accountName,
    managementPolicyName,
    properties
  );
  console.log(result);
}

storageAccountSetManagementPolicyForBlockAndAppendBlobs().catch(console.error);
```
