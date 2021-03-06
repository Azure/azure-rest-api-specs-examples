const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets additional information for the specified Azure Stack Edge/Data Box Gateway device.
 *
 * @summary Gets additional information for the specified Azure Stack Edge/Data Box Gateway device.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/ExtendedInfoPost.json
 */
async function extendedInfoPost() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const resourceGroupName = "GroupForEdgeAutomation";
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.devices.getExtendedInformation(deviceName, resourceGroupName);
  console.log(result);
}

extendedInfoPost().catch(console.error);
