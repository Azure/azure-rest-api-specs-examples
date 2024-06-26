const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an inventoryItem.
 *
 * @summary Deletes an inventoryItem.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/InventoryItems_Delete_MaximumSet_Gen.json
 */
async function inventoryItemsDeleteMaximumSet() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "79332E5A-630B-480F-A266-A941C015AB19";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "rgscvmm";
  const vmmServerName = "b";
  const inventoryItemResourceName = "EcECadfd-Eaaa-e5Ce-ebdA-badeEd3c6af1";
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential, subscriptionId);
  const result = await client.inventoryItems.delete(
    resourceGroupName,
    vmmServerName,
    inventoryItemResourceName,
  );
  console.log(result);
}
