const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Or Update InventoryItem.
 *
 * @summary Create Or Update InventoryItem.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateInventoryItem.json
 */
async function createInventoryItem() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "testrg";
  const vmmServerName = "ContosoVMMServer";
  const inventoryItemName = "12345678-1234-1234-1234-123456789abc";
  const body = { inventoryType: "Cloud" };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const result = await client.inventoryItems.create(
    resourceGroupName,
    vmmServerName,
    inventoryItemName,
    options
  );
  console.log(result);
}
