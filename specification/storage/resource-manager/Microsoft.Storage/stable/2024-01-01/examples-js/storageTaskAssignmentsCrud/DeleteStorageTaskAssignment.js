const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Delete the storage task assignment sub-resource
 *
 * @summary Delete the storage task assignment sub-resource
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/storageTaskAssignmentsCrud/DeleteStorageTaskAssignment.json
 */
async function deleteStorageTaskAssignment() {
  const subscriptionId =
    process.env["STORAGE_SUBSCRIPTION_ID"] || "1f31ba14-ce16-4281-b9b4-3e78da6e1616";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res4228";
  const accountName = "sto4445";
  const storageTaskAssignmentName = "myassignment1";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageTaskAssignments.beginDeleteAndWait(
    resourceGroupName,
    accountName,
    storageTaskAssignmentName,
  );
  console.log(result);
}
