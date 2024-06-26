const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Or Update InventoryItem.
 *
 * @summary Create Or Update InventoryItem.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/InventoryItems_Create_MinimumSet_Gen.json
 */
async function inventoryItemsCreateMinimumSet() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "79332E5A-630B-480F-A266-A941C015AB19";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "rgscvmm";
  const vmmServerName = ".";
  const inventoryItemResourceName = "bbFb0cBb-50ce-4bfc-3eeD-bC26AbCC257a";
  const resource = {};
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential, subscriptionId);
  const result = await client.inventoryItems.create(
    resourceGroupName,
    vmmServerName,
    inventoryItemResourceName,
    resource,
  );
  console.log(result);
}
