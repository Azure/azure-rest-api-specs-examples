const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update the restore point collection. Please refer to https://aka.ms/RestorePoints for more details. When updating a restore point collection, only tags may be modified.
 *
 * @summary the operation to create or update the restore point collection. Please refer to https://aka.ms/RestorePoints for more details. When updating a restore point collection, only tags may be modified.
 * x-ms-original-file: 2025-11-01/restorePointExamples/RestorePointCollection_CreateOrUpdate_ForCrossRegionCopy.json
 */
async function createOrUpdateARestorePointCollectionForCrossRegionCopy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePointCollections.createOrUpdate("myResourceGroup", "myRpc", {
    location: "norwayeast",
    source: {
      id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/sourceRpcName",
    },
    tags: { myTag1: "tagValue1" },
  });
  console.log(result);
}
