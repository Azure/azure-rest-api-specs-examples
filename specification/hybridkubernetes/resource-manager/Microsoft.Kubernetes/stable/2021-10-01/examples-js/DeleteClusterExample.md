Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-hybridkubernetes_2.0.1/sdk/hybridkubernetes/arm-hybridkubernetes/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ConnectedKubernetesClient } = require("@azure/arm-hybridkubernetes");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteClusterExample() {
  const subscriptionId = "1bfbb5d0-917e-4346-9026-1d3b344417f5";
  const resourceGroupName = "k8sc-rg";
  const clusterName = "testCluster";
  const credential = new DefaultAzureCredential();
  const client = new ConnectedKubernetesClient(credential, subscriptionId);
  const result = await client.connectedClusterOperations.beginDeleteAndWait(
    resourceGroupName,
    clusterName
  );
  console.log(result);
}

deleteClusterExample().catch(console.error);
```
