const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

async function moveCollectionsCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const moveCollectionName = "movecollection1";
  const body = {
    identity: { type: "SystemAssigned" },
    location: "eastus2",
    properties: { sourceRegion: "eastus", targetRegion: "westus" },
  };
  const options = { body: body };
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const result = await client.moveCollections.create(
    resourceGroupName,
    moveCollectionName,
    options
  );
  console.log(result);
}

moveCollectionsCreate().catch(console.error);
