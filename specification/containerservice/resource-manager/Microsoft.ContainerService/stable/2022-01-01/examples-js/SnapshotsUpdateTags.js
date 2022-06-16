const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateSnapshotTags() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "snapshot1";
  const parameters = { tags: { key2: "new-val2", key3: "val3" } };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.snapshots.updateTags(resourceGroupName, resourceName, parameters);
  console.log(result);
}

updateSnapshotTags().catch(console.error);
