const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the order related to the device.
 *
 * @summary Deletes the order related to the device.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/OrderDelete.json
 */
async function orderDelete() {
  const subscriptionId =
    process.env["DATABOXEDGE_SUBSCRIPTION_ID"] || "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const resourceGroupName = process.env["DATABOXEDGE_RESOURCE_GROUP"] || "GroupForEdgeAutomation";
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.orders.beginDeleteAndWait(deviceName, resourceGroupName);
  console.log(result);
}
