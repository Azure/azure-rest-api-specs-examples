const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the reservations summaries for the defined scope daily or monthly grain.
 *
 * @summary Lists the reservations summaries for the defined scope daily or monthly grain.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationSummariesDailyWithBillingProfileId.json
 */
async function reservationSummariesDailyWithBillingProfileId() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "providers/Microsoft.Billing/billingAccounts/12345:2468/billingProfiles/13579";
  const grain = "daily";
  const startDate = "2017-10-01";
  const endDate = "2017-11-20";
  const options = {
    startDate,
    endDate,
  };
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.reservationsSummaries.list(scope, grain, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

reservationSummariesDailyWithBillingProfileId().catch(console.error);
