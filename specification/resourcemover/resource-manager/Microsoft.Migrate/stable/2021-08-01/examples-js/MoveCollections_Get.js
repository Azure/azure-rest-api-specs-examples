const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

async function moveCollectionsGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const moveCollectionName = "movecollection1";
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const result = await client.moveCollections.get(resourceGroupName, moveCollectionName);
  console.log(result);
}

moveCollectionsGet().catch(console.error);
