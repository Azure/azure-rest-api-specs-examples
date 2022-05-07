Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-policy_5.0.1/sdk/policy/arm-policy/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation creates or updates a policy definition in the given management group with the given name.
 *
 * @summary This operation creates or updates a policy definition in the given management group with the given name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinitionAtManagementGroup.json
 */
async function createOrUpdateAPolicyDefinitionAtManagementGroupLevel() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyDefinitionName = "ResourceNaming";
  const managementGroupId = "MyManagementGroup";
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
  const result = await client.policyDefinitions.createOrUpdateAtManagementGroup(
    policyDefinitionName,
    managementGroupId,
    parameters
  );
  console.log(result);
}

createOrUpdateAPolicyDefinitionAtManagementGroupLevel().catch(console.error);
```
