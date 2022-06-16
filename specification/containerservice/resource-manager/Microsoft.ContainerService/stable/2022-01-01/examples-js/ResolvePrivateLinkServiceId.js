const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function resolveThePrivateLinkServiceIdForManagedCluster() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const parameters = { name: "management" };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.resolvePrivateLinkServiceId.post(
    resourceGroupName,
    resourceName,
    parameters
  );
  console.log(result);
}

resolveThePrivateLinkServiceIdForManagedCluster().catch(console.error);
