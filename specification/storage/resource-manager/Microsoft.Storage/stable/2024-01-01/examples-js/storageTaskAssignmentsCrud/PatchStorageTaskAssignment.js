const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update storage task assignment properties
 *
 * @summary Update storage task assignment properties
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/storageTaskAssignmentsCrud/PatchStorageTaskAssignment.json
 */
async function patchStorageTaskAssignment() {
  const subscriptionId =
    process.env["STORAGE_SUBSCRIPTION_ID"] || "1f31ba14-ce16-4281-b9b4-3e78da6e1616";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res4228";
  const accountName = "sto4445";
  const storageTaskAssignmentName = "myassignment1";
  const parameters = {
    properties: {
      description: "My Storage task assignment",
      enabled: true,
      executionContext: {
        target: { excludePrefix: [], prefix: ["prefix1", "prefix2"] },
        trigger: {
          type: "RunOnce",
          parameters: { startOn: new Date("2022-11-15T21:52:47.8145095Z") },
        },
      },
      report: { prefix: "container1" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageTaskAssignments.beginUpdateAndWait(
    resourceGroupName,
    accountName,
    storageTaskAssignmentName,
    parameters,
  );
  console.log(result);
}
