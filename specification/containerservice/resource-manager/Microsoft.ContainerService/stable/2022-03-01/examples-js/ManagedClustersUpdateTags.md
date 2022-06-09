```javascript
const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates tags on a managed cluster.
 *
 * @summary Updates tags on a managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-03-01/examples/ManagedClustersUpdateTags.json
 */
async function updateManagedClusterTags() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const parameters = { tags: { archv3: "", tier: "testing" } };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.beginUpdateTagsAndWait(
    resourceGroupName,
    resourceName,
    parameters
  );
  console.log(result);
}

updateManagedClusterTags().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-containerservice_16.1.0-beta.2/sdk/containerservice/arm-containerservice/README.md) on how to add the SDK to your project and authenticate.
