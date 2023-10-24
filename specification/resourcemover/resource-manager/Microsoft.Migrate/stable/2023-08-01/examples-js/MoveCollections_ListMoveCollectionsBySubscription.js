const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all the Move Collections in the subscription.
 *
 * @summary Get all the Move Collections in the subscription.
 * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_ListMoveCollectionsBySubscription.json
 */
async function moveCollectionsListMoveCollectionsBySubscription() {
  const subscriptionId = process.env["RESOURCEMOVER_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.moveCollections.listMoveCollectionsBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
