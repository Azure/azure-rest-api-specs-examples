Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the server DNS alias with the given name.
 *
 * @summary Deletes the server DNS alias with the given name.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerDnsAliasDelete.json
 */
async function deleteServerDnsAlias() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const serverName = "dns-alias-server";
  const dnsAliasName = "dns-alias-name-1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.serverDnsAliases.beginDeleteAndWait(
    resourceGroupName,
    serverName,
    dnsAliasName
  );
  console.log(result);
}

deleteServerDnsAlias().catch(console.error);
```
