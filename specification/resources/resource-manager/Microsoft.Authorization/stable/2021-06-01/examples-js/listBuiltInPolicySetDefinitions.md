```javascript
const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation retrieves a list of all the built-in policy set definitions that match the optional given $filter. If $filter='category -eq {value}' is provided, the returned list only includes all built-in policy set definitions whose category match the {value}.
 *
 * @summary This operation retrieves a list of all the built-in policy set definitions that match the optional given $filter. If $filter='category -eq {value}' is provided, the returned list only includes all built-in policy set definitions whose category match the {value}.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listBuiltInPolicySetDefinitions.json
 */
async function listBuiltInPolicySetDefinitions() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policySetDefinitions.listBuiltIn()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listBuiltInPolicySetDefinitions().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-policy_5.0.1/sdk/policy/arm-policy/README.md) on how to add the SDK to your project and authenticate.
