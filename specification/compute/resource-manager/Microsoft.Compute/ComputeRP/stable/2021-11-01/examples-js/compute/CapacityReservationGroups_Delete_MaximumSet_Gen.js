const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function capacityReservationGroupsDeleteMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const capacityReservationGroupName = "a";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservationGroups.delete(
    resourceGroupName,
    capacityReservationGroupName
  );
  console.log(result);
}

capacityReservationGroupsDeleteMaximumSetGen().catch(console.error);
