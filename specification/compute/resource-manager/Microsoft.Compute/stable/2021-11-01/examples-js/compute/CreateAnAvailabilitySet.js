const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function createAnAvailabilitySet() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const availabilitySetName = "myAvailabilitySet";
  const parameters = {
    location: "westus",
    platformFaultDomainCount: 2,
    platformUpdateDomainCount: 20,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.availabilitySets.createOrUpdate(
    resourceGroupName,
    availabilitySetName,
    parameters
  );
  console.log(result);
}

createAnAvailabilitySet().catch(console.error);
