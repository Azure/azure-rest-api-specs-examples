Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the subscription capabilities available for the specified location.
 *
 * @summary Gets the subscription capabilities available for the specified location.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/LocationCapabilityListByLocation.json
 */
async function listSubscriptionCapabilitiesInTheGivenLocation() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const locationName = "eastus";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.capabilities.listByLocation(locationName);
  console.log(result);
}

listSubscriptionCapabilitiesInTheGivenLocation().catch(console.error);
```
