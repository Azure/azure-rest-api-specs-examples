const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all of the capacity reservations in the specified capacity reservation group. Use the nextLink property in the response to get the next page of capacity reservations.
 *
 * @summary lists all of the capacity reservations in the specified capacity reservation group. Use the nextLink property in the response to get the next page of capacity reservations.
 * x-ms-original-file: 2026-03-01/capacityReservationExamples/BlockCapacityReservation_ListByReservationGroup.json
 */
async function listBlockCapacityReservationsInReservationGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.capacityReservations.listByCapacityReservationGroup(
    "myResourceGroup",
    "blockCapacityReservationGroup",
    { expand: "virtualMachines/$ref" },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
