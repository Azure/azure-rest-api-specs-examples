const { MarketplaceOrderingAgreements } = require("@azure/arm-marketplaceordering");
const { DefaultAzureCredential } = require("@azure/identity");

async function getMarketplaceTerms() {
  const subscriptionId = "subid";
  const offerType = "virtualmachine";
  const publisherId = "pubid";
  const offerId = "offid";
  const planId = "planid";
  const credential = new DefaultAzureCredential();
  const client = new MarketplaceOrderingAgreements(credential, subscriptionId);
  const result = await client.marketplaceAgreements.get(offerType, publisherId, offerId, planId);
  console.log(result);
}

getMarketplaceTerms().catch(console.error);
