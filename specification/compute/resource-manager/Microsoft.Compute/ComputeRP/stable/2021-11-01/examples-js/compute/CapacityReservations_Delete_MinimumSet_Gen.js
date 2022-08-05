const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function capacityReservationsDeleteMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const capacityReservationGroupName = "aaa";
  const capacityReservationName = "aaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservations.beginDeleteAndWait(
    resourceGroupName,
    capacityReservationGroupName,
    capacityReservationName
  );
  console.log(result);
}

capacityReservationsDeleteMinimumSetGen().catch(console.error);
