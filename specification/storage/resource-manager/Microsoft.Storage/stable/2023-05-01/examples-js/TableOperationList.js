const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all the tables under the specified storage account
 *
 * @summary Gets a list of all the tables under the specified storage account
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/TableOperationList.json
 */
async function tableOperationList() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res9290";
  const accountName = "sto328";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.tableOperations.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
