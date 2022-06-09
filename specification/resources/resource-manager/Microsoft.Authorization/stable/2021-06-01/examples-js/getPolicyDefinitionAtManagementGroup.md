```javascript
const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation retrieves the policy definition in the given management group with the given name.
 *
 * @summary This operation retrieves the policy definition in the given management group with the given name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicyDefinitionAtManagementGroup.json
 */
async function retrieveAPolicyDefinitionAtManagementGroupLevel() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyDefinitionName = "ResourceNaming";
  const managementGroupId = "MyManagementGroup";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.policyDefinitions.getAtManagementGroup(
    policyDefinitionName,
    managementGroupId
  );
  console.log(result);
}

retrieveAPolicyDefinitionAtManagementGroupLevel().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-policy_5.0.1/sdk/policy/arm-policy/README.md) on how to add the SDK to your project and authenticate.
