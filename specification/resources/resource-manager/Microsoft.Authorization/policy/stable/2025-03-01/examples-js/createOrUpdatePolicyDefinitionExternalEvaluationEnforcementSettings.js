const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to this operation creates or updates a policy definition in the given subscription with the given name.
 *
 * @summary this operation creates or updates a policy definition in the given subscription with the given name.
 * x-ms-original-file: 2025-03-01/createOrUpdatePolicyDefinitionExternalEvaluationEnforcementSettings.json
 */
async function createOrUpdateAPolicyDefinitionWithExternalEvaluationEnforcementSettings() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.policyDefinitions.createOrUpdate("RandomizeVMAllocation", {
    description:
      "Randomly disable VM allocation in eastus by having policy rule reference the outcome of invoking an external endpoint using the CoinFlip endpoint that returns random values.",
    displayName: "Randomize VM Allocation",
    externalEvaluationEnforcementSettings: {
      endpointSettings: { kind: "CoinFlip", details: { successProbability: 0.5 } },
      missingTokenAction: "audit",
      roleDefinitionIds: [
        "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/roleDefinitions/f0cc2aea-b517-48f6-8f9e-0c01c687907b",
      ],
    },
    metadata: { category: "VM" },
    mode: "Indexed",
    policyRule: {
      if: {
        allOf: [
          { equals: "Microsoft.Compute/virtualMachines", field: "type" },
          { equals: "eastus", field: "location" },
          { equals: "false", value: "[claims().isValid]" },
        ],
      },
      then: { effect: "deny" },
    },
  });
  console.log(result);
}
