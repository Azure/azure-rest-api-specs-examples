const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets the managementpolicy to the specified storage account.
 *
 * @summary Sets the managementpolicy to the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountSetManagementPolicy_LastAccessTimeBasedBlobActions.json
 */
async function storageAccountSetManagementPolicyLastAccessTimeBasedBlobActions() {
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
                delete: { daysAfterLastAccessTimeGreaterThan: 1000 },
                enableAutoTierToHotFromCool: true,
                tierToArchive: { daysAfterLastAccessTimeGreaterThan: 90 },
                tierToCool: { daysAfterLastAccessTimeGreaterThan: 30 },
              },
              snapshot: { delete: { daysAfterCreationGreaterThan: 30 } },
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

storageAccountSetManagementPolicyLastAccessTimeBasedBlobActions().catch(console.error);
