const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Service health event in the subscription by event tracking id
 *
 * @summary Service health event in the subscription by event tracking id
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/Event_GetBySubscriptionIdAndTrackingId.json
 */
async function securityAdvisoriesEventBySubscriptionIdAndTrackingId() {
  const subscriptionId = process.env["RESOURCEHEALTH_SUBSCRIPTION_ID"] || "subscriptionId";
  const filter = "properties/status eq 'Active'";
  const queryStartTime = "7/10/2022";
  const eventTrackingId = "eventTrackingId";
  const options = {
    filter,
    queryStartTime,
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential, subscriptionId);
  const result = await client.eventOperations.getBySubscriptionIdAndTrackingId(
    eventTrackingId,
    options
  );
  console.log(result);
}
