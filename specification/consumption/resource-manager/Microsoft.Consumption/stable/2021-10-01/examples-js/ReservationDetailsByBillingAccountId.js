const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the reservations details for the defined scope and provided date range.
 *
 * @summary Lists the reservations details for the defined scope and provided date range.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationDetailsByBillingAccountId.json
 */
async function reservationDetailsByBillingAccountId() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "providers/Microsoft.Billing/billingAccounts/12345";
  const filter = "properties/usageDate ge 2017-10-01 AND properties/usageDate le 2017-12-05";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.reservationsDetails.list(scope, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

reservationDetailsByBillingAccountId().catch(console.error);
