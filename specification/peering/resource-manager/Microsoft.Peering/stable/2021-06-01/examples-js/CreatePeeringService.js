const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function createAPeeringService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringServiceName = "peeringServiceName";
  const peeringService = {
    location: "eastus",
    peeringServiceLocation: "state1",
    peeringServiceProvider: "serviceProvider1",
    providerBackupPeeringLocation: "peeringLocation2",
    providerPrimaryPeeringLocation: "peeringLocation1",
  };
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peeringServices.createOrUpdate(
    resourceGroupName,
    peeringServiceName,
    peeringService
  );
  console.log(result);
}

createAPeeringService().catch(console.error);
