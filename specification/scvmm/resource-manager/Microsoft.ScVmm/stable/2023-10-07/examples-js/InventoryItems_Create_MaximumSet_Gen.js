const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Or Update InventoryItem.
 *
 * @summary Create Or Update InventoryItem.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/InventoryItems_Create_MaximumSet_Gen.json
 */
async function inventoryItemsCreateMaximumSet() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "79332E5A-630B-480F-A266-A941C015AB19";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "rgscvmm";
  const vmmServerName = "O";
  const inventoryItemResourceName = "1BdDc2Ab-bDd9-Ebd6-bfdb-C0dbbdB5DEDf";
  const resource = {
    kind: "M\\d_,V.",
    properties: { inventoryType: "Cloud" },
  };
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
