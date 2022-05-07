Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a managed server's threat detection policy.
 *
 * @summary Get a managed server's threat detection policy.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedServerSecurityAlertGet.json
 */
async function getAManagedServerThreatDetectionPolicy() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "securityalert-4799";
  const managedInstanceName = "securityalert-6440";
  const securityAlertPolicyName = "Default";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedServerSecurityAlertPolicies.get(
    resourceGroupName,
    managedInstanceName,
    securityAlertPolicyName
  );
  console.log(result);
}

getAManagedServerThreatDetectionPolicy().catch(console.error);
```
