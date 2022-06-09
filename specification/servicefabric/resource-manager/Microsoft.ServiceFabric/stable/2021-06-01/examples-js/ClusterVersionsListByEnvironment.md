```javascript
const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all available code versions for Service Fabric cluster resources by environment.
 *
 * @summary Gets all available code versions for Service Fabric cluster resources by environment.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterVersionsListByEnvironment.json
 */
async function listClusterVersionsByEnvironment() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const environment = "Windows";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.clusterVersions.listByEnvironment(location, environment);
  console.log(result);
}

listClusterVersionsByEnvironment().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicefabric_2.0.1/sdk/servicefabric/arm-servicefabric/README.md) on how to add the SDK to your project and authenticate.
