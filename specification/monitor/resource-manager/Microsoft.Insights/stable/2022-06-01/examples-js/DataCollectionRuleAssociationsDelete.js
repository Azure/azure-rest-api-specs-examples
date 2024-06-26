const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an association.
 *
 * @summary Deletes an association.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/DataCollectionRuleAssociationsDelete.json
 */
async function deleteAssociation() {
  const resourceUri =
    "subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm";
  const associationName = "myAssociation";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const result = await client.dataCollectionRuleAssociations.delete(resourceUri, associationName);
  console.log(result);
}
