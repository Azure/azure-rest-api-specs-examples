Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-policy_5.0.1/sdk/policy/arm-policy/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation retrieves a single policy assignment, given its name and the scope it was created at.
 *
 * @summary This operation retrieves a single policy assignment, given its name and the scope it was created at.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicyAssignmentWithIdentity.json
 */
async function retrieveAPolicyAssignmentWithASystemAssignedIdentity() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const policyAssignmentName = "EnforceNaming";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.policyAssignments.get(scope, policyAssignmentName);
  console.log(result);
}

retrieveAPolicyAssignmentWithASystemAssignedIdentity().catch(console.error);
```
