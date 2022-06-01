Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update a capacity reservation.
 *
 * @summary The operation to update a capacity reservation.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/capacityReservationExamples/CapacityReservation_Update_MaximumSet_Gen.json
 */
async function capacityReservationsUpdateMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const capacityReservationGroupName = "aaaaaaaaaa";
  const capacityReservationName = "aaaaaaaaaaaaaaaaaaa";
  const parameters = {
    instanceView: {
      statuses: [
        {
          code: "aaaaaaaaaaaaaaaaaaaaaaa",
          displayStatus: "aaaaaa",
          level: "Info",
          message: "a",
          time: new Date("2021-11-30T12:58:26.522Z"),
        },
      ],
      utilizationInfo: {},
    },
    sku: { name: "DSv3-Type1", capacity: 7, tier: "aaa" },
    tags: { key4974: "aaaaaaaaaaaaaaaa" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservations.beginUpdateAndWait(
    resourceGroupName,
    capacityReservationGroupName,
    capacityReservationName,
    parameters
  );
  console.log(result);
}

capacityReservationsUpdateMaximumSetGen().catch(console.error);
```
