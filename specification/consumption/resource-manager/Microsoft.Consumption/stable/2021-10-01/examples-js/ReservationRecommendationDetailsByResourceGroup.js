const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Details of a reservation recommendation for what-if analysis of reserved instances.
 *
 * @summary Details of a reservation recommendation for what-if analysis of reserved instances.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationRecommendationDetailsByResourceGroup.json
 */
async function reservationRecommendationsByResourceGroupLegacy() {
  const subscriptionId =
    process.env["CONSUMPTION_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope = "Single";
  const region = "westus";
  const term = "P3Y";
  const lookBackPeriod = "Last30Days";
  const product = "Standard_DS13_v2";
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
