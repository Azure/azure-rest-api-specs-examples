const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Service health event details in the tenant by event tracking id. This can be used to fetch sensitive properties for Security Advisory events
 *
 * @summary Service health event details in the tenant by event tracking id. This can be used to fetch sensitive properties for Security Advisory events
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/Event_fetchDetailsByTenantIdAndTrackingId.json
 */
async function eventDetailsByTenantIdAndTrackingId() {
  const subscriptionId =
    process.env["RESOURCEHEALTH_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const eventTrackingId = "eventTrackingId";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential, subscriptionId);
  const result = await client.eventOperations.fetchDetailsByTenantIdAndTrackingId(eventTrackingId);
  console.log(result);
}
