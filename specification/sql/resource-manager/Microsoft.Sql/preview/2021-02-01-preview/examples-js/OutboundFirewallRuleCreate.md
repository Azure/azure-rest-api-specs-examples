Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a outbound firewall rule with a given name.
 *
 * @summary Create a outbound firewall rule with a given name.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/OutboundFirewallRuleCreate.json
 */
async function approveOrRejectAOutboundFirewallRuleWithAGivenName() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const serverName = "sqlcrudtest-4645";
  const outboundRuleFqdn = "server.database.windows.net";
  const parameters = {};
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.outboundFirewallRules.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    outboundRuleFqdn,
    parameters
  );
  console.log(result);
}

approveOrRejectAOutboundFirewallRuleWithAGivenName().catch(console.error);
```
