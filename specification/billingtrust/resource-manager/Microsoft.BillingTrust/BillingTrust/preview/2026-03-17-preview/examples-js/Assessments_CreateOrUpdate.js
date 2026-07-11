const { BillingTrustClient } = require("@azure/arm-billing-trust");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an Assessment. Long-running operation — returns 200 (replace) or 201 (create) with the `Azure-AsyncOperation` polling header on both responses.
 *
 * @summary create or update an Assessment. Long-running operation — returns 200 (replace) or 201 (create) with the `Azure-AsyncOperation` polling header on both responses.
 * x-ms-original-file: 2026-03-17-preview/Assessments_CreateOrUpdate.json
 */
async function createOrUpdateTheEduAssessmentForAnEnrollment() {
  const credential = new DefaultAzureCredential();
  const client = new BillingTrustClient(credential);
  const result = await client.assessments.createOrUpdate(
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/billing-edu-rg/providers/Microsoft.Program/educationEnrollments/default",
    {
      properties: {
        assessmentType: "Edu",
        initialValues: [
          {
            kind: "eduQualification",
            domains: [
              {
                domainNames: ["students.contoso.edu", "faculty.contoso.edu"],
                tenantId: "11111111-1111-1111-1111-111111111111",
              },
            ],
          },
        ],
      },
    },
  );
  console.log(result);
}
