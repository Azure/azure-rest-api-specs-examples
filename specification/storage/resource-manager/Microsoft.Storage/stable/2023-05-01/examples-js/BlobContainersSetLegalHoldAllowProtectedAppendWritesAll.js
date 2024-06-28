const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets legal hold tags. Setting the same tag results in an idempotent operation. SetLegalHold follows an append pattern and does not clear out the existing tags that are not specified in the request.
 *
 * @summary Sets legal hold tags. Setting the same tag results in an idempotent operation. SetLegalHold follows an append pattern and does not clear out the existing tags that are not specified in the request.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/BlobContainersSetLegalHoldAllowProtectedAppendWritesAll.json
 */
async function setLegalHoldContainersWithAllowProtectedAppendWritesAll() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res4303";
  const accountName = "sto7280";
  const containerName = "container8723";
  const legalHold = {
    allowProtectedAppendWritesAll: true,
    tags: ["tag1", "tag2", "tag3"],
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.setLegalHold(
    resourceGroupName,
    accountName,
    containerName,
    legalHold,
  );
  console.log(result);
}
