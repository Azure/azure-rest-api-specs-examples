const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a reservation order alias.
 *
 * @summary Get a reservation order alias.
 * x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/ReservationOrderAliasGet.json
 */
async function reservationOrderAliasGet() {
  const reservationOrderAliasName = "reservationOrderAlias123";
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.reservationOrderAlias.get(reservationOrderAliasName);
  console.log(result);
}

reservationOrderAliasGet().catch(console.error);
