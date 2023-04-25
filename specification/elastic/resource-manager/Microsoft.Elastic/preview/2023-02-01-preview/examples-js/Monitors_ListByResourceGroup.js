const { MicrosoftElastic } = require("@azure/arm-elastic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all monitors under the specified resource group.
 *
 * @summary List all monitors under the specified resource group.
 * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/Monitors_ListByResourceGroup.json
 */
async function monitorsListByResourceGroup() {
  const subscriptionId =
    process.env["ELASTIC_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ELASTIC_RESOURCE_GROUP"] || "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftElastic(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.monitors.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
