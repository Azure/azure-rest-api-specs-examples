```javascript
const { MarketplaceOrderingAgreements } = require("@azure/arm-marketplaceordering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Save marketplace terms.
 *
 * @summary Save marketplace terms.
 * x-ms-original-file: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/SetMarketplaceTerms.json
 */
async function setMarketplaceTerms() {
  const subscriptionId = "subid";
  const offerType = "virtualmachine";
  const publisherId = "pubid";
  const offerId = "offid";
  const planId = "planid";
  const parameters = {
    accepted: false,
    licenseTextLink: "test.licenseLink",
    marketplaceTermsLink: "test.marketplaceTermsLink",
    plan: "planid",
    privacyPolicyLink: "test.privacyPolicyLink",
    product: "offid",
    publisher: "pubid",
    retrieveDatetime: "2017-08-15T11:33:07.12132Z",
    signature: "ASDFSDAFWEFASDGWERLWER",
  };
  const credential = new DefaultAzureCredential();
  const client = new MarketplaceOrderingAgreements(credential, subscriptionId);
  const result = await client.marketplaceAgreements.create(
    offerType,
    publisherId,
    offerId,
    planId,
    parameters
  );
  console.log(result);
}

setMarketplaceTerms().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-marketplaceordering_3.0.1/sdk/marketplaceordering/arm-marketplaceordering/README.md) on how to add the SDK to your project and authenticate.
