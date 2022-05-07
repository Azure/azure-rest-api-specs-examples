Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets the managementpolicy to the specified storage account.
 *
 * @summary Sets the managementpolicy to the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountSetManagementPolicy_LastTierChangeTimeActions.json
 */
async function storageAccountSetManagementPolicyLastTierChangeTimeActions() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res7687";
  const accountName = "sto9699";
  const managementPolicyName = "default";
  const properties = {
    policy: {
      rules: [
        {
          name: "olcmtest",
          type: "Lifecycle",
          definition: {
            actions: {
              baseBlob: {
                delete: { daysAfterModificationGreaterThan: 1000 },
                tierToArchive: {
                  daysAfterLastTierChangeGreaterThan: 120,
                  daysAfterModificationGreaterThan: 90,
                },
                tierToCool: { daysAfterModificationGreaterThan: 30 },
              },
              snapshot: {
                tierToArchive: {
                  daysAfterCreationGreaterThan: 30,
                  daysAfterLastTierChangeGreaterThan: 90,
                },
              },
              version: {
                tierToArchive: {
                  daysAfterCreationGreaterThan: 30,
                  daysAfterLastTierChangeGreaterThan: 90,
                },
              },
            },
            filters: {
              blobTypes: ["blockBlob"],
              prefixMatch: ["olcmtestcontainer"],
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

storageAccountSetManagementPolicyLastTierChangeTimeActions().catch(console.error);
```
