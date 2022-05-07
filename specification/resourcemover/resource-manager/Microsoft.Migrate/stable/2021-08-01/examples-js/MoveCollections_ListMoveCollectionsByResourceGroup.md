Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-resourcemover_2.0.1/sdk/resourcemover/arm-resourcemover/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all the Move Collections in the resource group.
 *
 * @summary Get all the Move Collections in the resource group.
 * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2021-08-01/examples/MoveCollections_ListMoveCollectionsByResourceGroup.json
 */
async function moveCollectionsListMoveCollectionsByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.moveCollections.listMoveCollectionsByResourceGroup(
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

moveCollectionsListMoveCollectionsByResourceGroup().catch(console.error);
```
