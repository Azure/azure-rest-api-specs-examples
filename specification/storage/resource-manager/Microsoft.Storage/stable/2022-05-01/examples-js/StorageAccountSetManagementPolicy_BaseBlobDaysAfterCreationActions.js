const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets the managementpolicy to the specified storage account.
 *
 * @summary Sets the managementpolicy to the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/StorageAccountSetManagementPolicy_BaseBlobDaysAfterCreationActions.json
 */
async function storageAccountSetManagementPolicyBaseBlobDaysAfterCreationActions() {
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
              baseBlob: {
                delete: { daysAfterCreationGreaterThan: 1000 },
                tierToArchive: { daysAfterCreationGreaterThan: 90 },
                tierToCool: { daysAfterCreationGreaterThan: 30 },
              },
            },
            filters: {
              blobTypes: ["blockBlob"],
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

storageAccountSetManagementPolicyBaseBlobDaysAfterCreationActions().catch(console.error);
