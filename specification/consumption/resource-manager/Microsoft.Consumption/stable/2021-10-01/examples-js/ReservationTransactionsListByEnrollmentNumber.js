const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of transactions for reserved instances on billing account scope
 *
 * @summary List of transactions for reserved instances on billing account scope
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationTransactionsListByEnrollmentNumber.json
 */
async function reservationTransactionsByEnrollmentNumber() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const filter = "properties/eventDate+ge+2020-05-20+AND+properties/eventDate+le+2020-05-30";
  const billingAccountId = "123456";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.reservationTransactions.list(billingAccountId, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

reservationTransactionsByEnrollmentNumber().catch(console.error);
