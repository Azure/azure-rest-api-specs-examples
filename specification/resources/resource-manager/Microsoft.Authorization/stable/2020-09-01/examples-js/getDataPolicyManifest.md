```javascript
const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation retrieves the data policy manifest with the given policy mode.
 *
 * @summary This operation retrieves the data policy manifest with the given policy mode.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-09-01/examples/getDataPolicyManifest.json
 */
async function retrieveADataPolicyManifestByPolicyMode() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyMode = "Microsoft.KeyVault.Data";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.dataPolicyManifests.getByPolicyMode(policyMode);
  console.log(result);
}

retrieveADataPolicyManifestByPolicyMode().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-policy_5.0.1/sdk/policy/arm-policy/README.md) on how to add the SDK to your project and authenticate.
