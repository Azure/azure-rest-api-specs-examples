const { MarketplaceOrderingAgreements } = require("@azure/arm-marketplaceordering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List marketplace agreements in the subscription.
 *
 * @summary List marketplace agreements in the subscription.
 * x-ms-original-file: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/ListMarketplaceTerms.json
 */
async function setMarketplaceTerms() {
  const subscriptionId = process.env["MARKETPLACEORDERING_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new MarketplaceOrderingAgreements(credential, subscriptionId);
  const result = await client.marketplaceAgreements.list();
  console.log(result);
}
