const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an attestation at subscription scope.
 *
 * @summary creates or updates an attestation at subscription scope.
 * x-ms-original-file: 2024-10-01/Attestations_CreateSubscriptionScope.json
 */
async function createAttestationAtSubscriptionScope() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "35ee058e-5fa0-414c-8145-3ebb8d09b6e2";
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.attestations.createOrUpdateAtSubscription(
    "790996e6-9871-4b1f-9cd9-ec42cd6ced1e",
    {
      complianceState: "Compliant",
      policyAssignmentId:
        "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5",
    },
  );
  console.log(result);
}
