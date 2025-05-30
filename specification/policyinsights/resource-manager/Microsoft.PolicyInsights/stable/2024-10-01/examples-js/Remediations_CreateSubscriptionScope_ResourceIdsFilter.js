const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a remediation at subscription scope.
 *
 * @summary Creates or updates a remediation at subscription scope.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/Remediations_CreateSubscriptionScope_ResourceIdsFilter.json
 */
async function createRemediationAtSubscriptionScopeWithResourceIdsFilter() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "35ee058e-5fa0-414c-8145-3ebb8d09b6e2";
  const remediationName = "storageRemediation";
  const parameters = {
    failureThreshold: { percentage: 0.1 },
    filters: {
      locations: ["eastus", "westus"],
      resourceIds: [
        "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/res2627/providers/Microsoft.Storage/storageAccounts/sto1125",
        "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/testcmk3/providers/Microsoft.Storage/storageAccounts/sto3699",
        "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/res9407/providers/Microsoft.Storage/storageAccounts/sto8596",
        "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/testcmk3/providers/Microsoft.Storage/storageAccounts/sto6637",
        "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/res8186/providers/Microsoft.Storage/storageAccounts/sto834",
        "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/testcmk3/providers/Microsoft.Storage/storageAccounts/sto9174",
      ],
    },
    parallelDeployments: 6,
    policyAssignmentId:
      "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5",
    policyDefinitionReferenceId: "8c8fa9e4",
    resourceCount: 42,
    resourceDiscoveryMode: "ExistingNonCompliant",
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.remediations.createOrUpdateAtSubscription(
    remediationName,
    parameters,
  );
  console.log(result);
}
