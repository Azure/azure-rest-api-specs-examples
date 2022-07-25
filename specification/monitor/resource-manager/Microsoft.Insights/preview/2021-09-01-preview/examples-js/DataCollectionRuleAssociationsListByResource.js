const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists associations for the specified resource.
 *
 * @summary Lists associations for the specified resource.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/DataCollectionRuleAssociationsListByResource.json
 */
async function listAssociationsForSpecifiedResource() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceUri =
    "subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dataCollectionRuleAssociations.listByResource(resourceUri)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAssociationsForSpecifiedResource().catch(console.error);
