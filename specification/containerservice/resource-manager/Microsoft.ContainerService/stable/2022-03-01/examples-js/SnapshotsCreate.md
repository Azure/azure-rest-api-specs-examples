Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-containerservice_16.1.0-beta.1/sdk/containerservice/arm-containerservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a snapshot.
 *
 * @summary Creates or updates a snapshot.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-03-01/examples/SnapshotsCreate.json
 */
async function createOrUpdateSnapshot() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "snapshot1";
  const parameters = {
    creationData: {
      sourceResourceId:
        "/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster1/agentPools/pool0",
    },
    location: "westus",
    tags: { key1: "val1", key2: "val2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.snapshots.createOrUpdate(resourceGroupName, resourceName, parameters);
  console.log(result);
}

createOrUpdateSnapshot().catch(console.error);
```
