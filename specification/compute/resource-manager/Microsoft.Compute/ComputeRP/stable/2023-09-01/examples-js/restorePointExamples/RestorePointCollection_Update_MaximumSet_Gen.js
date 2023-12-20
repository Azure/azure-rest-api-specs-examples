const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update the restore point collection.
 *
 * @summary The operation to update the restore point collection.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/restorePointExamples/RestorePointCollection_Update_MaximumSet_Gen.json
 */
async function restorePointCollectionUpdateMaximumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const restorePointCollectionName = "aaaaaaaaaaaaaaaaaaaa";
  const parameters = {
    source: {
      id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM",
    },
    tags: { key8536: "aaaaaaaaaaaaaaaaaaa" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePointCollections.update(
    resourceGroupName,
    restorePointCollectionName,
    parameters,
  );
  console.log(result);
}
