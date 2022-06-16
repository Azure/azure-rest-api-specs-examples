const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateManagedClusterTags() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const parameters = { tags: { archv3: "", tier: "testing" } };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.beginUpdateTagsAndWait(
    resourceGroupName,
    resourceName,
    parameters
  );
  console.log(result);
}

updateManagedClusterTags().catch(console.error);
