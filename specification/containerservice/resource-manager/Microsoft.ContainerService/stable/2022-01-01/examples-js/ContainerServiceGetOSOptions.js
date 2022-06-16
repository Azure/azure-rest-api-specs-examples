const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function getContainerServiceOSOptions() {
  const subscriptionId = "subid1";
  const location = "location1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.getOSOptions(location);
  console.log(result);
}

getContainerServiceOSOptions().catch(console.error);
