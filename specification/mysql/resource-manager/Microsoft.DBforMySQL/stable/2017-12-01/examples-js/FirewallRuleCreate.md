Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-mysql_5.0.1/sdk/mysql/arm-mysql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new firewall rule or updates an existing firewall rule.
 *
 * @summary Creates a new firewall rule or updates an existing firewall rule.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/FirewallRuleCreate.json
 */
async function firewallRuleCreate() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "TestGroup";
  const serverName = "testserver";
  const firewallRuleName = "rule1";
  const parameters = {
    endIpAddress: "255.255.255.255",
    startIpAddress: "0.0.0.0",
  };
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.firewallRules.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    firewallRuleName,
    parameters
  );
  console.log(result);
}

firewallRuleCreate().catch(console.error);
```
