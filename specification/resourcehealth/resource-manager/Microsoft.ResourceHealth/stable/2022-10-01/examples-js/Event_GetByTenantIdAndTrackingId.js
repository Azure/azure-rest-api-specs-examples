const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Service health event in the tenant by event tracking id
 *
 * @summary Service health event in the tenant by event tracking id
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/Event_GetByTenantIdAndTrackingId.json
 */
async function eventByTenantIdAndTrackingId() {
  const subscriptionId =
    process.env["RESOURCEHEALTH_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const filter = "properties/status eq 'Active'";
  const queryStartTime = "7/10/2022";
  const eventTrackingId = "eventTrackingId";
  const options = {
    filter,
    queryStartTime,
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential, subscriptionId);
  const result = await client.eventOperations.getByTenantIdAndTrackingId(eventTrackingId, options);
  console.log(result);
}
