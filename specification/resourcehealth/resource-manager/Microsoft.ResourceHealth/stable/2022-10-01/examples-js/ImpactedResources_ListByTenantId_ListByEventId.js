const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists impacted resources in the tenant by an event.
 *
 * @summary Lists impacted resources in the tenant by an event.
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/ImpactedResources_ListByTenantId_ListByEventId.json
 */
async function listEventsByTenantId() {
  const subscriptionId =
    process.env["RESOURCEHEALTH_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const eventTrackingId = "BC_1-FXZ";
  const filter = "targetRegion eq 'westus'";
  const options = {
    filter,
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.impactedResources.listByTenantIdAndEventId(
    eventTrackingId,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
