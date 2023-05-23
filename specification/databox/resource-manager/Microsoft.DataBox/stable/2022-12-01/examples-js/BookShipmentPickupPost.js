const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Book shipment pick up.
 *
 * @summary Book shipment pick up.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/BookShipmentPickupPost.json
 */
async function bookShipmentPickupPost() {
  const subscriptionId = process.env["DATABOX_SUBSCRIPTION_ID"] || "YourSubscriptionId";
  const resourceGroupName = process.env["DATABOX_RESOURCE_GROUP"] || "YourResourceGroupName";
  const jobName = "TestJobName1";
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
