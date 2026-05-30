const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to validate savings plan purchase.
 *
 * @summary validate savings plan purchase.
 * x-ms-original-file: 2025-12-01-preview/ConditionalCreditValidateCancel.json
 */
async function conditionalCreditValidateCancel() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.benefit.validate({
    benefits: [
      {
        benefitType: "ConditionalCredits",
        properties: {
          billingAccountResourceId:
            "/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:11111111-1111-1111-1111-111111111111_2025-10-28",
          entityType: "Primary",
          resourceId:
            "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/conditionalCredits/conditionalCredit_20250801",
          status: "Canceled",
          systemId: "CACO-SYSTEM-20250801",
        },
      },
    ],
  });
  console.log(result);
}
