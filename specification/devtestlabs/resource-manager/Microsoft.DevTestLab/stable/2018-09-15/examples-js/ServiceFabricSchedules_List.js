const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List schedules in a given service fabric.
 *
 * @summary List schedules in a given service fabric.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabricSchedules_List.json
 */
async function serviceFabricSchedulesList() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const userName = "@me";
  const serviceFabricName = "{serviceFrabicName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serviceFabricSchedules.list(
    resourceGroupName,
    labName,
    userName,
    serviceFabricName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

serviceFabricSchedulesList().catch(console.error);
