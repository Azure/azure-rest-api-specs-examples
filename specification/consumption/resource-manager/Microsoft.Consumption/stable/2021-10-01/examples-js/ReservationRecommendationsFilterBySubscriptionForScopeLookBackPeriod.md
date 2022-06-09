```javascript
const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of recommendations for purchasing reserved instances.
 *
 * @summary List of recommendations for purchasing reserved instances.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationRecommendationsFilterBySubscriptionForScopeLookBackPeriod.json
 */
async function reservationRecommendationsFilterBySubscriptionForScopeLookBackPeriodLegacy() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const filter = "properties/scope eq 'Single' AND properties/lookBackPeriod eq 'Last7Days'";
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.reservationRecommendations.list(scope, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

reservationRecommendationsFilterBySubscriptionForScopeLookBackPeriodLegacy().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-consumption_9.0.1/sdk/consumption/arm-consumption/README.md) on how to add the SDK to your project and authenticate.
