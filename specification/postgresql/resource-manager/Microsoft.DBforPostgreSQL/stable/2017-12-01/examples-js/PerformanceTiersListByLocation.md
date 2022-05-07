Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-postgresql_6.0.1/sdk/postgresql/arm-postgresql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PostgreSQLManagementClient } = require("@azure/arm-postgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the performance tiers at specified location in a given subscription.
 *
 * @summary List all the performance tiers at specified location in a given subscription.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/PerformanceTiersListByLocation.json
 */
async function performanceTiersList() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const locationName = "WestUS";
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.locationBasedPerformanceTier.list(locationName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

performanceTiersList().catch(console.error);
```
