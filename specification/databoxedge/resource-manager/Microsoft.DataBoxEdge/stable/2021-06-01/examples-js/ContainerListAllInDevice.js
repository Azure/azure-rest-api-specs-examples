const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the containers of a storage Account in a Data Box Edge/Data Box Gateway device.
 *
 * @summary Lists all the containers of a storage Account in a Data Box Edge/Data Box Gateway device.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/ContainerListAllInDevice.json
 */
async function containerListAllInDevice() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const storageAccountName = "storageaccount1";
  const resourceGroupName = "GroupForEdgeAutomation";
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.containers.listByStorageAccount(
    deviceName,
    storageAccountName,
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

containerListAllInDevice().catch(console.error);
