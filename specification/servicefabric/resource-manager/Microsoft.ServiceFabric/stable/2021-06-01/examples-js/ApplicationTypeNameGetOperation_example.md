Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicefabric_2.0.1/sdk/servicefabric/arm-servicefabric/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Service Fabric application type name resource created or in the process of being created in the Service Fabric cluster resource.
 *
 * @summary Get a Service Fabric application type name resource created or in the process of being created in the Service Fabric cluster resource.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationTypeNameGetOperation_example.json
 */
async function getAnApplicationType() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const clusterName = "myCluster";
  const applicationTypeName = "myAppType";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.applicationTypes.get(
    resourceGroupName,
    clusterName,
    applicationTypeName
  );
  console.log(result);
}

getAnApplicationType().catch(console.error);
```
