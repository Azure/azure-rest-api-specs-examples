const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation creates or updates a policy definition in the given subscription with the given name.
 *
 * @summary This operation creates or updates a policy definition in the given subscription with the given name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinition.json
 */
async function createOrUpdateAPolicyDefinition() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const policyDefinitionName = "ResourceNaming";
  const parameters = {
    description: "Force resource names to begin with given 'prefix' and/or end with given 'suffix'",
    displayName: "Enforce resource naming convention",
    metadata: { category: "Naming" },
    mode: "All",
    parameters: {
      prefix: {
        type: "String",
        metadata: { description: "Resource name prefix", displayName: "Prefix" },
      },
      suffix: {
        type: "String",
        metadata: { description: "Resource name suffix", displayName: "Suffix" },
      },
    },
    policyRule: {
      if: {
        not: {
          field: "name",
          like: "[concat(parameters('prefix'), '*', parameters('suffix'))]",
        },
      },
      then: { effect: "deny" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.policyDefinitions.createOrUpdate(policyDefinitionName, parameters);
  console.log(result);
}
