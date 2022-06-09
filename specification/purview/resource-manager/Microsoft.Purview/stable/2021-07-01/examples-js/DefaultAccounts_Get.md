```javascript
const { PurviewManagementClient } = require("@azure/arm-purview");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the default account for the scope.
 *
 * @summary Get the default account for the scope.
 * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/DefaultAccounts_Get.json
 */
async function defaultAccountsGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scopeTenantId = "12345678-1234-1234-12345678abc";
  const scopeType = "Tenant";
  const scope = "12345678-1234-1234-12345678abc";
  const options = { scope };
  const credential = new DefaultAzureCredential();
  const client = new PurviewManagementClient(credential, subscriptionId);
  const result = await client.defaultAccounts.get(scopeTenantId, scopeType, options);
  console.log(result);
}

defaultAccountsGet().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-purview_1.0.1/sdk/purview/arm-purview/README.md) on how to add the SDK to your project and authenticate.
