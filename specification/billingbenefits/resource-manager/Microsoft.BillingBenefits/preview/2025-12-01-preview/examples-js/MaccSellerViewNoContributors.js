const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list maccs by billing account
 *
 * @summary list maccs by billing account
 * x-ms-original-file: 2025-12-01-preview/MaccSellerViewNoContributors.json
 */
async function sellerViewNoContributors() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.sellerResource.list({
    filter: "properties/status ne 'Canceled'",
    billingAccountResourceId: "/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}",
  });
  console.log(result);
}
