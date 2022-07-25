const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Book shipment pick up.
 *
 * @summary Book shipment pick up.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/BookShipmentPickupPost.json
 */
async function bookShipmentPickupPost() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const resourceGroupName = "bvttoolrg6";
  const jobName = "TJ-636646322037905056";
  const shipmentPickUpRequest = {
    endTime: new Date("2019-09-22T18:30:00Z"),
    shipmentLocation: "Front desk",
    startTime: new Date("2019-09-20T18:30:00Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.jobs.bookShipmentPickUp(
    resourceGroupName,
    jobName,
    shipmentPickUpRequest
  );
  console.log(result);
}

bookShipmentPickupPost().catch(console.error);
