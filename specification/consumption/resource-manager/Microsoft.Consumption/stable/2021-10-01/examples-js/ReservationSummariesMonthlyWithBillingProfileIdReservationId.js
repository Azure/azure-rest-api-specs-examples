const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the reservations summaries for the defined scope daily or monthly grain.
 *
 * @summary Lists the reservations summaries for the defined scope daily or monthly grain.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationSummariesMonthlyWithBillingProfileIdReservationId.json
 */
async function reservationSummariesMonthlyWithBillingProfileIdReservationId() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "providers/Microsoft.Billing/billingAccounts/12345:2468/billingProfiles/13579";
  const grain = "monthly";
  const reservationId = "1c6b6358-709f-484c-85f1-72e862a0cf3b";
  const reservationOrderId = "9f39ba10-794f-4dcb-8f4b-8d0cb47c27dc";
  const options = {
    reservationId,
    reservationOrderId,
  };
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.reservationsSummaries.list(scope, grain, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

reservationSummariesMonthlyWithBillingProfileIdReservationId().catch(console.error);
