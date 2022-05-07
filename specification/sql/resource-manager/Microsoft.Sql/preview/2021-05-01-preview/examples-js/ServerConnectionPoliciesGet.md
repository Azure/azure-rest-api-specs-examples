Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a server connection policy
 *
 * @summary Gets a server connection policy
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ServerConnectionPoliciesGet.json
 */
async function getsAServerConnectionPolicy() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "rgtest-12";
  const serverName = "servertest-6285";
  const connectionPolicyName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.serverConnectionPolicies.get(
    resourceGroupName,
    serverName,
    connectionPolicyName
  );
  console.log(result);
}

getsAServerConnectionPolicy().catch(console.error);
```
