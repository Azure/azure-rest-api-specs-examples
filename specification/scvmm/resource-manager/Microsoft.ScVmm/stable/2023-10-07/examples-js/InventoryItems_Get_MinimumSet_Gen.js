const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Shows an inventory item.
 *
 * @summary Shows an inventory item.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/InventoryItems_Get_MinimumSet_Gen.json
 */
async function inventoryItemsGetMinimumSet() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "79332E5A-630B-480F-A266-A941C015AB19";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "rgscvmm";
  const vmmServerName = "_";
  const inventoryItemResourceName = "cacb8Ceb-efAC-bebb-ae7C-dec8C5Bb7100";
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential, subscriptionId);
  const result = await client.inventoryItems.get(
    resourceGroupName,
    vmmServerName,
    inventoryItemResourceName,
  );
  console.log(result);
}
