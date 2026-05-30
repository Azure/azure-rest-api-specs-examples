const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to validate savings plan purchase.
 *
 * @summary validate savings plan purchase.
 * x-ms-original-file: 2025-12-01-preview/MaccValidateUpdateSuppressShortfallOnMilestone.json
 */
async function maccValidateUpdateSuppressShortfallOnMilestone() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.benefit.validate({
    benefits: [
      {
        benefitType: "MACC",
        billingAccountResourceId: "/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}",
        entityType: "Primary",
        milestones: [
          {
            automaticShortfall: "Disabled",
            automaticShortfallSuppressReason: {
              code: "CustomerRequest",
              message: "Customer requested to suppress automatic shortfall",
            },
            milestoneId: "11111111-1111-1111-1111-111111111111",
          },
        ],
        productCode: "0001d726-0000-0160-330f-a0b98cdbbdc4",
        resourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.BillingBenefits/maccs/myBenefit",
        systemId: "13810867107109237",
      },
    ],
  });
  console.log(result);
}
