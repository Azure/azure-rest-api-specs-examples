```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List information about private endpoint connections under a disk access resource
 *
 * @summary List information about private endpoint connections under a disk access resource
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/ListPrivateEndpointConnectionsInADiskAccess.json
 */
async function getInformationAboutAPrivateEndpointConnectionUnderADiskAccessResource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskAccessName = "myDiskAccess";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.diskAccesses.listPrivateEndpointConnections(
    resourceGroupName,
    diskAccessName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getInformationAboutAPrivateEndpointConnectionUnderADiskAccessResource().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.
