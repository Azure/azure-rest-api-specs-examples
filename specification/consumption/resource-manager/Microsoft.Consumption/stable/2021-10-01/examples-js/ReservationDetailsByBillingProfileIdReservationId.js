const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the reservations details for the defined scope and provided date range.
 *
 * @summary Lists the reservations details for the defined scope and provided date range.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationDetailsByBillingProfileIdReservationId.json
 */
async function reservationDetailsByBillingProfileIdReservationId() {
  const subscriptionId =
    process.env["CONSUMPTION_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope = "providers/Microsoft.Billing/billingAccounts/12345:2468/billingProfiles/13579";
  const startDate = "2019-09-01";
  const endDate = "2019-10-31";
  const reservationId = "1c6b6358-709f-484c-85f1-72e862a0cf3b";
  const reservationOrderId = "9f39ba10-794f-4dcb-8f4b-8d0cb47c27dc";
  const options = {
    startDate,
    endDate,
    reservationId,
    reservationOrderId,
  };
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.reservationsDetails.list(scope, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
