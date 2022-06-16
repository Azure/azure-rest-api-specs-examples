const { MicrosoftDatadogClient } = require("@azure/arm-datadog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all monitors under the specified resource group.
 *
 * @summary List all monitors under the specified resource group.
 * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/Monitors_ListByResourceGroup.json
 */
async function monitorsListByResourceGroup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftDatadogClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.monitors.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

monitorsListByResourceGroup().catch(console.error);
