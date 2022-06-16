const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the StorageAccounts in a Data Box Edge/Data Box Gateway device.
 *
 * @summary Lists all the StorageAccounts in a Data Box Edge/Data Box Gateway device.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/StorageAccountGetAllInDevice.json
 */
async function storageAccountGetAllInDevice() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const resourceGroupName = "GroupForEdgeAutomation";
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.storageAccounts.listByDataBoxEdgeDevice(
    deviceName,
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

storageAccountGetAllInDevice().catch(console.error);
