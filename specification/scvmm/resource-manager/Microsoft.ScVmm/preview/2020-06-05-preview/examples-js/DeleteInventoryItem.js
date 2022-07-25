const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an inventoryItem.
 *
 * @summary Deletes an inventoryItem.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/DeleteInventoryItem.json
 */
async function deleteInventoryItem() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "testrg";
  const vmmServerName = "ContosoVMMServer";
  const inventoryItemName = "12345678-1234-1234-1234-123456789abc";
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const result = await client.inventoryItems.delete(
    resourceGroupName,
    vmmServerName,
    inventoryItemName
  );
  console.log(result);
}

deleteInventoryItem().catch(console.error);
