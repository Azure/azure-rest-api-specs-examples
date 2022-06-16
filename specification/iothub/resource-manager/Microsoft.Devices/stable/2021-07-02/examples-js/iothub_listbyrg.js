const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all the IoT hubs in a resource group.
 *
 * @summary Get all the IoT hubs in a resource group.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_listbyrg.json
 */
async function iotHubResourceListByResourceGroup() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.iotHubResource.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

iotHubResourceListByResourceGroup().catch(console.error);
