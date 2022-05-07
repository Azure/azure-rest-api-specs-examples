Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-policy_5.0.1/sdk/policy/arm-policy/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation creates or updates a policy set definition in the given subscription with the given name.
 *
 * @summary This operation creates or updates a policy set definition in the given subscription with the given name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicySetDefinition.json
 */
async function createOrUpdateAPolicySetDefinition() {
  const subscriptionId = "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const policySetDefinitionName = "CostManagement";
  const parameters = {
    description: "Policies to enforce low cost storage SKUs",
    displayName: "Cost Management",
    metadata: { category: "Cost Management" },
    parameters: {
      namePrefix: {
        type: "String",
        defaultValue: "myPrefix",
        metadata: { displayName: "Prefix to enforce on resource names" },
      },
    },
    policyDefinitions: [
      {
        parameters: {
          listOfAllowedSKUs: { value: ["Standard_GRS", "Standard_LRS"] },
        },
        policyDefinitionId:
          "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1",
        policyDefinitionReferenceId: "Limit_Skus",
      },
      {
        parameters: {
          prefix: { value: "[parameters('namePrefix')]" },
          suffix: { value: "-LC" },
        },
        policyDefinitionId:
          "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming",
        policyDefinitionReferenceId: "Resource_Naming",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.policySetDefinitions.createOrUpdate(
    policySetDefinitionName,
    parameters
  );
  console.log(result);
}

createOrUpdateAPolicySetDefinition().catch(console.error);
```
