Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a database's security alert policy.
 *
 * @summary Creates or updates a database's security alert policy.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseSecurityAlertCreateMin.json
 */
async function updateADatabaseThreatDetectionPolicyWithMinimalParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "securityalert-4799";
  const serverName = "securityalert-6440";
  const databaseName = "testdb";
  const securityAlertPolicyName = "Default";
  const parameters = { state: "Enabled" };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseSecurityAlertPolicies.createOrUpdate(
    resourceGroupName,
    serverName,
    databaseName,
    securityAlertPolicyName,
    parameters
  );
  console.log(result);
}

updateADatabaseThreatDetectionPolicyWithMinimalParameters().catch(console.error);
```
