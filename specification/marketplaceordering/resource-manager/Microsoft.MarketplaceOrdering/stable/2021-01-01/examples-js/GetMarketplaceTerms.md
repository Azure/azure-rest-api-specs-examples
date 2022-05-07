Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-marketplaceordering_3.0.1/sdk/marketplaceordering/arm-marketplaceordering/README.md) on how to add the SDK to your project and authenticate.

```javascript
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
```
