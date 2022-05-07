Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicefabric_2.0.1/sdk/servicefabric/arm-servicefabric/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a Service Fabric application resource with the specified name.
 *
 * @summary Update a Service Fabric application resource with the specified name.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationPatchOperation_example.json
 */
async function patchAnApplication() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const clusterName = "myCluster";
  const applicationName = "myApp";
  const parameters = {
    location: "eastus",
    metrics: [
      {
        name: "metric1",
        maximumCapacity: 3,
        reservationCapacity: 1,
        totalApplicationCapacity: 5,
      },
    ],
    removeApplicationCapacity: false,
    tags: {},
    typeVersion: "1.0",
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.applications.beginUpdateAndWait(
    resourceGroupName,
    clusterName,
    applicationName,
    parameters
  );
  console.log(result);
}

patchAnApplication().catch(console.error);
```
