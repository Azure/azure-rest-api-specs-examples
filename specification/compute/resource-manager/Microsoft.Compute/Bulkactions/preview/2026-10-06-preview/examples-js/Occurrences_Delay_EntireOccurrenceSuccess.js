const { ComputeClient } = require("@azure/arm-compute-bulkactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delays the specified occurrence for the specified resource IDs.
 *
 * @summary delays the specified occurrence for the specified resource IDs.
 * x-ms-original-file: 2026-10-06-preview/Occurrences_Delay_EntireOccurrenceSuccess.json
 */
async function _02DelayAllOperationsInARecurringScheduledActionOccurrence() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ComputeClient(credential, subscriptionId);
  const result = await client.occurrences.delay(
    "example-rg",
    "weekday-start",
    "77777777-7777-7777-7777-777777777777",
    { delay: "2026-09-15T09:00:00-07:00", resourceIds: [] },
  );
  console.log(result);
}
