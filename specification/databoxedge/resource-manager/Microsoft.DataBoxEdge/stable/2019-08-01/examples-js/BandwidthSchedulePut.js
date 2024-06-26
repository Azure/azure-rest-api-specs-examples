const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a bandwidth schedule.
 *
 * @summary Creates or updates a bandwidth schedule.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/BandwidthSchedulePut.json
 */
async function bandwidthSchedulePut() {
  const subscriptionId =
    process.env["DATABOXEDGE_SUBSCRIPTION_ID"] || "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const name = "bandwidth-1";
  const resourceGroupName = process.env["DATABOXEDGE_RESOURCE_GROUP"] || "GroupForEdgeAutomation";
  const parameters = {
    days: ["Sunday", "Monday"],
    rateInMbps: 100,
    start: "0:0:0",
    stop: "13:59:0",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.bandwidthSchedules.beginCreateOrUpdateAndWait(
    deviceName,
    name,
    resourceGroupName,
    parameters
  );
  console.log(result);
}
