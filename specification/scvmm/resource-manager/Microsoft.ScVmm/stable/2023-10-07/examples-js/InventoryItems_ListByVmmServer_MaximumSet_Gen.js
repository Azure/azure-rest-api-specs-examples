const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of inventoryItems in the given VmmServer.
 *
 * @summary Returns the list of inventoryItems in the given VmmServer.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/InventoryItems_ListByVmmServer_MaximumSet_Gen.json
 */
async function inventoryItemsListByVmmServerMaximumSet() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "79332E5A-630B-480F-A266-A941C015AB19";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "rgscvmm";
  const vmmServerName = "X";
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.inventoryItems.listByVmmServer(resourceGroupName, vmmServerName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
