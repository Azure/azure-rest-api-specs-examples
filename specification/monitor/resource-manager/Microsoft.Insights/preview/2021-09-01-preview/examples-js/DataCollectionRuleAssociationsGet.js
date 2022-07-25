const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the specified association.
 *
 * @summary Returns the specified association.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/DataCollectionRuleAssociationsGet.json
 */
async function getAssociation() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceUri =
    "subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm";
  const associationName = "myAssociation";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.dataCollectionRuleAssociations.get(resourceUri, associationName);
  console.log(result);
}

getAssociation().catch(console.error);
