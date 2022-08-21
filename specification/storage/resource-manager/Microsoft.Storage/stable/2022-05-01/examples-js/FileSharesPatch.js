const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates share properties as specified in request body. Properties not mentioned in the request will not be changed. Update fails if the specified share does not already exist.
 *
 * @summary Updates share properties as specified in request body. Properties not mentioned in the request will not be changed. Update fails if the specified share does not already exist.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/FileSharesPatch.json
 */
async function updateShares() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res3376";
  const accountName = "sto328";
  const shareName = "share6185";
  const fileShare = { metadata: { type: "image" } };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.update(
    resourceGroupName,
    accountName,
    shareName,
    fileShare
  );
  console.log(result);
}

updateShares().catch(console.error);
