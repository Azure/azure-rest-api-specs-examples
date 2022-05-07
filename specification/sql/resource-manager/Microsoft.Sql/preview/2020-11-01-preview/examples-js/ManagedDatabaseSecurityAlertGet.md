Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a managed database's security alert policy.
 *
 * @summary Gets a managed database's security alert policy.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSecurityAlertGet.json
 */
async function getADatabaseThreatDetectionPolicy() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "securityalert-6852";
  const managedInstanceName = "securityalert-2080";
  const databaseName = "testdb";
  const securityAlertPolicyName = "Default";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabaseSecurityAlertPolicies.get(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    securityAlertPolicyName
  );
  console.log(result);
}

getADatabaseThreatDetectionPolicy().catch(console.error);
```
