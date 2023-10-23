const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Service health event details in the tenant by event tracking id. This can be used to fetch sensitive properties for Security Advisory events
 *
 * @summary Service health event details in the tenant by event tracking id. This can be used to fetch sensitive properties for Security Advisory events
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/Event_fetchDetailsByTenantIdAndTrackingId.json
 */
async function eventDetailsByTenantIdAndTrackingId() {
  const eventTrackingId = "eventTrackingId";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential);
  const result = await client.eventOperations.fetchDetailsByTenantIdAndTrackingId(eventTrackingId);
  console.log(result);
}
