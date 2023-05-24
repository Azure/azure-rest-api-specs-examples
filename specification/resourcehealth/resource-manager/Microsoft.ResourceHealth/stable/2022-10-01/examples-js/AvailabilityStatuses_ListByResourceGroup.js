const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the current availability status for all the resources in the resource group.
 *
 * @summary Lists the current availability status for all the resources in the resource group.
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/AvailabilityStatuses_ListByResourceGroup.json
 */
async function listByResourceGroup() {
  const subscriptionId = process.env["RESOURCEHEALTH_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["RESOURCEHEALTH_RESOURCE_GROUP"] || "resourceGroupName";
  const expand = "recommendedactions";
  const options = {
    expand,
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availabilityStatuses.listByResourceGroup(
    resourceGroupName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
