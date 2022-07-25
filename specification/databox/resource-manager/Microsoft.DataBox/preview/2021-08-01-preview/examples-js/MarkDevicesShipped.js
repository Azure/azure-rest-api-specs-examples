const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Request to mark devices for a given job as shipped
 *
 * @summary Request to mark devices for a given job as shipped
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/MarkDevicesShipped.json
 */
async function markDevicesShipped() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const jobName = "SdkJob8367";
  const resourceGroupName = "SdkRg9836";
  const markDevicesShippedRequest = {
    deliverToDcPackageDetails: { carrierName: "DHL", trackingId: "123456" },
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

markDevicesShipped().catch(console.error);
