const { IotCentralClient } = require("@azure/arm-iotcentral");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all private link resources of a IoT Central Application.
 *
 * @summary Get all private link resources of a IoT Central Application.
 * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateLinks_List.json
 */
async function privateLinksList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const resourceName = "myIoTCentralApp";
  const credential = new DefaultAzureCredential();
  const client = new IotCentralClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinks.list(resourceGroupName, resourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

privateLinksList().catch(console.error);
