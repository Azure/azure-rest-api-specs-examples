Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-trafficmanager_6.0.1/sdk/trafficmanager/arm-trafficmanager/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the default Geographic Hierarchy used by the Geographic traffic routing method.
 *
 * @summary Gets the default Geographic Hierarchy used by the Geographic traffic routing method.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/GeographicHierarchy-GET-default.json
 */
async function geographicHierarchyGetDefault() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.geographicHierarchies.getDefault();
  console.log(result);
}

geographicHierarchyGetDefault().catch(console.error);
```
