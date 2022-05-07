Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-managementgroups_2.0.1/sdk/managementgroups/arm-managementgroups/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts backfilling subscriptions for the Tenant.
 *
 * @summary Starts backfilling subscriptions for the Tenant.
 * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/StartTenantBackfillRequest.json
 */
async function startTenantBackfill() {
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const result = await client.startTenantBackfill();
  console.log(result);
}

startTenantBackfill().catch(console.error);
```
