const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an attestation at resource group scope.
 *
 * @summary Creates or updates an attestation at resource group scope.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-09-01/examples/Attestations_CreateResourceGroupScope.json
 */
async function createAttestationAtResourceGroupScope() {
  const subscriptionId = "35ee058e-5fa0-414c-8145-3ebb8d09b6e2";
  const resourceGroupName = "myRg";
  const attestationName = "790996e6-9871-4b1f-9cd9-ec42cd6ced1e";
  const parameters = {
    assessmentDate: new Date("2021-06-10T00:00:00Z"),
    comments: "This subscription has passed a security audit.",
    complianceState: "Compliant",
    evidence: [
      {
        description: "The results of the security audit.",
        sourceUri: "https://gist.github.com/contoso/9573e238762c60166c090ae16b814011",
      },
    ],
    expiresOn: new Date("2021-06-15T00:00:00Z"),
    metadata: { departmentId: "NYC-MARKETING-1" },
    owner: "55a32e28-3aa5-4eea-9b5a-4cd85153b966",
    policyAssignmentId:
      "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5",
    policyDefinitionReferenceId: "0b158b46-ff42-4799-8e39-08a5c23b4551",
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.attestations.beginCreateOrUpdateAtResourceGroupAndWait(
    resourceGroupName,
    attestationName,
    parameters
  );
  console.log(result);
}

createAttestationAtResourceGroupScope().catch(console.error);
