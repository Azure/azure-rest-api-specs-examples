const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets additional information for the specified Data Box Edge/Data Box Gateway device.
 *
 * @summary Gets additional information for the specified Data Box Edge/Data Box Gateway device.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/GetUpdateExtendedInfo.json
 */
async function getUpdateExtendedInfo() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const resourceGroupName = "GroupForEdgeAutomation";
  const parameters = {};
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.devices.updateExtendedInformation(
    deviceName,
    resourceGroupName,
    parameters
  );
  console.log(result);
}

getUpdateExtendedInfo().catch(console.error);
