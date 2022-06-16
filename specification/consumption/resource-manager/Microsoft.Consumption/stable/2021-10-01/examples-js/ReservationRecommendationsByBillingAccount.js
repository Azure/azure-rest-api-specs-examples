const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of recommendations for purchasing reserved instances.
 *
 * @summary List of recommendations for purchasing reserved instances.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationRecommendationsByBillingAccount.json
 */
async function reservationRecommendationsByBillingAccountLegacy() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "providers/Microsoft.Billing/billingAccounts/123456";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.reservationRecommendations.list(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}

reservationRecommendationsByBillingAccountLegacy().catch(console.error);
