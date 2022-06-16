const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

async function moveCollectionsResolveDependencies() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const moveCollectionName = "movecollection1";
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const result = await client.moveCollections.beginResolveDependenciesAndWait(
    resourceGroupName,
    moveCollectionName
  );
  console.log(result);
}

moveCollectionsResolveDependencies().catch(console.error);
