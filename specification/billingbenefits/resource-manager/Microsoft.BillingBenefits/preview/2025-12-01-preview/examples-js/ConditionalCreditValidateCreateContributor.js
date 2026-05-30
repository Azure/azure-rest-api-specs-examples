const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to validate savings plan purchase.
 *
 * @summary validate savings plan purchase.
 * x-ms-original-file: 2025-12-01-preview/ConditionalCreditValidateCreateContributor.json
 */
async function conditionalCreditValidateCreateContributor() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.benefit.validate({
    benefits: [
      {
        benefitType: "ConditionalCredits",
        properties: {
          displayName: "Contributor CACO Test",
          entityType: "Contributor",
          primaryResourceId:
            "/subscriptions/eef82110-c91b-4395-9420-fcfcbefc5a47/resourceGroups/primaryRG/providers/Microsoft.BillingBenefits/conditionalCredits/primaryCACO",
          productCode: "000187f7-0000-0260-ab43-b8473ce57f1d",
          startAt: new Date("2025-09-01T00:00:00Z"),
        },
      },
    ],
  });
  console.log(result);
}
