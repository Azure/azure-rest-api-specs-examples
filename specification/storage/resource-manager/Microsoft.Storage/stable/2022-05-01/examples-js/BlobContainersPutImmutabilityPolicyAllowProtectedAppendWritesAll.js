const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an unlocked immutability policy. ETag in If-Match is honored if given but not required for this operation.
 *
 * @summary Creates or updates an unlocked immutability policy. ETag in If-Match is honored if given but not required for this operation.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/BlobContainersPutImmutabilityPolicyAllowProtectedAppendWritesAll.json
 */
async function createOrUpdateImmutabilityPolicyWithAllowProtectedAppendWritesAll() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res1782";
  const accountName = "sto7069";
  const containerName = "container6397";
  const parameters = {
    allowProtectedAppendWritesAll: true,
    immutabilityPeriodSinceCreationInDays: 3,
  };
  const options = {
    parameters,
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.createOrUpdateImmutabilityPolicy(
    resourceGroupName,
    accountName,
    containerName,
    options
  );
  console.log(result);
}

createOrUpdateImmutabilityPolicyWithAllowProtectedAppendWritesAll().catch(console.error);
