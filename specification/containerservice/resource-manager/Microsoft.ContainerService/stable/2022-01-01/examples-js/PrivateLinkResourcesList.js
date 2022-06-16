const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function listPrivateLinkResourcesByManagedCluster() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.privateLinkResources.list(resourceGroupName, resourceName);
  console.log(result);
}

listPrivateLinkResourcesByManagedCluster().catch(console.error);
