const { EasmMgmtClient } = require("@azure/arm-defendereasm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of workspaces in the given resource group.
 *
 * @summary Returns a list of workspaces in the given resource group.
 * x-ms-original-file: specification/riskiq/resource-manager/Microsoft.Easm/preview/2023-04-01-preview/examples/Workspaces_ListByResourceGroup.json
 */
async function workspaces() {
  const subscriptionId =
    process.env["DEFENDEREASM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["DEFENDEREASM_RESOURCE_GROUP"] || "dummyrg";
  const credential = new DefaultAzureCredential();
  const client = new EasmMgmtClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaces.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
