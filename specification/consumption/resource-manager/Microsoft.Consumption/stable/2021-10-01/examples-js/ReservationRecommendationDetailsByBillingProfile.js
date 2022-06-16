const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Details of a reservation recommendation for what-if analysis of reserved instances.
 *
 * @summary Details of a reservation recommendation for what-if analysis of reserved instances.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationRecommendationDetailsByBillingProfile.json
 */
async function reservationRecommendationsByBillingProfileModern() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "Shared";
  const region = "australiaeast";
  const term = "P1Y";
  const lookBackPeriod = "Last7Days";
  const product = "Standard_B2s";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const result = await client.reservationRecommendationDetails.get(
    scope,
    region,
    term,
    lookBackPeriod,
    product
  );
  console.log(result);
}

reservationRecommendationsByBillingProfileModern().catch(console.error);
