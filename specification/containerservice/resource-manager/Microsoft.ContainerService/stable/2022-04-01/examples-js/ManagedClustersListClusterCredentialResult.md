Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-containerservice_16.1.0/sdk/containerservice/arm-containerservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the cluster monitoring user credentials of a managed cluster.
 *
 * @summary Lists the cluster monitoring user credentials of a managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-04-01/examples/ManagedClustersListClusterCredentialResult.json
 */
async function getManagedCluster() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.listClusterMonitoringUserCredentials(
    resourceGroupName,
    resourceName
  );
  console.log(result);
}

getManagedCluster().catch(console.error);
```
