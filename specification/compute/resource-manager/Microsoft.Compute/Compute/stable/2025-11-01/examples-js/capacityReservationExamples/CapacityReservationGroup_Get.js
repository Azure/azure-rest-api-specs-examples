const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation that retrieves information about a capacity reservation group.
 *
 * @summary the operation that retrieves information about a capacity reservation group.
 * x-ms-original-file: 2025-11-01/capacityReservationExamples/CapacityReservationGroup_Get.json
 */
async function getACapacityReservationGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscriptionId}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservationGroups.get(
    "myResourceGroup",
    "myCapacityReservationGroup",
    { expand: "instanceView" },
  );
  console.log(result);
}
