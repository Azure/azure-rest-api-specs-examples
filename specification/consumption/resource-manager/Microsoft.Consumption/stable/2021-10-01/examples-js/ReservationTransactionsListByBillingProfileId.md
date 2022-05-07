Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-consumption_9.0.1/sdk/consumption/arm-consumption/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of transactions for reserved instances on billing account scope
 *
 * @summary List of transactions for reserved instances on billing account scope
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationTransactionsListByBillingProfileId.json
 */
async function reservationTransactionsByBillingProfileId() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const filter = "properties/eventDate ge 2020-05-20 AND properties/eventDate le 2020-05-30";
  const billingAccountId =
    "fcebaabc-fced-4284-a83d-79f83dee183c:45796ba8-988f-45ad-bea9-7b71fc6c7513_2018-09-30";
  const billingProfileId = "Z76D-SGAF-BG7-TGB";
  const options = {
    filter,
  };
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.reservationTransactions.listByBillingProfile(
    billingAccountId,
    billingProfileId,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

reservationTransactionsByBillingProfileId().catch(console.error);
```
