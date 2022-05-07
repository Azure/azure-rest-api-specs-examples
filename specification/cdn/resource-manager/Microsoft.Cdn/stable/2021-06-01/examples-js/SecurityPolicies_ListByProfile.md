Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-cdn_7.0.1/sdk/cdn/arm-cdn/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists security policies associated with the profile
 *
 * @summary Lists security policies associated with the profile
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/SecurityPolicies_ListByProfile.json
 */
async function securityPoliciesListByProfile() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.securityPolicies.listByProfile(resourceGroupName, profileName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

securityPoliciesListByProfile().catch(console.error);
```
