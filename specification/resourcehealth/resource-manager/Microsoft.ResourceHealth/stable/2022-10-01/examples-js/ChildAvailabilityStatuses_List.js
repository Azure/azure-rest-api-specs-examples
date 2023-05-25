const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the historical availability statuses for a single child resource. Use the nextLink property in the response to get the next page of availability status
 *
 * @summary Lists the historical availability statuses for a single child resource. Use the nextLink property in the response to get the next page of availability status
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/ChildAvailabilityStatuses_List.json
 */
async function getChildHealthHistoryByResource() {
  const subscriptionId =
    process.env["RESOURCEHEALTH_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceUri =
    "subscriptions/227b734f-e14f-4de6-b7fc-3190c21e69f6/resourceGroups/JUHACKETRHCTEST/providers/Microsoft.Compute/virtualMachineScaleSets/rhctest/virtualMachines/4";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.childAvailabilityStatuses.list(resourceUri)) {
    resArray.push(item);
  }
  console.log(resArray);
}
