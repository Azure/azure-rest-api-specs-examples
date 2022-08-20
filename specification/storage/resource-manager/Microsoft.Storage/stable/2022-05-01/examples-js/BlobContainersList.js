const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all containers and does not support a prefix like data plane. Also SRP today does not return continuation token.
 *
 * @summary Lists all containers and does not support a prefix like data plane. Also SRP today does not return continuation token.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/BlobContainersList.json
 */
async function listContainers() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9290";
  const accountName = "sto1590";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.blobContainers.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listContainers().catch(console.error);
