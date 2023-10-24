const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the Move Resources in the move collection.
 *
 * @summary Lists the Move Resources in the move collection.
 * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveResources_List.json
 */
async function moveResourcesList() {
  const subscriptionId = process.env["RESOURCEMOVER_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["RESOURCEMOVER_RESOURCE_GROUP"] || "rg1";
  const moveCollectionName = "movecollection1";
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.moveResources.list(resourceGroupName, moveCollectionName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
