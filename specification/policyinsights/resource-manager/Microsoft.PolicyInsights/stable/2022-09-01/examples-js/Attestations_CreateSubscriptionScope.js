const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an attestation at subscription scope.
 *
 * @summary Creates or updates an attestation at subscription scope.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-09-01/examples/Attestations_CreateSubscriptionScope.json
 */
async function createAttestationAtSubscriptionScope() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "35ee058e-5fa0-414c-8145-3ebb8d09b6e2";
  const attestationName = "790996e6-9871-4b1f-9cd9-ec42cd6ced1e";
  const parameters = {
    complianceState: "Compliant",
    policyAssignmentId:
      "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5",
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.attestations.beginCreateOrUpdateAtSubscriptionAndWait(
    attestationName,
    parameters
  );
  console.log(result);
}
