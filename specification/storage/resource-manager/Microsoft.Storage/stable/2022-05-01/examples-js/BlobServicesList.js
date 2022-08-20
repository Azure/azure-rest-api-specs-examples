const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List blob services of storage account. It returns a collection of one object named default.
 *
 * @summary List blob services of storage account. It returns a collection of one object named default.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/BlobServicesList.json
 */
async function listBlobServices() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4410";
  const accountName = "sto8607";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.blobServices.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listBlobServices().catch(console.error);
