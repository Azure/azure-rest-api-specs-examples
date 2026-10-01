const { ComputeClient } = require("@azure/arm-compute-bulkactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the specified occurrence.
 *
 * @summary gets the specified occurrence.
 * x-ms-original-file: 2026-10-06-preview/Occurrences_Get_ComprehensiveSuccess.json
 */
async function _02ReadARecurringScheduledActionOccurrenceWithMixedResults() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ComputeClient(credential, subscriptionId);
  const result = await client.occurrences.get(
    "example-rg",
    "weekday-start",
    "88888888-8888-8888-8888-888888888888",
  );
  console.log(result);
}
