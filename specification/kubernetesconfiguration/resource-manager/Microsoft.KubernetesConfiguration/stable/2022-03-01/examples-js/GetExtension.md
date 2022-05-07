Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-kubernetesconfiguration_5.0.0/sdk/kubernetesconfiguration/arm-kubernetesconfiguration/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SourceControlConfigurationClient } = require("@azure/arm-kubernetesconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets Kubernetes Cluster Extension.
 *
 * @summary Gets Kubernetes Cluster Extension.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/GetExtension.json
 */
async function getExtension() {
  const subscriptionId = "subId1";
  const resourceGroupName = "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const extensionName = "ClusterMonitor";
  const credential = new DefaultAzureCredential();
  const client = new SourceControlConfigurationClient(credential, subscriptionId);
  const result = await client.extensions.get(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
    extensionName
  );
  console.log(result);
}

getExtension().catch(console.error);
```
