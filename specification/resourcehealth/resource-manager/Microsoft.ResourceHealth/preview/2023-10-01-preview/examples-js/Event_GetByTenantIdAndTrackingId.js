const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Service health event in the tenant by event tracking id
 *
 * @summary Service health event in the tenant by event tracking id
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/Event_GetByTenantIdAndTrackingId.json
 */
async function eventByTenantIdAndTrackingId() {
  const filter = "properties/status eq 'Active'";
  const queryStartTime = "7/10/2022";
  const eventTrackingId = "eventTrackingId";
  const options = {
    filter,
    queryStartTime,
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential);
  const result = await client.eventOperations.getByTenantIdAndTrackingId(eventTrackingId, options);
  console.log(result);
}
