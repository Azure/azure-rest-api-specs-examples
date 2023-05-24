const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Request to mark devices for a given job as shipped
 *
 * @summary Request to mark devices for a given job as shipped
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/MarkDevicesShipped.json
 */
async function markDevicesShipped() {
  const subscriptionId = process.env["DATABOX_SUBSCRIPTION_ID"] || "YourSubscriptionId";
  const jobName = "TestJobName1";
  const resourceGroupName = process.env["DATABOX_RESOURCE_GROUP"] || "YourResourceGroupName";
  const markDevicesShippedRequest = {
    deliverToDcPackageDetails: {
      carrierName: "testCarrier",
      trackingId: "000000",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.jobs.markDevicesShipped(
    jobName,
    resourceGroupName,
    markDevicesShippedRequest
  );
  console.log(result);
}
