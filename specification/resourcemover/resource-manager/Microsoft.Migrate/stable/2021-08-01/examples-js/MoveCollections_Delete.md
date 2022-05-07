Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-resourcemover_2.0.1/sdk/resourcemover/arm-resourcemover/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a move collection.
 *
 * @summary Deletes a move collection.
 * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2021-08-01/examples/MoveCollections_Delete.json
 */
async function moveCollectionsDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const moveCollectionName = "movecollection1";
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const result = await client.moveCollections.beginDeleteAndWait(
    resourceGroupName,
    moveCollectionName
  );
  console.log(result);
}

moveCollectionsDelete().catch(console.error);
```
