Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-containerservice_16.1.0-beta.2/sdk/containerservice/arm-containerservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to To learn more about private clusters, see: https://docs.microsoft.com/azure/aks/private-clusters
 *
 * @summary To learn more about private clusters, see: https://docs.microsoft.com/azure/aks/private-clusters
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-03-01/examples/PrivateLinkResourcesList.json
 */
async function listPrivateLinkResourcesByManagedCluster() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.privateLinkResources.list(resourceGroupName, resourceName);
  console.log(result);
}

listPrivateLinkResourcesByManagedCluster().catch(console.error);
```
