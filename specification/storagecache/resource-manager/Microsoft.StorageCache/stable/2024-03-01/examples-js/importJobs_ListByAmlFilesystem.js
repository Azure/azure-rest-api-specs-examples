const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns all import jobs the user has access to under an AML File System.
 *
 * @summary Returns all import jobs the user has access to under an AML File System.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/importJobs_ListByAmlFilesystem.json
 */
async function importJobsListByAmlFilesystem() {
  const subscriptionId =
    process.env["STORAGECACHE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["STORAGECACHE_RESOURCE_GROUP"] || "scgroup";
  const amlFilesystemName = "fs1";
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.importJobs.listByAmlFilesystem(
    resourceGroupName,
    amlFilesystemName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
