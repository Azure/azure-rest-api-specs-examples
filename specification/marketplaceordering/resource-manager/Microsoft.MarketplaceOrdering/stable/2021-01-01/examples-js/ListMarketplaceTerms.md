Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-marketplaceordering_3.0.1/sdk/marketplaceordering/arm-marketplaceordering/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MarketplaceOrderingAgreements } = require("@azure/arm-marketplaceordering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List marketplace agreements in the subscription.
 *
 * @summary List marketplace agreements in the subscription.
 * x-ms-original-file: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/ListMarketplaceTerms.json
 */
async function setMarketplaceTerms() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new MarketplaceOrderingAgreements(credential, subscriptionId);
  const result = await client.marketplaceAgreements.list();
  console.log(result);
}

setMarketplaceTerms().catch(console.error);
```
