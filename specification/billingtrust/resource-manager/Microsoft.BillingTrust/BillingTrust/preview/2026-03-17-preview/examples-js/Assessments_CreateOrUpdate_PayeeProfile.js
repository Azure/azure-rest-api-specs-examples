const { BillingTrustClient } = require("@azure/arm-billing-trust");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an Assessment. Long-running operation — returns 200 (replace) or 201 (create) with the `Azure-AsyncOperation` polling header on both responses.
 *
 * @summary create or update an Assessment. Long-running operation — returns 200 (replace) or 201 (create) with the `Azure-AsyncOperation` polling header on both responses.
 * x-ms-original-file: 2026-03-17-preview/Assessments_CreateOrUpdate_PayeeProfile.json
 */
async function createOrUpdateThePayeeProfileAssessmentForABillingAccount() {
  const credential = new DefaultAzureCredential();
  const client = new BillingTrustClient(credential);
  const result = await client.assessments.createOrUpdate(
    "providers/Microsoft.Billing/billingAccounts/abc123:00000000-0000-0000-0000-000000000000_2019-05-31",
    { properties: { assessmentType: "PayeeProfile" } },
  );
  console.log(result);
}
