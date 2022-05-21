Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-containerservice_16.1.0-beta.2/sdk/containerservice/arm-containerservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a snapshot.
 *
 * @summary Deletes a snapshot.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-03-01/examples/SnapshotsDelete.json
 */
async function deleteSnapshot() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "snapshot1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.snapshots.delete(resourceGroupName, resourceName);
  console.log(result);
}

deleteSnapshot().catch(console.error);
```
