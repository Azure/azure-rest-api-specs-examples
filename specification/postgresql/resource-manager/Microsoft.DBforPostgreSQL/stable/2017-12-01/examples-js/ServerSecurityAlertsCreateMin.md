Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-postgresql_6.0.1/sdk/postgresql/arm-postgresql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PostgreSQLManagementClient } = require("@azure/arm-postgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a threat detection policy.
 *
 * @summary Creates or updates a threat detection policy.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerSecurityAlertsCreateMin.json
 */
async function updateAServerThreatDetectionPolicyWithMinimalParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "securityalert-4799";
  const serverName = "securityalert-6440";
  const securityAlertPolicyName = "Default";
  const parameters = {
    emailAccountAdmins: true,
    state: "Disabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementClient(credential, subscriptionId);
  const result = await client.serverSecurityAlertPolicies.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    securityAlertPolicyName,
    parameters
  );
  console.log(result);
}

updateAServerThreatDetectionPolicyWithMinimalParameters().catch(console.error);
```
