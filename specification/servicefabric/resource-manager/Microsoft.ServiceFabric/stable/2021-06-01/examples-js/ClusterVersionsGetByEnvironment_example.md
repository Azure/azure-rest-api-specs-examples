```javascript
const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about an available Service Fabric cluster code version by environment.
 *
 * @summary Gets information about an available Service Fabric cluster code version by environment.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterVersionsGetByEnvironment_example.json
 */
async function getClusterVersionByEnvironment() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const environment = "Windows";
  const clusterVersion = "6.1.480.9494";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.clusterVersions.getByEnvironment(
    location,
    environment,
    clusterVersion
  );
  console.log(result);
}

getClusterVersionByEnvironment().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicefabric_2.0.1/sdk/servicefabric/arm-servicefabric/README.md) on how to add the SDK to your project and authenticate.
