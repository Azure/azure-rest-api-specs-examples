const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

async function moveCollectionsDiscard() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const moveCollectionName = "movecollection1";
  const body = {
    moveResources: [
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1",
    ],
    validateOnly: false,
  };
  const options = { body: body };
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const result = await client.moveCollections.beginDiscardAndWait(
    resourceGroupName,
    moveCollectionName,
    options
  );
  console.log(result);
}

moveCollectionsDiscard().catch(console.error);
